package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Install(plug string) error {
	absObjPath, err := compile(plug)
	if err != nil {
		fmt.Print("error: ", plug, " could not be compiled (", err, ")\n")
	}
	// TODO: read iface string from metadata
	// TODO: use gRPC to call plg.RegisterPlugin(iface, absObjPath)
}

func compile(name string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	obj := name + ".so"
	src := name + ".go"
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o="+obj, src)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}

	absObjPath := filepath.Join(wd, obj)
	return absObjPath, nil
}
