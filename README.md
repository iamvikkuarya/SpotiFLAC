# SpotiFLAC (Enhanced Fork)

> Forked from [afkarxyz/SpotiFLAC](https://github.com/afkarxyz/SpotiFLAC) — all credit for the original project goes to the original author.

Get Spotify tracks in lossless FLAC from Tidal, Qobuz & Amazon Music — no account required.

![Windows](https://img.shields.io/badge/Windows-10%2B-0078D6?style=for-the-badge&logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI1MTIiIGhlaWdodD0iNTEyIiB2aWV3Qm94PSIwIDAgMjAgMjAiPjxwYXRoIGZpbGw9IiNmZmZmZmYiIGZpbGwtcnVsZT0iZXZlbm9kZCIgZD0iTTIwIDEwLjg3M1YyMEw4LjQ3OSAxOC41MzdsLjAwMS03LjY2NEgyMFptLTEzLjEyIDBsLS4wMDEgNy40NjFMMCAxNy40NjF2LTYuNTg4aDYuODhaTTIwIDkuMjczSDguNDhsLS4wMDEtNy44MUwyMCAwdjkuMjczWk02Ljg3OSAxLjY2NmwuMDAxIDcuNjA3SDBWMi41MzlsNi44NzktLjg3M1oiLz48L3N2Zz4=)

### [Download Latest Release](https://github.com/iamvikkuarya/SpotiFLAC/releases)

---

## What This Fork Adds

This fork builds on top of the original SpotiFLAC with two major features:

### Parallel Downloads (1-5 concurrent workers)
- Download multiple tracks simultaneously instead of one at a time
- Configurable from 1 (sequential) to 5 (maximum) concurrent downloads
- Ideal for batch downloading albums and playlists

### Output Format Selection
- **FLAC** — Lossless (default, same as original)
- **MP3** — Lossy, configurable bitrate: 128k / 192k / 256k / 320k
- **AAC (M4A)** — Lossy, configurable bitrate: 128k / 192k / 256k / 320k

Downloads are always fetched in lossless quality first, then automatically converted to your selected format with full metadata preservation (title, artist, album, cover art, lyrics).

---

## Screenshot

![Image](https://github.com/user-attachments/assets/adbdc056-bace-44a9-8ba6-898b4526b65a)

## Original Project

### [SpotiFLAC](https://github.com/afkarxyz/SpotiFLAC)

This fork is based on SpotiFLAC by [afkarxyz](https://github.com/afkarxyz). Check out the original project and their other work:

- [SpotiFLAC Next](https://github.com/spotiverse/SpotiFLAC-Next) — Hi-Res lossless FLACs
- [SpotiDownloader](https://github.com/afkarxyz/SpotiDownloader) — MP3 and FLAC via spotidownloader.com
- [SpotubeDL](https://spotubedl.com) — MP3/OGG/Opus downloads
- [SpotiFLAC Mobile](https://github.com/zarzet/SpotiFLAC-Mobile) — Android & iOS

## FAQ

### Is this software free?

_Yes. This software is completely free.
You do not need an account, login, or subscription.
All you need is an internet connection._

### Can using this software get my Spotify account suspended or banned?

_No.
This software has no connection to your Spotify account.
Spotify data is obtained through reverse engineering of the Spotify Web Player, not through user authentication._

### Where does the audio come from?

_The audio is fetched using third-party APIs._

### Why does metadata fetching sometimes fail?

_This usually happens because your IP address has been rate-limited.
You can wait and try again later, or use a VPN to bypass the rate limit._

### Why does Windows Defender or antivirus flag or delete the file?

_This is a false positive.
It likely happens because the executable is compressed using UPX._

_If you are concerned, you can fork the repository and build the software yourself from source._

## Disclaimer

This project is for **educational and private use only**. The developer does not condone or encourage copyright infringement.

**SpotiFLAC** is a third-party tool and is not affiliated with, endorsed by, or connected to Spotify, Tidal, Qobuz, Amazon Music, or any other streaming service.

You are solely responsible for:

1. Ensuring your use of this software complies with your local laws.
2. Reading and adhering to the Terms of Service of the respective platforms.
3. Any legal consequences resulting from the misuse of this tool.

The software is provided "as is", without warranty of any kind. The author assumes no liability for any bans, damages, or legal issues arising from its use.

## API Credits

- **Tidal**: [hifi-api](https://github.com/binimum/hifi-api)
- **Qobuz**: [dabmusic.xyz](https://dabmusic.xyz), [squid.wtf](https://squid.wtf), [jumo-dl](https://jumo-dl.pages.dev/)
