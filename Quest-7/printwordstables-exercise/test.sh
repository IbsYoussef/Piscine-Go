#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing printwordstables..."
echo ""

{
cat > temp_student.go << 'EOF'
package main
EOF

cat > temp_split.go << 'EOF'
package main

func SplitWhiteSpaces(s string) []string {
	result := []string{}
	word := ""
	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
EOF

sed '/^package student/d' student/printwordstables.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	a := SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}
EOF

cat > temp_solution.go << 'EOF'
package main
EOF

sed '/^package solutions/d' solutions/printwordstables.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	a := SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}
EOF

student_output=$(go run temp_student.go temp_split.go)
solution_output=$(go run temp_solution.go temp_split.go)

rm -f temp_student.go temp_solution.go temp_split.go

} 2>/dev/null

if [ "$student_output" = "$solution_output" ]; then
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