#!/usr/bin/env bash
set -euo pipefail

echo "Running unmatch..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/unmatch.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	test1 := []int{1, 2, 3, 1, 2, 3, 4}
	test2 := []int{5, 5, 7, 7}
	test3 := []int{1, 1, 1}
	test4 := []int{9}
	test5 := []int{2, 3, 2, 3, 5, 5, 8}

	fmt.Println("Unmatch([1,2,3,1,2,3,4]):", Unmatch(test1))
	fmt.Println("Unmatch([5,5,7,7]):", Unmatch(test2))
	fmt.Println("Unmatch([1,1,1]):", Unmatch(test3))
	fmt.Println("Unmatch([9]):", Unmatch(test4))
	fmt.Println("Unmatch([2,3,2,3,5,5,8]):", Unmatch(test5))
}
EOF

go run temp_run.go
rm -f temp_run.go