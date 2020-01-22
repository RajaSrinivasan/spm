package impl

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"../cfg"
	"../pkg"
)

var ShowOption bool

func verifyContents(filecfg *cfg.Config) error {

	pubkeyfilename := filepath.Join(pkg.ContentsDir, pkg.DefaultPublicKeyFileName)

	for _, c := range filecfg.Contents {
		bn := filepath.Base(c.From)
		cn := filepath.Join(pkg.ContentsDir, bn)
		content, err := os.Open(cn)
		if err != nil {
			log.Printf("%s\n", err)
			return err
		}
		content.Close()
		err = pkg.AuthenticateFile(cn, cn+".sig", pubkeyfilename)
		if err != nil {
			return err
		}
	}
	return nil
}

func showContents(filecfg *cfg.Config) {
	log.Printf("\nContents of the Package : %s\nID: %s\nVerion : %s\nCreated on the host : %s\nOn : %s\n",
		filecfg.Package.Name, filecfg.Package.ID, filecfg.Package.Version, filecfg.Package.Hostname, filecfg.Package.Created)
	log.Printf("Preinstall steps\n")
	for sn, step := range filecfg.Preinstall {
		log.Printf("%3d : %s\n", sn, step)
	}
	log.Printf("Contents of Package\n")
	filepath.Walk(pkg.ContentsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		fmt.Printf("%s\n", info.Name())
		return nil
	})

	log.Printf("Post steps\n")
	for sn, step := range filecfg.Postinstall {
		log.Printf("%3d : %s\n", sn, step)
	}
}

func Install(pkgfile string) {
	fullname, _ := filepath.Abs(pkgfile)
	log.Printf("Installing package %s\n", fullname)
	pkg.CreateWorkArea(Workarea)
	if !KeepWorkArea {
		defer pkg.CleanupWorkArea()
	}
	pkgbasename := filepath.Base(pkgfile)
	spmname := filepath.Join(pkg.WorkDir, pkgbasename)
	pkg.Decrypt(PkgPassword, fullname, spmname)
	log.Printf("Decrypted %s to create %s\n", fullname, spmname)

	pkg.UnPackfiles(spmname, pkg.ContentsDir)

	cfgfilename := filepath.Join(pkg.ContentsDir, ManifestFileName)
	filecfg, err := cfg.LoadConfig(cfgfilename)
	if err != nil {
		return
	}

	err = verifyContents(filecfg)
	if err != nil {
		return
	}

	if ShowOption {
		showContents(filecfg)
		return
	}
}
