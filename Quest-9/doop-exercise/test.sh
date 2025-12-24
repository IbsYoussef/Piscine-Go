#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

trap 'rm -f student/doop/doop solutions/doop/doop' EXIT

echo "Testing doop..."
echo ""

# Build programs
cd solutions/doop
go build 2>/dev/null
cd ../..

cd student/doop
go build 2>/dev/null
cd ../..

# Helper function to run test
run_test() {
    local test_name="$1"
    shift
    local args=("$@")
    
    local student_out=$(student/doop/doop "${args[@]}" 2>&1 || true)
    local solution_out=$(solutions/doop/doop "${args[@]}" 2>&1 || true)
    
    if [ "$student_out" = "$solution_out" ]; then
        echo -e "${GREEN}✓${NC} $test_name"
        return 0
    else
        echo -e "${RED}✗${NC} $test_name"
        echo "  Expected: $solution_out"
        echo "  Got:      $student_out"
        return 1
    fi
}

# Run tests
run_test "Addition (1 + 1)" 1 + 1
run_test "Subtraction (10 - 3)" 10 - 3
run_test "Multiplication (5 * 4)" 5 "*" 4
run_test "Division (20 / 4)" 20 / 4
run_test "Modulo (17 % 5)" 17 % 5
run_test "Negative result (1 - 10)" 1 - 10
run_test "Division by zero" 1 / 0
run_test "Modulo by zero" 1 % 0
run_test "Invalid operator" 1 p 1
run_test "Invalid number" hello + 1
run_test "Overflow (MaxInt + 1)" 9223372036854775807 + 1
run_test "No arguments"

echo ""
echo -e "${GREEN}Testing complete!${NC}"