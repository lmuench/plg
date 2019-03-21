package fs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Ls() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pluginsDir := filepath.Join(wd, "plugins")

	dir, err := os.Open(pluginsDir)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	names, err := dir.Readdirnames(-1)
	if err != nil {
		log.Fatal(err)
	}

	for id, name := range names {
		if filepath.Ext(name) == ".go" {
			fmt.Println(id, name[:len(name)-3])
		}
	}
}
