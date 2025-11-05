#!/usr/bin/env bash
set -euo pipefail

echo "Running basicatoi..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/basicatoi.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
EOF

go run temp_run.go
rm -f temp_run.go