### yt-downloader

**Dependencies:**

- [youtube-dl](https://github.com/ytdl-org/youtube-dl)
- [ffmpeg](https://ffmpeg.org/)

****

**Build:**

```go
go build yt-downloader.go
```



**Use:**

```bash
./yt-downloader -u 'url'
```

```bash
./yt-downloader -u 'url' -f mp3
```

```bash
./yt-downloader -u 'url' -o song.mp3
```



**Example:**

```bash
./yt-downloader -u 'https://www.youtube.com/watch?v=pPw_izFr5PA'
```

```bash
./yt-downloader -u 'https://www.youtube.com/watch?v=pPw_izFr5PA' -f mp4 
```

```bash
./yt-downloader -u 'https://www.youtube.com/watch?v=pPw_izFr5PA' -o gooba.mp3
```

```bash
./yt-downloader -u 'https://www.youtube.com/watch?v=pPw_izFr5PA' -f mp4 -o gooba.mp4
```

****



**Flag:**

- **'-u'** ---> define youtube video url
- **'-f'** ---> define file format(default: mp3)
- **'-o'** ---> define output filename(ex. 'song.mp3' or 'video.mp4')

**Format:**

- MP4
- MP3
- WAV
- WEBM

