# gitpuller is a simple Go program to keep your git repositories upt to date.  
This script is still under developement.  

Copyright 2019 by Shawn Rupp

1. [Installation](#installation)  
2. [Config](#config)
3. [Usage](#usage)  
4. [Credentials](#credentials)  


## Installation  
Just run install.sh. It will download any necesarry dependencies.  

	./install.sh  

## Config  
The config file contains a list of usernames for private repositories and a list of target repsories on the local machine.  

Make a file in the bin directory named "config.txt". Simply list the paths to the 
git repostories you want to manage with gitpuller (one per line).

## Usage  
./gitpuller pull  
./gitpuller status  

That's it.  

## Credentials  
See the following for [storing](https://git-scm.com/docs/git-credential-store) or [caching](https://git-scm.com/docs/git-credential-cache) git credentials to avoid re-entering them.  
