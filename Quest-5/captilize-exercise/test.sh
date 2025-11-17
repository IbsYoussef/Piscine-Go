#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing capitalize..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/capitalize.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
	fmt.Println(Capitalize("hello"))
	fmt.Println(Capitalize("HELLO WORLD"))
	fmt.Println(Capitalize("abc123def"))
	fmt.Println(Capitalize("a b c"))
	fmt.Println(Capitalize(""))
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package solutions/d' solutions/capitalize.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
	fmt.Println(Capitalize("hello"))
	fmt.Println(Capitalize("HELLO WORLD"))
	fmt.Println(Capitalize("abc123def"))
	fmt.Println(Capitalize("a b c"))
	fmt.Println(Capitalize(""))
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