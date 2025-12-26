#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing collatzcountdown..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/collatzcountdown.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	test1 := CollatzCountdown(12)
	test2 := CollatzCountdown(5)
	test3 := CollatzCountdown(1)
	test4 := CollatzCountdown(0)
	test5 := CollatzCountdown(27)

	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
	fmt.Println(test4)
	fmt.Println(test5)
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package solutions/d' solutions/collatzcountdown.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	test1 := CollatzCountdown(12)
	test2 := CollatzCountdown(5)
	test3 := CollatzCountdown(1)
	test4 := CollatzCountdown(0)
	test5 := CollatzCountdown(27)

	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
	fmt.Println(test4)
	fmt.Println(test5)
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
    echo "  Test 1 (12): $(echo "$student_output" | sed -n '1p')"
    echo "  Test 2 (5): $(echo "$student_output" | sed -n '2p')"
    echo "  Test 3 (1): $(echo "$student_output" | sed -n '3p')"
    echo "  Test 4 (0): $(echo "$student_output" | sed -n '4p')"
    echo "  Test 5 (27): $(echo "$student_output" | sed -n '5p')"
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected:"
    echo "  Test 1 (12): $(echo "$solution_output" | sed -n '1p')"
    echo "  Test 2 (5): $(echo "$solution_output" | sed -n '2p')"
    echo "  Test 3 (1): $(echo "$solution_output" | sed -n '3p')"
    echo "  Test 4 (0): $(echo "$solution_output" | sed -n '4p')"
    echo "  Test 5 (27): $(echo "$solution_output" | sed -n '5p')"
    echo ""
    echo "Got:"
    echo "  Test 1 (12): $(echo "$student_output" | sed -n '1p')"
    echo "  Test 2 (5): $(echo "$student_output" | sed -n '2p')"
    echo "  Test 3 (1): $(echo "$student_output" | sed -n '3p')"
    echo "  Test 4 (0): $(echo "$student_output" | sed -n '4p')"
    echo "  Test 5 (27): $(echo "$student_output" | sed -n '5p')"
    exit 1
fi