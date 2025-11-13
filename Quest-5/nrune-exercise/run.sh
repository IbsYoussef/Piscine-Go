#!/usr/bin/env bash
set -euo pipefail

echo "Running nrune..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "github.com/01-edu/z01"
EOF

sed '/^package student/d' student/nrune.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}
EOF

go run temp_run.go
rm -f temp_run.go