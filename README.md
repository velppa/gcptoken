# gcptoken
An app to retrieve Google Cloud Platform Bearer auth token.

## Installation

If you have Go runtime installed:

```bash
# install dependencies
go get -u golang.org/x/oauth2
go get -u golang.org/x/oauth2/google

# install binary
go get -u github.com/pavel-popov/gcptoken
```

If not, please check [Releases](releases) for a pre-built binary for your platform.


## Usage

Assuming you have downloaded GCP credentials file to `gcp.json`:

```bash
# obtain token
GCP_SCOPE=https://www.googleapis.com/auth/cloud-platform \
GCP_ACCOUNT_CREDENTIALS=gcp.json \
gcptoken

{"access_token":"ya29.c.ElpeB109jw35HKKpi6L0KXpc54QlQsC4_Wmzw6BQ3UfVKL_iENMWX59j97Vv6KpmDT8mIGNjj9JtJnS3Z5AymAVxSTpeqauJSb8V26TxeL8l-2z3LwX7pY0teWc","token_type":"Bearer","expiry":"2019-08-08T01:49:38.852183+02:00"}
```
