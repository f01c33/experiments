package main

import (
	"crypto/md5"
	"io/ioutil"
	"os"
	"time"

	"github.com/gabstv/ebiten-imgui/renderer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/inkyblackness/imgui-go/v2"
	lua "github.com/yuin/gopher-lua"
)

var (
	W, H int //= 800,600
	L    *lua.LState
)

func main() {
	mgr := renderer.New(nil)
	W, H = ebiten.ScreenSizeInFullscreen()
	W /= 2
	H /= 2
	ebiten.SetWindowSize(W, H)
	ebiten.SetScreenTransparent(true)
	// ebiten.SetWindowDecorated(false)
	ebiten.SetWindowResizable(true)
	ebiten.SetInitFocused(true)
	ebiten.SetRunnableOnUnfocused(true)
	// ebiten.SetWindowPosition(0, 0)
	gg := &G{
		mgr:   mgr,
		scale: 230,
	}
	L = lua.NewState()
	defer L.Close()
	go watch_base_changes()
	ebiten.RunGame(gg)
}

func watch_base_changes() {
	arg := "main.lua"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	hash := [16]byte{}
	for {
		time.Sleep(time.Second)
		data, err := ioutil.ReadFile(arg)
		if err != nil {
			panic(err)
		}
		newHash := md5.Sum(data)
		if newHash != hash {
			hash = newHash
			if err := L.DoFile(arg); err != nil {
				panic(err)
			}
		}
	}
}

type G struct {
	mgr *renderer.Manager
	// demo members:
	lines []string
	// clearColor [3]float32
	// floatVal   float32
	// counter    int
	command string
	scale   int32
	// selected   bool
}

func (g *G) Draw(screen *ebiten.Image) {
	// screen.Fill(color.RGBA{uint8(g.clearColor[0] * 255), uint8(g.clearColor[1] * 255), uint8(g.clearColor[2] * 255), 255})

	g.mgr.BeginFrame()
	imgui.SetNextWindowSize(imgui.Vec2{X: float32(W), Y: float32(H)})

	if imgui.Button("Quit") {
		os.Exit(0)
	}
	imgui.InputInt("font-scaling", &g.scale)
	imgui.CurrentIO().SetFontGlobalScale(float32(g.scale) / 100)
	// imgui.Text("ภาษาไทย测试조선말")                     // To display these, you'll need to register a compatible font
	// imgui.Text("Hello, world!")                       // Display some text
	// imgui.SliderFloat("float", &g.floatVal, 0.0, 1.0) // Edit 1 float using a slider from 0.0f to 1.0f
	// imgui.ColorEdit3("clear color", &g.clearColor)    // Edit 3 floats representing a color

	// imgui.Checkbox("Demo Window", &g.selected)
	// if g.selected {
	// 	imgui.ShowDemoWindow(&g.selected)
	// }

	// if imgui.Button("Button") { // Buttons return true when clicked (most widgets return true when edited/activated)
	// 	g.counter++
	// }
	// imgui.SameLine()
	// imgui.Text(fmt.Sprintf("counter = %d", g.counter))

	imgui.InputText("command", &g.command)
	if imgui.Button("Execute") {

		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		rescueStderr := os.Stderr
		re, we, _ := os.Pipe()
		os.Stderr = we
		if err := L.DoString(g.command); err != nil {
			g.lines = append(g.lines, "-- $ "+err.Error())
		} else {
			g.command = ""
		}
		g.lines = append(g.lines, g.command)
		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout
		we.Close()
		outE, _ := ioutil.ReadAll(re)
		os.Stderr = rescueStderr
		if len(out) > 0 {
			g.lines = append(g.lines, "-- > "+string(out))
		}
		if len(outE) > 0 {
			g.lines = append(g.lines, "-- & "+string(outE))
		}
		// imgui.Text(fmt.Sprint(L.Get(-1)))
	}
	for i := len(g.lines) - 1; i >= 0; i-- {
		imgui.Text(g.lines[i])
	}
	if len(g.lines) > 100 {
		g.lines = g.lines[len(g.lines)-100:]
	}
	// if !imgui.IsAnyItemFocused() {
	// 	ebiten.MinimizeWindow()
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyControlLeft) && ebiten.IsKeyPressed(ebiten.KeyF12) {
	// 	ebiten.RestoreWindow()
	// }
	g.mgr.EndFrame(screen)
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f", ebiten.CurrentTPS()))
}

func (g *G) Update() error {

	g.mgr.Update(1.0/60.0, float32(W), float32(H))
	return nil
}

func (g *G) Layout(outsideWidth, outsideHeight int) (int, int) {
	return W, H
}
