package pkg

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
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

	home, _ := homedir.Dir()
	hometemp := filepath.Join(home, "tmp")
	_, err := os.Stat(hometemp)
	if os.IsNotExist(err) {
		os.Mkdir(hometemp, os.ModePerm)
	}

	CreateWorkArea(hometemp)

	createFile(filepath.Join(ContentsDir, "a.sig"))
	createFile(filepath.Join(ContentsDir, "b.sig"))
	createFile(filepath.Join(ContentsDir, "c.sig"))

	createFile(filepath.Join(ArtifactsDir, "a.data"))
	createFile(filepath.Join(ArtifactsDir, "b.data"))
	createFile(filepath.Join(ArtifactsDir, "c.data"))

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
	copyFile("../systest/usb_ddk.pdf")
	copyFile("../systest//acsac.pdf")
	Packfiles(filepath.Join(WorkDir, "bigpack.tgz"), ContentsDir)

}

func TestUnPackfiles(t *testing.T) {
	TestPackfilesBig(t)
	UnPackfiles(filepath.Join(WorkDir, "bigpack.tgz"), "/tmp")
}
