package impl

import (
	"log"
	"os"
	"testing"
)

func TestBuild(t *testing.T) {

	KeepWorkArea = true
	PkgPassword = "Thisisagoodpassword"
	wd, _ := os.Getwd()
	defer os.Chdir(wd)

	os.Chdir("../../systest")
	Build("sp.yaml", "/tmp/sp.spm")
}

func TestMakePackageName(t *testing.T) {
	var names = []string{"redirect.yaml", "redirect.cfg", "redirect", "../tests/redirect.yaml"}
	for _, nm := range names {
		pn := makePackageName(nm)
		log.Printf("%s -> %s\n", nm, pn)
	}
}
