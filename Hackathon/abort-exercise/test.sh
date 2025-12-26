#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing abort..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/abort.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	test1 := Abort(2, 3, 8, 5, 7)
	test2 := Abort(1, 2, 3, 4, 5)
	test3 := Abort(9, 1, 5, 3, 7)

	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package solutions/d' solutions/abort.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	test1 := Abort(2, 3, 8, 5, 7)
	test2 := Abort(1, 2, 3, 4, 5)
	test3 := Abort(9, 1, 5, 3, 7)

	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
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
    echo "  Test 1 (2,3,8,5,7): $(echo "$student_output" | sed -n '1p')"
    echo "  Test 2 (1,2,3,4,5): $(echo "$student_output" | sed -n '2p')"
    echo "  Test 3 (9,1,5,3,7): $(echo "$student_output" | sed -n '3p')"
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