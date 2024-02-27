#!/bin/bash
cp go.mod go.sum backend/cmd/server
cd backend/cmd/server
go build -o ../../bin/server