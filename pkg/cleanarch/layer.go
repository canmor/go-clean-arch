package cleanarch

import (
	"os"
	"strings"
)

type Layer int

func (l Layer) String() (result string) {
	result = map[Layer]string{
		Domain:      "domain layer",
		Application: "application layer",
		Interface:   "interface layer",
	}[l]
	return
}

const (
	UnknownLayer Layer = iota
	Domain
	Application
	Interface
)

func strToLayer(pkg string) Layer {
	maps := map[string]Layer{
		"domain":       Domain,
		"domains":      Domain,
		"entity":       Domain,
		"entities":     Domain,
		"application":  Application,
		"applications": Application,
		"app":          Application,
		"use_case":     Application,
		"use_cases":    Application,
		"usecase":      Application,
		"usecases":     Application,
		"interface":    Interface,
		"adapter":      Interface,
	}
	return maps[pkg]
}

func NewLayer(packagePath string) Layer {
	pkg := strings.ToLower(packagePath)
	parts := strings.Split(pkg, string(os.PathSeparator))
	for i := 0; i < len(parts); i++ {
		layer := strToLayer(parts[len(parts)-1-i])
		if layer != UnknownLayer {
			return layer
		}
	}
	return UnknownLayer
}
