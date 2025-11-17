#!/usr/bin/env bash
set -euo pipefail

echo "Running isprintable..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/isprintable.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))
}
EOF

go run temp_run.go
rm -f temp_run.go