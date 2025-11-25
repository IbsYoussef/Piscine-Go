#!/usr/bin/env bash
set -euo pipefail

echo "Running splitwhitespaces..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/splitwhitespaces.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
EOF

go run temp_run.go
rm -f temp_run.go