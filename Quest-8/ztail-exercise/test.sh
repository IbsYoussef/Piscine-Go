#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

# Ensure cleanup happens no matter what
trap 'rm -f student/ztail/ztail solutions/ztail/ztail' EXIT

echo "Testing ztail..."
echo ""

# Build both programs
cd solutions/ztail
go build
cd ../..

cd student/ztail
go build
cd ../..

echo "Test 1: Single file"
student_out=$(student/ztail/ztail -c 4 test-files/file1.txt 2>&1)
solution_out=$(solutions/ztail/ztail -c 4 test-files/file1.txt 2>&1)

if [ "$student_out" = "$solution_out" ]; then
    echo -e "${GREEN}✓ Passed${NC}"
else
    echo -e "${RED}✗ Failed${NC}"
    echo "Expected: $solution_out"
    echo "Got: $student_out"
fi

echo ""
echo "Test 2: Multiple files"
student_out=$(student/ztail/ztail -c 4 test-files/file1.txt test-files/file2.txt 2>&1)
solution_out=$(solutions/ztail/ztail -c 4 test-files/file1.txt test-files/file2.txt 2>&1)

if [ "$student_out" = "$solution_out" ]; then
    echo -e "${GREEN}✓ Passed${NC}"
else
    echo -e "${RED}✗ Failed${NC}"
    echo "Expected:"
    echo "$solution_out"
    echo "Got:"
    echo "$student_out"
fi

echo ""
echo "Test 3: Error handling (exit status)"
student/ztail/ztail -c 4 test-files/nonexistent.txt >/dev/null 2>&1
student_exit=$?

if [ $student_exit -eq 1 ]; then
    echo -e "${GREEN}✓ Passed (exit status = 1)${NC}"
else
    echo -e "${RED}✗ Failed (exit status = $student_exit, expected 1)${NC}"
fi

echo ""
echo "Testing complete!"