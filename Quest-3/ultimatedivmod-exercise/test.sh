#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing ultimatedivmod..."
echo ""

cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/ultimatedivmod.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
	
	x := 50
	y := 7
	UltimateDivMod(&x, &y)
	fmt.Println(x)
	fmt.Println(y)
	
	m := 100
	n := 9
	UltimateDivMod(&m, &n)
	fmt.Println(m)
	fmt.Println(n)
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package solutions/d' solutions/ultimatedivmod.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
	
	x := 50
	y := 7
	UltimateDivMod(&x, &y)
	fmt.Println(x)
	fmt.Println(y)
	
	m := 100
	n := 9
	UltimateDivMod(&m, &n)
	fmt.Println(m)
	fmt.Println(n)
}
EOF

student_output=$(go run temp_student.go 2>&1)
student_exit=$?

solution_output=$(go run temp_solution.go 2>&1)
solution_exit=$?

rm -f temp_student.go temp_solution.go

if [ "$student_output" = "$solution_output" ] && [ $student_exit -eq $solution_exit ]; then
    echo -e "${GREEN}✓ Test passed!${NC}"
    echo ""
    echo "Output:"
    echo "$student_output"
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected:"
    echo "$solution_output"
    echo ""
    echo "Got:"
    echo "$student_output"
    exit 1
fi