package knowledge

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func Operation_1() {
	if runtime.GOOS == "darwin" {
		fmt.Println("Running on macOS")
		// Example command to get terminal name (might not be universally applicable)
		cmd := exec.Command("ps", "-ww", "-o", "comm")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error getting output:", err)
			return
		}

		// Split the output into lines
		lines := strings.Split(string(output), "\n")

		// Check each line for "bash"
		containsBash := false
		for _, line := range lines {
			if strings.Contains(line, "bash") {
				containsBash = true
				fmt.Println("Found 'bash' in the output:", line)
				break // Exit the loop if found
			}
		}

		if !containsBash {
			fmt.Println("'bash' not found in the output")
		}
	} else {
		fmt.Println("Not running on macOS")
	}
}
