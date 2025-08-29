package playlist

import (
	"fmt"
	"os"
)

type Playlist struct {
	paths       []string
	currentFile int
}

func NewPlaylist() (p *Playlist) {
	return &Playlist{
		paths:       make([]string, 0),
		currentFile: -1,
	}
}

func (p *Playlist) AddFile(file string) {
	p.paths = append(p.paths, file)
}

func (p *Playlist) GenerateM3uFile() (string, error) {
	tempFile, err := os.CreateTemp("", "playlist-*.m3u")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	tempFilePath := tempFile.Name()

	fmt.Fprint(tempFile, "#EXTM3U \n")

	for _, path := range p.paths {
		fmt.Fprint(tempFile, "\n", path)
	}

	return tempFilePath, err
}
