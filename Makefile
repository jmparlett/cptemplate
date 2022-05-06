CDIR = templates/c/*
GODIR = templates/go/*
PYDIR = templates/python/*
CPPDIR = templates/cpp/*
TDIR = templates/*

install: build move

build: cptemplate.go $(CDIR) $(GODIR) $(PYDIR) $(CPPDIR) templates/*
	go build cptemplate.go

move: 
	sudo mv cptemplate /usr/local/bin/cptemplate

