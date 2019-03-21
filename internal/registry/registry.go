package registry

import "plugin"

type Service struct {
	name   string
	symbol plugin.Symbol
}

type Services map[string]Service

type Plugins map[string]*plugin.Plugin
