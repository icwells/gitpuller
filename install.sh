#!/bin/bash

##############################################################################
# This script will install gitpuller to the GOBIN directory.
# 
# Required programs:	Go 1.11+
##############################################################################

IO="github.com/icwells/go-tools/iotools"
KP="gopkg.in/alecthomas/kingpin.v2"
MAIN="gitpuller"

# Get install location
SYS=$(ls $GOPATH/pkg | head -1)
PDIR=$GOPATH/pkg/$SYS

echo ""
echo "Preparing gitpuller package..."
echo "GOPATH identified as $GOPATH"
echo ""

# Get dependencies
for I in $IO $KP $PR ; do
	if [ ! -e "$PDIR/$I.a" ]; then
		echo "Installing $I..."
		go get -u $I
		echo ""
	fi
done

go build -o $GOBIN/$MAIN src/*.go

echo "Finished"
echo ""
