package main

import (
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width, height = 800, 600
)

func main() {
	runtime.LockOSThread()
}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}


}
