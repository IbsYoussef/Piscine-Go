#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing itoa..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/itoa.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	test1 := Itoa(12345)
	test2 := Itoa(0)
	test3 := Itoa(-1234)
	test4 := Itoa(987654321)
	test5 := Itoa(-999)

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

sed '/^package solutions/d' solutions/itoa.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	test1 := Itoa(12345)
	test2 := Itoa(0)
	test3 := Itoa(-1234)
	test4 := Itoa(987654321)
	test5 := Itoa(-999)

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
    echo "  Itoa(12345):     $(echo "$student_output" | sed -n '1p')"
    echo "  Itoa(0):         $(echo "$student_output" | sed -n '2p')"
    echo "  Itoa(-1234):     $(echo "$student_output" | sed -n '3p')"
    echo "  Itoa(987654321): $(echo "$student_output" | sed -n '4p')"
    echo "  Itoa(-999):      $(echo "$student_output" | sed -n '5p')"
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected:"
    echo "  Itoa(12345):     $(echo "$solution_output" | sed -n '1p')"
    echo "  Itoa(0):         $(echo "$solution_output" | sed -n '2p')"
    echo "  Itoa(-1234):     $(echo "$solution_output" | sed -n '3p')"
    echo "  Itoa(987654321): $(echo "$solution_output" | sed -n '4p')"
    echo "  Itoa(-999):      $(echo "$solution_output" | sed -n '5p')"
    echo ""
    echo "Got:"
    echo "  Itoa(12345):     $(echo "$student_output" | sed -n '1p')"
    echo "  Itoa(0):         $(echo "$student_output" | sed -n '2p')"
    echo "  Itoa(-1234):     $(echo "$student_output" | sed -n '3p')"
    echo "  Itoa(987654321): $(echo "$student_output" | sed -n '4p')"
    echo "  Itoa(-999):      $(echo "$student_output" | sed -n '5p')"
    exit 1
fi