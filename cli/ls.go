package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/lmuench/plg/plg"
)

func Ls() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir, err := os.Open(wd)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	filenames, err := dir.Readdirnames(-1)
	if err != nil {
		log.Fatal(err)
	}

	plugCnt := 0
	for _, filename := range filenames {
		if filepath.Ext(filename) == ".go" {
			plugCnt++
			name := filename[:len(filename)-3]
			data, err := ioutil.ReadFile(name + ".json")
			fmt.Println(plugCnt, "plugin:", name)
			if err != nil {
				fmt.Print("↳ error: metadata (", name, ".json) is missing!\n")
			} else {
				var plugMeta plg.PlugMetadata
				err = json.Unmarshal(data, &plugMeta)
				if err != nil {
					fmt.Print("↳ error: metadata parsing error (", err, ")\n")
				}
				for _, symbMeta := range plugMeta.Services {
					fmt.Print("↳ interface: ", symbMeta.IFace, ", symbol: ", symbMeta.Symb, "\n")
				}
			}
		}
	}
}
