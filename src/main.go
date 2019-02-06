// Performs git pull for all given repositories

package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	app    = kingpin.New("gitpuller", "Simple status manager for git repositories.")
	pull   = kingpin.Command("pull", "Performs git pull for all given repositories.")
	status = kingpin.Command("status", "Checks status of all repositories.")
)

func (p *puller) pullRepos() {
	// Pulls all given repositories
	for _, r := range p.repos {
		_ = os.Chdir(r)
		_, err := exec.Command("git pull").Output()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func filterStatus(status string) bool {
	// Returns true if status should be printed
	ret := true
	msg1 := "Your branch is up to date with "
	msg2 := "nothing to commit, working tree clean"
	if strings.Contains(status, msg1) == true && strings.Contains(status, msg2) == true {
		ret = false
	}
	return ret
}

func (p *puller) getStatus() {
	// Prints statuses for any repos which need to be committed/pushed
	for _, r := range p.repos {
		_ = os.Chdir(r)
		out, err := exec.Command("git status").Output()
		if err != nil {
			log.Fatal(err)
		}
		stdout := string(out)
		if filterStatus(stdout) == true {
			fmt.Print("\n", stdout, "\n")
		}
	}
}

func main() {
	puller := newPuller()
	switch kingpin.Parse() {
	case pull.FullCommand():
		puller.pullRepos()
	case status.FullCommand():
		puller.getStatus()
	}
}
