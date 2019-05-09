# gitpuller is a simple Go program to keep your git repositories up to date.  

Copyright 2019 by Shawn Rupp

1. [Installation](#installation)  
2. [Config](#config)
3. [Usage](#usage)  
4. [Credentials](#credentials)  


## Installation  

#### Installing Go and Setting Paths  
[Go version 1.11 or higher](https://golang.org/doc/install)  
Go requires a GOPATH environment variable to set to install packages, an Hyena requires the GOBIN variable to be set as well.  
Follow the directions [here](https://github.com/golang/go/wiki/SettingGOPATH) to set your GOPATH. Before you close your .bashrc or 
similar file, add the following lines after you deifne you GOPATH:  

	export GOBIN=$GOPATH/bin  
	export PATH=$PATH:$GOBIN  

### Download  
Download the repository into correct Go src directory (required for package imports):  

	cd $GOPATH/src  
	mkdir -p github.com/icwells  
	cd github.com/icwells  
	git clone https://github.com/icwells/gitpuller.git  
	cd gitpuller  

	./install.sh  

gitpuller will be installed to your GOBIN and will be added to your PATH.  

## Config  
Make a file in gitpuller directory named "config.txt". Simply list the paths to the git repostories you want to manage with gitpuller (one per line).

For simplicity's sake, the config file must remain in the gitpuller repository in your go/src directory so the binary can locate it.  

## Usage  

	gitpuller pull  
	gitpuller status  

That's it.  

## Credentials  
See the following for [storing](https://git-scm.com/docs/git-credential-store) or [caching](https://git-scm.com/docs/git-credential-cache) git credentials to avoid re-entering them.  
