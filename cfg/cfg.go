package cfg

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RajaSrinivasan/spm/pkg"
	uuid "github.com/google/uuid"
	yaml "gopkg.in/yaml.v3"
)

type Content struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}
type Package struct {
	Name     string    `yaml:"name"`
	ID       uuid.UUID `yaml:"id"`
	Version  string    `yaml:"version"`
	Hostname string    `yaml:"hostname"`
	Created  time.Time `yaml:"created"`
}

type Config struct {
	Package     Package   `yaml:"package"`
	Contents    []Content `yaml:"contents,flow"`
	Preinstall  []string  `yaml:"preinstall"`
	Postinstall []string  `yaml:"postinstall"`
}

func LoadConfig(cfgfile string) (*Config, error) {

	inpfile, err := os.Open(cfgfile)
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}
	cfg := new(Config)
	decoder := yaml.NewDecoder(inpfile)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}
	fmt.Printf("Loaded package File: %s Name : %s\n", cfgfile, cfg.Package.Name)

	if len(cfg.Preinstall) == 0 {
		log.Printf("No Preinstall steps to execute\n")
	}

	if len(cfg.Postinstall) == 0 {
		log.Printf("No Postinstall steps to execute\n")
	}

	return cfg, nil
}

func (cfg *Config) SaveManifest(manifestfile string) {
	cfg.Package.Created = time.Now()
	cfg.Package.Hostname, _ = os.Hostname()
	cfg.Package.ID = pkg.CreateUniqueId()
	ofile, _ := os.Create(manifestfile)
	encoder := yaml.NewEncoder(ofile)
	encoder.Encode(cfg)
	encoder.Close()
}
