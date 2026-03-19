---
layout: "ibm"
page_title: "IBM : ibm_logs-router_tenants"
description: |-
  Get information about logs-router_tenants
subcategory: "Logs Router"
---

# ibm_logs-router_tenants

Provides a read-only data source to retrieve information about logs-router_tenants. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs-router_tenants" "logs_router_tenants" {
	ibm_api_version = ibm_logs-router_tenant.logs_router_tenant_instance.ibm_api_version
	name = ibm_logs-router_tenant.logs_router_tenant_instance.name
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `ibm_api_version` - (Required, String) Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version.
  * Constraints: Length must be `10` characters. The value must match regular expression `/^[0-9]{4}-[0-9]{2}-[0-9]{2}$/`.
* `name` - (Optional, String) The name for this tenant. The name is regionally unique across all tenants in the account.
  * Constraints: The maximum length is `35` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z][a-zA-Z0-9-]*$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs-router_tenants.
* `tenants` - (List) List of tenants in the account.
  * Constraints: The maximum length is `1` item. The minimum length is `0` items.
Nested schema for **tenants**:
	* `created_at` - (String) Timestamp the tenant was originally created.
	* `crn` - (String) Cloud resource name of the tenant. Must be a valid CRN.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character.
	* `etag` - (String) Resource version identifier.
	  * Constraints: Length must be `66` characters. The value must match regular expression `/^(?:W\/)?"(?:[ !#-\\x7E\\x80-\\xFF]*|  [  ]|\\.)*"$/`.
	* `id` - (String) Unique ID of the tenant.
	  * Constraints: Length must be `36` characters.
	* `name` - (String) The name for this tenant. The name is regionally unique across all tenants in the account.
	  * Constraints: The maximum length is `35` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z][a-zA-Z0-9-]*$/`.
	* `targets` - (List) List of targets.
	  * Constraints: The maximum length is `2` items. The minimum length is `1` item.
	Nested schema for **targets**:
		* `created_at` - (String) Timestamp the target was originally created.
		* `etag` - (String) Resource version identifier.
		  * Constraints: Length must be `66` characters. The value must match regular expression `/^(?:W\/)?"(?:[ !#-\\x7E\\x80-\\xFF]*|  [  ]|\\.)*"$/`.
		* `id` - (String) Unique ID of the target.
		  * Constraints: Length must be `36` characters.
		* `log_sink_crn` - (String) Cloud resource name of the log-sink target instance. Must be a valid CRN.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character.
		* `name` - (String) The name for this tenant target. The name must be unique across all targets for this tenant.
		  * Constraints: The maximum length is `35` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z][a-zA-Z0-9-]*$/`.
		* `parameters` - (List) List of properties returned from a successful list operation for a log-sink of type <b>logs</b> (IBM Cloud Logs).
		Nested schema for **parameters**:
			* `host` - (String) Host name of the log-sink.
			  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-.:\/]+$/`.
			* `port` - (Integer) Network port of the log-sink.
			  * Constraints: The maximum value is `65535`. The minimum value is `1`.
		* `type` - (String) Type of log-sink. Identical to the <code>service-name</code> segment of <code>log_sink_crn</code>.
		  * Constraints: Allowable values are: `logs`.
		* `updated_at` - (String) Timestamp the target was last updated.
	* `updated_at` - (String) Timestamp the tenant was last updated.
	* `write_status` - (List) The status of the write attempt to the target with the provided endpoint parameters.
	Nested schema for **write_status**:
		* `last_failure` - (String) The timestamp of the failure.
		* `reason_for_last_failure` - (String) Detailed description of the cause of the failure.
		* `status` - (String) The status such as failed or success.

