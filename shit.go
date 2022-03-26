package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("--- Security Heuristic Information Tooling ---")

	files, err := filesInGit()
	if err != nil {
		panic(err)
	}

	viols := []violation{}

	fmt.Println("--- Checking Files ---")
	for _, f := range files {
		fmt.Println("checking: ", f)

		vv := defaultRuleset.evaluate(f)
		viols = append(viols, vv...)
	}

	fmt.Println("--- Generating CRAP Report ---")
	err = writeCRAPReport(viols)
	if err != nil {
		panic(err)
	}
}

func filesInGit() ([]string, error) {
	out, err := exec.Command("git", "ls-files").Output()
	if err != nil {
		return nil, fmt.Errorf("couldn't run git ls-files: %w", err)
	}

	return strings.Split(strings.TrimSuffix(string(out), "\n"), "\n"), nil
}
