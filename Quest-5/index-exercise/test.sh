#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing index..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/index.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
	fmt.Println(Index("test", ""))
	fmt.Println(Index("", "test"))
	fmt.Println(Index("Hello", "Hello"))
	fmt.Println(Index("Hello", "World"))
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package solutions/d' solutions/index.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
	fmt.Println(Index("test", ""))
	fmt.Println(Index("", "test"))
	fmt.Println(Index("Hello", "Hello"))
	fmt.Println(Index("Hello", "World"))
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