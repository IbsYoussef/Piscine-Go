#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing basicjoin..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/basicjoin.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	elems1 := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems1))
	
	elems2 := []string{"Go", "Lang"}
	fmt.Println(BasicJoin(elems2))
	
	elems3 := []string{}
	fmt.Println(BasicJoin(elems3))
	
	elems4 := []string{"single"}
	fmt.Println(BasicJoin(elems4))
	
	elems5 := []string{"a", "b", "c"}
	fmt.Println(BasicJoin(elems5))
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package solutions/d' solutions/basicjoin.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	elems1 := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems1))
	
	elems2 := []string{"Go", "Lang"}
	fmt.Println(BasicJoin(elems2))
	
	elems3 := []string{}
	fmt.Println(BasicJoin(elems3))
	
	elems4 := []string{"single"}
	fmt.Println(BasicJoin(elems4))
	
	elems5 := []string{"a", "b", "c"}
	fmt.Println(BasicJoin(elems5))
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