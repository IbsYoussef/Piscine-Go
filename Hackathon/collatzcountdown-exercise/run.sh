#!/usr/bin/env bash
set -euo pipefail

echo "Running collatzcountdown..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/collatzcountdown.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	test1 := CollatzCountdown(12)
	test2 := CollatzCountdown(5)
	test3 := CollatzCountdown(1)
	test4 := CollatzCountdown(0)
	test5 := CollatzCountdown(27)

	fmt.Println("Test 1 (12):", test1)
	fmt.Println("Test 2 (5):", test2)
	fmt.Println("Test 3 (1):", test3)
	fmt.Println("Test 4 (0):", test4)
	fmt.Println("Test 5 (27):", test5)
}
EOF

go run temp_run.go
rm -f temp_run.go