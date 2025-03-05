package main

import (
	"fmt"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Position mgl32.Vec3
	Pitch    float32
	Yaw      float32
	Up       mgl32.Vec3

	lastX, lastY float32
	FirstMouse    bool
}

func (m *Camera) GetFront() mgl32.Vec3 {
	x := float32(math.Cos(float64(mgl32.DegToRad(m.Pitch))) * math.Cos(float64(mgl32.DegToRad(m.Yaw))))
	y := float32(math.Sin(float64(mgl32.DegToRad(m.Pitch))))
	z := float32(math.Cos(float64(mgl32.DegToRad(m.Pitch))) * math.Sin(float64(mgl32.DegToRad(m.Yaw))))
	return mgl32.Vec3{x, y, z}.Normalize()
}

func (m *Camera) GetRight() mgl32.Vec3 {
	return m.Up.Cross(m.GetFront()).Normalize()
}

func (m *Camera) GetUp() mgl32.Vec3 {
	return m.GetFront().Cross(m.GetRight())
}

func (m *Camera) GetViewMatrix() mgl32.Mat4 {
	return mgl32.LookAtV(m.Position, m.Position.Add(m.GetFront()), m.Up)
}

func (m *Camera) ProcessMouseMovement(xpos, ypos float32) {
	if m.FirstMouse {
		m.lastX = xpos
		m.lastY = ypos
		m.FirstMouse = false
		return
	}

	xoffset := xpos - m.lastX
	yoffset := m.lastY - ypos
	m.lastX = xpos
	m.lastY = ypos

	sensitivity := float32(0.1)
	xoffset *= sensitivity
	yoffset *= sensitivity	

	m.Yaw += xoffset
	m.Pitch += yoffset

	if m.Pitch > 89.0 {
		m.Pitch = 89.0
	}
	if m.Pitch < -89.0 {
		m.Pitch = -89.0
	}
}