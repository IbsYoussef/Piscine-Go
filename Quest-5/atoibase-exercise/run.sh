#!/usr/bin/env bash
set -euo pipefail

echo "Running atoibase..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/atoibase.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
EOF

go run temp_run.go
rm -f temp_run.go