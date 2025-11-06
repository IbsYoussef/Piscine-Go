#!/usr/bin/env bash
set -euo pipefail

echo "Running fibonacci..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/fibonacci.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	arg1 := 4
	fmt.Println(Fibonacci(arg1))
}
EOF

go run temp_run.go
rm -f temp_run.go