package plg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"plugin"
)

func NewPLG() *PLG {
	return &PLG{registry: registry{}}
}

// GetSymbol returns a symbol for the given interface name
func (plg PLG) GetSymbol(iface string) (plugin.Symbol, bool) {
	services, ok := plg.registry[iface]
	if !ok {
		return nil, false
	}
	service := services[0]
	if service == nil {
		return nil, false
	}
	return service.Symb, true
}

func (plg PLG) RegisterPlugin(absObjPath string) error {
	plug, err := plugin.Open(absObjPath)
	if err != nil {
		fmt.Println("error: could not open plugin", absObjPath)
		return err
	}

	plugMeta, err := readPlugMetadata(absObjPath[:len(absObjPath)-3])
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, symbMeta := range plugMeta.Services {
		symb, err := plug.Lookup(symbMeta.Symb)
		if err != nil {
			fmt.Println(err)
		} else {
			plg.registerSymbol(symbMeta.IFace, symb)
		}
	}
	return nil
}

func readPlugMetadata(absPlugPath string) (*PlugMetadata, error) {
	data, err := ioutil.ReadFile(absPlugPath + ".json")
	if err != nil {
		return nil, errors.New("cannot read metadata (" + absPlugPath + ".json)")
	}
	var plugMeta PlugMetadata
	err = json.Unmarshal(data, &plugMeta)
	if err != nil {
		return nil, err
	}
	return &plugMeta, nil
}

func (plg PLG) registerSymbol(iface string, symb plugin.Symbol) {
	service := Service{
		IFace: iface,
		Symb:  symb,
	}
	plg.registry[iface] = append(plg.registry[iface], &service)
	fmt.Println("registered symbol with interface:", iface)
}
