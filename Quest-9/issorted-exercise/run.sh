#!/usr/bin/env bash
set -euo pipefail

echo "Running issorted..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"

func compare(a, b int) int {
	if a > b {
		return 1
	}
	if a == b {
		return 0
	}
	return -1
}
EOF

sed '/^package student/d' student/issorted.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	a3 := []int{5, 4, 3, 2, 1}
	a4 := []int{1, 1, 2, 2, 3}

	result1 := IsSorted(compare, a1)
	result2 := IsSorted(compare, a2)
	result3 := IsSorted(compare, a3)
	result4 := IsSorted(compare, a4)

	fmt.Println("Test 1 - Ascending [0,1,2,3,4,5]:", result1)
	fmt.Println("Test 2 - Mixed [0,2,1,3]:", result2)
	fmt.Println("Test 3 - Descending [5,4,3,2,1]:", result3)
	fmt.Println("Test 4 - Duplicates [1,1,2,2,3]:", result4)
}
EOF

go run temp_run.go
rm -f temp_run.go