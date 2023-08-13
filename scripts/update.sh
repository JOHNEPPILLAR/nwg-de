#!/bin/bash
clear

echo ""
echo "Updating Go packages..."
go get -u ./...
go mod tidy
echo ""

cp config/config /app/
