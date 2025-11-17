#!/usr/bin/env bash
set -euo pipefail

echo "Running islower..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/islower.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))
}
EOF

go run temp_run.go
rm -f temp_run.go