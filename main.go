package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func move_mouse() {
	x := rand.Intn(900)
	y := rand.Intn(500)
	robotgo.MoveSmooth(x, y, 1.0, 2.0)
	fmt.Printf("Moved mouse to: x=%d y=%d\n", x, y)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for true {
		move_mouse()
		time.Sleep(30 * time.Second)
	}
}
