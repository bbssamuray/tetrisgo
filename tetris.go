//REMINDER: (invisible) Top row is 0!

package main

import (
	"math/rand"
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const windowx, windowy int32 = 500, 723

const PlayAreax int32 = 360
const PlayAreay int32 = 720
const PlayAreaStartx int32 = 0
const PlayAreaStarty int32 = 0
const TotalBrickSize int32 = 36 // This includes borders
const BrickSize int32 = 30      // This doesn't
const BorderSize int32 = (TotalBrickSize - BrickSize) / 2

var scoreMultiplier float32 = 1.0
var gameOver bool = false

var PlayArea [21][10]brick

type brick struct {
	color  int8 //
	IsFull bool
}

type pawn struct {
	shapeId     int8
	originPoint raylib.Vector2
	bricksLocal []raylib.Vector2
}

func CheckFullLine() {
	for y := int32(0); y < 21; y++ {
		for x := int32(0); x < 10; x++ {
			if !PlayArea[y][x].IsFull {
				break
			}
			if x == 9 {
				scoreMultiplier -= 0.02
				for i := y; i > 0; i-- {
					PlayArea[i] = PlayArea[i-1]
				}
			}
		}
	}
}

func MovePawn(pawn *pawn, direction raylib.Vector2) (fallen bool) { //todo: Needs to adapt to new pawn format

	fallen = false

	for i := 0; i < 4; i++ {
		newLocation := raylib.Vector2{
			pawn.bricksLocal[i].X + pawn.originPoint.X + direction.X,
			pawn.bricksLocal[i].Y + pawn.originPoint.Y + direction.Y}
		if newLocation.X < 0 || newLocation.X > 9 {
			return false //Don't move if pawn is somehow out of bounds
		} else if (newLocation.Y > 20) || (newLocation.Y > 0 && PlayArea[int(newLocation.Y)][int(newLocation.X)].IsFull) {

			if direction.Y > 0 { //make the pawn solid if there is an obstacle below
				for x := 0; x < 4; x++ {
					if (pawn.bricksLocal[x].Y + pawn.originPoint.Y) < 0 { //If the solid block is above the ceiling call gameover
						gameOver = true
						return true
					}
					PlayArea[int(pawn.bricksLocal[x].Y+pawn.originPoint.Y)][int(pawn.bricksLocal[x].X+pawn.originPoint.X)].IsFull = true //make the current positions solid if next positions are blocked
				}

				fallen = true
				return fallen

			}
			if direction.X != 0 { //Don't move if the newLocation is blocked
				fallen = false
				return fallen
			}
		}
	}

	pawn.originPoint = raylib.Vector2Add(pawn.originPoint, direction) //Move pawn

	return fallen
}

func Createpawn(currentPawn *pawn, randNum int8) {

	currentPawn.bricksLocal = nil //We are clearing the slice here.
	currentPawn.originPoint = raylib.Vector2{5, -1}

	switch randNum {
	case 0: //Sshape
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, 0}) //origin Brick's local coordinate
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{-1, 0})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, -1})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{1, -1})
	case 1: //Zshape
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, 0})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{1, 0})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, -1})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{-1, -1})
	case 2: //Jshape
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, 0})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, -1})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, 1})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{-1, 1})
	case 3: //Lshape
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, 0})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, -1})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, 1})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{1, 1})
	case 4: // Tshape
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, 0})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{1, 0})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{-1, 0})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, -1})
	case 5: //Ishape
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, 0})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, -1})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, -2})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, 1})
	case 6: //Oshape
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, 0})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{1, 0})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{1, -1})
		currentPawn.bricksLocal = append(currentPawn.bricksLocal[:], raylib.Vector2{0, -1})

	}
}

func RotatePawn(pawn *pawn, rotation int) { //rotation = 1 is clockwise, rotation = -1 is anti-clockwise

	if rotation == 1 { //clockwise
		for i := 0; i < 4; i++ { //Turning clockwise
			temp := pawn.bricksLocal[i].X
			pawn.bricksLocal[i].X = pawn.bricksLocal[i].Y
			pawn.bricksLocal[i].Y = -1 * temp
		}
	} else {
		for i := 0; i < 4; i++ { //Turning anti-clockwise
			temp := pawn.bricksLocal[i].X
			pawn.bricksLocal[i].X = -1 * pawn.bricksLocal[i].Y
			pawn.bricksLocal[i].Y = temp
		}
	}

	for i := 0; i < 4; i++ { //Undo the changes if bricks are out of bounds
		//This is inefficent and can potentially cause problems
		//But it is the simplest solution i can think of
		newLocation := raylib.Vector2{
			pawn.bricksLocal[i].X + pawn.originPoint.X,
			pawn.bricksLocal[i].Y + pawn.originPoint.Y}
		if (newLocation.X < 0 || newLocation.X > 9) || (newLocation.Y > 20) || (newLocation.Y >= 0 && PlayArea[int(newLocation.Y)][int(newLocation.X)].IsFull) {
			RotatePawn(pawn, -rotation) //Rotate it to the opposite direction to undo
			return
		}
	}

}

