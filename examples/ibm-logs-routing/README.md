# Examples for IBM Cloud Logs Routing

These examples illustrate how to use the resources and data sources associated with IBM Cloud Logs Routing.

The following resources are supported:
* ibm_logs-router_tenant

The following data sources are supported:
* ibm_logs-router_tenants
* ibm_logs-router_targets

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## IBM Cloud Logs Routing resources

### Resource: ibm_logs-router_tenant

```hcl
resource "ibm_logs-router_tenant" "logs-router_tenant_instance" {
  ibm_api_version = var.logs-router_tenant_ibm_api_version
  name = var.logs-router_tenant_name
  targets = var.logs-router_tenant_targets
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| ibm_api_version | Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version. | `string` | true |
| name | The name for this tenant. The name is regionally unique across all tenants in the account. | `string` | true |
| targets | List of targets. | `list()` | true |

#### Outputs

| Name | Description |
|------|-------------|
| created_at | Timestamp the tenant was originally created. |
| updated_at | Timestamp the tenant was last updated. |
| crn | Cloud resource name of the tenant. Must be a valid CRN. |
| etag | Resource version identifier. |
| write_status | The status of the write attempt to the target with the provided endpoint parameters. |

## IBM Cloud Logs Routing data sources

### Data source: ibm_logs-router_tenants

```hcl
data "ibm_logs-router_tenants" "logs-router_tenants_instance" {
  ibm_api_version = var.logs-router_tenants_ibm_api_version
  name = var.logs-router_tenants_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibm_api_version | Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version. | `string` | true |
| name | The name for this tenant. The name is regionally unique across all tenants in the account. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| tenants | List of tenants in the account. |

### Data source: ibm_logs-router_targets

```hcl
data "ibm_logs-router_targets" "logs-router_targets_instance" {
  ibm_api_version = var.logs-router_targets_ibm_api_version
  tenant_id = var.logs-router_targets_tenant_id
  name = var.logs-router_targets_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibm_api_version | Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version. | `string` | true |
| tenant_id | The instance ID of the tenant. | `` | true |
| name | The name for this tenant target. The name must be unique across all targets for this tenant. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| targets | List of targets of a tenant. |

## Assumptions

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |
