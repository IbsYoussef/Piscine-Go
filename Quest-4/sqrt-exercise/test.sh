#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing sqrt..."
echo ""

cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/sqrt.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(16))
	fmt.Println(Sqrt(0))
	fmt.Println(Sqrt(-4))
	fmt.Println(Sqrt(25))
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package solutions/d' solutions/sqrt.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(16))
	fmt.Println(Sqrt(0))
	fmt.Println(Sqrt(-4))
	fmt.Println(Sqrt(25))
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