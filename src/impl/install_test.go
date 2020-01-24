package impl

import (
	"testing"
)

func TestInstall(t *testing.T) {
	//Install("../tests/goodpkg.spm")
	KeepWorkArea = true
	PkgPassword = "Thisisagoodpassword"
	Install("/tmp/sp.spm")
}

func TestInstallShow(t *testing.T) {
	//Install("../tests/goodpkg.spm")
	ShowOption = true
	PkgPassword = "Thisisagoodpassword"
	Install("/tmp/sp.spm")
}
