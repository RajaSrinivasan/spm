package impl

import (
	"testing"
)

func TestInstall(t *testing.T) {
	//Install("../tests/goodpkg.spm")
	KeepWorkArea = true
	PkgPassword = "Thisisagoodpassword"
	Install("../systest/sp.spm")
}
