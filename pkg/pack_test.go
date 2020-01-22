package pkg

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func copyFile(fn string) {
	infile, _ := os.Open(fn)
	defer infile.Close()
	outfilename := filepath.Join(ContentsDir, filepath.Base(fn))
	outfile, _ := os.Create(outfilename)
	defer outfile.Close()
	io.Copy(outfile, infile)
}

func createFile(fn string) {
	log.Printf("Creating %s\n", fn)
	f, _ := os.Create(fn)
	fmt.Fprintf(f, "This is a line of file %s", fn)
	f.Close()
}

func TestPackage(t *testing.T) {
	folder := "/Users/rajasrinivasan/tmp"
	CreateWorkArea(folder)
	createFile(filepath.Join(workArea, "package", "a.sig"))
	createFile(filepath.Join(workArea, "package", "b.sig"))
	createFile(filepath.Join(workArea, "package", "c.sig"))

	createFile(filepath.Join(workArea, "artifacts", "a.data"))
	createFile(filepath.Join(workArea, "artifacts", "b.data"))
	createFile(filepath.Join(workArea, "artifacts", "c.data"))

	Package("bn")
}

func TestPackfiles(t *testing.T) {

	CreateWorkArea("/tmp")

	createFile(filepath.Join(ContentsDir, "a.sig"))
	createFile(filepath.Join(ContentsDir, "b.sig"))
	createFile(filepath.Join(ContentsDir, "c.sig"))

	Packfiles(filepath.Join(WorkDir, "pack.tgz"), ContentsDir)
}

func TestPackfilesBig(t *testing.T) {
	CreateWorkArea("/tmp")
	copyFile("/Volumes/Dev1/Ref/Books/usb_ddk.pdf")
	copyFile("/Volumes/Dev1/Ref/Books/acsac.pdf")
	copyFile("/Volumes/Dev1/Ref/Books/manual.pdf")
	Packfiles("/Users/rajasrinivasan/Prj/work/bigpack.tgz", ContentsDir)
	CleanupWorkArea()
}

func TestUnPackfiles(t *testing.T) {
	CreateWorkArea("/tmp")
	UnPackfiles("/Users/rajasrinivasan/Prj/work/bigpack.tgz", ContentsDir)
}
