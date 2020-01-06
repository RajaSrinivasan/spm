package pkg

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func createFile(fn string) {
	log.Printf("Creating %s\n", fn)
	f, _ := os.Create(fn)
	fmt.Fprintf(f, "This is a line of file %s", fn)
	f.Close()
}
func TestPackage(t *testing.T) {
	folder := "/Users/rajasrinivasan/tmp"
	CreateWorkArea(folder)
	createFile(filepath.Join(workArea, "package", "a"))
	createFile(filepath.Join(workArea, "package", "b"))
	createFile(filepath.Join(workArea, "package", "c"))
	Package("bn")
}
