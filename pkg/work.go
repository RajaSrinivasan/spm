package pkg

import (
	"io/ioutil"
	"log"
	"os"

	uuid "github.com/google/uuid"
)

var workArea string
var uniqueId uuid.UUID

func init() {

}

func CreateWorkArea(root string) {
	dirname, err := ioutil.TempDir(root, "spm")
	if err != nil {
		log.Fatal(err)
	}
	workArea = dirname
	log.Printf("Workarea created %s\n", dirname)
}

func CleanupWorkArea() {
	os.RemoveAll(workArea)
	log.Printf("Removed %s\n", workArea)
}

func CreateUniqueId() {
	uniqueId = uuid.New()
	log.Printf("Unique Id created %s\n", uniqueId.String())
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
