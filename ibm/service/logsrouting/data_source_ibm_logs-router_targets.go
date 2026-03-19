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

func DataSourceIbmLogsRouterTargets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmLogsRouterTargetsRead,

		Schema: map[string]*schema.Schema{
			"ibm_api_version": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version.",
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance ID of the tenant.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name for this tenant target. The name must be unique across all targets for this tenant.",
			},
			"targets": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of targets of a tenant.",
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
		},
	}
}

func dataSourceIbmLogsRouterTargetsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRoutingClient, err := meta.(conns.ClientSession).LogsRoutingV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs-router_targets", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listTenantTargetsOptions := &logsroutingv1.ListTenantTargetsOptions{}

	listTenantTargetsOptions.SetIBMAPIVersion(d.Get("ibm_api_version").(string))
	listTenantTargetsOptions.SetTenantID(core.UUIDPtr(strfmt.UUID(d.Get("tenant_id").(string))))

	targetCollection, _, err := logsRoutingClient.ListTenantTargetsWithContext(context, listTenantTargetsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListTenantTargetsWithContext failed: %s", err.Error()), "(Data) ibm_logs-router_targets", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Use the provided filter argument and construct a new list with only the requested resource(s)
	var matchTargets []logsroutingv1.Target
	var name string
	var suppliedFilter bool

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
		suppliedFilter = true
		for _, data := range targetCollection.Targets {
			if *data.Name == name {
				matchTargets = append(matchTargets, data)
			}
		}
	} else {
		matchTargets = targetCollection.Targets
	}
	targetCollection.Targets = matchTargets

	if suppliedFilter {
		if len(targetCollection.Targets) == 0 {
			return flex.DiscriminatedTerraformErrorf(nil, fmt.Sprintf("no Targets found with name %s", name), "(Data) ibm_logs-router_targets", "read", "no-collection-found").GetDiag()
		}
		d.SetId(name)
	} else {
		d.SetId(dataSourceIbmLogsRouterTargetsID(d))
	}

	if !core.IsNil(targetCollection.Targets) {
		targets := []map[string]interface{}{}
		for _, targetsItem := range targetCollection.Targets {
			targetsItemMap, err := DataSourceIbmLogsRouterTargetsTargetToMap(targetsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs-router_targets", "read", "targets-to-map").GetDiag()
			}
			targets = append(targets, targetsItemMap)
		}
		if err = d.Set("targets", targets); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting targets: %s", err), "(Data) ibm_logs-router_targets", "read", "set-targets").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmLogsRouterTargetsID returns a reasonable ID for the list.
func dataSourceIbmLogsRouterTargetsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmLogsRouterTargetsTargetToMap(model logsroutingv1.TargetIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsroutingv1.TargetTypeLogs); ok {
		return DataSourceIbmLogsRouterTargetsTargetTypeLogsToMap(model.(*logsroutingv1.TargetTypeLogs))
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
			parametersMap, err := DataSourceIbmLogsRouterTargetsTargetParametersTypeLogsToMap(model.Parameters)
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

func DataSourceIbmLogsRouterTargetsTargetParametersTypeLogsToMap(model *logsroutingv1.TargetParametersTypeLogs) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["host"] = *model.Host
	modelMap["port"] = flex.IntValue(model.Port)
	return modelMap, nil
}

func DataSourceIbmLogsRouterTargetsTargetTypeLogsToMap(model *logsroutingv1.TargetTypeLogs) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	modelMap["name"] = *model.Name
	modelMap["etag"] = *model.Etag
	modelMap["type"] = *model.Type
	modelMap["created_at"] = model.CreatedAt.String()
	modelMap["updated_at"] = model.UpdatedAt.String()
	modelMap["log_sink_crn"] = *model.LogSinkCrn
	parametersMap, err := DataSourceIbmLogsRouterTargetsTargetParametersTypeLogsToMap(model.Parameters)
	if err != nil {
		return modelMap, err
	}
	modelMap["parameters"] = []map[string]interface{}{parametersMap}
	return modelMap, nil
}
