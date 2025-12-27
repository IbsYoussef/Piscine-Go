#!/usr/bin/env bash
set -euo pipefail

echo "Running itoa..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/itoa.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	test1 := Itoa(12345)
	test2 := Itoa(0)
	test3 := Itoa(-1234)
	test4 := Itoa(987654321)
	test5 := Itoa(-999)

	fmt.Println("Itoa(12345):", test1)
	fmt.Println("Itoa(0):", test2)
	fmt.Println("Itoa(-1234):", test3)
	fmt.Println("Itoa(987654321):", test4)
	fmt.Println("Itoa(-999):", test5)
}
EOF

go run temp_run.go
rm -f temp_run.go