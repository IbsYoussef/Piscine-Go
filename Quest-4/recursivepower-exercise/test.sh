#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing recursivepower..."
echo ""

cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/recursivepower.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	fmt.Println(RecursivePower(4, 3))
	fmt.Println(RecursivePower(2, 5))
	fmt.Println(RecursivePower(5, 0))
	fmt.Println(RecursivePower(3, -2))
	fmt.Println(RecursivePower(10, 1))
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package solutions/d' solutions/recursivepower.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	fmt.Println(RecursivePower(4, 3))
	fmt.Println(RecursivePower(2, 5))
	fmt.Println(RecursivePower(5, 0))
	fmt.Println(RecursivePower(3, -2))
	fmt.Println(RecursivePower(10, 1))
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