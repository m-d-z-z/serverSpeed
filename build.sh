#!/bin/bash
go build -ldflags "-s -w"
upx serverSpeed -9