#!/usr/bin/env bash
set -euo pipefail

echo "Running isalpha..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/isalpha.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
}
EOF

go run temp_run.go
rm -f temp_run.go