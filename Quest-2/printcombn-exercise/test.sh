#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing printcombn..."
echo ""

cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

# Remove package and import lines from student file
sed '/^package student/d; /^import/d' student/printcombn.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	PrintCombN(1)
	PrintCombN(2)
	PrintCombN(3)
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"
EOF

# Remove package and import lines from solution file
sed '/^package solutions/d; /^import/d' solutions/printcombn.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	PrintCombN(1)
	PrintCombN(2)
	PrintCombN(3)
}
EOF

student_output=$(go run temp_student.go 2>&1)
student_exit=$?

solution_output=$(go run temp_solution.go 2>&1)
solution_exit=$?

rm -f temp_student.go temp_solution.go

if [ "$student_output" = "$solution_output" ] && [ $student_exit -eq $solution_exit ]; then
    echo -e "${GREEN}✓ Test passed!${NC}"
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected (first 300 chars):"
    echo "$solution_output" | head -c 300
    echo "..."
    echo ""
    echo "Got (first 300 chars):"
    echo "$student_output" | head -c 300
    echo "..."
    exit 1
fi