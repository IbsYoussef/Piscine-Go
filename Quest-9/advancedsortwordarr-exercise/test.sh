#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing advancedsortwordarr..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "fmt"

func Compare(a, b string) int {
	if a > b {
		return 1
	}
	if a == b {
		return 0
	}
	return -1
}
EOF

sed '/^package student/d' student/advancedsortwordarr.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	test1 := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	test2 := []string{"z", "y", "x", "w"}
	test3 := []string{"hello", "world", "go", "programming"}

	AdvancedSortWordArr(test1, Compare)
	AdvancedSortWordArr(test2, Compare)
	AdvancedSortWordArr(test3, Compare)

	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"

func Compare(a, b string) int {
	if a > b {
		return 1
	}
	if a == b {
		return 0
	}
	return -1
}
EOF

sed '/^package solutions/d' solutions/advancedsortwordarr.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	test1 := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	test2 := []string{"z", "y", "x", "w"}
	test3 := []string{"hello", "world", "go", "programming"}

	AdvancedSortWordArr(test1, Compare)
	AdvancedSortWordArr(test2, Compare)
	AdvancedSortWordArr(test3, Compare)

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
```