#!/usr/bin/env bash
set -euo pipefail

echo "Running advancedsortwordarr..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"

// Compare function for testing
func Compare(a, b string) int {
	if a > b {
		return 1
	}
	if a == b {
		return 0
	}
	return -1
}
EOF

sed '/^package student/d' student/advancedsortwordarr.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	test1 := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	test2 := []string{"z", "y", "x", "w"}
	test3 := []string{"hello", "world", "go", "programming"}

	AdvancedSortWordArr(test1, Compare)
	AdvancedSortWordArr(test2, Compare)
	AdvancedSortWordArr(test3, Compare)

	fmt.Println("Test 1:", test1)
	fmt.Println("Test 2:", test2)
	fmt.Println("Test 3:", test3)
}
EOF

go run temp_run.go
rm -f temp_run.go