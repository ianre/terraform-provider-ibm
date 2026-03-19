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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-router-go-sdk/logsroutingv1"
)

func ResourceIbmLogsRouterTenant() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmLogsRouterTenantCreate,
		ReadContext:   resourceIbmLogsRouterTenantRead,
		UpdateContext: resourceIbmLogsRouterTenantUpdate,
		DeleteContext: resourceIbmLogsRouterTenantDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"ibm_api_version": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs-router_tenant", "ibm_api_version"),
				Description:  "Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs-router_tenant", "name"),
				Description:  "The name for this tenant. The name is regionally unique across all tenants in the account.",
			},
			"targets": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of targets.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique ID of the target.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name for this tenant target. The name must be unique across all targets for this tenant.",
						},
						"etag": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Resource version identifier.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Type of log-sink. Identical to the <code>service-name</code> segment of <code>log_sink_crn</code>.",
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Timestamp the target was originally created.",
						},
						"updated_at": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Timestamp the target was last updated.",
						},
						"log_sink_crn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Cloud resource name of the log-sink target instance. Must be a valid CRN.",
						},
						"parameters": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "List of properties returned from a successful list operation for a log-sink of type <b>logs</b> (IBM Cloud Logs).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"host": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Host name of the log-sink.",
									},
									"port": &schema.Schema{
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Network port of the log-sink.",
									},
								},
							},
						},
					},
				},
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
			"etag": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource version identifier.",
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
							Optional:    true,
							Computed:    true,
							Description: "Detailed description of the cause of the failure.",
						},
						"last_failure": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The timestamp of the failure.",
						},
					},
				},
			},
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIbmLogsRouterTenantValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "ibm_api_version",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9]{4}-[0-9]{2}-[0-9]{2}$`,
			MinValueLength:             10,
			MaxValueLength:             10,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z][a-zA-Z0-9-]*$`,
			MinValueLength:             1,
			MaxValueLength:             35,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_logs-router_tenant", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmLogsRouterTenantCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRoutingClient, err := meta.(conns.ClientSession).LogsRoutingV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs-router_tenant", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createTenantOptions := &logsroutingv1.CreateTenantOptions{}

	createTenantOptions.SetIBMAPIVersion(d.Get("ibm_api_version").(string))
	createTenantOptions.SetName(d.Get("name").(string))
	var targets []logsroutingv1.TargetPrototypeIntf
	for _, v := range d.Get("targets").([]interface{}) {
		value := v.(map[string]interface{})
		targetsItem, err := ResourceIbmLogsRouterTenantMapToTargetPrototype(value)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs-router_tenant", "create", "parse-targets").GetDiag()
		}
		targets = append(targets, targetsItem)
	}
	createTenantOptions.SetTargets(targets)

	tenant, _, err := logsRoutingClient.CreateTenantWithContext(context, createTenantOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateTenantWithContext failed: %s", err.Error()), "ibm_logs-router_tenant", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*tenant.ID)

	return resourceIbmLogsRouterTenantRead(context, d, meta)
}

func resourceIbmLogsRouterTenantRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRoutingClient, err := meta.(conns.ClientSession).LogsRoutingV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs-router_tenant", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getTenantDetailOptions := &logsroutingv1.GetTenantDetailOptions{}

	getTenantDetailOptions.SetTenantID(d.Id())
	getTenantDetailOptions.SetIBMAPIVersion(d.Get("ibm_api_version").(string))

	tenant, response, err := logsRoutingClient.GetTenantDetailWithContext(context, getTenantDetailOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetTenantDetailWithContext failed: %s", err.Error()), "ibm_logs-router_tenant", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("name", tenant.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs-router_tenant", "read", "set-name").GetDiag()
	}
	targets := []map[string]interface{}{}
	for _, targetsItem := range tenant.Targets {
		targetsItemMap, err := ResourceIbmLogsRouterTenantTargetToMap(targetsItem)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs-router_tenant", "read", "targets-to-map").GetDiag()
		}
		targets = append(targets, targetsItemMap)
	}
	if err = d.Set("targets", targets); err != nil {
		err = fmt.Errorf("Error setting targets: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs-router_tenant", "read", "set-targets").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(tenant.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs-router_tenant", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("updated_at", flex.DateTimeToString(tenant.UpdatedAt)); err != nil {
		err = fmt.Errorf("Error setting updated_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs-router_tenant", "read", "set-updated_at").GetDiag()
	}
	if err = d.Set("crn", tenant.Crn); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs-router_tenant", "read", "set-crn").GetDiag()
	}
	if err = d.Set("etag", tenant.Etag); err != nil {
		err = fmt.Errorf("Error setting etag: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs-router_tenant", "read", "set-etag").GetDiag()
	}
	writeStatusMap, err := ResourceIbmLogsRouterTenantWriteStatusToMap(tenant.WriteStatus)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs-router_tenant", "read", "write_status-to-map").GetDiag()
	}
	if err = d.Set("write_status", []map[string]interface{}{writeStatusMap}); err != nil {
		err = fmt.Errorf("Error setting write_status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs-router_tenant", "read", "set-write_status").GetDiag()
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_logs-router_tenant", "read", "set-etag").GetDiag()
	}

	return nil
}

func resourceIbmLogsRouterTenantUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRoutingClient, err := meta.(conns.ClientSession).LogsRoutingV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs-router_tenant", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateTenantOptions := &logsroutingv1.UpdateTenantOptions{}

	updateTenantOptions.SetTenantID(d.Id())

	hasChange := false

	patchVals := &logsroutingv1.TenantPatch{}
	if d.HasChange("ibm_api_version") {
		updateTenantOptions.SetIBMAPIVersion(d.Get("ibm_api_version").(string))
		hasChange = true
	}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		hasChange = true
	}
	updateTenantOptions.SetIfMatch(d.Get("etag").(string))

	if hasChange {
		// Fields with `nil` values are omitted from the generic map,
		// so we need to re-add them to support removing arguments
		// in merge-patch operations sent to the service.
		updateTenantOptions.TenantPatch = ResourceIbmLogsRouterTenantTenantPatchAsPatch(patchVals, d)

		_, _, err = logsRoutingClient.UpdateTenantWithContext(context, updateTenantOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateTenantWithContext failed: %s", err.Error()), "ibm_logs-router_tenant", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmLogsRouterTenantRead(context, d, meta)
}

func resourceIbmLogsRouterTenantDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRoutingClient, err := meta.(conns.ClientSession).LogsRoutingV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs-router_tenant", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteTenantOptions := &logsroutingv1.DeleteTenantOptions{}

	deleteTenantOptions.SetTenantID(d.Id())
	deleteTenantOptions.SetIBMAPIVersion(d.Get("ibm_api_version").(string))

	_, err = logsRoutingClient.DeleteTenantWithContext(context, deleteTenantOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteTenantWithContext failed: %s", err.Error()), "ibm_logs-router_tenant", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIbmLogsRouterTenantMapToTargetPrototype(modelMap map[string]interface{}) (logsroutingv1.TargetPrototypeIntf, error) {
	model := &logsroutingv1.TargetPrototype{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["log_sink_crn"] != nil && modelMap["log_sink_crn"].(string) != "" {
		model.LogSinkCrn = core.StringPtr(modelMap["log_sink_crn"].(string))
	}
	if modelMap["parameters"] != nil && len(modelMap["parameters"].([]interface{})) > 0 {
		ParametersModel, err := ResourceIbmLogsRouterTenantMapToTargetParametersTypeLogsPrototype(modelMap["parameters"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Parameters = ParametersModel
	}
	return model, nil
}

func ResourceIbmLogsRouterTenantMapToTargetParametersTypeLogsPrototype(modelMap map[string]interface{}) (*logsroutingv1.TargetParametersTypeLogsPrototype, error) {
	model := &logsroutingv1.TargetParametersTypeLogsPrototype{}
	model.Host = core.StringPtr(modelMap["host"].(string))
	model.Port = core.Int64Ptr(int64(modelMap["port"].(int)))
	return model, nil
}

func ResourceIbmLogsRouterTenantMapToTargetPrototypeTargetTypeLogsPrototype(modelMap map[string]interface{}) (*logsroutingv1.TargetPrototypeTargetTypeLogsPrototype, error) {
	model := &logsroutingv1.TargetPrototypeTargetTypeLogsPrototype{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.LogSinkCrn = core.StringPtr(modelMap["log_sink_crn"].(string))
	ParametersModel, err := ResourceIbmLogsRouterTenantMapToTargetParametersTypeLogsPrototype(modelMap["parameters"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Parameters = ParametersModel
	return model, nil
}

func ResourceIbmLogsRouterTenantTargetToMap(model logsroutingv1.TargetIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsroutingv1.TargetTypeLogs); ok {
		return ResourceIbmLogsRouterTenantTargetTypeLogsToMap(model.(*logsroutingv1.TargetTypeLogs))
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
			parametersMap, err := ResourceIbmLogsRouterTenantTargetParametersTypeLogsToMap(model.Parameters)
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

func ResourceIbmLogsRouterTenantTargetParametersTypeLogsToMap(model *logsroutingv1.TargetParametersTypeLogs) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["host"] = *model.Host
	modelMap["port"] = flex.IntValue(model.Port)
	return modelMap, nil
}

func ResourceIbmLogsRouterTenantTargetTypeLogsToMap(model *logsroutingv1.TargetTypeLogs) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	modelMap["name"] = *model.Name
	modelMap["etag"] = *model.Etag
	modelMap["type"] = *model.Type
	modelMap["created_at"] = model.CreatedAt.String()
	modelMap["updated_at"] = model.UpdatedAt.String()
	modelMap["log_sink_crn"] = *model.LogSinkCrn
	parametersMap, err := ResourceIbmLogsRouterTenantTargetParametersTypeLogsToMap(model.Parameters)
	if err != nil {
		return modelMap, err
	}
	modelMap["parameters"] = []map[string]interface{}{parametersMap}
	return modelMap, nil
}

func ResourceIbmLogsRouterTenantWriteStatusToMap(model *logsroutingv1.WriteStatus) (map[string]interface{}, error) {
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

func ResourceIbmLogsRouterTenantTenantPatchAsPatch(patchVals *logsroutingv1.TenantPatch, d *schema.ResourceData) map[string]interface{} {
	patch, _ := patchVals.AsPatch()
	var path string

	path = "name"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["name"] = nil
	} else if !exists {
		delete(patch, "name")
	}

	return patch
}
