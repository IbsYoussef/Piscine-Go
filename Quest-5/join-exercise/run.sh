#!/usr/bin/env bash
set -euo pipefail

echo "Running join..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/join.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}
EOF

go run temp_run.go
rm -f temp_run.go