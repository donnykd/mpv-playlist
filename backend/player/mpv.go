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
	socket string
	conn   *mpvipc.Connection
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
		err = conn.Open()

		time.Sleep(300 * time.Millisecond)
		if err != nil {
			log.Fatalf("error opening connection: %v", err)
		}
	}()

	wg.Wait()

	return &Player{
		socket: socket,
		conn:   conn,
		playlist: playlist.NewPlaylist(),
	}
}

func (p *Player) Play(file string) {
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

	p.Play(playlistPath)
}

func (p *Player) AddFile(file string) {
	p.playlist.AddFile(file) 
}

func mpvExists() bool {
	_, err := exec.LookPath("mpv")
	return err == nil
}
