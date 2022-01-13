#!/bin/bash

go mod download
go build -o kaznet-status
sudo chmod +x kaznet-status
sudo ./kaznet-status