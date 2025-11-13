#!/usr/bin/env bash
set -euo pipefail

echo "Running firstrune..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "github.com/01-edu/z01"
EOF

sed '/^package student/d' student/firstrune.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	z01.PrintRune(FirstRune("Hello!"))
    z01.PrintRune('\n')
	z01.PrintRune(FirstRune("Salut!"))
    z01.PrintRune('\n')
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}
EOF

go run temp_run.go
rm -f temp_run.go