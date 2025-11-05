#!/usr/bin/env bash
set -euo pipefail

echo "Running printstr..."
echo ""

cat > temp_run.go << 'EOF'
package main
EOF

sed '/^package student/d' student/printstr.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	PrintStr("Hello World!")
}
EOF

go run temp_run.go
echo ""
rm -f temp_run.go