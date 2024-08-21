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
	goFiles = append(goFiles, "run")

	allFiles, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		fmt.Scanln()
		os.Exit(1)
	}

	for f := range slices.Values(allFiles) {
		splitFName := strings.Split(f.Name(), ".")
		if !f.IsDir() && splitFName[len(splitFName)-1] == "go" {
			goFiles = append(goFiles, f.Name())
		}
	}

	cmd := exec.Command("go", goFiles...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}
