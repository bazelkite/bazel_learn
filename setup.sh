#!/bin/bash
set -e

echo "🎯 Bazel Training Example Setup"
echo "================================"
echo ""

# Check for Bazelisk/Bazel
if command -v bazel &> /dev/null; then
    echo "✅ Bazel found: $(bazel version 2>&1 | head -1)"
elif command -v bazelisk &> /dev/null; then
    echo "✅ Bazelisk found"
    # Create symlink if it doesn't exist
    if ! command -v bazel &> /dev/null; then
        echo "   Creating 'bazel' symlink to bazelisk..."
        ln -sf $(which bazelisk) /usr/local/bin/bazel 2>/dev/null || echo "   (May need sudo)"
    fi
else
    echo "❌ Bazel/Bazelisk not found"
    echo ""
    echo "Please install Bazelisk:"
    echo "  brew install bazelisk"
    exit 1
fi

echo ""

# Check for Go
if command -v go &> /dev/null; then
    echo "✅ Go found: $(go version)"
else
    echo "⚠️  Go not found (needed for traditional pipeline)"
    echo "   Install from: https://go.dev/doc/install"
fi

echo ""
echo "🧪 Testing Bazel setup..."
echo ""

# Test Bazel query
echo "→  Ensuring Bazel installed:"
bazel version

echo ""
echo "✨ Setup complete!"
echo ""
echo "Next steps:"
echo "  1. Try: bazel test //..."
echo "  2. Try: bazel run //services/hello:hello"
echo "  3. Read: README.md for more examples"
echo "  4. Study: BAZEL_BUILDKITE_TRAINING.md for training materials"
