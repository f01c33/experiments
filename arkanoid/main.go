package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jakecoffman/cp"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	ball     = ebiten.NewImage(16, 16)
	paddle   = ebiten.NewImage(16*5, 16)
	wall     = ebiten.NewImage(8, screenHeight)
	wallWide = ebiten.NewImage(screenWidth, 8)
)

func init() {
	ball, _, _ = ebitenutil.NewImageFromFile("./ball.png")
	ball.Fill(color.White)
	paddle.Fill(color.White)
	wall.Fill(color.White)
	wallWide.Fill(color.White)
}

type Game struct {
	space  *cp.Space
	paddle *cp.Shape
	ball   *cp.Shape
	wall   *cp.Shape
	wall2  *cp.Shape
	wall3  *cp.Shape
}

func NewGame() *Game {
	space := cp.NewSpace()
	space.Iterations = 1

	// The space will contain a very large number of similarly sized objects.
	// This is the perfect candidate for using the spatial hash.
	// Generally you will never need to do this.
	// space.UseSpatialHash(2.0, 10000)
	var paddle *cp.Body = space.AddBody(cp.NewBody(10, cp.INFINITY))
	var shape *cp.Shape = cp.NewBox(paddle, 16*5, 16, 0)
	shape = space.AddShape(shape)
	shape.Body().SetPosition(cp.Vector{X: screenWidth / 2, Y: screenHeight - 16*2})
	shape.SetElasticity(1.0)
	shape.SetFriction(0.0)

	var ball *cp.Body = space.AddBody(cp.NewBody(1.0, 1.0))
	var ballShape *cp.Shape = cp.NewCircle(ball, 10, cp.Vector{})
	ball.SetPosition(cp.Vector{X: screenWidth / 2, Y: 16 * 2})
	ballShape = space.AddShape(ballShape)
	ball.SetVelocity(0, 300)
	ballShape.SetElasticity(1.0)

	var wall *cp.Body = space.AddBody(cp.NewBody(cp.INFINITY, cp.INFINITY))
	var wallShape *cp.Shape = space.AddShape(cp.NewBox(wall, 8, screenHeight, 0))
	wall.SetPosition(cp.Vector{X: 0, Y: 0})
	wallShape = space.AddShape(wallShape)
	wallShape.SetElasticity(1.0)

	var wall2 *cp.Body = space.AddBody(cp.NewBody(cp.INFINITY, cp.INFINITY))
	var wall2Shape *cp.Shape = space.AddShape(cp.NewBox(wall2, 8, screenHeight, 0))
	wall2.SetPosition(cp.Vector{X: screenWidth - 8, Y: 0})
	wall2Shape = space.AddShape(wall2Shape)
	wall2Shape.SetElasticity(1.0)

	var wall3 *cp.Body = space.AddBody(cp.NewBody(cp.INFINITY, cp.INFINITY))
	var wall3Shape *cp.Shape = space.AddShape(cp.NewBox(wall3, screenWidth, 8, 0))
	wall3.SetPosition(cp.Vector{X: 0, Y: 0})
	wall3Shape = space.AddShape(wall3Shape)
	wall3Shape.SetElasticity(1.0)

	// var wallShape2 *cp.Shape = space.AddShape(cp.NewBox(wall, screenWidth, 8, 0))
	// wallShape2 = space.AddShape(wallShape2)
	// wallShape2.SetElasticity(1.0)

	// var wallShape3 *cp.Shape = space.AddShape(cp.NewBox(wall, 8, screenHeight, 0))
	// wallShape3 = space.AddShape(wallShape3)
	// wallShape.Body().SetPosition(cp.Vector{X: screenWidth - 16.0, Y: 0})
	// wallShape3.SetElasticity(1.0)

	// 	var body *cp.Body
	// 	var shape *cp.Shape

	// 	for y := 0; y < imageHeight; y++ {
	// 		for x := 0; x < imageWidth; x++ {
	// 			if getPixel(uint(x), uint(y)) == 0 {
	// 				continue
	// 			}

	// 			xJitter := 0.05 * rand.Float64()
	// 			yJitter := 0.05 * rand.Float64()

	// 			shape = makeBall(2.0*(float64(x)+imageWidth/2+xJitter)-75, 2*(imageHeight/2.0+float64(y)+yJitter)+150)
	// 			space.AddBody(shape.Body())
	// 			space.AddShape(shape)
	// 		}
	// 	}

	// 	body = space.AddBody(cp.NewBody(1e9, cp.INFINITY))
	// 	body.SetPosition(cp.Vector{X: -1000, Y: 225})
	// 	body.SetVelocity(400, 0)

	// 	shape = space.AddShape(cp.NewCircle(body, 8, cp.Vector{}))
	// 	shape.SetElasticity(0)
	// 	shape.SetFriction(0)

	return &Game{
		space:  space,
		paddle: shape,
		ball:   ballShape,
		wall:   wallShape,
		wall2:  wall2Shape,
		wall3:  wall3Shape,
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.paddle.Body().SetVelocity(-700, 0)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.paddle.Body().SetVelocity(700, 0)
	} else {
		g.paddle.Body().SetVelocity(0, 0)
	}
	g.space.Step(1.0 / float64(ebiten.MaxTPS()))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(200.0/255.0, 200.0/255.0, 200.0/255.0, 1)

	g.space.EachShape(func(s *cp.Shape) {
		op.GeoM.Reset()
		op.GeoM.Translate(s.Body().Position().X, s.Body().Position().Y)
		if s.HashId() == g.paddle.HashId() {
			screen.DrawImage(paddle, op)
		} else if s.HashId() == g.wall.HashId() {
			screen.DrawImage(wall, op)
		} else if s.HashId() == g.wall2.HashId() {
			screen.DrawImage(wall, op)
		} else if s.HashId() == g.wall3.HashId() {
			screen.DrawImage(wallWide, op)
		} else {
			// op.GeoM.Translate(0, 0)
			// op.GeoM.Rotate(math.Pi / 2.0)
			// op.GeoM.Translate(s.Body().Position().X, s.Body().Position().Y)
			screen.DrawImage(ball, op)
		}
	})

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebiten")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
