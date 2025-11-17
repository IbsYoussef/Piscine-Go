#!/usr/bin/env bash
set -euo pipefail

echo "Running printnbrbase..."
echo ""

cat > temp_run.go << 'EOF'
package main
EOF

# Remove package line but keep imports
sed '/^package student/d' student/printnbrbase.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}
EOF

go run temp_run.go
rm -f temp_run.go