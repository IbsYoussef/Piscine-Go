#!/usr/bin/env bash
set -euo pipefail

echo "Running compact..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/compact.go >> temp_run.go

cat >> temp_run.go << 'EOF'

const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	fmt.Println("=== Before Compacting ===")
	for _, v := range a {
		fmt.Println(v)
	}

	size := Compact(&a)
	
	fmt.Println()
	fmt.Println("Size after compacting:", size)
	fmt.Println()
	fmt.Println("=== After Compacting ===")
	for _, v := range a {
		fmt.Println(v)
	}
}
EOF

go run temp_run.go
rm -f temp_run.go