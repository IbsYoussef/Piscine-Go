#!/usr/bin/env bash
set -euo pipefail

echo "Running maxwordcountn..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/maxwordcountn.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	result := MaxWordCountN(`
	Orange is the sun sliding to the horizon after a summer day. Orange is the sound of dribbling basketball. Orange is the smell of a tiger lily petal. Orange is the taste of thirst-quenching Nehi Soda. Orange is the color of peach marmalade on a side of toast. Orange is the sound of a carrot popping out of the ground.
	`, 3)
	
	fmt.Println(result)
}
EOF

go run temp_run.go
rm -f temp_run.go