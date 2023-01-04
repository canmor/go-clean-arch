package cleanarch

import (
	"fmt"
	"strings"
)

type ImportChecker struct {
	thisLayer Layer
}

func NewImportChecker(path string) *ImportChecker {
	isTestPackage := strings.HasSuffix(path, "_test")
	if isTestPackage {
		return nil
	}
	layer := NewLayer(path)
	if layer == UnknownLayer {
		return nil
	}
	return &ImportChecker{layer}
}

func (c ImportChecker) Check(path string) (err error) {
	violated := false
	importLayer := NewLayer(path)
	if c.thisLayer == Application && importLayer == Interface {
		violated = true
	}
	if c.thisLayer == Domain &&
		(importLayer == Interface || importLayer == Application) {
		violated = true
	}
	if violated {
		err = fmt.Errorf("Importing <%s> aka: %s from <%s>", importLayer, path, c.thisLayer)
	}
	return
}
