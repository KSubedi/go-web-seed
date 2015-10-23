#!/bin/bash

CompileDaemon -build="go build -o app" -command="./app" -exclude-dir=".git" -exclude-dir="static" -exclude-dir="node_modules" -exclude-dir="bower_components"
