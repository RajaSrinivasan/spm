package cfg

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	LoadConfig("../../tests/spm.yaml")
}

func TestSaveManifest(t *testing.T) {
	cfg, _ := LoadConfig("../../tests/spm.yaml")
	cfg.SaveManifest("../../tests/Packagefile")
	cfg2, _ := LoadConfig("../../tests/Packagefile")
	cfg2.SaveManifest("../../tests/Packagefile2")
	LoadConfig("../../tests/badspm.yaml")

}
