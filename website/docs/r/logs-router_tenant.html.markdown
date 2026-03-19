---
layout: "ibm"
page_title: "IBM : ibm_logs-router_tenant"
description: |-
  Manages logs-router_tenant.
subcategory: "Logs Router"
---

# ibm_logs-router_tenant

Create, update, and delete logs-router_tenants with this resource.

## Example Usage

```hcl
resource "ibm_logs-router_tenant" "logs_router_tenant_instance" {
  ibm_api_version = "ibm_api_version"
  name = "my-logging-tenant"
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
```

## Argument Reference

You can specify the following arguments for this resource.

* `ibm_api_version` - (Required, String) Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version.
  * Constraints: Length must be `10` characters. The value must match regular expression `/^[0-9]{4}-[0-9]{2}-[0-9]{2}$/`.
* `name` - (Required, String) The name for this tenant. The name is regionally unique across all tenants in the account.
  * Constraints: The maximum length is `35` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z][a-zA-Z0-9-]*$/`.
* `targets` - (Required, List) List of targets.
  * Constraints: The maximum length is `2` items. The minimum length is `1` item.
Nested schema for **targets**:
	* `created_at` - (Required, String) Timestamp the target was originally created.
	* `etag` - (Required, String) Resource version identifier.
	  * Constraints: Length must be `66` characters. The value must match regular expression `/^(?:W\/)?"(?:[ !#-\\x7E\\x80-\\xFF]*|  [  ]|\\.)*"$/`.
	* `id` - (Required, String) Unique ID of the target.
	  * Constraints: Length must be `36` characters.
	* `log_sink_crn` - (Optional, String) Cloud resource name of the log-sink target instance. Must be a valid CRN.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character.
	* `name` - (Required, String) The name for this tenant target. The name must be unique across all targets for this tenant.
	  * Constraints: The maximum length is `35` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z][a-zA-Z0-9-]*$/`.
	* `parameters` - (Optional, List) List of properties returned from a successful list operation for a log-sink of type <b>logs</b> (IBM Cloud Logs).
	Nested schema for **parameters**:
		* `host` - (Required, String) Host name of the log-sink.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-.:\/]+$/`.
		* `port` - (Required, Integer) Network port of the log-sink.
		  * Constraints: The maximum value is `65535`. The minimum value is `1`.
	* `type` - (Required, String) Type of log-sink. Identical to the <code>service-name</code> segment of <code>log_sink_crn</code>.
	  * Constraints: Allowable values are: `logs`.
	* `updated_at` - (Required, String) Timestamp the target was last updated.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs-router_tenant.
* `created_at` - (String) Timestamp the tenant was originally created.
* `crn` - (String) Cloud resource name of the tenant. Must be a valid CRN.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character.
* `etag` - (String) Resource version identifier.
  * Constraints: Length must be `66` characters. The value must match regular expression `/^(?:W\/)?"(?:[ !#-\\x7E\\x80-\\xFF]*|  [  ]|\\.)*"$/`.
* `updated_at` - (String) Timestamp the tenant was last updated.
* `write_status` - (List) The status of the write attempt to the target with the provided endpoint parameters.
Nested schema for **write_status**:
	* `last_failure` - (String) The timestamp of the failure.
	* `reason_for_last_failure` - (String) Detailed description of the cause of the failure.
	* `status` - (String) The status such as failed or success.

* `etag` - ETag identifier for logs-router_tenant.

## Import

You can import the `ibm_logs-router_tenant` resource by using `id`. Unique ID of the tenant.
For more information, see [the documentation](http://cloud.ibm.com)

# Syntax
<pre>
$ terraform import ibm_logs-router_tenant.logs_router_tenant &lt;id&gt;
</pre>

# Example
```
$ terraform import ibm_logs-router_tenant.logs_router_tenant 8717db99-2cfb-4ba6-a033-89c994c2e9f0
```
