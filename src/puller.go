// Defines puller struct and methods

package main

import (
	"fmt"
	"github.com/Songmu/prompter"
	"github.com/icwells/go-tools/iotools"
	"os"
	"path"
	"strings"
)

type puller struct {
	config string
	names  map[string][]string
	repos  []string
}

func (p *puller) setRepos(credentials bool) {
	// Sets input from config file
	r := false
	f := iotools.OpenFile(p.config)
	defer f.Close()
	scanner := iotools.GetScanner(f)
	for scanner.Scan() {
		s := strings.TrimSpace(string(scanner.Text()))
		if r == true && iotools.Exists(s) == true {
			// Only store directories which exist
			p.repos = append(p.repos, s)
		} else {
			if strings.Contains(s, "TargetRepositories") == true {
				r = true
			} else if credentials == true && strings.Contains(s, "#") == false && len(s) > 0 {
				// Store username and password for private repsotories
				spl := strings.Split(s, "\t")
				pw := prompter.Password(fmt.Sprintf("\tEnter password for %s", spl[0]))
				p.names[spl[0]] = []string{spl[1], pw}
			}
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
	p.names = make(map[string][]string)
	p.setRepos(credentials)
	return p
}
