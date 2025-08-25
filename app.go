package main

import (
	"context"
	"fmt"

	"github.com/donnykd/mpv-playlist/backend"
)

// App struct
type App struct {
	ctx    context.Context
	player backend.Player
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		player: *backend.NewPlayer(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.player.Play("/home/khalid/Downloads/ssstwitter.com_1754318376612.mp4")
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
