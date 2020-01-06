package pkg

import (
	"io/ioutil"
	"log"
	"os"
)

var workArea string

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
