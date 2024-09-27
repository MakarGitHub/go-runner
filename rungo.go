package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func main() {
	goFiles := make([]string, 0)

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-b":
			goFiles = append(goFiles, "build")
		case "-help":
			fmt.Println("-b - Build go files")
			fmt.Scanln()
			return
		}
	} else {
		goFiles = append(goFiles, "run")
	}

	allFiles, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		fmt.Scanln()
		os.Exit(1)
	}

	for f := range slices.Values(allFiles) {
		splitFName := strings.Split(f.Name(), ".")
		if !f.IsDir() && splitFName[len(splitFName)-1] == "go" && !strings.HasSuffix(splitFName[0], "_test") {
			goFiles = append(goFiles, f.Name())
		}
	}

	cmd := exec.Command("go", goFiles...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}
