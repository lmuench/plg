package plg

import "plugin"

type PLG struct {
	registry registry
}
type registry map[string][]*Service

type Service struct {
	IFace  string
	Symbol plugin.Symbol
}

func NewPLG(plugPath string) PLG {
	return PLG{
		registry: registry{},
	}
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
	return service.Symbol, true
}

func (plg PLG) RegisterPlugin(iface string, plug plugin.Plugin) error {
	symb, err := plug.Lookup(iface)
	if err != nil {
		return err
	}
	plg.registerSymbol(iface, symb)
	return nil
}

func (plg PLG) registerSymbol(iface string, symb plugin.Symbol) {
	service := Service{
		IFace:  iface,
		Symbol: symb,
	}
	plg.registry[iface] = append(plg.registry[iface], &service)
}
