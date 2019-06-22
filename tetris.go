//REMINDER: (invisible) Top row is 0!

package main

import (
	"fmt"
	"math/rand"
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const windowx, windowy int32 = 500, 723 // orig :363, 723

const PlayAreax int32 = 360
const PlayAreay int32 = 720
const PlayAreaStartx int32 = 0
const PlayAreaStarty int32 = 0
const TotalBrickSize int32 = 36 // This includes borders
const BrickSize int32 = 30      // This doesn't
const BorderSize int32 = (TotalBrickSize - BrickSize) / 2

var shapes = [7]string{"sshape", "zshape", "jshape", "lshape", "tshape", "ishape", "oshape"}

var PlayArea [21][10]brick

type brick struct {
	color  int8 //
	IsFull bool
}

func CheckFullLine() {
	for y := int32(0); y < 21; y++ {
		for x := int32(0); x < 10; x++ {
			if !PlayArea[y][x].IsFull {
				break
			}
			if x == 9 {
				for i := y; i > 0; i-- {
					PlayArea[i] = PlayArea[i-1]
				}
			}
		}
	}
}

func MovePawn(pawn []raylib.Vector2, key int8) (fallen bool) {
	fallen = false
	var pathClear bool = true
	//##############################################################################################
	if key == 0 {
		for i := 0; i < 4; i++ {
			if int(pawn[i].Y)+1 >= 2 && pawn[i].Y+1 < 21 {
				if PlayArea[int(pawn[i].Y)+1][int(pawn[i].X)].IsFull {
					pathClear = false
				}
			}
			if int(pawn[i].Y) == 20 {
				pathClear = false
			}
		}
		if pathClear {
			for i := 0; i < 4; i++ {
				pawn[i].Y = pawn[i].Y + 1
			}
		} else {
			fallen = true
			return
		}
		//##############################################################################################
	} else if key == 1 { //1 = left
		for i := 0; i < 4; i++ {
			if pawn[i].X == 0 {
				pathClear = false
			} else if pawn[i].Y > 0 && pawn[i].Y < 21 {
				if PlayArea[int(pawn[i].Y)][int(pawn[i].X-1)].IsFull {
					pathClear = false
				}
			}
		}
		if pathClear {
			for i := 0; i < 4; i++ {
				pawn[i].X = pawn[i].X - 1
			}
		}
		//##############################################################################################
	} else if key == 2 { //2 = Right
		for i := 0; i < 4; i++ {
			if pawn[i].X == 9 {
				pathClear = false
			} else if pawn[i].Y > 0 && pawn[i].Y < 21 {
				if PlayArea[int(pawn[i].Y)][int(pawn[i].X+1)].IsFull {
					pathClear = false
				}
			}
		}
		if pathClear {
			for i := 0; i < 4; i++ {
				pawn[i].X = pawn[i].X + 1
			}
		}
	} else if key == 3 { //3 = soft drop
		fallen = MovePawn(pawn[:], 0)
		return
	} else if key == 4 { //4 = hard drop

		fallen = MovePawn(pawn[:], 0)
		for {
			if !fallen {
				fallen = MovePawn(pawn[:], 0)
			} else {
				break
			}
		}
		return
	}
	return
}

func Createpawn(pawn []raylib.Vector2, randNum int8) {
	switch randNum {
	case 0: //Sshape
		pawn[0].X = 5
		pawn[0].Y = -1
		pawn[1].X = 6
		pawn[1].Y = -1
		pawn[2].X = 4
		pawn[2].Y = 0
		pawn[3].X = 5
		pawn[3].Y = 0
	case 1: //Zshape
		pawn[0].X = 5
		pawn[0].Y = -1
		pawn[1].X = 4
		pawn[1].Y = -1
		pawn[2].X = 5
		pawn[2].Y = 0
		pawn[3].X = 6
		pawn[3].Y = 0
	case 2: //Jshape
		pawn[0].X = 5
		pawn[0].Y = -1
		pawn[1].X = 5
		pawn[1].Y = -2
		pawn[2].X = 4
		pawn[2].Y = 0
		pawn[3].X = 5
		pawn[3].Y = 0
	case 3: //Lshape
		pawn[0].X = 5
		pawn[0].Y = -1
		pawn[1].X = 5
		pawn[1].Y = -2
		pawn[2].X = 5
		pawn[2].Y = 0
		pawn[3].X = 6
		pawn[3].Y = 0
	case 4: // Tshape
		pawn[0].X = 5
		pawn[0].Y = -1
		pawn[1].X = 4
		pawn[1].Y = -1
		pawn[2].X = 6
		pawn[2].Y = -1
		pawn[3].X = 5
		pawn[3].Y = 0
	case 5: //Ishape
		pawn[0].X = 5
		pawn[0].Y = -2
		pawn[1].X = 5
		pawn[1].Y = -3
		pawn[2].X = 5
		pawn[2].Y = -1
		pawn[3].X = 5
		pawn[3].Y = 0
	case 6: //Oshape //Doesn't matter since we won't rotate this
		pawn[0].X = 5
		pawn[0].Y = -1
		pawn[1].X = 6
		pawn[1].Y = -1
		pawn[2].X = 5
		pawn[2].Y = 0
		pawn[3].X = 6
		pawn[3].Y = 0

	case 8:
		for i := int8(0); i < 4; i++ {
			pawn[i].X = 0
			pawn[i].Y = 0
		}
	default:
	}
}

func RotatePawn(pawn []raylib.Vector2, currentShape *int8, currentRotation *int8) { //Check if currentrot is same in main YOU ARE NOT USİNG POINTERS (i think)
	fmt.Println("Before rot: ", *currentRotation, "    Pawn pos: ", pawn)
	switch *currentShape { //Bet there is a better way of doing this... But WHO cares!
	case 0: // Sshape
		switch *currentRotation {
		case 0:
			if pawn[1].X-2 >= 0 && pawn[2].Y-2 >= 0 {
				if !(PlayArea[int(pawn[1].Y)][int(pawn[1].X)-2].IsFull || PlayArea[int(pawn[2].Y)-2][int(pawn[2].X)].IsFull || PlayArea[int(pawn[3].Y)][int(pawn[3].X)].IsFull) {
					pawn[1].X = pawn[1].X - 2

					pawn[2].Y = pawn[2].Y - 2
					*currentRotation = 1
				}
			}
		case 1:
			if pawn[1].X+2 < 21 && pawn[2].Y+2 < 21 {
				if !(PlayArea[int(pawn[1].Y)][int(pawn[1].X)+2].IsFull || PlayArea[int(pawn[2].Y)+2][int(pawn[2].X)].IsFull || PlayArea[int(pawn[3].Y)][int(pawn[3].X)].IsFull) {
					pawn[1].X = pawn[1].X + 2

					pawn[2].Y = pawn[2].Y + 2
					*currentRotation = 0
				}
			}
		}
	case 1: //Zshape
		switch *currentRotation {
		case 0:
			if pawn[3].X-2 >= 0 && pawn[2].Y-2 >= 0 {
				if !(PlayArea[int(pawn[1].Y)][int(pawn[1].X)].IsFull || PlayArea[int(pawn[2].Y)-2][int(pawn[2].X)].IsFull || PlayArea[int(pawn[3].Y)][int(pawn[3].X)-2].IsFull) {
					pawn[3].X = pawn[3].X - 2
					pawn[2].Y = pawn[2].Y - 2
					*currentRotation = 1
				}
			}
		case 1:
			if pawn[3].X+2 < 10 && pawn[2].X+2 < 21 {
				if !(PlayArea[int(pawn[1].Y)][int(pawn[1].X)].IsFull || PlayArea[int(pawn[2].Y)+2][int(pawn[2].X)].IsFull || PlayArea[int(pawn[3].Y)][int(pawn[3].X)+2].IsFull) {
					pawn[3].X = pawn[3].X + 2
					pawn[2].Y = pawn[2].Y + 2
					*currentRotation = 0
				}
			}
		}
	case 2: //Jshape
		switch *currentRotation {
		case 0:
			if pawn[1].X+1 < 10 && pawn[1].Y+1 < 21 || pawn[2].Y-1 >= 0 && pawn[3].X+1 < 10 {
				if !(PlayArea[int(pawn[1].Y)+1][int(pawn[1].X)+1].IsFull || PlayArea[int(pawn[2].Y)-1][int(pawn[2].X)].IsFull || PlayArea[int(pawn[3].Y)][int(pawn[3].X)+1].IsFull) {
					pawn[1].X = pawn[1].X + 1
					pawn[1].Y = pawn[1].Y + 1

					pawn[2].Y = pawn[2].Y - 1

					pawn[3].X = pawn[3].X + 1
					*currentRotation = 1
				}
			}
		case 1:
			if pawn[1].Y-1 >= 0 || pawn[2].X+1 < 10 && pawn[2].Y-1 >= 0 && pawn[3].X-1 >= 0 {
				if !(PlayArea[int(pawn[1].Y)-1][int(pawn[1].X)].IsFull || PlayArea[int(pawn[2].Y)-1][int(pawn[2].X)+1].IsFull || PlayArea[int(pawn[3].Y)][int(pawn[3].X)-1].IsFull) {
					pawn[1].Y = pawn[1].Y - 1

					pawn[2].X = pawn[2].X + 1
					pawn[2].Y = pawn[2].Y - 1

					pawn[3].X = pawn[3].X - 1
					*currentRotation = 2
				}
			}
		case 2:
			if pawn[1].Y+1 < 21 && pawn[2].X-1 >= 0 && pawn[3].X-1 >= 0 && pawn[3].Y-1 >= 0 {
				if !(PlayArea[int(pawn[1].Y)+1][int(pawn[1].X)].IsFull || PlayArea[int(pawn[2].Y-2)][int(pawn[2].X)].IsFull || PlayArea[int(pawn[3].Y)-1][int(pawn[3].X)-1].IsFull) {
					pawn[1].Y = pawn[1].Y + 1

					pawn[2].X = pawn[2].X - 1

					pawn[3].X = pawn[3].X - 1
					pawn[3].Y = pawn[3].Y - 1
					*currentRotation = 3
				}
			}
		case 3:
			if pawn[1].X-1 >= 0 && pawn[1].Y-1 >= 0 && pawn[2].Y+2 < 21 && pawn[3].X+1 < 10 && pawn[3].Y < 21 {
				if !(PlayArea[int(pawn[1].Y)-1][int(pawn[1].X)-1].IsFull || PlayArea[int(pawn[2].Y)+2][int(pawn[2].X)].IsFull || PlayArea[int(pawn[3].Y)+1][int(pawn[3].X)+1].IsFull) {
					pawn[1].X = pawn[1].X - 1
					pawn[1].Y = pawn[1].Y - 1

					pawn[2].Y = pawn[2].Y + 2

					pawn[3].X = pawn[3].X + 1
					pawn[3].Y = pawn[3].Y + 1
					*currentRotation = 0
				}
			}
		}

	case 3: //Lshape //CHECK THE EXPCEPTIONS
		switch *currentRotation {
		case 0:
			if pawn[1].X-1 >= 0 && pawn[1].Y+1 < 21 && pawn[2].X-1 >= 0 && pawn[3].Y-1 >= 0 {
				if !(PlayArea[int(pawn[1].Y)+1][int(pawn[1].X)-1].IsFull || PlayArea[int(pawn[2].Y)][int(pawn[2].X)-1].IsFull || PlayArea[int(pawn[3].Y)-1][int(pawn[3].X)].IsFull) {
					pawn[1].X = pawn[1].X - 1
					pawn[1].Y = pawn[1].Y + 1

					pawn[2].X = pawn[2].X - 1

					pawn[3].Y = pawn[3].Y - 1
					*currentRotation = 1
				}
			}
		case 1:
			if pawn[1].Y-1 >= 0 && pawn[2].X+1 < 10 && pawn[3].X-1 >= 0 && pawn[3].Y-1 >= 0 {
				if !(PlayArea[int(pawn[1].Y)-1][int(pawn[1].X)].IsFull || PlayArea[int(pawn[2].Y)][int(pawn[2].X)].IsFull || PlayArea[int(pawn[3].Y)][int(pawn[3].X)].IsFull) {
					pawn[1].Y = pawn[1].Y - 1

					pawn[2].X = pawn[2].X + 1

					pawn[3].X = pawn[3].X - 1
					pawn[3].Y = pawn[3].Y - 1
					*currentRotation = 2
				}
			}
		case 2:
			if pawn[1].Y+1 < 21 && pawn[2].X+1 < 21 && pawn[2].Y-1 >= 0 && pawn[3].X+1 < 10 {
				if !(PlayArea[int(pawn[1].Y)+1][int(pawn[1].X)].IsFull || PlayArea[int(pawn[2].Y-1)][int(pawn[2].X+1)].IsFull || PlayArea[int(pawn[3].Y)][int(pawn[3].X+1)].IsFull) {
					pawn[1].Y = pawn[1].Y + 1

					pawn[2].X = pawn[2].X + 1
					pawn[2].Y = pawn[2].Y - 1

					pawn[3].X = pawn[3].X + 1
					*currentRotation = 3
				}
			}

		case 3:
			if pawn[1].X+1 < 10 && pawn[1].Y-1 >= 0 && pawn[2].X-1 >= 0 && pawn[2].Y+1 < 21 && pawn[3].Y+2 < 21 {
				if !(PlayArea[int(pawn[1].Y)-1][int(pawn[1].X)+1].IsFull || PlayArea[int(pawn[2].Y)+1][int(pawn[2].X)-1].IsFull || PlayArea[int(pawn[3].Y)+2][int(pawn[3].X)].IsFull) {
					pawn[1].X = pawn[1].X + 1
					pawn[1].Y = pawn[1].Y - 1

					pawn[2].X = pawn[2].X - 1
					pawn[2].Y = pawn[2].Y + 1

					pawn[3].Y = pawn[3].Y + 2
					*currentRotation = 0
				}
			}
		}
	case 4: //Tshape //CHECK EXPECTIONS

		switch *currentRotation {
		case 0:
			if pawn[2].X-1 >= 0 && pawn[2].Y-1 >= 0 {
				if !(PlayArea[int(pawn[1].Y)][int(pawn[1].X)].IsFull || PlayArea[int(pawn[2].Y)-1][int(pawn[2].X)-1].IsFull || PlayArea[int(pawn[3].Y)][int(pawn[3].X)].IsFull) {
					pawn[2].X = pawn[2].X - 1
					pawn[2].Y = pawn[2].Y - 1
					*currentRotation = 1
				}
			}
		case 1:
			if pawn[3].Y-1 >= 0 && pawn[3].X+1 < 10 {
				if !(PlayArea[int(pawn[1].Y)][int(pawn[1].X)].IsFull || PlayArea[int(pawn[2].Y)][int(pawn[2].X)].IsFull || PlayArea[int(pawn[3].Y)-1][int(pawn[3].X)+1].IsFull) {
					pawn[3].Y = pawn[3].Y - 1
					pawn[3].X = pawn[3].X + 1
					*currentRotation = 2
				}
			}
		case 2:
			if pawn[1].X+1 < 10 && pawn[1].Y+1 < 21 {
				if !(PlayArea[int(pawn[1].Y)+1][int(pawn[1].X)+1].IsFull || PlayArea[int(pawn[2].Y)][int(pawn[2].X)].IsFull || PlayArea[int(pawn[3].Y)][int(pawn[3].X)].IsFull) {
					pawn[1].X = pawn[1].X + 1
					pawn[1].Y = pawn[1].Y + 1
					*currentRotation = 3
				}
			}
		case 3:
			if pawn[2].X+1 < 10 && pawn[2].Y+1 < 21 && pawn[1].X-1 >= 0 && pawn[1].Y-1 >= 0 && pawn[3].X-1 >= 0 && pawn[3].Y+1 < 21 {
				if !(PlayArea[int(pawn[1].Y)-1][int(pawn[1].X)-1].IsFull || PlayArea[int(pawn[2].Y)+1][int(pawn[2].X)+1].IsFull || PlayArea[int(pawn[3].Y)+1][int(pawn[3].X)-1].IsFull) {
					pawn[2].X = pawn[2].X + 1
					pawn[2].Y = pawn[2].Y + 1

					pawn[1].X = pawn[1].X - 1
					pawn[1].Y = pawn[1].Y - 1

					pawn[3].X = pawn[3].X - 1
					pawn[3].Y = pawn[3].Y + 1
					*currentRotation = 0
				}
			}
		}

	case 5:
		switch *currentRotation {
		case 0:
			if pawn[1].X-1 >= 0 && pawn[1].Y+1 < 21 && pawn[2].X+1 < 10 && pawn[2].Y-1 >= 0 && pawn[3].X+2 < 10 && pawn[3].Y-2 >= 0 {
				if !(PlayArea[int(pawn[1].Y)+1][int(pawn[1].X)-1].IsFull || PlayArea[int(pawn[2].Y)-1][int(pawn[2].X)+1].IsFull || PlayArea[int(pawn[3].Y)-2][int(pawn[3].X)+2].IsFull) {
					pawn[1].X = pawn[1].X - 1
					pawn[1].Y = pawn[1].Y + 1

					pawn[2].X = pawn[2].X + 1
					pawn[2].Y = pawn[2].Y - 1

					pawn[3].X = pawn[3].X + 2
					pawn[3].Y = pawn[3].Y - 2
					*currentRotation = 1
				}
			}
		case 1:
			if pawn[1].X+1 < 10 && pawn[1].Y-1 >= 0 && pawn[2].X-1 >= 0 && pawn[2].Y+1 < 21 && pawn[3].X-2 >= 0 && pawn[3].Y+2 < 21 {
				if !(PlayArea[int(pawn[1].Y)-1][int(pawn[1].X)+1].IsFull || PlayArea[int(pawn[2].Y)+1][int(pawn[2].X)-1].IsFull || PlayArea[int(pawn[3].Y)+2][int(pawn[3].X)-2].IsFull) {
					pawn[1].X = pawn[1].X + 1
					pawn[1].Y = pawn[1].Y - 1

					pawn[2].X = pawn[2].X - 1
					pawn[2].Y = pawn[2].Y + 1

					pawn[3].X = pawn[3].X - 2
					pawn[3].Y = pawn[3].Y + 2
					*currentRotation = 0
				}
			}
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

	var pawn [4]raylib.Vector2
	var nextShapes [5]int8

	var isDeleted bool = false
	var frame int64
	var gameOver bool = false
	var isFallen bool = false

	var cs int8 = 0
	var cr int8 = 0

	var currentShape *int8 = &cs
	var currentRotation *int8 = &cr

	rand.Seed(time.Now().UTC().UnixNano())

	for y := int32(11); y < 21; y++ {
		for x := int32(0); x < 10; x++ {
			PlayArea[y][x].IsFull = false
		}
	}

	for i := int8(0); i < 5; i++ { //
		nextShapes[i] = 1 //int8(rand.Intn(7))
	}

	updateNextShape(nextShapes[:])

	PlayArea[20][5].IsFull = true
	PlayArea[15][3].IsFull = false
	PlayArea[13][1].IsFull = false
	PlayArea[11][7].IsFull = false

	raylib.InitWindow(windowx, windowy, "Tetris!")

	raylib.SetTargetFPS(60)

	CheckFullLine()
	Createpawn(pawn[:], nextShapes[0])
	*currentShape = nextShapes[0]
	*currentRotation = 0
	updateNextShape(nextShapes[:])

	for !raylib.WindowShouldClose() { //Game loop

		if !gameOver { //Checking for keys
			if raylib.IsKeyPressed(raylib.KeyJ) { //J is left

				MovePawn(pawn[:], 1)

			}
			if raylib.IsKeyPressed(raylib.KeyL) { //J is left
				MovePawn(pawn[:], 2)
			}
			if raylib.IsKeyDown(raylib.KeyM) { //M is soft drop
				isFallen = MovePawn(pawn[:], 3)
				if isFallen {
					for i := 0; i < 4; i++ {
						if pawn[i].Y < 0 {
							gameOver = true
							fmt.Println("Game over called in soft drop!")
							break
						} else {
							PlayArea[int(pawn[i].Y)][int(pawn[i].X)].IsFull = true
						}
					}
					if !gameOver {
						Createpawn(pawn[:], nextShapes[0])
						*currentShape = nextShapes[0]
						*currentRotation = 0
						updateNextShape(nextShapes[:])
					}
				}

			}
			if raylib.IsKeyPressed(raylib.KeySpace) { //Space is hard drop
				MovePawn(pawn[:], 4)

				for i := 0; i < 4; i++ {

					if pawn[i].Y < 0 {
						fmt.Println("Game over called in hard drop!")
						gameOver = true
					}
					if pawn[i].Y > 0 {
						PlayArea[int(pawn[i].Y)][int(pawn[i].X)].IsFull = true
					}
				}
				isFallen = false
				Createpawn(pawn[:], nextShapes[0])
				*currentShape = nextShapes[0]
				*currentRotation = 0
				updateNextShape(nextShapes[:])

			}
			if raylib.IsKeyPressed(raylib.KeyK) { //K is rotate left

				RotatePawn(pawn[:], currentShape, currentRotation)
				fmt.Println(*currentRotation)

			}
		}

		//Ensuring top row is always empty
		for i := int32(0); i < 10; i++ {
			PlayArea[0][i].IsFull = false
		}

		frame++
		if !gameOver {
			if frame%6000 == 0 {
				isFallen = MovePawn(pawn[:], 0)

				if isFallen {

					for i := 0; i < 4; i++ {

						if pawn[i].Y < 0 {
							fmt.Println("Game over called in main!")
							gameOver = true
						}
						if pawn[i].Y > 0 {
							PlayArea[int(pawn[i].Y)][int(pawn[i].X)].IsFull = true
						}
					}
					if !gameOver {
						Createpawn(pawn[:], nextShapes[0])
						*currentShape = nextShapes[0]
						*currentRotation = 0
						updateNextShape(nextShapes[:])
					} else {
						if !isDeleted {
							Createpawn(pawn[:], nextShapes[0])
							*currentShape = nextShapes[0]
							*currentRotation = 0
							updateNextShape(nextShapes[:])
							isFallen = false
							isDeleted = true
						}
					}

				} else {

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

		for i := 0; i < 4; i++ { //Draw the current falling piece //AKA "The Mighty Pawn"
			if pawn[i].Y > 0 {
				raylib.DrawRectangle(int32(pawn[i].X)*TotalBrickSize+BorderSize, (int32(pawn[i].Y)-1)*TotalBrickSize+BorderSize, BrickSize+BorderSize, BrickSize+BorderSize, raylib.Color{0, 255, 0, 255})
			}
		}

		for i := 0; i < 21; i++ {
			raylib.DrawRectangle(PlayAreaStartx, PlayAreaStarty+int32(i)*TotalBrickSize, 360, BorderSize, raylib.Color{155, 155, 155, 255}) // horizontal lines
		}

		for i := 0; i < 11; i++ {
			raylib.DrawRectangle(PlayAreaStartx+int32(i)*TotalBrickSize, PlayAreaStarty, BorderSize, 720, raylib.Color{155, 155, 155, 255}) // vertical lines
		}

		raylib.EndDrawing()
	}

	raylib.CloseWindow()
}
