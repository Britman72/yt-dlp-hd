# yt-dlp-hd

A lightweight wrapper around `yt-dlp` to enforce high resolution (upto 4K) downloading of YouTube videos via Playnite.

## Installation

1. Download the latest release and move the files to a suitable location ie. C:\Playnite\yt-dlp-wrapper.
2. Do NOT overwrite or place in the same folder as the original yt-dlp.exe!
3. Update the included yt-dlp.ini file as per the settings below.
4. In ExtraMetadataLoader's settings, point the YT-DLP path to the wrapper's folder.
5. Test download a YouTube video trailer for an existing game. Ensure the video you want to download is high res (not all are).

## Configuration

Edit the `yt-dlp.ini` file in the wrapper's folder:

### INI Settings

- maxres - Set the maximum resolution for the video. Available options are: Best, 4k, 1080p, 720p, 480p. If the resolution you specifiy isn't available to download then it will grab the next best quality.
- yt-dlp-path - Set the path to the folder where the original yt-dlp.exe is stored (do not include yt-dlp.exe in the path)
- ffmpeg-path - Set the path to the folder where the ffmpeg.exe binary is installed.
- debug - Set to "true" to log the output to yt-dlp.log or false for no logging.

#### Example INI File:

```ini
maxres=1080p
yt-dlp-path=C:\Playnite\tools\yt-dlp
ffmpeg-path=C:\Playnite\tools\ffmpeg\bin
debug=false
```

## Removing/Uninstalling

To remove or uninstall simply delete the wrapper folder and repoint YT-DLP in ExtraMetadata's add-on setting to the original .exe.
