package plg

import "plugin"

type PLG struct {
	registry registry
}
type registry map[string][]*Service

type Service struct {
	IFace string
	Symb  plugin.Symbol
}

type PlugMetadata struct {
	Services []SymbMetadata
}

type SymbMetadata struct {
	IFace string `json:"interface"`
	Symb  string `json:"symbol"`
}
