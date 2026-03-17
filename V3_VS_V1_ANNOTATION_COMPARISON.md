# V3 vs V1 API Annotation Comparison

## Summary
After analyzing the V3 API definition (`logs-router-v3.yaml`) and comparing it to the V1 API definition (`logs-router-service-api.v1.json`), I've identified the correct Terraform annotation patterns used by IBM Cloud.

## Key Findings

### 1. Annotation Naming Convention

**V3 API uses these annotation names:**
- `x-terraform-resource-name` - Defines the Terraform resource name
- `x-terraform-resource-operations` - Maps CRUD operations to API methods
- `x-terraform-resource-id` - Specifies the resource ID field (optional, for special cases)
- `x-terraform-datasource-name` - Defines the Terraform data source name
- `x-terraform-datasource-filter` - Specifies the filter field for data sources

**V1 API currently has NO Terraform annotations** - This is why Terraform generation only created examples/README.

### 2. V3 API Structure

The V3 API has proper annotations on three main schemas:

#### Target Resource (lines 2013-2018)
```yaml
x-terraform-resource-name: logs_router_target
x-terraform-resource-operations:
  create: create_target
  read: get_target
  update: update_target
  delete: delete_target
```

#### Route Resource (line 2344)
```yaml
x-terraform-resource-name: logs_router_route
# Note: No explicit operations mapping - uses default CRUD operation names
```

#### Settings Resource (lines 2448-2454)
```yaml
x-terraform-resource-id: primary_metadata_region
x-terraform-resource-name: logs_router_settings
x-terraform-resource-operations:
  create: update_settings
  delete: update_settings
  read: get_settings
  update: update_settings
```

#### Targets Data Source (lines 2080-2081)
```yaml
x-terraform-datasource-name: logs_router_targets
x-terraform-datasource-filter: name
```

### 3. V1 API Structure

The V1 API has a different schema structure:
- Uses `Tenant` instead of `Target`
- Has different operation names (e.g., `update_tenant` vs `update_target`)
- Currently has NO Terraform annotations at all

### 4. Code Generation Configuration

**V3 API (lines 21-30):**
```yaml
x-codegen-config:
  improvedNameFormattingV2: true
  go:
    apiPackage: github.com/IBM/platform-services-go-sdk
  java:
    apiPackage: com.ibm.cloud.platform_services
  python:
    apiPackage: ibm_platform_services
  terraform:
    apiEndpoint: IBMCLOUD_LOGS_ROUTING_API_ENDPOINT
```

**V1 API (lines 15-22):**
```json
"x-codegen-config": {
  "go": {
    "apiPackage": "github.com/IBM/logs-router-go-sdk/ibmcloudlogsroutingv0"
  },
  "terraform": {
    "subcategory": "Logs Router"
  }
}
```

## Differences Between V3 and V1

| Aspect | V3 API | V1 API |
|--------|--------|--------|
| **Format** | YAML | JSON |
| **Main Resource** | Target | Tenant |
| **SDK Package** | platform-services-go-sdk (shared) | logs-router-go-sdk (standalone) |
| **Terraform Annotations** | ✅ Present | ❌ Missing |
| **Resource Operations** | Explicitly mapped | Not defined |
| **Data Sources** | ✅ Defined | ❌ Missing |
| **API Endpoint Config** | Custom env var | Not specified |

## Required Changes for V1 API

To make V1 API work with Terraform generation, we need to add:

1. **Tenant Resource Annotations:**
```json
"x-terraform-resource-name": "logs_router_tenant",
"x-terraform-resource-operations": {
  "create": "create_tenant",
  "read": "get_tenant_detail",
  "update": "update_tenant",
  "delete": "delete_tenant"
}
```

2. **Tenant Collection Data Source:**
```json
"x-terraform-datasource-name": "logs_router_tenants",
"x-terraform-datasource-filter": "name"
```

3. **Update x-codegen-config:**
```json
"terraform": {
  "subcategory": "Logs Router",
  "apiEndpoint": "IBMCLOUD_LOGS_ROUTING_API_ENDPOINT"
}
```

## Annotation Pattern Rules

Based on V3 analysis:

1. **Resource annotations go on the main schema** (e.g., `Tenant`, `Target`, `Route`)
2. **Data source annotations go on collection schemas** (e.g., `TargetCollection`)
3. **Operation mapping is optional** - if not provided, generator uses default names
4. **Resource ID field is optional** - only needed for special cases (like Settings)
5. **Filter field for data sources** - specifies which field to filter by

## Next Steps

1. ✅ Update the `add-terraform-annotations.sh` script to use correct annotation names
2. ⏳ Add annotations to V1 API definition
3. ⏳ Regenerate Terraform provider
4. ⏳ Verify resource and data source files are created
5. ⏳ Test the generated Terraform code

## Script Updates Needed

The current script uses incorrect annotation names:
- ❌ `x-resource-name` → ✅ `x-terraform-resource-name`
- ❌ `x-data-source-name` → ✅ `x-terraform-datasource-name`
- ❌ `x-data-source-collection` → ✅ `x-terraform-datasource-filter`

Additionally, need to add:
- `x-terraform-resource-operations` mapping
- `apiEndpoint` to terraform config