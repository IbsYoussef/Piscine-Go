#!/usr/bin/env bash
set -euo pipefail

echo "Running alphacount..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/alphacount.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}
EOF

go run temp_run.go
rm -f temp_run.go