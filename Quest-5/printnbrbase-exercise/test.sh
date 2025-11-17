#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing printnbrbase..."
echo ""

{
cat > temp_student.go << 'EOF'
package main
EOF

# Remove package line but keep imports
sed '/^package student/d' student/printnbrbase.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
	PrintNbrBase(0, "01")
	z01.PrintRune('\n')
	PrintNbrBase(255, "0123456789ABCDEF")
	z01.PrintRune('\n')
}
EOF

cat > temp_solution.go << 'EOF'
package main
EOF

# Remove package line but keep imports
sed '/^package solutions/d' solutions/printnbrbase.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
	PrintNbrBase(0, "01")
	z01.PrintRune('\n')
	PrintNbrBase(255, "0123456789ABCDEF")
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