#!/usr/bin/env bash
set -euo pipefail

echo "Running activebits..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/activebits.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	test1 := ActiveBits(7)
	test2 := ActiveBits(10)
	test3 := ActiveBits(15)
	test4 := ActiveBits(0)
	test5 := ActiveBits(1)

	fmt.Println("ActiveBits(7):", test1)
	fmt.Println("ActiveBits(10):", test2)
	fmt.Println("ActiveBits(15):", test3)
	fmt.Println("ActiveBits(0):", test4)
	fmt.Println("ActiveBits(1):", test5)
}
EOF

go run temp_run.go
rm -f temp_run.go