#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing nrune..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "github.com/01-edu/z01"
EOF

sed '/^package student/d' student/nrune.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
	z01.PrintRune(NRune("Go", 1))
	z01.PrintRune(NRune("Programming", 7))
	z01.PrintRune(NRune("Test", 0))
	z01.PrintRune(NRune("Test", 10))
	z01.PrintRune('\n')
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "github.com/01-edu/z01"
EOF

sed '/^package solutions/d' solutions/nrune.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
	z01.PrintRune(NRune("Go", 1))
	z01.PrintRune(NRune("Programming", 7))
	z01.PrintRune(NRune("Test", 0))
	z01.PrintRune(NRune("Test", 10))
	z01.PrintRune('\n')
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