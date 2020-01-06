package pkg

import (
	"testing"
)

func TestCreateWorkArea(t *testing.T) {

	CreateWorkArea("/tmp")
	CleanupWorkArea()
	CreateWorkArea("/Users/rajasrinivasan/tmp")
	CleanupWorkArea()

}
