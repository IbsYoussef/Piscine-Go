#!/usr/bin/env bash
set -euo pipefail

echo "Running recursivepower..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/recursivepower.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(RecursivePower(4, 3))
}
EOF

go run temp_run.go
rm -f temp_run.go