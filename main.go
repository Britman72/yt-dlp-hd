package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func logLine(line string) {
	if !debugMode {
		return
	}
	logFile := "yt-dlp.log"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	f, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	f.WriteString(fmt.Sprintf("%s  %s\n", timestamp, line))
}

func readINI() map[string]string {
	config := map[string]string{
		"maxres":      "best",
		"yt-dlp-path": "",
		"ffmpeg-path": "",
		"debug":       "true", // default to true if not provided
	}
	data, err := ioutil.ReadFile("yt-dlp.ini")
	if err != nil {
		logLine("INI not found, using default settings")
		return config
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		kv := strings.SplitN(strings.TrimSpace(line), "=", 2)
		if len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			val := strings.TrimSpace(kv[1])
			config[strings.ToLower(key)] = val
		}
	}
	debugVal := strings.ToLower(config["debug"])
	if debugVal == "false" {
		debugMode = false
	}

	return config
}

func buildFormatString(maxres string) string {
	switch strings.ToLower(maxres) {
	case "1080p":
		return "bestvideo[height<=1080]+bestaudio/best[height<=1080]"
	case "4k":
		return "bestvideo[height<=2160]+bestaudio/best[height<=2160]"
	default:
		return "bestvideo+bestaudio/best"
	}
}

var debugMode = true

func main() {
	args := os.Args[1:]
	config := readINI()

	maxres := config["maxres"]
	format := buildFormatString(maxres)
	ytDlpPath := filepath.Join(config["yt-dlp-path"], "yt-dlp.exe")
	ffmpegPath := config["ffmpeg-path"]

	logLine("Original args: " + strings.Join(args, " "))

	var finalArgs []string
	skipNext := false

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if skipNext {
			skipNext = false
			continue
		}

		if arg == "-f" && i+1 < len(args) && args[i+1] == "mp4" {
			logLine("Stripped -f mp4")
			skipNext = true
			continue
		}

		finalArgs = append(finalArgs, arg)
	}

	// Add corrected format
	finalArgs = append(finalArgs, "-f", format)

	// Add ffmpeg location and merge format
	if ffmpegPath != "" {
		finalArgs = append(finalArgs, "--ffmpeg-location", ffmpegPath)
		logLine("Set ffmpeg path: " + ffmpegPath)
	}
	finalArgs = append(finalArgs, "--merge-output-format", "mp4")
	finalArgs = append(finalArgs, "--no-keep-video")

	// Fix static -o output (e.g., VideoTemp -> VideoTemp.%(ext)s)
	for i := 0; i < len(finalArgs)-1; i++ {
		if finalArgs[i] == "-o" {
			out := finalArgs[i+1]

			// If it does not contain a template token...
			if !strings.Contains(out, "%(") {
				// If it ends in .mp4, strip it (we'll reattach via template)
				if strings.HasSuffix(strings.ToLower(out), ".mp4") {
					out = strings.TrimSuffix(out, ".mp4")
					logLine("Stripped .mp4 extension from output path")
				}
				// Append .%(ext)s template
				out += ".%(ext)s"
				finalArgs[i+1] = out
				logLine("Adjusted output to: " + out)
			}
		}
	}

	logLine("Final yt-dlp args: " + strings.Join(finalArgs, " "))

	cmd := exec.Command(ytDlpPath, finalArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir, _ = filepath.Abs(".")
	err := cmd.Run()
	if err != nil {
		logLine("Error running yt-dlp: " + err.Error())
		os.Exit(1)
	}
}
