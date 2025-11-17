#!/usr/bin/env bash
set -euo pipefail

echo "Running printnbrinorder..."
echo ""

cat > temp_run.go << 'EOF'
package main
EOF

sed '/^package student/d' student/printnbrinorder.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}
EOF

go run temp_run.go
echo ""
rm -f temp_run.go