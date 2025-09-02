// Package playlist provides a simple playlist structure for managing collection of video files.
package playlist

import (
	"fmt"
	"os"
	"slices"
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

func (p *Playlist) Clear() {
	p.paths = nil
	p.currentFile = -1
}

func (p *Playlist) RemoveFile(index int) error {
	if index < 0 || index >= len(p.paths) {
		return fmt.Errorf("index out range")
	}

	p.paths = slices.Delete(p.paths, index, index+1)
	return nil
}

func (p *Playlist) IsEmpty() bool {
	return len(p.paths) == 0
}

func (p *Playlist) Count() int {
	return len(p.paths)
}

func (p *Playlist) GetFiles() []string {
	return slices.Clone(p.paths)
}
