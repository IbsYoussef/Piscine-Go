#!/usr/bin/env bash

echo "Running your solution (testing with -123, 0, 456):"
echo ""

cat > temp_run.go << 'EOF'
package main

import "github.com/01-edu/z01"
EOF

# Remove package and import lines
sed '/^package student/d; /^import/d' student/printnbr.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	PrintNbr(-123)
	z01.PrintRune('\n')
	PrintNbr(0)
	z01.PrintRune('\n')
	PrintNbr(456)
	z01.PrintRune('\n')
}
EOF

go run temp_run.go
rm -f temp_run.go