CDIR = templates/c/*
GODIR = templates/go/*
PYDIR = templates/python/*
CPPDIR = templates/cpp/*
TDIR = templates/*

cptemplate: cptemplate.go $(CDIR) $(GODIR) $(PYDIR) $(CPPDIR) templates/*
	go build cptemplate.go 

install:
	mv cptemplate /usr/bin/cptemplate
