package animation

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// Hello prints Hello world
func Hello() {
	fmt.Println("Hello world!!!")
	// Automatically start Doge animation after printing Hello
	AnimateDoge()
}

// Doge character frames for animation
var dogeFrames = []string{
	`       ___
      / _ \ 
     / / \ \
    /___/_\_\
     Wow! Such Doge!`,
	`        ___
       / _ \ 
      / / \ \
     /___/_\_\
      Much Move! Wow!`,
	`         ___
        / _ \ 
       / / \ \
      /___/_\_\
       So Animation!`,
}

// clearScreen clears the terminal screen based on OS
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// AnimateDoge runs the Doge animation in the terminal
func AnimateDoge() {
	fmt.Println("Starting Doge Animation! Press Ctrl+C to stop.")
	spaces := 0
	direction := 1 // 1 for right, -1 for left

	for {
		// Clear the terminal
		clearScreen()

		// Select frame based on position
		frameIndex := (spaces / 5) % len(dogeFrames)
		if frameIndex < 0 {
			frameIndex = -frameIndex
		}

		// Print spaces to shift Doge
		for i := 0; i < spaces; i++ {
			fmt.Print(" ")
		}

		// Print the current Doge frame
		fmt.Println(dogeFrames[frameIndex])

		// Update position
		spaces += direction

		// Reverse direction at boundaries
		if spaces >= 20 || spaces <= 0 {
			direction = -direction
		}

		// Delay to control animation speed
		time.Sleep(100 * time.Millisecond)
	}
}
