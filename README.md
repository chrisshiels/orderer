# orderer - Dependency Orderer

Utility to order files according to their dependencies.


## Build

    host$ export GOPATH=~/workspace
    host$ mkdir -p $GOPATH
    host$ cd $GOPATH
    host$ go get -v github.com/chrisshiels/orderer
    host$ go test -v github.com/chrisshiels/orderer


## Usage

    host$ ./bin/orderer
    Usage:  orderer [ -v ] --filetype filetype file ...
      -filetype string
    	    File type:  commented, rpm
      -v	Verbose


## Ordering RPM Spec Files

Here orderer processes RPM spec files, parses Name: and BuildRequires:
RPM tags, and then outputs the RPM spec files ordered according to their
dependencies.

    host$ ./bin/orderer --filetype rpm \
        ./src/github.com/chrisshiels/orderer/testdata/rpm/*.spec
    ./src/github.com/chrisshiels/orderer/testdata/rpm/grip.spec
    ./src/github.com/chrisshiels/orderer/testdata/rpm/id3lib.spec
    ./src/github.com/chrisshiels/orderer/testdata/rpm/lame.spec
    ./src/github.com/chrisshiels/orderer/testdata/rpm/mpg123.spec
    ./src/github.com/chrisshiels/orderer/testdata/rpm/ffmpeg.spec
    ./src/github.com/chrisshiels/orderer/testdata/rpm/vlc.spec
    ./src/github.com/chrisshiels/orderer/testdata/rpm/audacious.spec
    ./src/github.com/chrisshiels/orderer/testdata/rpm/audacious-plugins.spec


## Usage - RPM Spec Files

Here orderer processes commented files, parses comments of the form
'orderer: name dependencies: [ name ... ]', and then outputs the commented
files ordered according to their dependencies.

    host$ ./bin/orderer --filetype commented \
        ./src/github.com/chrisshiels/orderer/testdata/commented/*
    ./src/github.com/chrisshiels/orderer/testdata/commented/e
    ./src/github.com/chrisshiels/orderer/testdata/commented/d
    ./src/github.com/chrisshiels/orderer/testdata/commented/c
    ./src/github.com/chrisshiels/orderer/testdata/commented/b
    ./src/github.com/chrisshiels/orderer/testdata/commented/a
