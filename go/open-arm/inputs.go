package main

import "github.com/go-gl/glfw/v3.3/glfw"

func processInputs(window *glfw.Window, cam *Camera, deltaTime float32) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}

	// Camera controls
	cameraSpeed := 2.5 * deltaTime
	if window.GetKey(glfw.KeyW) == glfw.Press {
		cam.Position = cam.Position.Add(cam.GetFront().Mul(cameraSpeed))
	}
	if window.GetKey(glfw.KeyS) == glfw.Press {
		cam.Position = cam.Position.Sub(cam.GetFront().Mul(cameraSpeed))
	}
	if window.GetKey(glfw.KeyA) == glfw.Press {
		cam.Position = cam.Position.Add(cam.GetRight().Mul(cameraSpeed))
	}
	if window.GetKey(glfw.KeyD) == glfw.Press {
		cam.Position = cam.Position.Sub(cam.GetRight().Mul(cameraSpeed))
	}
	if window.GetKey(glfw.KeySpace) == glfw.Press {
		cam.Position = cam.Position.Add(cam.GetUp().Mul(cameraSpeed))
	}
	if window.GetKey(glfw.KeyLeftControl) == glfw.Press {
		cam.Position = cam.Position.Sub(cam.GetUp().Mul(cameraSpeed))
	}
}
