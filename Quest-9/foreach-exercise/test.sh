#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing foreach..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	
	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	
	for _, r := range digits {
		z01.PrintRune(r)
	}
}
EOF

sed '/^package student/d' student/foreach.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
	z01.PrintRune('\n')
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	
	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	
	for _, r := range digits {
		z01.PrintRune(r)
	}
}
EOF

sed '/^package solutions/d' solutions/foreach.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
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
    echo "Output: $student_output"
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected: $solution_output"
    echo "Got: $student_output"
    exit 1
fi