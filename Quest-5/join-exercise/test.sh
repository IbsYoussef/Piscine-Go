#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing join..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/join.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	toConcat1 := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat1, ":"))
	
	toConcat2 := []string{"Go", "Lang"}
	fmt.Println(Join(toConcat2, "-"))
	
	toConcat3 := []string{}
	fmt.Println(Join(toConcat3, ","))
	
	toConcat4 := []string{"single"}
	fmt.Println(Join(toConcat4, ","))
	
	toConcat5 := []string{"a", "b", "c"}
	fmt.Println(Join(toConcat5, ""))
	
	toConcat6 := []string{"one", "two", "three"}
	fmt.Println(Join(toConcat6, " | "))
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package solutions/d' solutions/join.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	toConcat1 := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat1, ":"))
	
	toConcat2 := []string{"Go", "Lang"}
	fmt.Println(Join(toConcat2, "-"))
	
	toConcat3 := []string{}
	fmt.Println(Join(toConcat3, ","))
	
	toConcat4 := []string{"single"}
	fmt.Println(Join(toConcat4, ","))
	
	toConcat5 := []string{"a", "b", "c"}
	fmt.Println(Join(toConcat5, ""))
	
	toConcat6 := []string{"one", "two", "three"}
	fmt.Println(Join(toConcat6, " | "))
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