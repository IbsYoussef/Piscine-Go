#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

trap 'rm -f student/ispowerof2/ispowerof2 solutions/ispowerof2/ispowerof2' EXIT

echo "Testing ispowerof2..."
echo ""

# Build programs
cd solutions/ispowerof2
go build 2>/dev/null
cd ../..

cd student/ispowerof2
go build 2>/dev/null
cd ../..

# Helper function
run_test() {
    local test_name="$1"
    shift
    local args=("$@")
    
    local student_out=$(student/ispowerof2/ispowerof2 "${args[@]}" 2>&1 || true)
    local solution_out=$(solutions/ispowerof2/ispowerof2 "${args[@]}" 2>&1 || true)
    
    if [ "$student_out" = "$solution_out" ]; then
        echo -e "${GREEN}✓${NC} $test_name"
        return 0
    else
        echo -e "${RED}✗${NC} $test_name"
        echo "  Expected: '$solution_out'"
        echo "  Got:      '$student_out'"
        return 1
    fi
}

# Run tests
run_test "Power of 2: 2" 2
run_test "Power of 2: 64" 64
run_test "Power of 2: 1" 1
run_test "Power of 2: 16" 16
run_test "NOT power of 2: 513" 513
run_test "NOT power of 2: 3" 3
run_test "NOT power of 2: 24" 24
run_test "No arguments"
run_test "Multiple arguments" 64 1024

echo ""
echo -e "${GREEN}Testing complete!${NC}"