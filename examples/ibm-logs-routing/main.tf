provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision logs-router_tenant resource instance
resource "ibm_logs-router_tenant" "logs-router_tenant_instance" {
  ibm_api_version = var.logs-router_tenant_ibm_api_version
  name = var.logs-router_tenant_name
  targets {
    id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
    name = "my-log-sink"
    etag = "822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049"
    type = "logs"
    created_at = "2024-06-20T18:30:00.143Z"
    updated_at = "2024-06-20T18:30:00.143Z"
    log_sink_crn = "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::"
    parameters {
      host = "www.example.com"
      port = 1
    }
  }
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs-router_tenants data source
data "ibm_logs-router_tenants" "logs-router_tenants_instance" {
  ibm_api_version = var.logs-router_tenants_ibm_api_version
  name = var.logs-router_tenants_name
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs-router_targets data source
data "ibm_logs-router_targets" "logs-router_targets_instance" {
  ibm_api_version = var.logs-router_targets_ibm_api_version
  tenant_id = var.logs-router_targets_tenant_id
  name = var.logs-router_targets_name
}
*/
