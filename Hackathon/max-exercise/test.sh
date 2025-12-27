#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing max..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/max.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	test1 := []int{23, 123, 1, 11, 55, 93}
	test2 := []int{-5, -10, -3, -20}
	test3 := []int{42}
	test4 := []int{}
	test5 := []int{0, -1, 5, -10, 100}

	fmt.Println(Max(test1))
	fmt.Println(Max(test2))
	fmt.Println(Max(test3))
	fmt.Println(Max(test4))
	fmt.Println(Max(test5))
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package solutions/d' solutions/max.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	test1 := []int{23, 123, 1, 11, 55, 93}
	test2 := []int{-5, -10, -3, -20}
	test3 := []int{42}
	test4 := []int{}
	test5 := []int{0, -1, 5, -10, 100}

	fmt.Println(Max(test1))
	fmt.Println(Max(test2))
	fmt.Println(Max(test3))
	fmt.Println(Max(test4))
	fmt.Println(Max(test5))
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
    echo "  Positive numbers: $(echo "$student_output" | sed -n '1p')"
    echo "  Negative numbers: $(echo "$student_output" | sed -n '2p')"
    echo "  Single element:   $(echo "$student_output" | sed -n '3p')"
    echo "  Empty slice:      $(echo "$student_output" | sed -n '4p')"
    echo "  Mixed numbers:    $(echo "$student_output" | sed -n '5p')"
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected:"
    echo "  Positive numbers: $(echo "$solution_output" | sed -n '1p')"
    echo "  Negative numbers: $(echo "$solution_output" | sed -n '2p')"
    echo "  Single element:   $(echo "$solution_output" | sed -n '3p')"
    echo "  Empty slice:      $(echo "$solution_output" | sed -n '4p')"
    echo "  Mixed numbers:    $(echo "$solution_output" | sed -n '5p')"
    echo ""
    echo "Got:"
    echo "  Positive numbers: $(echo "$student_output" | sed -n '1p')"
    echo "  Negative numbers: $(echo "$student_output" | sed -n '2p')"
    echo "  Single element:   $(echo "$student_output" | sed -n '3p')"
    echo "  Empty slice:      $(echo "$student_output" | sed -n '4p')"
    echo "  Mixed numbers:    $(echo "$student_output" | sed -n '5p')"
    exit 1
fi