package player

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"

	"github.com/dexterlb/mpvipc"
	"github.com/donnykd/mpv-playlist/backend/playlist"
)

type Player struct {
	socket   string
	conn     *mpvipc.Connection
	playlist *playlist.Playlist
}

func NewPlayer() (p *Player) {
	socket := "/tmp/mpv_rpc"
	var conn *mpvipc.Connection

	if !mpvExists() {
		log.Fatal("MPV player not installed on this machine")
	}

	cmd := exec.Command("mpv", "--idle", "--input-ipc-server="+socket)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		conn = mpvipc.NewConnection(socket)

		maxRetries := 10
		for i := range maxRetries {
			err = conn.Open()
			if err == nil {
				return
			}

			waitTime := min(50*time.Millisecond<<i, time.Second)

			time.Sleep(waitTime)
		}

		log.Fatalf("error opening connection after %d retries: %v", maxRetries, err)
	}()

	wg.Wait()

	return &Player{
		socket:   socket,
		conn:     conn,
		playlist: playlist.NewPlaylist(),
	}
}

func (p *Player) play(file string) {
	_, err := p.conn.Call("loadfile", file)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("current file playing: %s", file)
}

func (p *Player) PlayAll() {
	playlistPath, err := p.playlist.GenerateM3uFile()
	if err != nil {
		log.Fatal(err)
	}

	p.play(playlistPath)
}

func (p *Player) AddFile(file string) {
	if !isFileValid(file) {
		log.Fatal("File is not valid")
	}

	path, err := normalizePath(file)
	if err != nil {
		log.Fatal(err)
	}
	p.playlist.AddFile(path)
}

func mpvExists() bool {
	_, err := exec.LookPath("mpv")
	return err == nil
}
