# SDK and Terraform Generation Analysis

## Summary

✅ **Go SDK Generation**: SUCCESS - Generated correctly  
❌ **Terraform Provider Generation**: INCOMPLETE - Only examples generated, no resources or data sources

## Issue Analysis

### What Happened

1. **Go SDK Generation** worked perfectly:
   - Generated `ibm_cloud_logs_routing_v0.go`
   - Generated tests
   - All expected files created in `/Users/ianre/go/src/github.com/IBM/logs-router-go-sdk`

2. **Terraform Provider Generation** only created:
   - Examples in `examples/ibm-ibm-cloud-logs-routing/`
   - A README in `ibm/service/ibmcloudlogsrouting/README.md`
   - **NO resource files** (resource_ibm_*.go)
   - **NO data source files** (data_source_ibm_*.go)
   - **NO test files**

### Root Cause

The API definition file `/Users/ianre/go/src/github.ibm.com/cloud-api-docs/logs-router/logs-router-service-api.v1.json` is **missing required Terraform annotations**.

The generator needs these annotations on schema objects to know which schemas should become Terraform resources and data sources:
- `x-resource-name` - Marks a schema as a Terraform resource
- `x-data-source-name` - Marks a schema as a Terraform data source
- `x-codegen-config` with `go.apiPackage` - Specifies the Go SDK package to import

## Required Fixes

### 1. Add Go SDK Package Configuration

Add this to the `info` section of the API definition:

```json
{
  "info": {
    "title": "IBM Cloud Logs Routing",
    "version": "0.0.1",
    "x-codegen-config": {
      "go": {
        "apiPackage": "github.com/IBM/logs-router-go-sdk/ibmcloudlogsroutingv0"
      },
      "terraform": {
        "subcategory": "Logs Router"
      }
    }
  }
}
```

### 2. Add Resource Annotations

Find the `Tenant` schema in `components.schemas` and add `x-resource-name`:

```json
{
  "components": {
    "schemas": {
      "Tenant": {
        "type": "object",
        "description": "A tenant configuration",
        "x-resource-name": "logs_router_tenant",
        "x-terraform-include-tags": false,
        "properties": {
          "id": {
            "type": "string",
            "description": "The tenant ID",
            "readOnly": true
          },
          "name": {
            "type": "string",
            "description": "The name of the tenant"
          },
          "targets": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/Target"
            }
          }
        }
      }
    }
  }
}
```

### 3. Add Data Source Annotations

Find the `TenantCollection` schema (or create one for list operations) and add `x-data-source-name`:

```json
{
  "TenantCollection": {
    "type": "object",
    "description": "A collection of tenants",
    "x-data-source-name": "logs_router_tenants",
    "x-data-source-collection": "tenants",
    "properties": {
      "tenants": {
        "type": "array",
        "items": {
          "$ref": "#/components/schemas/Tenant"
        }
      }
    }
  }
}
```

### 4. Add Target Resource Annotations

If you want Target as a separate resource:

```json
{
  "Target": {
    "type": "object",
    "description": "A target configuration",
    "x-resource-name": "logs_router_target",
    "properties": {
      "id": {
        "type": "string",
        "readOnly": true
      },
      "name": {
        "type": "string"
      }
    }
  }
}
```

## Complete Annotation Reference

### Essential Terraform Annotations

| Annotation | Where | Purpose | Example |
|------------|-------|---------|---------|
| `x-resource-name` | Schema | Define Terraform resource | `"x-resource-name": "logs_router_tenant"` |
| `x-data-source-name` | Schema | Define Terraform data source | `"x-data-source-name": "logs_router_tenants"` |
| `x-codegen-config.go.apiPackage` | info | Go SDK package path | See above |
| `x-terraform-include-tags` | Schema | Add tags support | `"x-terraform-include-tags": false` |
| `x-data-source-collection` | Schema | Collection property name | `"x-data-source-collection": "tenants"` |

### Optional But Useful Annotations

