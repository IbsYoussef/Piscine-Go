#!/usr/bin/env bash
set -euo pipefail

echo "Running index..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/index.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}
EOF

go run temp_run.go
rm -f temp_run.go