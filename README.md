# yt-dlp-wrapper

A lightweight Go wrapper around `yt-dlp` to enforce consistent format selection and merging behavior when downloading YouTube videos via PlayNite.

## Features

- Automatically merges best video+audio
- Respects a configurable max resolution (`best`, `1080p`, `4k`)

## Installation

1. Download the latest release and move the files to a suitable location ie. C:\Playnite\yt-dlp-wrapper.
2. Do NOT overwrite or place in the same folder as the original yt-dlp.exe!

## Configuration

Create or update the `yt-dlp.ini` file in the same folder:

### INI Settings

- maxres - Set the maximum resolution for the video ie. 4k, 1080p, 720p, 480p.
- yt-dlp-path - Set the path to the folder where the original yt-dlp.exe is stored (do not include yt-dlp.exe in the path)
- ffmpeg-path - Set the path to the folder where the ffmpeg.exe binary is installed.
- debug - Set to "true" to log the output to yt-dlp.log or false for no logging.

#### Example INI File:

```ini
maxres=1080p
yt-dlp-path=C:\Playnite\tools\yt-dlp
ffmpeg-path=C:\Playnite\tools\ffmpeg\bin
debug=false