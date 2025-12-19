#!/usr/bin/env bash
set -euo pipefail

echo "Running countif..."
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

sed '/^package student/d' student/countif.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This","1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
EOF

go run temp_run.go
rm -f temp_run.go