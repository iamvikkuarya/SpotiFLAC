package backend

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ParallelConfig holds settings for parallel downloads
type ParallelConfig struct {
	MaxConcurrent int  `json:"max_concurrent"`
	Enabled       bool `json:"enabled"`
}

// StreamingURLResult holds pre-fetched streaming URLs for a track
type StreamingURLResult struct {
	SpotifyID string `json:"spotify_id"`
	TidalURL  string `json:"tidal_url,omitempty"`
	AmazonURL string `json:"amazon_url,omitempty"`
	QobuzISRC string `json:"qobuz_isrc,omitempty"`
	Error     string `json:"error,omitempty"`
	HasURLs   bool   `json:"has_urls"`
}

// BatchURLResult holds results from batch URL fetching
type BatchURLResult struct {
	Results      map[string]*StreamingURLResult `json:"results"`
	TotalFetched int                            `json:"total_fetched"`
	TotalFailed  int                            `json:"total_failed"`
	FetchTimeMs  int64                          `json:"fetch_time_ms"`
}

// ParallelDownloadState tracks the state of parallel downloads
type ParallelDownloadState struct {
	activeWorkers int32
	maxWorkers    int32
	mutex         sync.RWMutex
}

var parallelState = &ParallelDownloadState{
	maxWorkers: 1, // Default: 1 (sequential, safest)
}

// GetActiveWorkers returns the number of currently active download workers
func GetActiveWorkers() int {
	return int(atomic.LoadInt32(&parallelState.activeWorkers))
}

// GetMaxWorkers returns the maximum number of concurrent workers
func GetMaxWorkers() int {
	return int(atomic.LoadInt32(&parallelState.maxWorkers))
}

// SetMaxWorkers sets the maximum concurrent downloads (1-5)
func SetMaxWorkers(max int) {
	if max < 1 {
		max = 1
	}
	if max > 5 {
		max = 5
	}
	atomic.StoreInt32(&parallelState.maxWorkers, int32(max))
}

// PreFetchStreamingURLs fetches streaming URLs for multiple tracks with rate limiting
// This is Phase 1 of parallel downloads - must be sequential due to song.link rate limits
func PreFetchStreamingURLs(spotifyIDs []string, progressCallback func(current, total int)) *BatchURLResult {
	startTime := time.Now()
	results := make(map[string]*StreamingURLResult)

	client := NewSongLinkClient()
	totalFailed := 0

	for i, spotifyID := range spotifyIDs {
		if progressCallback != nil {
			progressCallback(i+1, len(spotifyIDs))
		}

		result := &StreamingURLResult{
			SpotifyID: spotifyID,
		}

		// Get URLs from song.link (this has built-in rate limiting)
		availability, err := client.CheckTrackAvailability(spotifyID, "")
		if err != nil {
			result.Error = fmt.Sprintf("Failed to get URLs: %v", err)
			totalFailed++
		} else {
			if availability.TidalURL != "" {
				result.TidalURL = availability.TidalURL
				result.HasURLs = true
			}
			if availability.AmazonURL != "" {
				result.AmazonURL = availability.AmazonURL
				result.HasURLs = true
			}
			if availability.Qobuz {
				// For Qobuz, we need the ISRC which we'll get during download
				result.QobuzISRC = spotifyID // Placeholder, actual ISRC fetched later
				result.HasURLs = true
			}
		}

		results[spotifyID] = result
	}

	return &BatchURLResult{
		Results:      results,
		TotalFetched: len(spotifyIDs) - totalFailed,
		TotalFailed:  totalFailed,
		FetchTimeMs:  time.Since(startTime).Milliseconds(),
	}
}

// ParallelDownloadProgress tracks progress of parallel downloads
type ParallelDownloadProgress struct {
	Phase           string  `json:"phase"` // "prefetch" or "download"
	Current         int     `json:"current"`
	Total           int     `json:"total"`
	ActiveDownloads int     `json:"active_downloads"`
	Completed       int     `json:"completed"`
	Failed          int     `json:"failed"`
	Skipped         int     `json:"skipped"`
	SpeedMBps       float64 `json:"speed_mbps"`
}

var (
	parallelProgress     ParallelDownloadProgress
	parallelProgressLock sync.RWMutex
)

// SetParallelProgress updates the parallel download progress
func SetParallelProgress(progress ParallelDownloadProgress) {
	parallelProgressLock.Lock()
	parallelProgress = progress
	parallelProgressLock.Unlock()
}

// GetParallelProgress returns the current parallel download progress
func GetParallelProgress() ParallelDownloadProgress {
	parallelProgressLock.RLock()
	defer parallelProgressLock.RUnlock()

	progress := parallelProgress
	progress.ActiveDownloads = GetActiveWorkers()
	return progress
}

// ResetParallelProgress resets the parallel download progress
func ResetParallelProgress() {
	parallelProgressLock.Lock()
	parallelProgress = ParallelDownloadProgress{}
	parallelProgressLock.Unlock()
}
