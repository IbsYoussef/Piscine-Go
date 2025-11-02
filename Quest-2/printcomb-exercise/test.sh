#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing printcomb..."
echo ""

# Create test for student
cat > temp_student.go << 'EOF'
package main
EOF

sed 's/package student//' student/printcomb.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	PrintComb()
}
EOF

# Create test for solution
cat > temp_solution.go << 'EOF'
package main
EOF

sed 's/package solutions//' solutions/printcomb.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	PrintComb()
}
EOF

# Run both
student_output=$(go run temp_student.go 2>&1)
student_exit=$?

solution_output=$(go run temp_solution.go 2>&1)
solution_exit=$?

# Cleanup
rm -f temp_student.go temp_solution.go

# Compare
if [ "$student_output" = "$solution_output" ] && [ $student_exit -eq $solution_exit ]; then
    echo -e "${GREEN}✓ Test passed!${NC}"
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected:"
    echo "$solution_output" | head -c 200
    echo "..."
    echo ""
    echo "Got:"
    echo "$student_output" | head -c 200
    echo "..."
    exit 1
fi