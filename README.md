# yt-dlp-hd

A lightweight wrapper around `yt-dlp` to enforce high resolution (upto 4K) downloading of YouTube videos via Playnite's Extra Metadata Loader add-on.

## Installation

1. Download the latest release and extract the zip file into folder "yt-dlp-hd".
2. Place the folder in a suitable folder such as C:\Playnite\video-tools.
3. Update the included yt-dlp.ini file as per the settings below.
4. In ExtraMetadataLoader's settings, point the YT-DLP path to the wrapper's folder.
5. Restart Playnite (important!)
6. Test download a YouTube video trailer for an existing game. Ensure the video you want to download is high res (not all are).

You can watch a video of the setup process here: https://www.youtube.com/watch?v=zXiarzc5iJA

## Configuration

Edit the `yt-dlp.ini` file in the wrapper's folder:

### INI Settings

- maxres - Set the maximum resolution for the video. Available options are: Best, 4k, 1080p, 720p, 480p. If the resolution you specifiy isn't available to download then it will grab the next best quality.
- yt-dlp-path - Set the path to the folder where the original yt-dlp.exe binary is stored.
- ffmpeg-path - Set the path to the folder where the ffmpeg.exe binary is installed.
- debug - Set to "true" to log the output to yt-dlp.log or false for no logging.

#### Example INI File:

```ini
maxres=1080p
yt-dlp-path=C:\Playnite\video-tools\yt-dlp
ffmpeg-path=C:\Playnite\video-tools\ffmpeg\bin
debug=false
```

## Removing/Uninstalling

To remove or uninstall simply delete the yt-dlp-hd folder and repoint YT-DLP in ExtraMetadata's add-on setting to the original .exe.