| Annotation | Purpose |
|------------|---------|
| `x-terraform-computed` | Mark field as computed (read-only in Terraform) |
| `x-terraform-force-new` | Force resource recreation on change |
| `x-terraform-sensitive` | Mark field as sensitive |
| `x-terraform-user-account` | Mark parameter as user account ID |
| `x-resource-operations` | Explicitly specify CRUD operations |

## Example: Minimal Working API Definition

Here's what a minimal working section would look like:

```json
{
  "openapi": "3.0.0",
  "info": {
    "title": "IBM Cloud Logs Routing",
    "version": "0.0.1",
    "x-codegen-config": {
      "go": {
        "apiPackage": "github.com/IBM/logs-router-go-sdk/ibmcloudlogsroutingv0"
      },
      "terraform": {
        "subcategory": "Logs Router"
      }
    }
  },
  "components": {
    "schemas": {
      "Tenant": {
        "type": "object",
        "description": "A tenant configuration",
        "x-resource-name": "logs_router_tenant",
        "properties": {
          "id": {
            "type": "string",
            "description": "Tenant ID",
            "readOnly": true
          },
          "name": {
            "type": "string",
            "description": "Tenant name"
          },
          "crn": {
            "type": "string",
            "description": "Cloud Resource Name",
            "readOnly": true
          },
          "targets": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/Target"
            }
          }
        }
      },
      "TenantCollection": {
        "type": "object",
        "x-data-source-name": "logs_router_tenants",
        "x-data-source-collection": "tenants",
        "properties": {
          "tenants": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/Tenant"
            }
          }
        }
      }
    }
  }
}
```

## Steps to Fix

1. **Backup the current API definition**:
   ```bash
   cp /Users/ianre/go/src/github.ibm.com/cloud-api-docs/logs-router/logs-router-service-api.v1.json \
      /Users/ianre/go/src/github.ibm.com/cloud-api-docs/logs-router/logs-router-service-api.v1.json.backup
   ```

2. **Edit the API definition** to add the required annotations:
   - Add `x-codegen-config` to `info` section
   - Add `x-resource-name` to `Tenant` schema
   - Add `x-data-source-name` to list response schemas
   - Add any other needed annotations

3. **Validate the API definition**:
   ```bash
   npx @ibm-cloud/openapi-ruleset-validator \
     /Users/ianre/go/src/github.ibm.com/cloud-api-docs/logs-router/logs-router-service-api.v1.json
   ```

4. **Regenerate Terraform provider**:
   ```bash
   cd /Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router
   ./scripts/generate-terraform.sh
   ```

5. **Verify generation**:
   ```bash
   ls -la ibm/service/ibmcloudlogsrouting/
   # Should now see:
   # - resource_ibm_logs_router_tenant.go
   # - resource_ibm_logs_router_tenant_test.go
   # - data_source_ibm_logs_router_tenants.go
   # - data_source_ibm_logs_router_tenants_test.go
   ```

## Expected Output After Fix

After adding the annotations and regenerating, you should see:

```
ibm/service/ibmcloudlogsrouting/
├── README.md
├── resource_ibm_logs_router_tenant.go
├── resource_ibm_logs_router_tenant_test.go
├── data_source_ibm_logs_router_tenants.go
├── data_source_ibm_logs_router_tenants_test.go
└── (other generated files)
```

## Documentation References

- [Terraform Provider Generation Wiki](https://github.ibm.com/CloudEngineering/openapi-sdkgen/wiki/Terraform-Provider-Generation)
- [API Annotations for Terraform](https://github.ibm.com/CloudEngineering/openapi-sdkgen/wiki/API-Document-annotations-for-Terraform-&-Ansible-generation)
- [Config Options](https://github.ibm.com/CloudEngineering/openapi-sdkgen/wiki/Config-Options)

## Notes

- The Go SDK generated successfully because it doesn't require these special annotations
- Terraform generation requires explicit marking of which schemas should become resources/data sources
- This is by design - not all API schemas should become Terraform resources
- The generator can't guess which schemas are resources vs. internal models

## Next Steps

1. Add the required annotations to the API definition
2. Regenerate the Terraform provider
3. Test the generated code
4. Iterate on annotations as needed (computed fields, force-new, etc.)

---

**Generated**: 2026-03-16  
**Status**: Awaiting API definition updates