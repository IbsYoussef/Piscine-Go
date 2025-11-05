#!/usr/bin/env bash

echo "Running your solution:"
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/swap.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
EOF

go run temp_run.go
rm -f temp_run.go