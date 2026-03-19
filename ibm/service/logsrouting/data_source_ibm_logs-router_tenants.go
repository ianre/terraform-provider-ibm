// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.112.0-f88e9264-20260220-115155
 */

package logsrouting

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-router-go-sdk/logsroutingv1"
)

func DataSourceIbmLogsRouterTenants() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmLogsRouterTenantsRead,

		Schema: map[string]*schema.Schema{
			"ibm_api_version": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name for this tenant. The name is regionally unique across all tenants in the account.",
			},
			"tenants": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of tenants in the account.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique ID of the tenant.",
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp the tenant was originally created.",
						},
						"updated_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp the tenant was last updated.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud resource name of the tenant. Must be a valid CRN.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this tenant. The name is regionally unique across all tenants in the account.",
						},
						"etag": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource version identifier.",
						},
						"targets": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of targets.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Unique ID of the target.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name for this tenant target. The name must be unique across all targets for this tenant.",
									},
									"etag": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Resource version identifier.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Type of log-sink. Identical to the <code>service-name</code> segment of <code>log_sink_crn</code>.",
									},
									"created_at": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Timestamp the target was originally created.",
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Timestamp the target was last updated.",
									},
									"log_sink_crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cloud resource name of the log-sink target instance. Must be a valid CRN.",
									},
									"parameters": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "List of properties returned from a successful list operation for a log-sink of type <b>logs</b> (IBM Cloud Logs).",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"host": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Host name of the log-sink.",
												},
												"port": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Network port of the log-sink.",
												},
											},
										},
									},
								},
							},
						},
						"write_status": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The status of the write attempt to the target with the provided endpoint parameters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status such as failed or success.",
									},
									"reason_for_last_failure": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Detailed description of the cause of the failure.",
									},
									"last_failure": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The timestamp of the failure.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmLogsRouterTenantsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRoutingClient, err := meta.(conns.ClientSession).LogsRoutingV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs-router_tenants", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listTenantsOptions := &logsroutingv1.ListTenantsOptions{}

	listTenantsOptions.SetIBMAPIVersion(d.Get("ibm_api_version").(string))

	tenantCollection, _, err := logsRoutingClient.ListTenantsWithContext(context, listTenantsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListTenantsWithContext failed: %s", err.Error()), "(Data) ibm_logs-router_tenants", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Use the provided filter argument and construct a new list with only the requested resource(s)
	var matchTenants []logsroutingv1.Tenant
	var name string
	var suppliedFilter bool

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
		suppliedFilter = true
		for _, data := range tenantCollection.Tenants {
			if *data.Name == name {
				matchTenants = append(matchTenants, data)
			}
		}
	} else {
		matchTenants = tenantCollection.Tenants
	}
	tenantCollection.Tenants = matchTenants

	if suppliedFilter {
		if len(tenantCollection.Tenants) == 0 {
			return flex.DiscriminatedTerraformErrorf(nil, fmt.Sprintf("no Tenants found with name %s", name), "(Data) ibm_logs-router_tenants", "read", "no-collection-found").GetDiag()
		}
		d.SetId(name)
	} else {
		d.SetId(dataSourceIbmLogsRouterTenantsID(d))
	}

	if !core.IsNil(tenantCollection.Tenants) {
		tenants := []map[string]interface{}{}
		for _, tenantsItem := range tenantCollection.Tenants {
			tenantsItemMap, err := DataSourceIbmLogsRouterTenantsTenantToMap(&tenantsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs-router_tenants", "read", "tenants-to-map").GetDiag()
			}
			tenants = append(tenants, tenantsItemMap)
		}
		if err = d.Set("tenants", tenants); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tenants: %s", err), "(Data) ibm_logs-router_tenants", "read", "set-tenants").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmLogsRouterTenantsID returns a reasonable ID for the list.
func dataSourceIbmLogsRouterTenantsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmLogsRouterTenantsTenantToMap(model *logsroutingv1.Tenant) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	modelMap["created_at"] = model.CreatedAt.String()
	modelMap["updated_at"] = model.UpdatedAt.String()
	modelMap["crn"] = *model.Crn
	modelMap["name"] = *model.Name
	modelMap["etag"] = *model.Etag
	targets := []map[string]interface{}{}
	for _, targetsItem := range model.Targets {
		targetsItemMap, err := DataSourceIbmLogsRouterTenantsTargetToMap(targetsItem)
		if err != nil {
			return modelMap, err
		}
		targets = append(targets, targetsItemMap)
	}
	modelMap["targets"] = targets
	writeStatusMap, err := DataSourceIbmLogsRouterTenantsWriteStatusToMap(model.WriteStatus)
	if err != nil {
		return modelMap, err
	}
	modelMap["write_status"] = []map[string]interface{}{writeStatusMap}
	return modelMap, nil
}

func DataSourceIbmLogsRouterTenantsTargetToMap(model logsroutingv1.TargetIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsroutingv1.TargetTypeLogs); ok {
		return DataSourceIbmLogsRouterTenantsTargetTypeLogsToMap(model.(*logsroutingv1.TargetTypeLogs))
	} else if _, ok := model.(*logsroutingv1.Target); ok {
		modelMap := make(map[string]interface{})
		model := model.(*logsroutingv1.Target)
		modelMap["id"] = model.ID.String()
		modelMap["name"] = *model.Name
		modelMap["etag"] = *model.Etag
		modelMap["type"] = *model.Type
		modelMap["created_at"] = model.CreatedAt.String()
		modelMap["updated_at"] = model.UpdatedAt.String()
		if model.LogSinkCrn != nil {
			modelMap["log_sink_crn"] = *model.LogSinkCrn
		}
		if model.Parameters != nil {
			parametersMap, err := DataSourceIbmLogsRouterTenantsTargetParametersTypeLogsToMap(model.Parameters)
			if err != nil {
				return modelMap, err
			}
			modelMap["parameters"] = []map[string]interface{}{parametersMap}
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized logsroutingv1.TargetIntf subtype encountered")
	}
}

func DataSourceIbmLogsRouterTenantsTargetParametersTypeLogsToMap(model *logsroutingv1.TargetParametersTypeLogs) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["host"] = *model.Host
	modelMap["port"] = flex.IntValue(model.Port)
	return modelMap, nil
}

func DataSourceIbmLogsRouterTenantsTargetTypeLogsToMap(model *logsroutingv1.TargetTypeLogs) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	modelMap["name"] = *model.Name
	modelMap["etag"] = *model.Etag
	modelMap["type"] = *model.Type
	modelMap["created_at"] = model.CreatedAt.String()
	modelMap["updated_at"] = model.UpdatedAt.String()
	modelMap["log_sink_crn"] = *model.LogSinkCrn
	parametersMap, err := DataSourceIbmLogsRouterTenantsTargetParametersTypeLogsToMap(model.Parameters)
	if err != nil {
		return modelMap, err
	}
	modelMap["parameters"] = []map[string]interface{}{parametersMap}
	return modelMap, nil
}

func DataSourceIbmLogsRouterTenantsWriteStatusToMap(model *logsroutingv1.WriteStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["status"] = *model.Status
	if model.ReasonForLastFailure != nil {
		modelMap["reason_for_last_failure"] = *model.ReasonForLastFailure
	}
	if model.LastFailure != nil {
		modelMap["last_failure"] = model.LastFailure.String()
	}
	return modelMap, nil
}
