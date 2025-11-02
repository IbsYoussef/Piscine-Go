#!/usr/bin/env bash

echo "Running your solution (testing with -1, 0, 1):"
echo ""

# Create test for student (same approach as test.sh)
cat > temp_run.go << 'EOF'
package main
EOF

# Add student code without package declaration
sed 's/package student//' student/isnegative.go >> temp_run.go

# Add main function to test
cat >> temp_run.go << 'EOF'

func main() {
	IsNegative(-1)
	IsNegative(0)
	IsNegative(1)
}
EOF

# Run it
go run temp_run.go

# Cleanup
rm -f temp_run.go