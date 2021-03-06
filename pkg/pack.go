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
	err := Packfiles(filepath.Join(WorkDir, bn+".spm"), ContentsDir)
	if err != nil {
		return err
	}
	err = Packfiles(filepath.Join(WorkDir, bn+"_art.spm"), ArtifactsDir)
	return err
}

func UnPackfiles(fn, dir string) error {

	pkgfilename := fn
	pkgfile, err := os.Open(pkgfilename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer pkgfile.Close()

	gzreader, err := gzip.NewReader(pkgfile)
	if err != nil {
		log.Fatal(err)
	}
	defer gzreader.Close()

	tarReader := tar.NewReader(gzreader)

	for {
		hdr, err := tarReader.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Extracting %s\n", hdr.Name)
		outname := filepath.Join(dir, hdr.Name)
		outfile, err := os.Create(outname)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := io.Copy(outfile, tarReader); err != nil {
			log.Fatal(err)
		}
		outfile.Close()
	}

	return nil
}
