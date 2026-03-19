---
layout: "ibm"
page_title: "IBM : ibm_logs-router_targets"
description: |-
  Get information about logs-router_targets
subcategory: "Logs Router"
---

# ibm_logs-router_targets

Provides a read-only data source to retrieve information about logs-router_targets. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs-router_targets" "logs_router_targets" {
	ibm_api_version = "ibm_api_version"
	name = "my-log-sink"
	tenant_id = 8717db99-2cfb-4ba6-a033-89c994c2e9f0
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `ibm_api_version` - (Required, String) Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version.
  * Constraints: Length must be `10` characters. The value must match regular expression `/^[0-9]{4}-[0-9]{2}-[0-9]{2}$/`.
* `name` - (Optional, String) The name for this tenant target. The name must be unique across all targets for this tenant.
  * Constraints: The maximum length is `35` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z][a-zA-Z0-9-]*$/`.
* `tenant_id` - (Required, Forces new resource, String) The instance ID of the tenant.
  * Constraints: Length must be `36` characters.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs-router_targets.
* `targets` - (List) List of targets of a tenant.
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

