#!/usr/bin/env bash
set -euo pipefail

echo "Running rotatevowels..."
echo ""

# Navigate to student folder and build
cd student/rotatevowels
go build

echo "Test 1:"
./rotatevowels "Hello World"

echo "Test 2:"
./rotatevowels "HEllO World" "problem solved"

echo "Test 3:"
./rotatevowels "str" "shh" "psst"

echo "Test 4:"
./rotatevowels "happy thoughts" "good luck"

echo "Test 5:"
./rotatevowels "aEi" "Ou"

echo "Test 6: (no args)"
./rotatevowels

# Clean up
rm -f rotatevowels
cd ../..