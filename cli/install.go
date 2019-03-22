package cli

import (
	"os"
	"os/exec"
)

func Install(plug string) {
	compile(plug)
}

func compile(name string) (string, error) {
	srcPath := name + ".go"
	objectPath := name + ".so"
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o="+objectPath, srcPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return objectPath, nil
}
