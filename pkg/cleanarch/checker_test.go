package cleanarch

import "testing"

func TestImportChecker_IgnoreTest(t *testing.T) {
	c := NewImportChecker("app/domain/model_test")
	if c != nil {
		t.Errorf("test package should be ignored, but get c='%+v'", c)
	}
}
