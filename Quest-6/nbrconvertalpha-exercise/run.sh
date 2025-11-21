#!/usr/bin/env bash
set -euo pipefail

echo "Running nbrconvertalpha..."
echo ""

# Navigate to student folder and build
cd student/nbrconvertalpha
go build

# Test cases
echo "Test 1: No args"
./nbrconvertalpha

echo "Test 2: hello"
./nbrconvertalpha 8 5 12 12 15

echo "Test 3: legen dary"
./nbrconvertalpha 12 5 7 5 14 56 4 1 18 25

echo "Test 4: invalid inputs"
./nbrconvertalpha 32 86 h

echo "Test 5: uppercase"
./nbrconvertalpha --upper 8 5 25

# Clean up
rm -f nbrconvertalpha
cd ../..