// Defines puller struct and methods

package gitpuller

import (
	"fmt"
	"github.com/icwells/go-tools/iotools"
	"os"
	"path"
	"strings"
)

type puller struct {
	config string
	repos  []string
}

func (p *puller) setRepos(credentials bool) {
	// Sets input from config file
	f := iotools.OpenFile(p.config)
	defer f.Close()
	scanner := iotools.GetScanner(f)
	for scanner.Scan() {
		s := strings.TrimSpace(string(scanner.Text()))
		if iotools.Exists(s) == true {
			// Only store directories which exist
			p.repos = append(p.repos, s)
		}
	}
}

func (p *puller) setConfig() {
	// Locates config file
	ex, _ := os.Executable()
	pt, _ := path.Split(ex)
	p.config = path.Join(pt, "config.txt")
	if iotools.Exists(p.config) == false {
		fmt.Print("\n\t[Error] Cannot find config file. Exiting.\n")
		os.Exit(1)
	}
}

func newPuller(credentials bool) puller {
	// Initializes new puller
	var p puller
	p.setConfig()
	p.setRepos(credentials)
	return p
}
