#!/usr/bin/env bash
set -euo pipefail

trap 'rm -f student/pilot/pilot' EXIT

echo "Running pilot..."
echo ""

# Build and run student program
cd student/pilot
go build
./pilot
cd ../..

echo ""