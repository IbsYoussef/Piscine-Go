#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

trap 'rm -f student/cat/cat solutions/cat/cat' EXIT

echo "Testing cat..."
echo ""

cd solutions/cat
go build
cd ../..

cd student/cat
go build
cd ../..

echo "Test 1: Single file"
student_out=$(student/cat/cat test-files/quest8.txt 2>&1)
solution_out=$(solutions/cat/cat test-files/quest8.txt 2>&1)

if [ "$student_out" = "$solution_out" ]; then
    echo -e "${GREEN}✓ Passed${NC}"
else
    echo -e "${RED}✗ Failed${NC}"
    echo "Expected: $solution_out"
    echo "Got: $student_out"
fi

echo ""
echo "Test 2: Multiple files"
student_out=$(student/cat/cat test-files/quest8.txt test-files/quest8T.txt 2>&1)
solution_out=$(solutions/cat/cat test-files/quest8.txt test-files/quest8T.txt 2>&1)

if [ "$student_out" = "$solution_out" ]; then
    echo -e "${GREEN}✓ Passed${NC}"
else
    echo -e "${RED}✗ Failed${NC}"
fi

echo ""
echo "Test 3: Piped input (stdin mode)"
student_out=$(echo "Hello from pipe!" | timeout 2 student/cat/cat 2>&1)
solution_out=$(echo "Hello from pipe!" | timeout 2 solutions/cat/cat 2>&1)

if [ "$student_out" = "$solution_out" ]; then
    echo -e "${GREEN}✓ Passed${NC}"
else
    echo -e "${RED}✗ Failed${NC}"
    echo "Expected: $solution_out"
    echo "Got: $student_out"
fi

echo ""
echo "Test 4: Error handling"
student/cat/cat test-files/nonexistent.txt >/dev/null 2>&1
student_exit=$?

if [ $student_exit -eq 1 ]; then
    echo -e "${GREEN}✓ Passed (exit status = 1)${NC}"
else
    echo -e "${RED}✗ Failed (exit status = $student_exit, expected 1)${NC}"
fi

echo ""
echo "Testing complete!"