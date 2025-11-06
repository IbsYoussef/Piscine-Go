#!/usr/bin/env bash
set -euo pipefail

echo "Running findnextprime..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/findnextprime.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
}
EOF

go run temp_run.go
rm -f temp_run.go