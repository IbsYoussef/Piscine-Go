#!/usr/bin/env bash
set -euo pipefail

echo "Running isprime..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/isprime.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}
EOF

go run temp_run.go
rm -f temp_run.go