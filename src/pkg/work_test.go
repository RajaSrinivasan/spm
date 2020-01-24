package pkg

import (
	"log"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func TestCreateWorkArea(t *testing.T) {

	CreateWorkArea("/tmp")
	CleanupWorkArea()
	home, err := homedir.Dir()
	if err == nil {
		hometemp := filepath.Join(home, "tmp")
		CreateWorkArea(hometemp)
	} else {
		log.Printf("%s", err)
	}

}

func TestCreateUniqueId(t *testing.T) {
	CreateUniqueId()
}

func TestSetUniqueId(t *testing.T) {
	SetUniqueId("abc")
	SetUniqueId("69a8f2e4-d0c9-44bc-9638-cb5f5138d927")
}
