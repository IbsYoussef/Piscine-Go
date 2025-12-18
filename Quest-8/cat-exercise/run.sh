#!/usr/bin/env bash
set -euo pipefail

trap 'rm -f student/cat/cat' EXIT

echo "Running cat..."
echo ""

cd student/cat
go build
cd ../..

echo "=== Test 1: Single file ==="
student/cat/cat test-files/quest8.txt
echo ""

echo "=== Test 2: Multiple files ==="
student/cat/cat test-files/quest8.txt test-files/quest8T.txt
echo ""

echo "=== Test 3: File with error ==="
student/cat/cat test-files/quest8.txt test-files/nonexistent.txt || echo "Exit status: $?"
echo ""

echo "=== Test 4: Piped input (stdin mode) ==="
echo "Hello from pipe!" | student/cat/cat
echo ""

echo "=== Test 5: File piped to cat ==="
cat test-files/quest8.txt | student/cat/cat
echo ""