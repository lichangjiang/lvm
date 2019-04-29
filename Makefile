#!/usr/bin/env bash
VERSION := v1.0.0
GOPATH := $(HOME)/go


.PHONY: all
all : docker.push


install :
	go clean
	go install
	mkdir -p ./drivers
	cp $(HOME)/go/bin/lvm ./drivers


build.img : install
	sudo docker build --no-cache -t riverlcj/lvm:$(VERSION) ./


docker.push : build.img
#	sudo docker tag lvm:$(VERSION) riverlcj/lvm:$(VERSION)
	sudo docker push riverlcj/lvm:$(VERSION)

clean :
	rm -rf ./drivers
	rm -f $(HOME)/go//bin/lvm

