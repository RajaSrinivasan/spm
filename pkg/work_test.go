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

func TestCreateUniqueId(t *testing.T) {
	CreateUniqueId()
}

func TestSetUniqueId(t *testing.T) {
	SetUniqueId("abc")
	SetUniqueId("69a8f2e4-d0c9-44bc-9638-cb5f5138d927")
}
