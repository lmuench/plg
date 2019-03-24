package plg

import (
	"fmt"
	"plugin"
)

type PLG struct {
	registry registry
}
type registry map[string][]*Service

type Service struct {
	IFace string
	Symb  plugin.Symbol
}

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

func (plg PLG) RegisterPlugin(iface string, absObjPath string) error {
	plug, err := plugin.Open(absObjPath)
	if err != nil {
		fmt.Println("error: could not open plugin", absObjPath)
		return err
	}
	symb, err := plug.Lookup(iface)
	if err != nil {
		return err
	}
	plg.registerSymbol(iface, symb)
	return nil
}

func (plg PLG) registerSymbol(iface string, symb plugin.Symbol) {
	service := Service{
		IFace: iface,
		Symb:  symb,
	}
	plg.registry[iface] = append(plg.registry[iface], &service)
	fmt.Println("registered symbol with interface:", iface)
}
