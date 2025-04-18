---
layout: "ibm"
page_title: "IBM : ibm_sm_imported_certificate"
description: |-
  Manages ImportedCertificate.
subcategory: "Secrets Manager"
---

# ibm_sm_imported_certificate

Provides a resource for ImportedCertificate. This allows ImportedCertificate to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_sm_imported_certificate" "sm_imported_certificate" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  name 			= "secret-name"
  custom_metadata = {"key":"value"}
  description = "Extended description for this secret."
  labels = ["my-label"]
  secret_group_id = ibm_sm_secret_group.sm_secret_group.secret_group_id
  certificate = "-----BEGIN CERTIFICATE-----\nMIIE3jCCBGSgAwIBAgIUZfTbf3adn87l5J2Q2Aw+6Vk/qhowCgYIKoZIzj0EAwIw\n-----END CERTIFICATE-----"
}

An imported certificate with managed CSR:

resource "ibm_sm_imported_certificate" "sm_imported_certificate" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  name 			= "secret-name"
  custom_metadata = {"key":"value"}
  description = "Extended description for this secret."
  labels = ["my-label"]
  secret_group_id = ibm_sm_secret_group.sm_secret_group.secret_group_id
  managed_csr {
    alt_names = "altname"
    client_flag = true
    code_signing_flag = false
    common_name = "example.com"
    country = ["US"]
    email_protection_flag = false
    exclude_cn_from_sans = false
    ext_key_usage = "timestamping"
    ext_key_usage_oids = "1.3.6.1.5.5.7.3.67"
    ip_sans = "127.0.0.1"
    key_bits = 2048
    key_type = "rsa"
    key_usage = "DigitalSignature"
    locality = ["Boston"]
    organization = ["IBM"]
    other_sans = "1.3.6.1.4.1.311.21.2.3;utf8:*.example.com"
    ou = ["ILSL"]
    require_cn = true
    rotate_keys = false
    server_flag = true
    uri_sans = "https://www.example.com/test"
    user_ids = "user"
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
  * Constraints: Allowable values are: `private`, `public`.
* `certificate` - (Optional, String) The PEM-encoded contents of your certificate. You can manually rotate the secret by modifying this argument, together with the optional arguments `intermediate` and `private_key`. Modifying the certificate creates a new version of the secret. If the secret is used to generate a Certificate Signing Reques (CSR) no certificate should be provided initially. Add the certificate value only after the CSR is signed.
  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.
* `custom_metadata` - (Optional, Map) The secret metadata that a user can customize.
* `description` - (Optional, String) An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/(.*?)/`.
* `expiration_date` - (Optional, Forces new resource, String) The date a secret is expired. The date format follows RFC 3339.
* `intermediate` - (Computed, String) (Optional) The PEM-encoded intermediate certificate to associate with the root certificate.
  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.
* `labels` - (Optional, List) Labels that you can use to search for secrets in your instance.Up to 30 labels can be created.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `30` items. The minimum length is `0` items.
* `managed_csr` - (Optional, List) The data specified to create the CSR and the private key.
  Nested scheme for **managed_csr**:
  * `alt_names` - (Optional, String) With the Subject Alternative Name field, you can specify additional hostnames to be protected by a single SSL certificate.
  * `client_flag` - (Optional, Boolean) This field indicates whether certificate is flagged for client use. The default is `true`.
  * `code_signing_flag` - (Optional, Boolean) This field indicates whether certificate is flagged for code signing use. The default is `true`.
  * `common_name` - (Optional, String) The Common Name (CN) represents the server name protected by the SSL certificate.
  * `csr` - (Computed, String) The certificate signing request generated based on the parameters in the `managed_csr` data.
  * `country` - (Optional, List) The Country (C) values to define in the subject field of the resulting certificate.
  * `email_protection_flag` - (Optional, String) This field indicates whether certificate is flagged for email protection use. The default is `false`.
  * `exclude_cn_from_sans` - (Optional, String) This parameter controls whether the common name is excluded from Subject Alternative Names (SANs). The default is `false`.
  * `ext_key_usage` - (Optional, String) The allowed extended key usage constraint on certificate, in a comma-delimited list.
  * `ext_key_usage_oids` - (Optional, String) A comma-delimited list of extended key usage Object Identifiers (OIDs).
  * `ip_sans` - (Optional, String) The IP Subject Alternative Names to define for the certificate, in a comma-delimited list.
  * `key_bits` - (Optional, Integer) The number of bits to use to generate the private key.
  * `key_type` - (Optional, String) The type of private key to generate. The default is `rsa`.
  * `key_usage` - (Optional, String) The allowed key usage constraint to define for certificate, in a comma-delimited list.
  * `locality` - (Optional, List) The Locality (L) values to define in the subject field of the resulting certificate.
  * `organization` - (Optional, List) The Organization (O) values to define in the subject field of the resulting certificate.
  * `other_sans` - (Optional, String) The custom Object Identifier (OID) or UTF8-string Subject Alternative Names to define for the certificate, in a comma-delimited list.
  * `ou` - (Optional, List) The Organizational Unit (OU) values to define in the subject field of the resulting certificate.
  * `policy_identifiers` - (Optional, String) A comma-delimited list of policy Object Identifiers (OIDs).
  * `postal_code` - (Optional, List) The postal code values to define in the subject field of the resulting certificate.
  * `province` - (Optional, List) The Province (ST) values to define in the subject field of the resulting certificate.
  * `require_cn` - (Optional, Boolean) If set to false, makes the common_name field optional while generating a certificate. The default is `true`.
  * `rotate_keys` - (Optional, Boolean) This field indicates whether the private key will be rotated. The default is `false`.
  * `server_flag` - (Optional, Boolean) This field indicates whether certificate is flagged for server use. The default is `true`.
  * `street_address` - (Optional, List) The street address values to define in the subject field of the resulting certificate.
  * `uri_sans` - (Optional, String) The URI Subject Alternative Names to define for the certificate, in a comma-delimited list.
  * `user_ids` - (Optional, String) Specifies the list of requested User ID (OID 0.9.2342.19200300.100.1.1) Subject values to be placed on the signed certificate.
* `name` - (Required, String) The human-readable name of your secret.
  * Constraints: The maximum length is `256` characters. The minimum length is `2` characters. The value must match regular expression `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$`.
* `private_key` - (Computed, String) (Optional) The PEM-encoded private key to associate with the certificate.
  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.
* `secret_group_id` - (Optional, Forces new resource, String) A UUID identifier, or `default` secret group.
  * Constraints: The maximum length is `36` characters. The minimum length is `7` characters. The value must match regular expression `/^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|default)$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `secret_id` - The unique identifier of the ImportedCertificate.
