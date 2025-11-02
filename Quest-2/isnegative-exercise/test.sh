#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing isnegative..."
echo ""

# Create test for student
cat > temp_student_run.go << 'EOF'
package main
EOF

sed 's/package student//' student/isnegative.go >> temp_student_run.go

cat >> temp_student_run.go << 'EOF'

func main() {
	IsNegative(-5)
	IsNegative(0)
	IsNegative(10)
	IsNegative(-1)
}
EOF

# Create test for solution
cat > temp_solution_run.go << 'EOF'
package main
EOF

sed 's/package solutions//' solutions/isnegative.go >> temp_solution_run.go

cat >> temp_solution_run.go << 'EOF'

func main() {
	IsNegative(-5)
	IsNegative(0)
	IsNegative(10)
	IsNegative(-1)
}
EOF

# Run both
student_output=$(go run temp_student_run.go 2>&1)
student_exit=$?

solution_output=$(go run temp_solution_run.go 2>&1)
solution_exit=$?

# Cleanup
rm -f temp_student_run.go temp_solution_run.go

# Compare
if [ "$student_output" = "$solution_output" ] && [ $student_exit -eq $solution_exit ]; then
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