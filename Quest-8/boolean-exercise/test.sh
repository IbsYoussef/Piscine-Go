#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing boolean..."
echo ""

{
# Build student program
cd student/boolean
go build

student_test1=$(./boolean "not" "odd")
student_test2=$(./boolean "not even")
student_test3=$(./boolean)
student_test4=$(./boolean "one" "two" "three")

rm -f boolean
cd ../..

# Build solution program
cd solutions/boolean
go build

solution_test1=$(./boolean "not" "odd")
solution_test2=$(./boolean "not even")
solution_test3=$(./boolean)
solution_test4=$(./boolean "one" "two" "three")

rm -f boolean
cd ../..

} 2>/dev/null

if [ "$student_test1" = "$solution_test1" ] && \
   [ "$student_test2" = "$solution_test2" ] && \
   [ "$student_test3" = "$solution_test3" ] && \
   [ "$student_test4" = "$solution_test4" ]; then
    echo -e "${GREEN}✓ All tests passed!${NC}"
    echo ""
    echo "Test 1 (even): $student_test1"
    echo "Test 2 (odd): $student_test2"
    echo "Test 3 (even): $student_test3"
    echo "Test 4 (odd): $student_test4"
    exit 0
else
    echo -e "${RED}✗ Tests failed${NC}"
    echo ""
    if [ "$student_test1" != "$solution_test1" ]; then
        echo "Test 1 - Expected: $solution_test1, Got: $student_test1"
    fi
    if [ "$student_test2" != "$solution_test2" ]; then
        echo "Test 2 - Expected: $solution_test2, Got: $student_test2"
    fi
    if [ "$student_test3" != "$solution_test3" ]; then
        echo "Test 3 - Expected: $solution_test3, Got: $student_test3"
    fi
    if [ "$student_test4" != "$solution_test4" ]; then
        echo "Test 4 - Expected: $solution_test4, Got: $student_test4"
    fi
    exit 1
fi