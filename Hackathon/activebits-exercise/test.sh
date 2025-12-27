#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing activebits..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/activebits.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	test1 := ActiveBits(7)
	test2 := ActiveBits(10)
	test3 := ActiveBits(15)
	test4 := ActiveBits(0)
	test5 := ActiveBits(1)

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

sed '/^package solutions/d' solutions/activebits.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	test1 := ActiveBits(7)
	test2 := ActiveBits(10)
	test3 := ActiveBits(15)
	test4 := ActiveBits(0)
	test5 := ActiveBits(1)

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
    echo "  ActiveBits(7):  $(echo "$student_output" | sed -n '1p')"
    echo "  ActiveBits(10): $(echo "$student_output" | sed -n '2p')"
    echo "  ActiveBits(15): $(echo "$student_output" | sed -n '3p')"
    echo "  ActiveBits(0):  $(echo "$student_output" | sed -n '4p')"
    echo "  ActiveBits(1):  $(echo "$student_output" | sed -n '5p')"
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected:"
    echo "  ActiveBits(7):  $(echo "$solution_output" | sed -n '1p')"
    echo "  ActiveBits(10): $(echo "$solution_output" | sed -n '2p')"
    echo "  ActiveBits(15): $(echo "$solution_output" | sed -n '3p')"
    echo "  ActiveBits(0):  $(echo "$solution_output" | sed -n '4p')"
    echo "  ActiveBits(1):  $(echo "$solution_output" | sed -n '5p')"
    echo ""
    echo "Got:"
    echo "  ActiveBits(7):  $(echo "$student_output" | sed -n '1p')"
    echo "  ActiveBits(10): $(echo "$student_output" | sed -n '2p')"
    echo "  ActiveBits(15): $(echo "$student_output" | sed -n '3p')"
    echo "  ActiveBits(0):  $(echo "$student_output" | sed -n '4p')"
    echo "  ActiveBits(1):  $(echo "$student_output" | sed -n '5p')"
    exit 1
fi