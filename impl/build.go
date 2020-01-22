package impl

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"../cfg"
	"../pkg"
)

var Workarea = "/tmp"
var KeepWorkArea bool
var PkgPassword string

const ManifestFileName = "Packagefile"

func copyFile(from, to string) error {
	log.Printf("Copying file %s to %s\n", from, to)
	fullname, _ := filepath.Abs(from)
	infile, err := os.Open(from)
	if err != nil {
		log.Fatalf("Failed to open %s\n", fullname)
		return err
	}
	defer infile.Close()
	outfilename := filepath.Join(to, filepath.Base(from))
	outfile, err := os.Create(outfilename)
	if err != nil {
		log.Printf("Failed to create %s\n", outfilename)
		return err
	}
	defer outfile.Close()
	io.Copy(outfile, infile)
	return nil
}

func assembleFiles(workarea string, pkgconfig *cfg.Config) error {
	target := workarea
	nfiles := len(pkgconfig.Contents)
	for no := 0; no < nfiles; no++ {
		f := pkgconfig.Contents[no]
		copyFile(f.From, target)
	}
	return nil
}
func makePackageName(cfgname string) string {
	bn := filepath.Base(cfgname)
	cfgtype := filepath.Ext(cfgname)
	pkgname := bn[:len(bn)-len(cfgtype)] + ".spm"
	return pkgname
}

func Build(cfgfile string, outfile string) {
	log.Printf("Building package for configuration file %s\n", cfgfile)
	pkg.CreateWorkArea(Workarea)
	if !KeepWorkArea {
		defer pkg.CleanupWorkArea()
	}

	pkgconfig, err := cfg.LoadConfig(cfgfile)
	if err != nil {
		return
	}

	assembleFiles(pkg.ContentsDir, pkgconfig)
	var pkgcontent = cfg.Content{From: ManifestFileName, To: ManifestFileName}
	pkgconfig.Contents = append(pkgconfig.Contents, pkgcontent)

	pvtkeyfile := filepath.Join(pkg.WorkDir, pkg.DefaultPrivateKeyFileName)
	pubkeyfile := filepath.Join(pkg.ContentsDir, pkg.DefaultPublicKeyFileName)
	pkg.GenerateKeys(pvtkeyfile, pubkeyfile)
	log.Printf("Created keypair %s and %s\n", pvtkeyfile, pubkeyfile)

	contfiles, _ := ioutil.ReadDir(pkg.ContentsDir)
	var contnames []string
	for _, fi := range contfiles {
		cfn, _ := filepath.Abs(filepath.Join(pkg.ContentsDir, fi.Name()))
		log.Printf("Content file %s\n", cfn)
		contnames = append(contnames, cfn)
	}
	log.Printf("Files: %v\n", contnames)
	pkg.SignFiles(contnames, pvtkeyfile)

	pkgfilename := filepath.Join(pkg.ContentsDir, ManifestFileName)
	pkgconfig.SaveManifest(pkgfilename)
	log.Printf("Saved manifest %s\n", pkgfilename)

	sigfilename := pkgfilename + ".sig"
	pkg.SignFile(pkgfilename, sigfilename, pvtkeyfile)
	log.Printf("Signed the Package file. Generated %s\n", sigfilename)

	spmbasename := makePackageName(cfgfile)
	spmname := filepath.Join(pkg.WorkDir, spmbasename)
	pkg.Packfiles(spmname, pkg.ContentsDir)
	log.Printf("Created %s\n", spmname)

	encspmname := outfile
	fi, err := os.Stat(outfile)
	if err == nil {
		if fi.IsDir() {
			encspmname = filepath.Join(outfile, spmbasename)
		}
	}
	if len(PkgPassword) < 1 {
		log.Printf("No password provided for finalization. Cannot create %s\n", encspmname)
		return
	}
	pkg.Encrypt(PkgPassword, spmname, encspmname)
	log.Printf("Created %s\n", encspmname)

}
