#!/usr/bin/env bash
set -euo pipefail

echo "Running boolean..."
echo ""

# Navigate to student folder and build
cd student/boolean
go build

echo "Test 1: Even arguments"
./boolean "not" "odd"

echo "Test 2: Odd arguments"
./boolean "not even"

echo "Test 3: No arguments"
./boolean

echo "Test 4: Three arguments"
./boolean "one" "two" "three"

# Clean up
rm -f boolean
cd ../..