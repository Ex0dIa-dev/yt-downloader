/*
	MADE BY Ex0dIa-dev
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func init() {
	flag.StringVar(&url, "u", "", "insert url")
	flag.StringVar(&format, "f", "mp3", "mp4 or mp3")
	flag.StringVar(&output, "o", "", "output filename")
}

var url, format, output string

func main() {

	flag.Parse()

	if url == "" {
		fmt.Println("[-]Please insert url.")
		os.Exit(1)
	}

	switch format {
	case "mp4":
		//getting direct url from youtube-dl
		mp4_url, mp3_url := GetDownloadUrl(url)
		tmp_audio_webm := "tmp_audio.webm"
		tmp_video := "tmp_video.mp4"

		//downloading audio and video
		fmt.Println("[+]Downloading...")
		go DownloadVideo(mp4_url, tmp_video)
		DownloadAudio(mp3_url, tmp_audio_webm)

		fmt.Println("[+]Download complete.")

		//converting audio.webm to audio.mp3
		fmt.Println("[+]Processing file...")
		tmp_audio_mp3 := "tmp_audio.mp3"
		WebmToMp3(tmp_audio_webm, tmp_audio_mp3)

		//merging mp4 and mp3

		if output == "" {
			output = GetVideoTitle(url)
			output = output + ".mp4"
		}

		MergeAudioVideo(output, tmp_video, tmp_audio_mp3)

		fmt.Println("[+]Done.")

	case "mp3":
		_, mp3_url := GetDownloadUrl(url)
		tmp_audio_webm := "tmp_audio.webm"

		//downloading audio
		fmt.Println("[+]Downloading...")
		DownloadAudio(mp3_url, tmp_audio_webm)

		fmt.Println("[+]Download complete.")

		//converting audio.webm to audio.mp3
		fmt.Println("[+]Processing file...")

		if output == "" {
			output = GetVideoTitle(url)
			output = output + ".mp3"
		}

		WebmToMp3(tmp_audio_webm, output)

		fmt.Println("[+]Done.")

	}

}

//get the mp3 and mp4 downlaod url from youtube-dl
func GetDownloadUrl(url string) (string, string) {

	if FileExists("url.txt") {
		err := os.Remove("url.txt")
		checkerr(err)
	}

	file, err := os.Create("url.txt")
	checkerr(err)

	cmd := exec.Command("youtube-dl", "--get-url", url)
	cmd.Stdout = file
	err = cmd.Run()
	checkerr(err)

	file.Close()

	file, err = os.Open("url.txt")
	checkerr(err)

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	mp4_url := lines[0]
	mp3_url := lines[1]

	err = os.Remove("url.txt")
	checkerr(err)

	return mp4_url, mp3_url
}

func GetVideoTitle(url string) string {

	out, err := exec.Command("youtube-dl", "-e", url).Output()
	if err != nil {
		log.Fatal(err)
	}

	title := string(out)
	title = strings.Replace(title, "\n", "", -1)

	return title
}

func DownloadVideo(url, filename string) {

	if FileExists(filename) {
		err := os.Remove(filename)
		checkerr(err)
	}

	file, err := os.Create(filename)
	checkerr(err)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		file.Close()
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if _, err = io.Copy(file, resp.Body); err != nil {
		file.Close()
		log.Fatal(err)
	}

	file.Close()

}

func DownloadAudio(url, filename string) {

	if FileExists(filename) {
		err := os.Remove(filename)
		checkerr(err)
	}

	file, err := os.Create(filename)
	checkerr(err)

	resp, err := http.Get(url)
	if err != nil {
		file.Close()
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if _, err = io.Copy(file, resp.Body); err != nil {
		file.Close()
		log.Fatal(err)
	}

	file.Close()

}

func WebmToMp3(in_filename, out_filename string) {

	cmd := exec.Command("ffmpeg", "-i", in_filename, "-vn", "-ab", "128k", "-ar", "44100", "-y", out_filename)
	err := cmd.Run()
	checkerr(err)

	err = os.Remove(in_filename)
	checkerr(err)
}

func MergeAudioVideo(output_filename, mp4_path, mp3_path string) {

	cmd := exec.Command("ffmpeg", "-i", mp4_path, "-i", mp3_path, "-map", "0:v", "-map", "1:a", "-c:v", "copy", "-c:a", "copy", "-y", output_filename)
	err := cmd.Run()
	checkerr(err)

	err = os.Remove(mp4_path)
	checkerr(err)
	err = os.Remove(mp3_path)
	checkerr(err)
}

func FileExists(filename string) bool {

	stat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !stat.IsDir()
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
