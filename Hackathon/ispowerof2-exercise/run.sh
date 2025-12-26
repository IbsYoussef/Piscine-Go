#!/usr/bin/env bash
set -euo pipefail

trap 'rm -f student/ispowerof2/ispowerof2' EXIT

echo "Running ispowerof2..."
echo ""

# Build student program
cd student/ispowerof2
go build
cd ../..

echo "=== Test 1: 2 (power of 2) ==="
student/ispowerof2/ispowerof2 2

echo ""
echo "=== Test 2: 64 (power of 2) ==="
student/ispowerof2/ispowerof2 64

echo ""
echo "=== Test 3: 513 (NOT power of 2) ==="
student/ispowerof2/ispowerof2 513

echo ""
echo "=== Test 4: 1 (power of 2) ==="
student/ispowerof2/ispowerof2 1

echo ""
echo "=== Test 5: 16 (power of 2) ==="
student/ispowerof2/ispowerof2 16

echo ""
echo "=== Test 6: No arguments ==="
student/ispowerof2/ispowerof2 || echo "(no output - correct)"

echo ""
echo "=== Test 7: Multiple arguments ==="
student/ispowerof2/ispowerof2 64 1024 || echo "(no output - correct)"

echo ""