#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing ultimatepointone..."
echo ""

cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/ultimatepointone.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	a := 0
	b := &a
	n := &b
	UltimatePointOne(&n)
	fmt.Println(a)
	
	x := 42
	y := &x
	z := &y
	UltimatePointOne(&z)
	fmt.Println(x)
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package solutions/d' solutions/ultimatepointone.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	a := 0
	b := &a
	n := &b
	UltimatePointOne(&n)
	fmt.Println(a)
	
	x := 42
	y := &x
	z := &y
	UltimatePointOne(&z)
	fmt.Println(x)
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