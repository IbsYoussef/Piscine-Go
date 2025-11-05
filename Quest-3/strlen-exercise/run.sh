#!/usr/bin/env bash
set -euo pipefail

echo "Running strlen..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/strlen.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
EOF

go run temp_run.go
rm -f temp_run.go