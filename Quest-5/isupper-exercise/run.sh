#!/usr/bin/env bash
set -euo pipefail

echo "Running isupper..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/isupper.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
}
EOF

go run temp_run.go
rm -f temp_run.go