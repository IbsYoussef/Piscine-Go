#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing rotatevowels..."
echo ""

{
# Build student program
cd student/rotatevowels
go build

# Run test cases
student_test1=$(./rotatevowels "Hello World")
student_test2=$(./rotatevowels "HEllO World" "problem solved")
student_test3=$(./rotatevowels "str" "shh" "psst")
student_test4=$(./rotatevowels "happy thoughts" "good luck")
student_test5=$(./rotatevowels "aEi" "Ou")
student_test6=$(./rotatevowels)

rm -f rotatevowels
cd ../..

} 2>/dev/null

# Expected outputs
expected_test1="Hollo Werld"
expected_test2="Hello Werld problom sOlvEd"
expected_test3="str shh psst"
expected_test4="huppy thooghts guod lack"
expected_test5="uOi Ea"
expected_test6=""

# Compare results
if [ "$student_test1" = "$expected_test1" ] && \
   [ "$student_test2" = "$expected_test2" ] && \
   [ "$student_test3" = "$expected_test3" ] && \
   [ "$student_test4" = "$expected_test4" ] && \
   [ "$student_test5" = "$expected_test5" ] && \
   [ "$student_test6" = "$expected_test6" ]; then
    echo -e "${GREEN}✓ All tests passed!${NC}"
    echo ""
    echo "Test 1: $student_test1"
    echo "Test 2: $student_test2"
    echo "Test 3: $student_test3"
    echo "Test 4: $student_test4"
    echo "Test 5: $student_test5"
    exit 0
else
    echo -e "${RED}✗ Tests failed${NC}"
    echo ""
    if [ "$student_test1" != "$expected_test1" ]; then
        echo "Test 1 - Expected: '$expected_test1', Got: '$student_test1'"
    fi
    if [ "$student_test2" != "$expected_test2" ]; then
        echo "Test 2 - Expected: '$expected_test2', Got: '$student_test2'"
    fi
    if [ "$student_test3" != "$expected_test3" ]; then
        echo "Test 3 - Expected: '$expected_test3', Got: '$student_test3'"
    fi
    if [ "$student_test4" != "$expected_test4" ]; then
        echo "Test 4 - Expected: '$expected_test4', Got: '$student_test4'"
    fi
    if [ "$student_test5" != "$expected_test5" ]; then
        echo "Test 5 - Expected: '$expected_test5', Got: '$student_test5'"
    fi
    if [ "$student_test6" != "$expected_test6" ]; then
        echo "Test 6 - Expected: (empty), Got: '$student_test6'"
    fi
    exit 1
fi