package main

var vertexShaderSource = `
#version 460 core
layout (location = 0) in vec3 aPos;			// Vertex position
layout (location = 1) in vec2 aTextCoord;	// Vertex color

out vec2 TexCoord;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

void main() {
	gl_Position = projection * view * model * vec4(aPos, 1.0);		// Set the vertex position
	TexCoord = aTextCoord;
}
` + "\x00"

// Fragment Shader Source Code
var fragmentShaderSource = `
#version 460 core
out vec4 fragColor;		// Output color
                      
in vec3 ourColor;     	// Input color from vertex shader
in vec2 TexCoord;

uniform sampler2D texture1;
uniform sampler2D texture2;

void main() {
	fragColor = mix(texture(texture1, TexCoord), texture(texture2, TexCoord), 0.2);     // Set fragment color
}
` + "\x00" // Null-terminate the string
