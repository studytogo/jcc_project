#!/bin/bash

createGo(){
  sleep 1
  echo "package routers" > routers/a.go
}

createGo & bee run