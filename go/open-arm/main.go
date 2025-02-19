package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// initOpenGL initializes OpenGL and returns an initialized program.
func initOpenGL() uint32 {
	// Initialize OpenGL bindings
	if err := gl.Init(); err != nil {
		log.Fatalln("Failed to initialize OpenGL:", err)
	}

	// Print OpenGL version
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version:", version)

	gl.Viewport(0, 0, 800, 600)

	// Compile shaders
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		log.Fatalln(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		log.Fatalln(err)
	}

	// Link the shaders into a program
	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.LINK_STATUS, &logLength)
		l := strings.Repeat("\x00", int(logLength)+1)
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(l))
		log.Fatalf("failed to link program: %v", l)
	}

	// Clean up the shaders (they're no longer needed once linked into a program)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program
}

// makeVao initializes and return a Vertex Arra object with a triangle
func makeVao() uint32 {
	// Define the vertices and colors of the striangle
	vertices := []float32{
		// Position			// Color			// Texture coordinates
		0.5, 0.5, 0.0,		1.0, 0.0, 0.0,		1.0, 1.0,	// top right
		0.5, -0.5, 0.0,		0.0, 1.0, 0.0,		1.0, 0.0,	// bottom right
		-0.5, -0.5, 0.0,	0.0, 0.0, 1.0,		0.0, 0.0,	// bottom left
		-0.5, 0.5, 0.0,		1.0, 1.0, 0.0,		0.0, 1.0,	// top left
	}

	indices := []uint32{
		0, 1, 3,	// first triangle
		1, 2, 3,	// second triangle
	}

	// Create a Vertex Buffer Object and Vertex Array Object
	var vbo, vao, ebo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)

	// Bind the Vertex Array Object
	gl.BindVertexArray(vao)

	// Bind and fill the VBO with vertex data
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	// Set up vertex attribute pointers
	// Position attribute (location = 0)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 8*4, 0)
	gl.EnableVertexAttribArray(0)

	// Color attribute (location = 1)
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 8*4, 3*4)
	gl.EnableVertexAttribArray(1)

	// Texture loca attribute (location = 2)
	gl.VertexAttribPointerWithOffset(2, 2, gl.FLOAT, false, 8*4, 6*4)
	gl.EnableVertexAttribArray(2)

	// Unbind the VBO and VAO
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return vao
}

// mainLoop runs the main rendering loop.
func mainLoop(window *glfw.Window, program uint32, vao uint32) {
	// Set the background color (RGBA)
	gl.ClearColor(0.2, 0.3, 0.4, 1.0)

	// Load the texture1
	texture1, err := newTexture("container.jpg")
	if err != nil {
		log.Fatalln(err)
	}

	// Load the texture
	texture2, err := newTexture("awesomeface.png")
	if err != nil {
		log.Fatalln(err)
	}

	// Main loop
	for !window.ShouldClose() {
		// Process inputs
		processInputs(window)

		// Clear the screen
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Use the program
		gl.UseProgram(program)

		// Bind the VAO
		gl.BindVertexArray(vao)

		// Bind the texture
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, texture1)
		gl.ActiveTexture(gl.TEXTURE1)
		gl.BindTexture(gl.TEXTURE_2D, texture2)

		gl.Uniform1i(gl.GetUniformLocation(program, gl.Str("texture1\x00")), 0)
		gl.Uniform1i(gl.GetUniformLocation(program, gl.Str("texture2\x00")), 1)

		// Draw the triangle
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)

		// Swap buffers and poll events
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func main() {
	// Ensure that OpenGL runs on the main thread
	runtime.LockOSThread()

	// Initialize GLFW and create a window
	window := initGlfw()
	defer glfw.Terminate()

	// Initiliase OpenGL
	program := initOpenGL()
	vao := makeVao()

	// Enter the main loop
	mainLoop(window, program, vao)
}