func updateNextShape(shapes []int8) {
	for i := int8(0); i < 4; i++ {
		shapes[i] = shapes[i+1]
	}
	shapes[4] = int8(rand.Intn(7))
}

func main() {

	var pawn pawn
	var nextShapes [5]int8

	var frame int64

	var isFallen bool = false
	var gameSpeed float32

	var cs int8 = 0

	var currentShape *int8 = &cs

	rand.Seed(time.Now().UTC().UnixNano())

	for i := int8(0); i < 5; i++ { //
		nextShapes[i] = int8(rand.Intn(7))
	}

	updateNextShape(nextShapes[:])

	raylib.InitWindow(windowx, windowy, "Tetris!")

	raylib.SetTargetFPS(60)

	CheckFullLine()
	Createpawn(&pawn, nextShapes[0])
	*currentShape = nextShapes[0]
	updateNextShape(nextShapes[:])

	for !raylib.WindowShouldClose() { //Game loop
		if !gameOver { //Checking for keys
			if raylib.IsKeyPressed(raylib.KeyJ) { //J is left
				MovePawn(&pawn, raylib.Vector2{-1, 0})
			}
			if raylib.IsKeyPressed(raylib.KeyL) { //L is Right
				MovePawn(&pawn, raylib.Vector2{1, 0})
			}
			if raylib.IsKeyDown(raylib.KeyM) { //M is soft drop

				if frame%2 == 0 {
					isFallen = MovePawn(&pawn, raylib.Vector2{0, 1})
					if isFallen && !gameOver {

						Createpawn(&pawn, nextShapes[0])
						*currentShape = nextShapes[0]
						updateNextShape(nextShapes[:])

					}
				}

			}
			if raylib.IsKeyPressed(raylib.KeySpace) { //Space is hard drop
				for !isFallen {
					isFallen = MovePawn(&pawn, raylib.Vector2{0, 1})
				}
				isFallen = false
				Createpawn(&pawn, nextShapes[0])
				*currentShape = nextShapes[0]
				updateNextShape(nextShapes[:])

			}
			if raylib.IsKeyPressed(raylib.KeyK) { //K is rotate clockwise
				RotatePawn(&pawn, 1)
			}
			if raylib.IsKeyPressed(raylib.KeyU) { //U is rotate anti-clockwise
				RotatePawn(&pawn, -1)
			}

			gameSpeed = 60 * scoreMultiplier
			if frame%int64(gameSpeed) == 0 {
				isFallen = MovePawn(&pawn, raylib.Vector2{0, 1})

				if isFallen {
					if !gameOver {
						Createpawn(&pawn, nextShapes[0])
						*currentShape = nextShapes[0]
						updateNextShape(nextShapes[:])
					}
				}
			}
		}
		CheckFullLine()

		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.Gray)

		raylib.DrawRectangle(PlayAreaStartx, PlayAreaStarty, PlayAreax, PlayAreay, raylib.Color{0, 0, 0, 255}) // Inner black area

		for y := int32(1); y < 21; y++ {
			for x := int32(0); x < 10; x++ {
				if PlayArea[y][x].IsFull {
					raylib.DrawRectangle(x*TotalBrickSize+BorderSize, (y-1)*TotalBrickSize+BorderSize, BrickSize+BorderSize, BrickSize+BorderSize, raylib.Color{255, 255, 0, 255})
				}
			}
		}

		for i := 0; i < 4; i++ { //Draw the current falling piece

			if pawn.bricksLocal[i].Y+pawn.originPoint.Y > 0 {
				raylib.DrawRectangle(
					int32(pawn.bricksLocal[i].X+pawn.originPoint.X)*TotalBrickSize+BorderSize,
					(int32(pawn.bricksLocal[i].Y+pawn.originPoint.Y)-1)*TotalBrickSize+BorderSize,
					BrickSize+BorderSize, BrickSize+BorderSize, raylib.Color{0, 255, 0, 255})
			}
		}

		for i := 0; i < 21; i++ {
			raylib.DrawRectangle(PlayAreaStartx, PlayAreaStarty+int32(i)*TotalBrickSize, 360, BorderSize, raylib.Color{155, 155, 155, 255}) // horizontal lines
		}

		for i := 0; i < 11; i++ {
			raylib.DrawRectangle(PlayAreaStartx+int32(i)*TotalBrickSize, PlayAreaStarty, BorderSize, 720, raylib.Color{155, 155, 155, 255}) // vertical lines
		}

		raylib.EndDrawing()
		frame++
	}

	raylib.CloseWindow()
}
