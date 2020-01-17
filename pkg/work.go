package pkg

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	uuid "github.com/google/uuid"
)

var workArea string
var uniqueId uuid.UUID

func init() {

}

func CreateWorkArea(root string) string {
	dirname, err := ioutil.TempDir(root, "spm")
	if err != nil {
		log.Fatal(err)
	}
	workArea = dirname
	log.Printf("Workarea created %s\n", dirname)
	os.Mkdir(filepath.Join(workArea, "package"), os.ModePerm)
	os.Mkdir(filepath.Join(workArea, "artifacts"), os.ModePerm)
	log.Printf("Created dir %s and %s\n", filepath.Join(workArea, "package"), filepath.Join(workArea, "artifacts"))
	return workArea
}

func CleanupWorkArea() {
	os.RemoveAll(workArea)
	log.Printf("Removed %s\n", workArea)
}

func CreateUniqueId() uuid.UUID {
	uniqueId = uuid.New()
	log.Printf("Unique Id created %s\n", uniqueId.String())
	return uniqueId
}

func SetUniqueId(id string) {
	uniqueid, err := uuid.Parse(id)
	if err == nil {
		log.Printf("Set the Unique id for package from: %s\n", id)
		uniqueId = uniqueid
	} else {
		log.Printf("Error: %s\n%s\n", err, id)
	}
}
