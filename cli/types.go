package cli

type PlugMetadata struct {
	Services []SymbMetadata
}

type SymbMetadata struct {
	IFace string `json:"interface"`
	Symb  string `json:"symbol"`
}
