#!/usr/bin/env bash
set -euo pipefail

echo "Running strrev..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/strrev.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
EOF

go run temp_run.go
rm -f temp_run.go