#!/usr/bin/env bash
set -euo pipefail

echo "Running sortwordarr..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/sortwordarr.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	test1 := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	test2 := []string{"z", "y", "x", "w"}
	test3 := []string{"hello", "world", "go", "programming"}

	SortWordArr(test1)
	SortWordArr(test2)
	SortWordArr(test3)

	fmt.Println("Test 1:", test1)
	fmt.Println("Test 2:", test2)
	fmt.Println("Test 3:", test3)
}
EOF

go run temp_run.go
rm -f temp_run.go