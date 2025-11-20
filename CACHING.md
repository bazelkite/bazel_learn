# Bazel Remote Caching with bazel-remote

## Why Use Remote Caching?

Without remote caching, every Buildkite agent starts from scratch. With remote caching:
- ✅ Builds are **5-10x faster** on CI
- ✅ Agents share build artifacts
- ✅ Rebuilds only what changed

## Quick Start (5 minutes)

### Step 1: Deploy bazel-remote

On any server your Buildkite agents can reach:

```bash
docker run -d \
  --name bazel-cache \
  -v /mnt/bazel-cache:/data \
  -p 9092:9092 \
  --restart unless-stopped \
  buchgr/bazel-remote-cache:latest \
  --max_size 100
```

That's it! You now have a cache server running on port 9092.

### Step 2: Configure Bazel

Create `.bazelrc.remote` in your repo:

```bash
# Remote cache configuration
build:remote --remote_cache=grpc://your-cache-server.com:9092
build:remote --remote_timeout=60s
build:remote --remote_download_minimal
```

### Step 3: Update Buildkite Pipeline

Add `--config=remote` to your bazel commands:

```yaml
steps:
  - label: ":bazel: Build & Test"
    command: |
      bazel test --config=remote //... --test_output=errors
```

Done! Your builds now use the remote cache.

## How It Works

```
┌─────────────┐         ┌─────────────┐         ┌─────────────┐
│  Agent 1    │────────▶│ bazel-remote│◀────────│  Agent 2    │
│             │  cache  │   server    │  cache  │             │
│ Build job   │  hit/   │             │  hit/   │ Test job    │
│             │  miss   │ 100GB disk  │  miss   │             │
└─────────────┘         └─────────────┘         └─────────────┘
```

**First build:** Agent uploads build outputs to cache (slower)  
**Subsequent builds:** Agent downloads from cache instead of rebuilding (faster)

## Production Tips

### Security (Optional)

Add basic auth with htpasswd:

```bash
# Generate password file
htpasswd -c .htpasswd buildkite

# Run with auth
docker run -d \
  --name bazel-cache \
  -v /mnt/bazel-cache:/data \
  -v $(pwd)/.htpasswd:/etc/bazel-remote/htpasswd:ro \
  -p 9092:9092 \
  buchgr/bazel-remote-cache:latest \
  --max_size 100 \
  --htpasswd_file=/etc/bazel-remote/htpasswd
```

Update `.bazelrc.remote`:
```bash
build:remote --remote_cache=grpc://username:password@your-cache-server.com:9092
```

### Use S3/GCS Backend (Optional)

For unlimited storage, proxy to S3:

```bash
docker run -d \
  --name bazel-cache \
  -v /mnt/bazel-cache:/data \
  -p 9092:9092 \
  buchgr/bazel-remote-cache:latest \
  --max_size 20 \
  --s3.bucket=my-bazel-cache \
  --s3.region=us-east-1 \
  --s3.auth_method=iam_role
```

Local disk acts as a fast tier, S3 as backup.

### Monitor Cache Performance

Check cache status:
```bash
curl http://your-cache-server.com:8080/status
```

Returns:
```json
{
  "CurrSize": 45000000000,
  "MaxSize": 100000000000,
  "NumFiles": 50000,
  "ServerTime": 1234567890
}
```

## Resources

- [bazel-remote GitHub](https://github.com/buchgr/bazel-remote)
- [Bazel Remote Caching Docs](https://bazel.build/remote/caching)
