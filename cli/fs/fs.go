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

	for id, filename := range filenames {
		if filepath.Ext(filename) == ".go" {
			name := filename[:len(filename)-3]
			data, err := ioutil.ReadFile(name + ".json")
			if err != nil {
				fmt.Print(id, " ", name, ": metadata (", name, ".json) is missing!\n")
			} else {
				var metadata PlugMetadata
				err = json.Unmarshal(data, &metadata)
				if err != nil {
					fmt.Println(id, name, ": metadata parsing error ", err)
				}
				fmt.Println(id, name)
				for n, ias := range metadata.Services {
					fmt.Print("↳ ", id, ".", n+1, " interface: ", ias.IFace, ", symbol: ", ias.Symb, "\n")
				}
			}
		}
	}
}
