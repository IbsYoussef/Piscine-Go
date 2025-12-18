#!/usr/bin/env bash
set -euo pipefail

# Ensure cleanup happens no matter what
trap 'rm -f student/ztail/ztail' EXIT

echo "Running ztail..."
echo ""

# Build student program
cd student/ztail
go build
cd ../..

echo "=== Test 1: Single file (-c 4) ==="
student/ztail/ztail -c 4 test-files/file1.txt
echo ""

echo "=== Test 2: Multiple files (-c 4) ==="
student/ztail/ztail -c 4 test-files/file1.txt test-files/file2.txt
echo ""

echo "=== Test 3: Error handling ==="
student/ztail/ztail -c 4 test-files/file1.txt test-files/nonexistent.txt test-files/file2.txt || echo "Exit status: $?"
echo ""