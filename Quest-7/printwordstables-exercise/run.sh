#!/usr/bin/env bash
set -euo pipefail

echo "Running printwordstables..."
echo ""

cat > temp_run.go << 'EOF'
package main
EOF

# Need to include SplitWhiteSpaces for the test
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

sed '/^package student/d' student/printwordstables.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	a := SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}
EOF

go run temp_run.go temp_split.go
rm -f temp_run.go temp_split.go