#!/usr/bin/env bash
set -euo pipefail

echo "Running lastrune..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "github.com/01-edu/z01"
EOF

sed '/^package student/d' student/lastrune.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
    z01.PrintRune(LastRune("Hello!"))
    z01.PrintRune(LastRune("Salut!"))
    z01.PrintRune(LastRune("Ola!"))
    z01.PrintRune('\n')
}
EOF

go run temp_run.go
rm -f temp_run.go