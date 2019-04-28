#!/usr/bin/env bash


HOST := 172.17.8.101
PORT := 30002
VERSION := v0.1.0
GOPATH := $(HOME)/go


.PHONY: all
all : docker.push


install :
	go clean
	go install
	mkdir -p ./drivers
	cp $(HOME)/go/bin/lvm ./drivers


build.img : install
	sudo docker build -t lvm:$(VERSION) ./


docker.push : build.img
	sudo docker tag lvm:$(VERSION) $(HOST):$(PORT)/flexmnt/lvm:$(VERSION)
	sudo docker push $(HOST):$(PORT)/flexmnt/lvm:$(VERSION)

clean :
	rm -rf ./drivers
	rm -f $(HOME)/go//bin/lvm

