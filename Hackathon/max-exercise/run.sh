#!/usr/bin/env bash
set -euo pipefail

echo "Running max..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/max.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	test1 := []int{23, 123, 1, 11, 55, 93}
	test2 := []int{-5, -10, -3, -20}
	test3 := []int{42}
	test4 := []int{}
	test5 := []int{0, -1, 5, -10, 100}

	fmt.Println("Max([23, 123, 1, 11, 55, 93]):", Max(test1))
	fmt.Println("Max([-5, -10, -3, -20]):", Max(test2))
	fmt.Println("Max([42]):", Max(test3))
	fmt.Println("Max([]):", Max(test4))
	fmt.Println("Max([0, -1, 5, -10, 100]):", Max(test5))
}
EOF

go run temp_run.go
rm -f temp_run.go