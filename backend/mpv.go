package backend

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/dexterlb/mpvipc"
)

func Play(file string){
	
	socket := "/tmp/mpv_rpc"
	
	if !mpv_exists(){
		log.Fatal("MPV player not installed on this machine")
	}
	
	cmd := exec.Command("mpv", "--idle", "--input-ipc-server=" + socket)
	err := cmd.Start()
	if err != nil{
		log.Fatalf("could not start mpv: %v", err)
	}
	
	conn := mpvipc.NewConnection(socket)
	err = conn.Open()
	if err != nil {
		log.Fatalf("error opening connection: %v", err)
	}
	defer conn.Close()
	
	_, err = conn.Call("loadfile", file)	
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("current file playing: %s", file)
}

func mpv_exists() bool {
	_, err := exec.LookPath("mpv")
	return err == nil
}