package cfg

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	LoadConfig("../tests/spm.yaml")
}

func TestSaveManifest(t *testing.T) {
	cfg := LoadConfig("../tests/spm.yaml")
	cfg.SaveManifest("../tests/Packagefile")
	cfg2 := LoadConfig("../tests/Packagefile")
	cfg2.SaveManifest("../tests/Packagefile2")
}
