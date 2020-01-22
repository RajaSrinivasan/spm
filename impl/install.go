package impl

import (
	"log"
	"os"
	"path/filepath"

	"../cfg"
	"../pkg"
)

func verifyContents() (*cfg.Config, error) {
	cfgfilename := filepath.Join(pkg.ContentsDir, ManifestFileName)
	filecfg, err := cfg.LoadConfig(cfgfilename)
	if err != nil {
		return nil, err
	}
	pubkeyfilename := filepath.Join(pkg.ContentsDir, pkg.DefaultPublicKeyFileName)

	for _, c := range filecfg.Contents {
		bn := filepath.Base(c.From)
		cn := filepath.Join(pkg.ContentsDir, bn)
		content, err := os.Open(cn)
		if err != nil {
			log.Printf("%s\n", err)
			return nil, err
		}
		content.Close()
		err = pkg.AuthenticateFile(cn, cn+".sig", pubkeyfilename)
		if err != nil {
			return nil, err
		}
	}
	return filecfg, nil
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

	_, err := verifyContents()
	if err != nil {
		return
	}

}
