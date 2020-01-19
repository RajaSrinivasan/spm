package pkg

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func addFile(nm string, info os.FileInfo, tf *tar.Writer) error {
	if !info.Mode().IsRegular() {
		return nil
	}

	log.Printf("Adding %s Size %d\n", info.Name(), info.Size())
	hdr, _ := tar.FileInfoHeader(info, info.Name())
	tf.WriteHeader(hdr)

	f, _ := os.Open(nm)
	defer f.Close()

	io.Copy(tf, f)
	tf.Flush()

	return nil
}

func Packfiles(fn, dir string) error {

	pkgfilename := fn
	pkgfile, err := os.Create(pkgfilename)
	if err != nil {
		log.Printf("Error Creating %s\n%s", pkgfilename, err)
		return err
	}
	log.Printf("Created %s\n", pkgfilename)
	defer pkgfile.Close()

	gzwriter := gzip.NewWriter(pkgfile)
	defer gzwriter.Close()

	tarWriter := tar.NewWriter(gzwriter)
	defer tarWriter.Close()

	err = filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			errtemp := addFile(path, info, tarWriter)
			if errtemp != nil {
				log.Printf("Error adding %s\n%s\n", info.Name(), errtemp)
				return err
			}
			return nil
		})
	tarWriter.Flush()

	return nil
}

func Package(bn string) error {
	err := Packfiles(filepath.Join(workArea, bn+".spm"), filepath.Join(workArea, "package"))
	if err != nil {
		return err
	}
	err = Packfiles(filepath.Join(workArea, bn+"_art.spm"), filepath.Join(workArea, "artifacts"))
	return err
}
