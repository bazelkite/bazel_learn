# Bazel + Buildkite Training

A hands-on training repository demonstrating how Bazel and Buildkite work together to create intelligent, scalable CI/CD pipelines.

## 🚀 Quick Start

**New to Bazel?** Run the interactive tutorial:

```bash
./tutor
```

Or head over to our [quickstart guide](QUICKSTART.md)

This 10-minute guided session teaches you Bazel fundamentals through hands-on exercises.

## What You'll Learn

This repository demonstrates the difference between traditional CI and Bazel-powered CI:

- **Traditional approach**: Manually defined steps, rebuilds everything on every change
- **Bazel approach**: Dynamic pipelines that only build/test what's affected by changes

## Prerequisites

Install Bazelisk (Bazel version manager):

```bash
# macOS
brew install bazelisk

# Linux
curl -Lo /usr/local/bin/bazelisk \
  https://github.com/bazelbuild/bazelisk/releases/latest/download/bazelisk-linux-amd64
chmod +x /usr/local/bin/bazelisk
```

Verify installation:

```bash
bazel version
```

## Try It Yourself

### Basic Commands

```bash
# List all targets in the workspace
bazel query //...

# Run all tests
bazel test //...

# Build and run a service
bazel run //services/hello:hello
```

### The Magic: Smart Testing

See what's affected when you change a library:

```bash
# What depends on the utils library?
bazel query "rdeps(//..., //lib/utils:utils)"

# Test only what's affected
bazel test $(bazel query "kind('.*_test', rdeps(//..., //lib/utils:utils))")
```

This is how Bazel enables intelligent CI pipelines!

## Project Structure

```text
.
├── services/
│   ├── hello/          # Example microservice
│   └── goodbye/        # Another microservice
├── lib/
│   └── utils/          # Shared library
└── .buildkite/
    ├── pipeline.yaml         # Traditional (manual steps)
    └── pipeline.bazel.yaml   # Bazel-powered (dynamic)
```

## Compare the Pipelines

### Traditional Pipeline

```yaml
# Hardcoded steps - manual maintenance required
steps:
  - label: "Test Hello"
    command: "cd services/hello && go test ./..."
  - label: "Test Goodbye"
    command: "cd services/goodbye && go test ./..."
  # ... manually list every service
```

**Problems**: No caching, tests everything, high maintenance

### Bazel Pipeline

```yaml
# Discovers what to test dynamically
steps:
  - label: "Generate Dynamic Pipeline"
    command: |
      # Ask Bazel what changed and what's affected
      AFFECTED=$(bazel query "rdeps(//..., set($CHANGED_FILES))")
      # Generate steps for only affected targets
      for target in $AFFECTED; do
        echo "  - command: bazel test $target"
      done | buildkite-agent pipeline upload
```

**Benefits**: Intelligent caching, selective testing, zero maintenance

## Key Takeaways

1. **Bazel knows your dependency graph** - it can tell you exactly what's affected by any change
2. **Use `bazel query` to generate dynamic pipelines** - only build/test what matters
3. **Remote caching is critical** - share build artifacts across your team's CI runs
4. **Buildkite's unlimited concurrency** pairs perfectly with Bazel's parallel execution

## Next Steps

- **Run the tutorial**: `./tutor` (recommended starting point)
- **Compare pipelines**: Check `.buildkite/pipeline.yaml` vs `.buildkite/pipeline.bazel.yaml`
- **Learn more**: See [Bazel docs](https://bazel.build/) for deeper concepts

## Common Commands Reference

```bash
# Query
bazel query //...                              # List all targets
bazel query "kind('.*_test', //...)"           # Find all tests
bazel query "rdeps(//..., //lib/utils:utils)"  # What depends on utils?

# Build & Test
bazel build //...                    # Build everything
bazel test //...                     # Test everything
bazel run //services/hello:hello     # Run a service

# Clean
bazel clean                          # Clean build outputs
```
