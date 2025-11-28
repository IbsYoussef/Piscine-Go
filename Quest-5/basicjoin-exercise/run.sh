#!/usr/bin/env bash
set -euo pipefail

echo "Running basicjoin..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/basicjoin.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}
EOF

go run temp_run.go
rm -f temp_run.go