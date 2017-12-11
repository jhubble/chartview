#!/bin/bash

if ! [ -x "$(command -v npm)" ]; then
  echo 'Error: npm is not installed. Please see https://nodejs.org for nodejs/npm downloads for your OS' >&2
  exit 1
fi

if ! [ -x "$(command -v go)" ]; then
  echo 'Error: Go is not installed. Please see https://goland.org for go language for your OS' >&2
  exit 1
fi

echo Installing prerequisities
npm install -g brunch
npm install

echo building
brunch build

echo Starting server. Go to http://localhost:8081 to see the app
go run simpleServer.go
