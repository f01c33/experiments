package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.6-compatibility/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	hook "github.com/robotn/gohook"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
}

func drawLine(x1, y1, x2, y2 int32) {
	gl.Color3f(1.0, 0, 0)
	gl.LineWidth(2.0)
	gl.Begin(gl.LINES)
	gl.Vertex2i(x1, y1)
	gl.Vertex2i(x2, y2)
	gl.End()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	pm := glfw.GetPrimaryMonitor()
	vm := pm.GetVideoMode()
	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)
	glfw.WindowHint(glfw.Floating, glfw.True)
	glfw.WindowHint(glfw.Decorated, glfw.False)

	// if it's the same size of the window it doesn't work (at least on windows), idk why
	window, err := glfw.CreateWindow(vm.Width-1, vm.Height-1, "much wow", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)

	if err := gl.Init(); err != nil {
		panic(err)
	}
	prog := gl.CreateProgram()
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)

	w, h := window.GetSize()

	gl.Ortho(0, // left
		float64(w), // right
		float64(h), // bottom
		0,          // top
		0,          // zNear
		1,          // zFar
	)

	// window.SetCursorPosCallback(func(win *glfw.Window, xpos float64, ypos float64) {
	// })

	gl.UseProgram(prog)

	EvChan := hook.Start()
	defer hook.End()

	// for ev := range EvChan {
	// 	fmt.Println("hook: ", ev)
	// }

	for !window.ShouldClose() {
		glfw.PollEvents()
		for len(EvChan) > 1 {
			<-EvChan
		}
		select {
		case ev := <-EvChan:
			if ev.Kind == hook.MouseMove {
				gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
				drawLine(int32(ev.X), 0, int32(ev.X), int32(h))
				drawLine(0, int32(ev.Y), int32(w), int32(ev.Y))
			}
		}
		window.SwapBuffers()
	}
}
