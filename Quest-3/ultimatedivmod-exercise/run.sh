#!/usr/bin/env bash
set -euo pipefail

echo "Running ultimatedivmod..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/ultimatedivmod.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
EOF

go run temp_run.go
rm -f temp_run.go