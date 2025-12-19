#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing issorted..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "fmt"

func compare(a, b int) int {
	if a > b {
		return 1
	}
	if a == b {
		return 0
	}
	return -1
}
EOF

sed '/^package student/d' student/issorted.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	a3 := []int{5, 4, 3, 2, 1}
	a4 := []int{1, 1, 2, 2, 3}

	result1 := IsSorted(compare, a1)
	result2 := IsSorted(compare, a2)
	result3 := IsSorted(compare, a3)
	result4 := IsSorted(compare, a4)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"

func compare(a, b int) int {
	if a > b {
		return 1
	}
	if a == b {
		return 0
	}
	return -1
}
EOF

sed '/^package solutions/d' solutions/issorted.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	a3 := []int{5, 4, 3, 2, 1}
	a4 := []int{1, 1, 2, 2, 3}

	result1 := IsSorted(compare, a1)
	result2 := IsSorted(compare, a2)
	result3 := IsSorted(compare, a3)
	result4 := IsSorted(compare, a4)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
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
    echo "  Test 1 (Ascending):  $(echo "$student_output" | sed -n '1p')"
    echo "  Test 2 (Mixed):      $(echo "$student_output" | sed -n '2p')"
    echo "  Test 3 (Descending): $(echo "$student_output" | sed -n '3p')"
    echo "  Test 4 (Duplicates): $(echo "$student_output" | sed -n '4p')"
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected:"
    echo "  Test 1 (Ascending):  $(echo "$solution_output" | sed -n '1p')"
    echo "  Test 2 (Mixed):      $(echo "$solution_output" | sed -n '2p')"
    echo "  Test 3 (Descending): $(echo "$solution_output" | sed -n '3p')"
    echo "  Test 4 (Duplicates): $(echo "$solution_output" | sed -n '4p')"
    echo ""
    echo "Got:"
    echo "  Test 1 (Ascending):  $(echo "$student_output" | sed -n '1p')"
    echo "  Test 2 (Mixed):      $(echo "$student_output" | sed -n '2p')"
    echo "  Test 3 (Descending): $(echo "$student_output" | sed -n '3p')"
    echo "  Test 4 (Duplicates): $(echo "$student_output" | sed -n '4p')"
    exit 1
fi