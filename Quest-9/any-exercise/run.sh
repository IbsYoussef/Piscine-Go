#!/usr/bin/env bash
set -euo pipefail

echo "Running any..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

# Include IsNumeric helper
cat >> temp_run.go << 'EOF'

func IsNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
EOF

sed '/^package student/d' student/any.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Any(IsNumeric, a1)
	result2 := Any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
EOF

go run temp_run.go
rm -f temp_run.go