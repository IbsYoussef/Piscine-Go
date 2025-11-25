#!/usr/bin/env bash
set -euo pipefail

echo "Running appendrange..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/appendrange.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
EOF

go run temp_run.go
rm -f temp_run.go