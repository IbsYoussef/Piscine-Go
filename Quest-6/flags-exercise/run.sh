#!/usr/bin/env bash
set -euo pipefail

echo "Running flags..."
echo ""

# Navigate to student folder and build
cd student/flags
go build

echo "Test 1: Insert and order"
./flags --insert=4321 --order asdad
echo ""

echo "Test 2: Insert only"
./flags --insert=4321 asdad
echo ""

echo "Test 3: String only"
./flags asdad
echo ""

echo "Test 4: Order only"
./flags --order 43a21
echo ""

echo "Test 5: No arguments (help)"
./flags
echo ""

# Clean up
rm -f flags
cd ../..