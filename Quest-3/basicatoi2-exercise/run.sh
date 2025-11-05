#!/usr/bin/env bash
set -euo pipefail

echo "Running basicatoi2..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/basicatoi2.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}
EOF

go run temp_run.go
rm -f temp_run.go