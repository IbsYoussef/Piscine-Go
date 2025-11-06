#!/usr/bin/env bash
set -euo pipefail

echo "Running iterativepower..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/iterativepower.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(IterativePower(4, 3))
}
EOF

go run temp_run.go
rm -f temp_run.go