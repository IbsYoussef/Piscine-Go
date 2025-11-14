#!/usr/bin/env bash
set -euo pipefail

echo "Running compare..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/compare.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}
EOF

go run temp_run.go
rm -f temp_run.go