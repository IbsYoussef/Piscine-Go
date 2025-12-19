#!/usr/bin/env bash
set -euo pipefail

echo "Running foreach..."
echo ""

cat > temp_run.go << 'EOF'
package main
EOF

# Include PrintNbr helper
cat >> temp_run.go << 'EOF'

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	
	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	
	for _, r := range digits {
		z01.PrintRune(r)
	}
}
EOF

sed '/^package student/d' student/foreach.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
	z01.PrintRune('\n')
}
EOF

go run temp_run.go
rm -f temp_run.go