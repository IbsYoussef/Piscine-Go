#!/usr/bin/env bash
set -euo pipefail

echo "Running split..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/split.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
EOF

go run temp_run.go
rm -f temp_run.go