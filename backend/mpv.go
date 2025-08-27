package backend

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"

	"github.com/dexterlb/mpvipc"
)

type Player struct {
	socket string
	conn   *mpvipc.Connection
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

		retries := 10
		delay := 10 * time.Millisecond

		for i := range retries{
			
			err = conn.Open()
			if err == nil {
				break
			}

			if i == retries-1 {
				log.Fatalf("error opening connection: %v", err)
			}

			time.Sleep(delay)
		}
	}()

	wg.Wait()

	return &Player{
		socket: socket,
		conn:   conn,
	}
}

func (p *Player) Play(file string) {
	defer p.conn.Close()

	_, err := p.conn.Call("loadfile", file)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("current file playing: %s", file)
}

func mpvExists() bool {
	_, err := exec.LookPath("mpv")
	return err == nil
}
