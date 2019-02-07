#!/bin/bash

##############################################################################
# This script will install scripts for the gitpuller package.
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
for I in $IO $KP ; do
	if [ ! -e "$PDIR/$I.a" ]; then
		echo "Installing $I..."
		go get -u $I
		echo ""
	fi
done

# compOncDB 
echo "Building main..."
go build -o bin/$MAIN src/*.go

echo "Finished"
echo ""
