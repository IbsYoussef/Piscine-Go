#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing printnbr..."
echo ""

cat > temp_student.go << 'EOF'
package main

import "github.com/01-edu/z01"
EOF

# Remove package and import lines from student
sed '/^package student/d; /^import/d' student/printnbr.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	PrintNbr(-123)
	z01.PrintRune('\n')
	PrintNbr(0)
	z01.PrintRune('\n')
	PrintNbr(456)
	z01.PrintRune('\n')
	PrintNbr(-9223372036854775808) // min int
	z01.PrintRune('\n')
	PrintNbr(9223372036854775807) // max int
	z01.PrintRune('\n')
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "github.com/01-edu/z01"
EOF

# Remove package and import lines from solution
sed '/^package solutions/d; /^import/d' solutions/printnbr.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	PrintNbr(-123)
	z01.PrintRune('\n')
	PrintNbr(0)
	z01.PrintRune('\n')
	PrintNbr(456)
	z01.PrintRune('\n')
	PrintNbr(-9223372036854775808) // min int
	z01.PrintRune('\n')
	PrintNbr(9223372036854775807) // max int
	z01.PrintRune('\n')
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