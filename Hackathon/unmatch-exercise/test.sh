#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing unmatch..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/unmatch.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	test1 := []int{1, 2, 3, 1, 2, 3, 4}
	test2 := []int{5, 5, 7, 7}
	test3 := []int{1, 1, 1}
	test4 := []int{9}
	test5 := []int{2, 3, 2, 3, 5, 5, 8}

	fmt.Println(Unmatch(test1))
	fmt.Println(Unmatch(test2))
	fmt.Println(Unmatch(test3))
	fmt.Println(Unmatch(test4))
	fmt.Println(Unmatch(test5))
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package solutions/d' solutions/unmatch.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	test1 := []int{1, 2, 3, 1, 2, 3, 4}
	test2 := []int{5, 5, 7, 7}
	test3 := []int{1, 1, 1}
	test4 := []int{9}
	test5 := []int{2, 3, 2, 3, 5, 5, 8}

	fmt.Println(Unmatch(test1))
	fmt.Println(Unmatch(test2))
	fmt.Println(Unmatch(test3))
	fmt.Println(Unmatch(test4))
	fmt.Println(Unmatch(test5))
}
EOF

student_output=$(go run temp_student.go)
solution_output=$(go run temp_solution.go)

rm -f temp_student.go temp_solution.go

} 2>/dev/null

if [ "$student_output" = "$solution_output" ]; then
    echo -e "${GREEN}✓ Test passed!${NC}"
    echo ""
    echo "Output:"
    echo "  Test 1 (one unpaired):    $(echo "$student_output" | sed -n '1p')"
    echo "  Test 2 (all paired):      $(echo "$student_output" | sed -n '2p')"
    echo "  Test 3 (odd occurrences): $(echo "$student_output" | sed -n '3p')"
    echo "  Test 4 (single element):  $(echo "$student_output" | sed -n '4p')"
    echo "  Test 5 (one unpaired):    $(echo "$student_output" | sed -n '5p')"
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected:"
    echo "  Test 1: $(echo "$solution_output" | sed -n '1p')"
    echo "  Test 2: $(echo "$solution_output" | sed -n '2p')"
    echo "  Test 3: $(echo "$solution_output" | sed -n '3p')"
    echo "  Test 4: $(echo "$solution_output" | sed -n '4p')"
    echo "  Test 5: $(echo "$solution_output" | sed -n '5p')"
    echo ""
    echo "Got:"
    echo "  Test 1: $(echo "$student_output" | sed -n '1p')"
    echo "  Test 2: $(echo "$student_output" | sed -n '2p')"
    echo "  Test 3: $(echo "$student_output" | sed -n '3p')"
    echo "  Test 4: $(echo "$student_output" | sed -n '4p')"
    echo "  Test 5: $(echo "$student_output" | sed -n '5p')"
    exit 1
fi