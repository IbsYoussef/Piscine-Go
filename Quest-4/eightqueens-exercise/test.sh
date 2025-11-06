#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing eightqueens..."
echo ""
echo "Comparing student output with solution..."
echo ""

cat > temp_student.go << 'EOF'
package main

import "github.com/01-edu/z01"
EOF

sed '/^package student/d' student/eightqueens.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	EightQueens()
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "github.com/01-edu/z01"
EOF

sed '/^package solutions/d' solutions/eightqueens.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	EightQueens()
}
EOF

student_output=$(go run temp_student.go 2>&1)
student_exit=$?

solution_output=$(go run temp_solution.go 2>&1)
solution_exit=$?

rm -f temp_student.go temp_solution.go

if [ "$student_output" = "$solution_output" ] && [ $student_exit -eq $solution_exit ]; then
    echo -e "${GREEN}✓ Test passed!${NC}"
    echo ""
    echo "All 92 solutions match!"
    echo ""
    echo "First 5 solutions:"
    echo "$student_output" | head -5
    echo "..."
    echo "Last 5 solutions:"
    echo "$student_output" | tail -5
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    
    # Count lines to help debug
    student_lines=$(echo "$student_output" | wc -l)
    solution_lines=$(echo "$solution_output" | wc -l)
    
    echo "Student output: $student_lines lines"
    echo "Expected output: $solution_lines lines"
    echo ""
    
    if [ $student_lines -ne $solution_lines ]; then
        echo "Line count mismatch! Should have 92 solutions."
    else
        echo "Line count matches but content differs."
        echo ""
        echo "First difference:"
        diff <(echo "$solution_output") <(echo "$student_output") | head -10
    fi
    
    exit 1
fi