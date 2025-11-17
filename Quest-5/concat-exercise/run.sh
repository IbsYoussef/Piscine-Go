#!/usr/bin/env bash
set -euo pipefail

echo "Running concat..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/concat.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(Concat("Hello!", " How are you?"))
}
EOF

go run temp_run.go
rm -f temp_run.go