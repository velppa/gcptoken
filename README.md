# gcptoken
Simple token to retrieve Google Cloud Platform Bearer auth token


## Usage

Assuming you have downloaded GCP credentials file to `gcp.json`:

```bash

# install dependencies
go get -u golang.org/x/oauth2
go get -u golang.org/x/oauth2/google

# install binary
go get -u github.com/pavel-popov/gcptoken

# obtain token
GCP_SCOPE=https://www.googleapis.com/auth/cloud-platform \
GCP_ACCOUNT_CREDENTIALS=gcp.json \
gcptoken
```
