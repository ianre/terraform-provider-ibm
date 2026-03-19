variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for logs-router_tenant
variable "logs-router_tenant_ibm_api_version" {
  description = "Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version."
  type        = string
  default     = "ibm_api_version"
}
variable "logs-router_tenant_name" {
  description = "The name for this tenant. The name is regionally unique across all tenants in the account."
  type        = string
  default     = "my-logging-tenant"
}

// Data source arguments for logs-router_tenants
variable "logs-router_tenants_ibm_api_version" {
  description = "Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version."
  type        = string
  default     = "ibm_api_version"
}
variable "logs-router_tenants_name" {
  description = "The name for this tenant. The name is regionally unique across all tenants in the account."
  type        = string
  default     = "my-logging-tenant"
}

// Data source arguments for logs-router_targets
variable "logs-router_targets_ibm_api_version" {
  description = "Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version."
  type        = string
  default     = "ibm_api_version"
}
variable "logs-router_targets_tenant_id" {
  description = "The instance ID of the tenant."
  type        = 
  default     = 8717db99-2cfb-4ba6-a033-89c994c2e9f0
}
variable "logs-router_targets_name" {
  description = "The name for this tenant target. The name must be unique across all targets for this tenant."
  type        = string
  default     = "my-log-sink"
}
