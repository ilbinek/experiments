package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
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
	/*vertices := []float32{
		// Position			// Color			// Texture coordinates
		0.5, 0.5, 0.0,		1.0, 0.0, 0.0,		1.0, 1.0,	// top right
		0.5, -0.5, 0.0,		0.0, 1.0, 0.0,		1.0, 0.0,	// bottom right
		-0.5, -0.5, 0.0,	0.0, 0.0, 1.0,		0.0, 0.0,	// bottom left
		-0.5, 0.5, 0.0,		1.0, 1.0, 0.0,		0.0, 1.0,	// top left
	}*/

	vertices := []float32{
		-0.5, -0.5, -0.5,  0.0, 0.0,
		 0.5, -0.5, -0.5,  1.0, 0.0,
		 0.5,  0.5, -0.5,  1.0, 1.0,
		 0.5,  0.5, -0.5,  1.0, 1.0,
		-0.5,  0.5, -0.5,  0.0, 1.0,
		-0.5, -0.5, -0.5,  0.0, 0.0,
	
		-0.5, -0.5,  0.5,  0.0, 0.0,
		 0.5, -0.5,  0.5,  1.0, 0.0,
		 0.5,  0.5,  0.5,  1.0, 1.0,
		 0.5,  0.5,  0.5,  1.0, 1.0,
		-0.5,  0.5,  0.5,  0.0, 1.0,
		-0.5, -0.5,  0.5,  0.0, 0.0,
	
		-0.5,  0.5,  0.5,  1.0, 0.0,
		-0.5,  0.5, -0.5,  1.0, 1.0,
		-0.5, -0.5, -0.5,  0.0, 1.0,
		-0.5, -0.5, -0.5,  0.0, 1.0,
		-0.5, -0.5,  0.5,  0.0, 0.0,
		-0.5,  0.5,  0.5,  1.0, 0.0,
	
		 0.5,  0.5,  0.5,  1.0, 0.0,
		 0.5,  0.5, -0.5,  1.0, 1.0,
		 0.5, -0.5, -0.5,  0.0, 1.0,
		 0.5, -0.5, -0.5,  0.0, 1.0,
		 0.5, -0.5,  0.5,  0.0, 0.0,
		 0.5,  0.5,  0.5,  1.0, 0.0,
	
		-0.5, -0.5, -0.5,  0.0, 1.0,
		 0.5, -0.5, -0.5,  1.0, 1.0,
		 0.5, -0.5,  0.5,  1.0, 0.0,
		 0.5, -0.5,  0.5,  1.0, 0.0,
		-0.5, -0.5,  0.5,  0.0, 0.0,
		-0.5, -0.5, -0.5,  0.0, 1.0,
	
		-0.5,  0.5, -0.5,  0.0, 1.0,
		 0.5,  0.5, -0.5,  1.0, 1.0,
		 0.5,  0.5,  0.5,  1.0, 0.0,
		 0.5,  0.5,  0.5,  1.0, 0.0,
		-0.5,  0.5,  0.5,  0.0, 0.0,
		-0.5,  0.5, -0.5,  0.0, 1.0,
	}

	/*indices := []uint32{
		0, 1, 3,	// first triangle
		1, 2, 3,	// second triangle
	}*/

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

	/*gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)
	*/

	// Set up vertex attribute pointers
	// Position attribute (location = 0)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 5*4, 0)
	gl.EnableVertexAttribArray(0)

	// Texture loca attribute (location = 1)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)
	gl.EnableVertexAttribArray(1)

	// Unbind the VBO and VAO
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return vao
}

// mainLoop runs the main rendering loop.
func mainLoop(window *glfw.Window, program uint32, vao uint32) {
	// Set the background color (RGBA)
	gl.ClearColor(0.2, 0.3, 0.4, 1.0)
	gl.Enable(gl.DEPTH_TEST)

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

	/*cubePositions := []mgl32.Vec3{
		{0.0, 0.0, 0.0},
		{2.0, 5.0, -15.0},
		{-1.5, -2.2, -2.5},
		{-3.8, -2.0, -12.3},
		{2.4, -0.4, -3.5},
		{-1.7, 3.0, -7.5},
		{1.3, -2.0, -2.5},
		{1.5, 2.0, -2.5},
		{1.5, 0.2, -1.5},
		{-1.3, 1.0, -1.5},
	}*/

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

		model := mgl32.HomogRotate3D(mgl32.DegToRad(float32(glfw.GetTime()*40)), mgl32.Vec3{0.5, 1, 0})
		view := mgl32.Translate3D(0, 0, -3)
		width, height := window.GetSize()
		aspect := float32(width) / float32(height)
		projection := mgl32.Perspective(mgl32.DegToRad(45), aspect, 0.1, 100.0)

		fmt.Println("model:\n", model)
		fmt.Println("view:\n", view)
		fmt.Println("projection:\n", projection)

		// Create a transformation matrix
		modelLoc := gl.GetUniformLocation(program, gl.Str("model\x00"))
		viewLoc := gl.GetUniformLocation(program, gl.Str("view\x00"))
		projectionLoc := gl.GetUniformLocation(program, gl.Str("projection\x00"))
		gl.UniformMatrix4fv(modelLoc, 1, false, (*float32)(gl.Ptr(&model[0])))
		gl.UniformMatrix4fv(viewLoc, 1, false, (*float32)(gl.Ptr(&view[0])))
		gl.UniformMatrix4fv(projectionLoc, 1, false, (*float32)(gl.Ptr(&projection[0])))

		/*for cubePosition := range cubePositions {
			model = mgl32.Translate3D(cubePositions[cubePosition].X(), cubePositions[cubePosition].Y(), cubePositions[cubePosition].Z())
			model = model.Mul4(mgl32.HomogRotate3D(float32(cubePosition * 20), mgl32.Vec3{1, 0.3, 0.5}))
			gl.UniformMatrix4fv(modelLoc, 1, false, (*float32)(gl.Ptr(&model[0])))
			gl.DrawArrays(gl.TRIANGLES, 0, 36)
		}*/

		gl.DrawArrays(gl.TRIANGLES, 0, 36)

		// Draw the triangle
		//gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)

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
