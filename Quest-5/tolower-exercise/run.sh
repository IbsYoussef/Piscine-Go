#!/usr/bin/env bash
set -euo pipefail

echo "Running tolower..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/tolower.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}
EOF

go run temp_run.go
rm -f temp_run.go