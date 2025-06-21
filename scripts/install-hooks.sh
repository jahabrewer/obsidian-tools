#!/bin/bash
set -e

echo "Installing git hooks..."

# Create the pre-push hook
cat > .git/hooks/pre-push << 'EOF'
#!/bin/bash
set -e

echo "🔍 Running pre-push checks..."

# Run all checks
make pre-push

echo "✅ Pre-push checks completed successfully!"
EOF

# Make the hook executable
chmod +x .git/hooks/pre-push

echo "✅ Git hooks installed successfully!"
echo ""
echo "The pre-push hook will now run:"
echo "  - go fmt (code formatting)"
echo "  - golangci-lint (linting)"
echo "  - go test (tests)"
echo ""
echo "To skip the hook (not recommended), use: git push --no-verify" 