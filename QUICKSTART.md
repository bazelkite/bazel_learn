# Quick Start Guide

Get up and running in 5 minutes.

## Check Your Environment

We have a script which will check your environment for the required tools.

```bash
./setup
```

## 1. Install Bazelisk

```bash
# macOS
brew install bazelisk

# Linux
curl -Lo /usr/local/bin/bazelisk \
  https://github.com/bazelbuild/bazelisk/releases/latest/download/bazelisk-linux-amd64
chmod +x /usr/local/bin/bazelisk
```

Verify:

```bash
bazel version
```

## 2. Try Basic Commands

```bash
# List all targets in the workspace
bazel query //...

# Run all tests
bazel test //...

# Run a service
bazel run //services/hello:hello
```

## 3. See the Magic: Dependency Intelligence

```bash
# What would be affected if we changed lib/utils?
bazel query "rdeps(//..., //lib/utils:utils)"

# Test only what's affected by that change
bazel test $(bazel query "kind('.*_test', rdeps(//..., //lib/utils:utils))")
```

This is the foundation of intelligent CI pipelines!

## 4. Experience Caching

```bash
# First build (downloads toolchains)
time bazel build //...

# Second build (everything cached - notice the speed!)
time bazel build //...
```

## Common Commands

```bash
# Query
bazel query //...                              # List all targets
bazel query "kind('.*_test', //...)"           # Find all tests
bazel query "rdeps(//..., //lib/utils:utils)"  # What depends on utils?

# Build & Test
bazel build //...                   # Build everything
bazel test //...                    # Test everything
bazel run //services/hello:hello    # Run a service

# Clean
bazel clean                         # Clean build outputs
```

## Understanding Target Names

```text
//services/hello:hello
│         │      │
│         │      └─ Target name
│         └──────── Package path
└────────────────── Root (workspace)
```

**First build is slow?**

- This is normal - Bazel downloads toolchains and dependencies
- Subsequent builds will be much faster (cached)

## Next Steps

- **Try the interactive tutorial**: `./tutor`
- **Compare pipelines**: Look at [.buildkite/pipeline.standard.yaml](.buildkite/pipeline.standard.yaml) vs [.buildkite/pipeline.bazel.yaml](.buildkite/pipeline.bazel.yaml)
- **Full documentation**: See [README.md](README.md)
