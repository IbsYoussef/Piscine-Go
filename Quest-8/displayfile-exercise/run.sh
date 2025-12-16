#!/usr/bin/env bash
set -euo pipefail

echo "Running displayfile..."
echo ""

# Create test file
echo "Almost there!!" > quest8.txt

# Navigate to student folder and build
cd student/displayfile
go build

echo "Test 1: No arguments"
./displayfile

echo ""
echo "Test 2: Too many arguments"
./displayfile quest8.txt main.go

echo ""
echo "Test 3: Display file content"
./displayfile ../../quest8.txt

# Clean up
rm -f displayfile
cd ../..
rm -f quest8.txt