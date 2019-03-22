package fs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type PlugMetadata struct {
	Services []IFaceAndSymb
}

type IFaceAndSymb struct {
	IFace string `json:"interface"`
	Symb  string `json:"symbol"`
}

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
				var metadata PlugMetadata
				err = json.Unmarshal(data, &metadata)
				if err != nil {
					fmt.Print("↳ error: metadata parsing error (", err, ")\n")
				}
				for _, ias := range metadata.Services {
					fmt.Print("↳ interface: ", ias.IFace, ", symbol: ", ias.Symb, "\n")
				}
			}
		}
	}
}
