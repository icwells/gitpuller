#!/bin/bash

##############################################################################
# This script will install gitpuller to the GOBIN directory.
# 
# Required programs:	Go 1.11+
##############################################################################

echo ""
echo "Installing gitpuller package..."
echo "GOBIN identified as $GOBIN"

go install -i src/*.go

echo "Finished"
echo ""
