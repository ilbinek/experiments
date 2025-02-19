package main

var vertexShaderSource = `
#version 460 core
layout (location = 0) in vec3 aPos;		// Vertex position
layout (location = 1) in vec3 aColor;	// Vertex color
layout (location = 2) in vec2 aTextCoord;	// Vertex color

out vec3 ourColor;
out vec2 TexCoord;

void main() {
	gl_Position = vec4(aPos, 1.0);		// Set the vertex position
	ourColor = aColor;					// Set the vertex color
	TexCoord = aTextCoord;
}
` + "\x00"

// Fragment Shader Source Code
var fragmentShaderSource = `
#version 460 core
out vec4 fragColor;
                      // Output color
in vec3 ourColor;                        // Input color from vertex shader
in vec2 TexCoord;

uniform sampler2D texture1;
uniform sampler2D texture2;

void main() {
	fragColor = mix(texture(texture1, TexCoord), texture(texture2, TexCoord), 0.2);     // Set fragment color
}
` + "\x00" // Null-terminate the string
