#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing flags..."
echo ""

{
# Build student program
cd student/flags
go build

# Run test cases
student_test1=$(./flags --insert=4321 --order asdad)
student_test2=$(./flags --insert=4321 asdad)
student_test3=$(./flags asdad)
student_test4=$(./flags --order 43a21)

rm -f flags
cd ../..

} 2>/dev/null

# Expected outputs
expected_test1="1234aadds"
expected_test2="asdad4321"
expected_test3="asdad"
expected_test4="1234a"

# Compare results
if [ "$student_test1" = "$expected_test1" ] && \
   [ "$student_test2" = "$expected_test2" ] && \
   [ "$student_test3" = "$expected_test3" ] && \
   [ "$student_test4" = "$expected_test4" ]; then
    echo -e "${GREEN}✓ All tests passed!${NC}"
    echo ""
    echo "Test 1 (--insert=4321 --order asdad): $student_test1"
    echo "Test 2 (--insert=4321 asdad): $student_test2"
    echo "Test 3 (asdad): $student_test3"
    echo "Test 4 (--order 43a21): $student_test4"
    exit 0
else
    echo -e "${RED}✗ Tests failed${NC}"
    echo ""
    if [ "$student_test1" != "$expected_test1" ]; then
        echo "Test 1 - Expected: $expected_test1, Got: $student_test1"
    fi
    if [ "$student_test2" != "$expected_test2" ]; then
        echo "Test 2 - Expected: $expected_test2, Got: $student_test2"
    fi
    if [ "$student_test3" != "$expected_test3" ]; then
        echo "Test 3 - Expected: $expected_test3, Got: $student_test3"
    fi
    if [ "$student_test4" != "$expected_test4" ]; then
        echo "Test 4 - Expected: $expected_test4, Got: $student_test4"
    fi
    exit 1
fi