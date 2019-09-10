# This is how we want to name the binary output
OUTPUT=main

# These are the values we want to pass for Version and BuildTime
GITTAG=`git rev-parse HEAD`
BUILD_TIME=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X github.com/sereiner/parrot/parrot.VERSION=${GITTAG} -X github.com/sereiner/parrot/parrot.BuildTime=${BUILD_TIME}"

local:
	go build ${LDFLAGS} -tags local

prod:
	go build ${LDFLAGS} -tags prod
