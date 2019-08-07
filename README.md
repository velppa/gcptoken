# gcptoken
Simple token to retrieve Google Cloud Platform Bearer auth token


## Usage

Assuming you have downloaded GCP credentials file to `gcp.json`:

```bash
go install github.com/pavel-popov/gcptoken
GCP_SCOPE=https://www.googleapis.com/auth/cloud-platform \
  GCP_ACCOUNT_CREDENTIALS=gcp.json \
  gcptoken
```
