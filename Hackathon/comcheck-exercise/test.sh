#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

trap 'rm -f student/comcheck/comcheck solutions/comcheck/comcheck' EXIT

echo "Testing comcheck..."
echo ""

# Build programs
cd solutions/comcheck
go build 2>/dev/null
cd ../..

cd student/comcheck
go build 2>/dev/null
cd ../..

# Helper function
run_test() {
    local test_name="$1"
    shift
    local args=("$@")
    
    local student_out=$(student/comcheck/comcheck "${args[@]}" 2>&1 || true)
    local solution_out=$(solutions/comcheck/comcheck "${args[@]}" 2>&1 || true)
    
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
run_test "Match 'galaxy'" "I" "Will" "Enter" "the" "galaxy"
run_test "Match 'galaxy 01'" "galaxy 01" "do" "you" "hear" "me"
run_test "Match '01'" "test" "01" "hello"
run_test "No match" "hello" "world"
run_test "Multiple matches" "galaxy" "01"

echo ""
echo -e "${GREEN}Testing complete!${NC}"