* `common_name` - (Forces new resource, String) The Common Name (AKA CN) represents the server name protected by the SSL certificate.
  * Constraints: The maximum length is `64` characters. The minimum length is `4` characters. The value must match regular expression `/^(\\*\\.)?(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])\\.?$/`.
* `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.
* `created_by` - (String) The unique identifier that is associated with the entity that created the secret.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.
* `crn` - (String) A CRN that uniquely identifies an IBM Cloud resource.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
* `csr` - (String) The certificate signing request generated based on the parameters in the `managed_csr` data. The value may differ from the `csr` attribute within `managed_csr` if the `managed_csr` attributes have been modified.
* `downloaded` - (Boolean) Indicates whether the secret data that is associated with a secret version was retrieved in a call to the service API.
* `intermediate_included` - (Boolean) Indicates whether the certificate was imported with an associated intermediate certificate.
* `issuer` - (Forces new resource, String) The distinguished name that identifies the entity that signed and issued the certificate.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
* `key_algorithm` - (String) The identifier for the cryptographic algorithm used to generate the public key that is associated with the certificate.
  * Constraints: The maximum length is `10` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
* `locks_total` - (Integer) The number of locks of the secret.
  * Constraints: The maximum value is `1000`. The minimum value is `0`.
* `private_key_included` - (Boolean) Indicates whether the certificate was imported with an associated private key.
* `secret_type` - (String) The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.
  * Constraints: Allowable values are: `arbitrary`, `imported_cert`, `public_cert`, `iam_credentials`, `kv`, `username_password`, `private_cert`.
* `serial_number` - (String) The unique serial number that was assigned to a certificate by the issuing certificate authority.
  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[^a-fA-F0-9]/`.
* `signing_algorithm` - (String) The identifier for the cryptographic algorithm that was used by the issuing certificate authority to sign a certificate.
  * Constraints: The maximum length is `64` characters. The minimum length is `4` characters.
* `state` - (Integer) The secret state that is based on NIST SP 800-57. States are integers and correspond to the `Pre-activation = 0`, `Active = 1`,  `Suspended = 2`, `Deactivated = 3`, and `Destroyed = 5` values.
  * Constraints: Allowable values are: `0`, `1`, `2`, `3`, `5`.
* `state_description` - (String) A text representation of the secret state.
  * Constraints: Allowable values are: `pre_activation`, `active`, `suspended`, `deactivated`, `destroyed`.
* `updated_at` - (String) The date when a resource was recently modified. The date format follows RFC 3339.
* `validity` - (List) The date and time that the certificate validity period begins and ends.
Nested scheme for **validity**:
    * `not_after` - (String) The date-time format follows RFC 3339.
    * `not_before` - (String) The date-time format follows RFC 3339.
* `versions_total` - (Integer) The number of versions of the secret.
  * Constraints: The maximum value is `50`. The minimum value is `0`.

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:
```
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` and `iaas_classic_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`
  - Select `Classic Infrastructure API Keys` option from view dropdown for `iaas_classic_api_key`
2. For iaas_classic_username
  - Go to [Users](https://cloud.ibm.com/iam/users)
  - Click on user.
  - Find user name in the `VPN password` section under `User Details` tab

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_sm_imported_certificate` resource by using `region`, `instance_id`, and `secret_id`.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager)

# Syntax
```bash
$ terraform import ibm_sm_imported_certificate.sm_imported_certificate <region>/<instance_id>/<secret_id>
```

# Example
```bash
$ terraform import ibm_sm_imported_certificate.sm_imported_certificate us-east/6ebc4224-e983-496a-8a54-f40a0bfa9175/b49ad24d-81d4-5ebc-b9b9-b0937d1c84d5
```
