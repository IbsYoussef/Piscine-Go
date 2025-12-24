#!/usr/bin/env bash
set -euo pipefail

trap 'rm -f student/doop/doop' EXIT

echo "Running doop..."
echo ""

# Build student program
cd student/doop
go build
cd ../..

echo "=== Test 1: Addition (1 + 1) ==="
student/doop/doop 1 + 1

echo ""
echo "=== Test 2: Subtraction (10 - 3) ==="
student/doop/doop 10 - 3

echo ""
echo "=== Test 3: Multiplication (5 * 4) ==="
student/doop/doop 5 "*" 4

echo ""
echo "=== Test 4: Division (20 / 4) ==="
student/doop/doop 20 / 4

echo ""
echo "=== Test 5: Modulo (17 % 5) ==="
student/doop/doop 17 % 5

echo ""
echo "=== Test 6: Division by zero (1 / 0) ==="
student/doop/doop 1 / 0

echo ""
echo "=== Test 7: Modulo by zero (1 % 0) ==="
student/doop/doop 1 % 0

echo ""
echo "=== Test 8: Invalid operator (1 p 1) ==="
student/doop/doop 1 p 1 || echo "(no output - correct)"

echo ""
echo "=== Test 9: Overflow (MaxInt + 1) ==="
student/doop/doop 9223372036854775807 + 1 || echo "(no output - correct)"

echo ""
echo "=== Test 10: No arguments ==="
student/doop/doop || echo "(no output - correct)"

echo ""