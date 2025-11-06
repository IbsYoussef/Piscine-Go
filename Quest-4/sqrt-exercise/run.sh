#!/usr/bin/env bash
set -euo pipefail

echo "Running sqrt..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/sqrt.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}
EOF

go run temp_run.go
rm -f temp_run.go