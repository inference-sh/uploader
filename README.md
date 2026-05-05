# uploader

Lightweight CLI for uploading files to S3-compatible storage. Single binary, no dependencies.

Works with AWS S3, Cloudflare R2, MinIO, DigitalOcean Spaces, Backblaze B2, and any S3-compatible service.

## Install

**Go:**
```bash
go install github.com/inference-sh/uploader@latest
```

**GitHub Releases:**

Download the binary for your platform from [Releases](https://github.com/inference-sh/uploader/releases/latest), then:

```bash
chmod +x uploader-*
sudo mv uploader-* /usr/local/bin/uploader
```

## Upload

```bash
uploader s3 \
  --source ./build/artifact.tar.gz \
  --dest releases/v1.0.0/artifact.tar.gz \
  --endpoint https://your-endpoint.r2.cloudflarestorage.com \
  --region auto \
  --access-key-id YOUR_KEY \
  --access-key-secret YOUR_SECRET \
  --bucket my-bucket \
  --url https://dist.example.com
```

All S3 flags have env var equivalents — set them once, then just pass `--source` and `--dest`:

| Flag | Env var |
|------|---------|
| `--endpoint` | `S3_ENDPOINT` |
| `--region` | `S3_REGION` |
| `--access-key-id` | `S3_ACCESS_KEY_ID` |
| `--access-key-secret` | `S3_SECRET_ACCESS_KEY` |
| `--bucket` | `S3_BUCKET` |
| `--url` | `S3_URL` |

```bash
export S3_ENDPOINT=https://your-endpoint.r2.cloudflarestorage.com
export S3_REGION=auto
export S3_ACCESS_KEY_ID=YOUR_KEY
export S3_SECRET_ACCESS_KEY=YOUR_SECRET
export S3_BUCKET=my-bucket
export S3_URL=https://dist.example.com

uploader s3 --source ./file.tar.gz --dest path/to/file.tar.gz
```

Flags always take precedence over env vars.

## Download

```bash
uploader download --url https://dist.example.com/file.tar.gz --dest ./file.tar.gz
```

Downloads are atomic — writes to a temp file first, then moves to the destination on success.

## CI example

```yaml
- name: Upload to R2
  env:
    S3_ENDPOINT: ${{ secrets.S3_ENDPOINT }}
    S3_ACCESS_KEY_ID: ${{ secrets.S3_ACCESS_KEY_ID }}
    S3_SECRET_ACCESS_KEY: ${{ secrets.S3_SECRET_ACCESS_KEY }}
    S3_BUCKET: ${{ secrets.S3_BUCKET }}
    S3_URL: ${{ secrets.S3_URL }}
    S3_REGION: auto
  run: |
    uploader s3 --source ./build/artifact.tar.gz --dest releases/artifact.tar.gz
```

## How it works

Generates a presigned PUT URL via the S3 API, then uploads directly over HTTP. This works with providers like Cloudflare R2 that don't support POST multipart uploads.

## License

[MIT](LICENSE)
