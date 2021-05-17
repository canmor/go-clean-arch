package cleanarch

import "testing"

func TestLayerOfPackage(t *testing.T) {
	tests := []struct {
		name          string
		expectedLayer Layer
	}{
		{"domain", Domain},
		{"domains", Domain},
		{"entity", Domain},
		{"entities", Domain},
		{"app", Application},
		{"application", Application},
		{"applications", Application},
		{"usecase", Application},
		{"usecases", Application},
		{"use_case", Application},
		{"use_cases", Application},
		{"interface", Interface},
		{"adapter", Interface},
		{"wrong", UnknownLayer},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if NewLayer(test.name) != test.expectedLayer {
				t.Errorf("package '%s' should be regconized as %s", test.name, test.expectedLayer)
			}
		})
	}
}

func TestLayerOfNestedPackage(t *testing.T) {
	tests := []struct {
		name          string
		expectedLayer Layer
	}{
		{"domain/util", Domain},
		{"app/domain/util", Domain},
		{"app/bar/util", Application},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if NewLayer(test.name) != test.expectedLayer {
				t.Errorf("package '%s' should be regconized as %s", test.name, test.expectedLayer)
			}
		})
	}
}
