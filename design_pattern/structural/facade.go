package main

func main() {

}

type Video struct {
	*MP4
	*MPEG
	*FFMPEG
}

type MP4 struct{}
type MPEG struct{}
type FFMPEG struct{}

func NewVideo(mp4 *MP4, mpeg *MPEG, ffmpeg *FFMPEG) *Video {
	return &Video{
		MP4:    mp4,
		MPEG:   mpeg,
		FFMPEG: ffmpeg,
	}
}
