#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing maxwordcountn..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/maxwordcountn.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	result := MaxWordCountN(`
	Orange is the sun sliding to the horizon after a summer day. Orange is the sound of dribbling basketball. Orange is the smell of a tiger lily petal. Orange is the taste of thirst-quenching Nehi Soda. Orange is the color of peach marmalade on a side of toast. Orange is the sound of a carrot popping out of the ground.
	`, 3)
	
	// Print in deterministic order for testing
	if val, ok := result["the"]; ok {
		fmt.Printf("the:%d ", val)
	}
	if val, ok := result["of"]; ok {
		fmt.Printf("of:%d ", val)
	}
	if val, ok := result["Orange"]; ok {
		fmt.Printf("Orange:%d", val)
	}
	fmt.Println()
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package solutions/d' solutions/maxwordcountn.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	result := MaxWordCountN(`
	Orange is the sun sliding to the horizon after a summer day. Orange is the sound of dribbling basketball. Orange is the smell of a tiger lily petal. Orange is the taste of thirst-quenching Nehi Soda. Orange is the color of peach marmalade on a side of toast. Orange is the sound of a carrot popping out of the ground.
	`, 3)
	
	// Print in deterministic order for testing
	if val, ok := result["the"]; ok {
		fmt.Printf("the:%d ", val)
	}
	if val, ok := result["of"]; ok {
		fmt.Printf("of:%d ", val)
	}
	if val, ok := result["Orange"]; ok {
		fmt.Printf("Orange:%d", val)
	}
	fmt.Println()
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
    echo "Got:      $student_output"
    exit 1
fi