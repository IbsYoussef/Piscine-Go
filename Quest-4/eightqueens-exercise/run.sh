#!/usr/bin/env bash
set -euo pipefail

echo "Running eightqueens..."
echo ""
echo "This will print all 92 solutions (may take a moment)..."
echo ""

cat > temp_run.go << 'EOF'
package main
EOF

# Copy the entire student file (including its imports)
sed '/^package student/d' student/eightqueens.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	EightQueens()
}
EOF

go run temp_run.go
rm -f temp_run.go