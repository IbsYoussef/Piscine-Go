#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing compact..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/compact.go >> temp_student.go

cat >> temp_student.go << 'EOF'

const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package solutions/d' solutions/compact.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
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
    echo "$student_output" | sed 's/^/  /'
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected:"
    echo "$solution_output" | sed 's/^/  /'
    echo ""
    echo "Got:"
    echo "$student_output" | sed 's/^/  /'
    exit 1
fi