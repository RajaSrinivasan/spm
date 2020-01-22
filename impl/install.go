package impl

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

func execCommand(cmdstr string) error {
	cmdsplit := strings.Fields(cmdstr)
	cmd := exec.Command(cmdsplit[0], cmdsplit[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Executing %s\n", cmdstr)
		log.Fatal(err)
	}
	log.Printf("%s\n", output)
	return nil
}

func executeSteps(steps []string) error {
	for _, stepstr := range steps {
		err := execCommand(stepstr)
		if err != nil {
			return err
		}
	}
	return nil
}

func installFile(from, to string) error {
	cmd := fmt.Sprintf("cp -p %s %s", from, to)
	err := execCommand(cmd)
	return err
}

func installFiles(filecfg *cfg.Config) {
	for _, c := range filecfg.Contents {
		fname := filepath.Base(c.From)
		fromname := filepath.Join(pkg.ContentsDir, fname)
		err := installFile(fromname, c.To)
		if err != nil {
			log.Fatal(err)
		}
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
	log.Printf("Executing Preinstall steps\n")
	err = executeSteps(filecfg.Preinstall)
	if err != nil {
		log.Fatal(err)
	}
	installFiles(filecfg)
	log.Printf("Executing Postinstall steps\n")
	err = executeSteps(filecfg.Postinstall)
	if err != nil {
		log.Fatal(err)
	}
}
