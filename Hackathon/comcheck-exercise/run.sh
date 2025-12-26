#!/usr/bin/env bash
set -euo pipefail

trap 'rm -f student/comcheck/comcheck' EXIT

echo "Running comcheck..."
echo ""

# Build student program
cd student/comcheck
go build
cd ../..

echo "=== Test 1: galaxy match ==="
student/comcheck/comcheck "I" "Will" "Enter" "the" "galaxy"

echo ""
echo "=== Test 2: galaxy 01 match ==="
student/comcheck/comcheck "galaxy 01" "do" "you" "hear" "me"

echo ""
echo "=== Test 3: 01 match ==="
student/comcheck/comcheck "test" "01" "hello"

echo ""
echo "=== Test 4: no match ==="
student/comcheck/comcheck "hello" "world" || echo "(no output - correct)"

echo ""
echo "=== Test 5: multiple matches ==="
student/comcheck/comcheck "galaxy" "01"

echo ""