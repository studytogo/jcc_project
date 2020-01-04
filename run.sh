#!/bin/bash

createGo(){
  sleep 3
  echo "package routers" > routers/a.go
}

createGo & bee run -gendoc=true
