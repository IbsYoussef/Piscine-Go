#!/usr/bin/env bash
set -euo pipefail

echo "Running ultimatepointone..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/ultimatepointone.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	a := 0
	b := &a
	n := &b
	UltimatePointOne(&n)
	fmt.Println(a)
}
EOF

go run temp_run.go
rm -f temp_run.go