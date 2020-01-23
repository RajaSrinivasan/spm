SOURCES=$(wildcard *.go)
EXEC=spm
BINARIES=../bin

all: $(SOURCES)
	go build -o $(BINARIES)/$(EXEC)

clean:
	$(RM) $(BINARIES)/$(EXEC)

dependencies:
	go get -u -v github.com/spf13/cobra
	go get -u -v github.com/mitchellh/go-homedir
	go get -u -v github.com/spf13/viper
	go get -u -v github.com/google/uuid
	go get -u -v golang.org/x/crypto/ssh
	go get -u -v gopkg.in/yaml.v3
