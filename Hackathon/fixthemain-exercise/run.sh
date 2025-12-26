#!/usr/bin/env bash
set -euo pipefail

trap 'rm -f student/fixthemain/fixthemain' EXIT

echo "Running fixthemain..."
echo ""

# Build and run student program
cd student/fixthemain
go build
./fixthemain
cd ../..

echo ""