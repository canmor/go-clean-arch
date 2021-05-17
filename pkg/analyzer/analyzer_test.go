package analyzer_test

import (
	"github.com/canmor/go-clean-arch/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"path/filepath"
	"testing"
)

var testdata string

func init() {
	wd, _ := os.Getwd()
	testdata = filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
}

func TestNotCleanArch(t *testing.T) {
	analysistest.Run(t, testdata, analyzer.Analyzer, "clean_arch/util")
}

func TestAppImportDomain(t *testing.T) {
	analysistest.Run(t, testdata, analyzer.Analyzer, "clean_arch/app")
}

func TestAppImportAdapter(t *testing.T) {
	analysistest.Run(t, testdata, analyzer.Analyzer, "clean_arch/app/import_adapter")
}

func TestDomainImportAdapter(t *testing.T) {
	analysistest.Run(t, testdata, analyzer.Analyzer, "clean_arch/domain/import_adapter")
}

func TestDomainImportApp(t *testing.T) {
	analysistest.Run(t, testdata, analyzer.Analyzer, "clean_arch/domain/import_app")
}