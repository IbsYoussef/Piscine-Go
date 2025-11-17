#!/usr/bin/env bash
set -euo pipefail

echo "Running isnumeric..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/isnumeric.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}
EOF

go run temp_run.go
rm -f temp_run.go