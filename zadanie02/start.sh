#!/bin/bash
set -e

APP_NAME="play-scala-api" 

ngrok config add-authtoken YOUR TOKEN

ngrok http http://localhost:9000

./app/bin/$APP_NAME

