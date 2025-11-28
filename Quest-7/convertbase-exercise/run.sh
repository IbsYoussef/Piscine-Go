#!/usr/bin/env bash
set -euo pipefail

echo "Running convertbase..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/convertbase.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	result := ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
}
EOF

go run temp_run.go
rm -f temp_run.go