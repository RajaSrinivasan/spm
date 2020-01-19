package impl

import (
	"log"
	"testing"
)

func TestBuild(t *testing.T) {
	Build("../tests/spm.yaml")
	Build("../tests/goodpkg.yaml")
}

func TestMakePackageName(t *testing.T) {
	var names = []string{"redirect.yaml", "redirect.cfg", "redirect", "../tests/redirect.yaml"}
	for _, nm := range names {
		pn := makePackageName(nm)
		log.Printf("%s -> %s\n", nm, pn)
	}

}
