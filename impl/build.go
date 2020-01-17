package impl

import (
	"log"
	"../pkg"
)

func Build(cfgfile string) {
	log.Printf("Building package for configuration file %s\n", cfgfile)
	pkg.CreateWorkArea()
	pkg.CreateUniqueId()
}
