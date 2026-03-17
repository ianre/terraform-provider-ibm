I've investigated both issues with the SDK/Terraform generation. Here are my findings and solutions:

## Issue 1: Regional URL Support - "service does not support regional URLs"

**Root Cause:**
The API definition uses **server variables** (line 35-54 in the API definition) with `{region}` placeholder, but the SDK generator is NOT recognizing this as regional endpoint support. Instead, it's treating it as a parameterized URL.

**Current Behavior:**
- The generated SDK has `GetServiceURLForRegion()` returning an error: "service does not support regional URLs" (line 128 in the SDK)
- However, it DOES have `ConstructServiceURL()` function (line 142) that accepts region variables
- The API definition correctly defines regional URLs with the `{region}` variable

**Why This Happens:**
According to the wiki documentation I reviewed:
- **Regional endpoints** require the `x-ibm-region` extension on each server object
- **Server variables** (what you're using) create a `ConstructServiceURL()` function instead
- These are two DIFFERENT features!

**Solution:**
You need to add `x-ibm-region` extensions to your API definition. Modify the `servers` section in `/Users/ianre/go/src/github.ibm.com/cloud-api-docs/logs-router/logs-router-service-api.v1.json`:

```json
"servers": [
  {
    "url": "https://management.private.au-syd.logs-router.cloud.ibm.com/v1",
    "description": "Australia Sydney region",
    "x-ibm-region": "au-syd"
  },
  {
    "url": "https://management.private.br-sao.logs-router.cloud.ibm.com/v1",
    "description": "Brazil Sao Paulo region",
    "x-ibm-region": "br-sao"
  },
  {
    "url": "https://management.private.ca-tor.logs-router.cloud.ibm.com/v1",
    "description": "Canada Toronto region",
    "x-ibm-region": "ca-tor"
  },
  {
    "url": "https://management.private.eu-de.logs-router.cloud.ibm.com/v1",
    "description": "Germany Frankfurt region",
    "x-ibm-region": "eu-de"
  },
  {
    "url": "https://management.private.eu-es.logs-router.cloud.ibm.com/v1",
    "description": "Spain Madrid region",
    "x-ibm-region": "eu-es"
  },
  {
    "url": "https://management.private.eu-gb.logs-router.cloud.ibm.com/v1",
    "description": "United Kingdom London region",
    "x-ibm-region": "eu-gb"
  },
  {
    "url": "https://management.private.jp-tok.logs-router.cloud.ibm.com/v1",
    "description": "Japan Tokyo region",
    "x-ibm-region": "jp-tok"
  },
  {
    "url": "https://management.private.jp-osa.logs-router.cloud.ibm.com/v1",
    "description": "Japan Osaka region",
    "x-ibm-region": "jp-osa"
  },
  {
    "url": "https://management.private.us-south.logs-router.cloud.ibm.com/v1",
    "description": "US South region",
    "x-ibm-region": "us-south"
  },
  {
    "url": "https://management.private.us-east.logs-router.cloud.ibm.com/v1",
    "description": "US East region",
    "x-ibm-region": "us-east"
  }
]
```

This will generate a proper `GetServiceURLForRegion()` function with a map of regions to URLs.

## Issue 2: "IBM" vs "Ibm" Capitalization

**Root Cause:**
The SDK generator has default acronym handling that capitalizes common acronyms like HTTP, URL, ID, etc., but "IBM" is NOT in the default list (see line 81 in IBMDefaultCodegen documentation).

**Solution:**
Add the `x-acronyms` extension to your API definition's `info` section:

```json
"info": {
  "title": "IBM Cloud Logs Routing",
  "description": "IBM Cloud Logs Routing is an IBM cloud platform service...",
  "version": "0.0.1",
  "x-acronyms": ["ibm"],
  "x-codegen-config": {
    ...
  }
}
```

According to the Config-Options.md wiki (line 842-849):
- `x-acronyms` provides custom acronyms for a service
- Each acronym should be lowercase
- Type: string array
- Scope: info level

This will ensure "IBM" is treated as an acronym and capitalized properly throughout the generated code (e.g., "IBM" instead of "Ibm").

## Next Steps:

1. Update the API definition with both changes
2. Run `./scripts/generate-go-sdk.sh` to regenerate the SDK
3. Run `./scripts/generate-terraform.sh` to regenerate Terraform provider
4. Verify that:
   - `GetServiceURLForRegion()` now returns proper regional URLs
   - "IBM" is capitalized correctly throughout the code