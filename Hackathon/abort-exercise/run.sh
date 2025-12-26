#!/usr/bin/env bash
set -euo pipefail

echo "Running abort..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/abort.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	test1 := Abort(2, 3, 8, 5, 7)
	test2 := Abort(1, 2, 3, 4, 5)
	test3 := Abort(9, 1, 5, 3, 7)

	fmt.Println("Test 1 (2,3,8,5,7):", test1)
	fmt.Println("Test 2 (1,2,3,4,5):", test2)
	fmt.Println("Test 3 (9,1,5,3,7):", test3)
}
EOF

go run temp_run.go
rm -f temp_run.go