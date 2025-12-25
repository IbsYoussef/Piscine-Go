#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing rot14..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "github.com/01-edu/z01"
EOF

sed '/^package student/d' student/rot14.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	test1 := Rot14("Hello! How are You?")
	test2 := Rot14("abc")
	test3 := Rot14("xyz")

	for _, r := range test1 {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')

	for _, r := range test2 {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')

	for _, r := range test3 {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "github.com/01-edu/z01"
EOF

sed '/^package solutions/d' solutions/rot14.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	test1 := Rot14("Hello! How are You?")
	test2 := Rot14("abc")
	test3 := Rot14("xyz")

	for _, r := range test1 {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')

	for _, r := range test2 {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')

	for _, r := range test3 {
		z01.PrintRune(r)
	}
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
    echo "  Test 1: $(echo "$student_output" | sed -n '1p')"
    echo "  Test 2: $(echo "$student_output" | sed -n '2p')"
    echo "  Test 3: $(echo "$student_output" | sed -n '3p')"
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected:"
    echo "  Test 1: $(echo "$solution_output" | sed -n '1p')"
    echo "  Test 2: $(echo "$solution_output" | sed -n '2p')"
    echo "  Test 3: $(echo "$solution_output" | sed -n '3p')"
    echo ""
    echo "Got:"
    echo "  Test 1: $(echo "$student_output" | sed -n '1p')"
    echo "  Test 2: $(echo "$student_output" | sed -n '2p')"
    echo "  Test 3: $(echo "$student_output" | sed -n '3p')"
    exit 1
fi