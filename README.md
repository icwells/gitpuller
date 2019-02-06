# gitpuller is a simple Go program to keep your git repositories upt to date.  
This script is still under developement.  

Copyright 2019 by Shawn Rupp

1. [Installation](#installation)  
2. [Config](#config)
3. [Usage](#usage)  


## Installation  
Just run install.sh. It will download any necesarry dependencies.  

	./install.sh  

## Config  
The config file contains a list of usernames for private repositories and a list of target repsories on the local machine.  

Make a copy of the example_config.txt file in the bin directory and remove the "example_". For any private repostories which require 
a password to pull, enter the name of the website under the first header (i.e. "gitlab.com") followed by a tab and your username for 
that site. When the config file is read, your password for the site will be requested once. Leave this section empty if you don't have 
any private repositories. After the "TargetRepositories" line, simply list the paths to the git repostories you want to manage with 
gitpuller (one per line).

## Usage  
./gitpuller pull  
./gitpuller status  

That's it.  
