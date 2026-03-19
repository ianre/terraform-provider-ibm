// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.112.0-f88e9264-20260220-115155
 */

package logsrouting_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logsrouting"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-router-go-sdk/logsroutingv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsRouterTenantsDataSourceBasic(t *testing.T) {
	tenantIBMAPIVersion := fmt.Sprintf("tf_ibm_api_version_%d", acctest.RandIntRange(10, 100))
	tenantName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsRouterTenantsDataSourceConfigBasic(tenantIBMAPIVersion, tenantName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs-router_tenants.logs_router_tenants_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs-router_tenants.logs_router_tenants_instance", "ibm_api_version"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsRouterTenantsDataSourceConfigBasic(tenantIBMAPIVersion string, tenantName string) string {
	return fmt.Sprintf(`
		resource "ibm_logs-router_tenant" "logs_router_tenant_instance" {
			ibm_api_version = "%s"
			name = "%s"
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

		data "ibm_logs-router_tenants" "logs_router_tenants_instance" {
			ibm_api_version = ibm_logs-router_tenant.logs_router_tenant_instance.ibm_api_version
			name = ibm_logs-router_tenant.logs_router_tenant_instance.name
		}
	`, tenantIBMAPIVersion, tenantName)
}

func TestDataSourceIbmLogsRouterTenantsTenantToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		targetParametersTypeLogsModel := make(map[string]interface{})
		targetParametersTypeLogsModel["host"] = "www.example.com"
		targetParametersTypeLogsModel["port"] = int(8080)

		targetModel := make(map[string]interface{})
		targetModel["id"] = "BFB84FFB-6C2D-4204-A354-829F180B1D8B"
		targetModel["name"] = "my-logs-log-sink"
		targetModel["etag"] = "c7e475a63f3e9dcd74e432ead845be0821e839e08b087ad5fafd87907d21dbfd"
		targetModel["type"] = "logs"
		targetModel["created_at"] = "2024-06-20T18:30:00.143156Z"
		targetModel["updated_at"] = "2024-06-20T18:30:00.143156Z"
		targetModel["log_sink_crn"] = "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:42DAB2C2-BEC2-43E3-8E90-F0259087A884::"
		targetModel["parameters"] = []map[string]interface{}{targetParametersTypeLogsModel}

		writeStatusModel := make(map[string]interface{})
		writeStatusModel["status"] = "success"
		writeStatusModel["reason_for_last_failure"] = "Logs endpoint is not reachable. Received status code: 403"
		writeStatusModel["last_failure"] = "2024-10-14T10:49:09.000Z"

		model := make(map[string]interface{})
		model["id"] = "8717db99-2cfb-4ba6-a033-89c994c2e9f0"
		model["created_at"] = "2024-06-20T18:30:00.143Z"
		model["updated_at"] = "2024-06-20T18:30:00.143Z"
		model["crn"] = "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:b10d649e-5cad-4274-a223-40f8e4dbdea6::"
		model["name"] = "my-logging-tenant"
		model["etag"] = "822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049"
		model["targets"] = []map[string]interface{}{targetModel}
		model["write_status"] = []map[string]interface{}{writeStatusModel}

		assert.Equal(t, result, model)
	}

	targetParametersTypeLogsModel := new(logsroutingv1.TargetParametersTypeLogs)
	targetParametersTypeLogsModel.Host = core.StringPtr("www.example.com")
	targetParametersTypeLogsModel.Port = core.Int64Ptr(int64(8080))

	targetModel := new(logsroutingv1.TargetTypeLogs)
	targetModel.ID = CreateMockUUID("BFB84FFB-6C2D-4204-A354-829F180B1D8B")
	targetModel.Name = core.StringPtr("my-logs-log-sink")
	targetModel.Etag = core.StringPtr("c7e475a63f3e9dcd74e432ead845be0821e839e08b087ad5fafd87907d21dbfd")
	targetModel.Type = core.StringPtr("logs")
	targetModel.CreatedAt = CreateMockDateTime("2024-06-20T18:30:00.143156Z")
	targetModel.UpdatedAt = CreateMockDateTime("2024-06-20T18:30:00.143156Z")
	targetModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:42DAB2C2-BEC2-43E3-8E90-F0259087A884::")
	targetModel.Parameters = targetParametersTypeLogsModel

	writeStatusModel := new(logsroutingv1.WriteStatus)
	writeStatusModel.Status = core.StringPtr("success")
	writeStatusModel.ReasonForLastFailure = core.StringPtr("Logs endpoint is not reachable. Received status code: 403")
	writeStatusModel.LastFailure = CreateMockDateTime("2024-10-14T10:49:09.000Z")

	model := new(logsroutingv1.Tenant)
	model.ID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
	model.CreatedAt = CreateMockDateTime("2024-06-20T18:30:00.143Z")
	model.UpdatedAt = CreateMockDateTime("2024-06-20T18:30:00.143Z")
	model.Crn = core.StringPtr("crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:b10d649e-5cad-4274-a223-40f8e4dbdea6::")
	model.Name = core.StringPtr("my-logging-tenant")
	model.Etag = core.StringPtr("822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049")
	model.Targets = []logsroutingv1.TargetIntf{targetModel}
	model.WriteStatus = writeStatusModel

	result, err := logsrouting.DataSourceIbmLogsRouterTenantsTenantToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsRouterTenantsTargetToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		targetParametersTypeLogsModel := make(map[string]interface{})
		targetParametersTypeLogsModel["host"] = "www.example.com"
		targetParametersTypeLogsModel["port"] = int(8080)

		model := make(map[string]interface{})
		model["id"] = "c40e55a5-0833-4748-b032-b8e8cfe6e135"
		model["name"] = "my-log-sink"
		model["etag"] = "822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049"
		model["type"] = "logs"
		model["created_at"] = "2024-06-20T18:30:00.143Z"
		model["updated_at"] = "2024-06-20T18:30:00.143Z"
		model["log_sink_crn"] = "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::"
		model["parameters"] = []map[string]interface{}{targetParametersTypeLogsModel}

		assert.Equal(t, result, model)
	}

	targetParametersTypeLogsModel := new(logsroutingv1.TargetParametersTypeLogs)
	targetParametersTypeLogsModel.Host = core.StringPtr("www.example.com")
	targetParametersTypeLogsModel.Port = core.Int64Ptr(int64(8080))

	model := new(logsroutingv1.Target)
	model.ID = CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
	model.Name = core.StringPtr("my-log-sink")
	model.Etag = core.StringPtr("822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049")
	model.Type = core.StringPtr("logs")
	model.CreatedAt = CreateMockDateTime("2024-06-20T18:30:00.143Z")
	model.UpdatedAt = CreateMockDateTime("2024-06-20T18:30:00.143Z")
	model.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
	model.Parameters = targetParametersTypeLogsModel

	result, err := logsrouting.DataSourceIbmLogsRouterTenantsTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsRouterTenantsTargetParametersTypeLogsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["host"] = "www.example.com"
		model["port"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(logsroutingv1.TargetParametersTypeLogs)
	model.Host = core.StringPtr("www.example.com")
	model.Port = core.Int64Ptr(int64(1))

	result, err := logsrouting.DataSourceIbmLogsRouterTenantsTargetParametersTypeLogsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsRouterTenantsTargetTypeLogsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		targetParametersTypeLogsModel := make(map[string]interface{})
		targetParametersTypeLogsModel["host"] = "www.example.com"
		targetParametersTypeLogsModel["port"] = int(8080)

		model := make(map[string]interface{})
		model["id"] = "c40e55a5-0833-4748-b032-b8e8cfe6e135"
		model["name"] = "my-log-sink"
		model["etag"] = "822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049"
		model["type"] = "logs"
		model["created_at"] = "2024-06-20T18:30:00.143Z"
		model["updated_at"] = "2024-06-20T18:30:00.143Z"
		model["log_sink_crn"] = "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::"
		model["parameters"] = []map[string]interface{}{targetParametersTypeLogsModel}

		assert.Equal(t, result, model)
	}

	targetParametersTypeLogsModel := new(logsroutingv1.TargetParametersTypeLogs)
	targetParametersTypeLogsModel.Host = core.StringPtr("www.example.com")
	targetParametersTypeLogsModel.Port = core.Int64Ptr(int64(8080))

	model := new(logsroutingv1.TargetTypeLogs)
	model.ID = CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
	model.Name = core.StringPtr("my-log-sink")
	model.Etag = core.StringPtr("822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049")
	model.Type = core.StringPtr("logs")
	model.CreatedAt = CreateMockDateTime("2024-06-20T18:30:00.143Z")
	model.UpdatedAt = CreateMockDateTime("2024-06-20T18:30:00.143Z")
	model.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
	model.Parameters = targetParametersTypeLogsModel

	result, err := logsrouting.DataSourceIbmLogsRouterTenantsTargetTypeLogsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsRouterTenantsWriteStatusToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["status"] = "success"
		model["reason_for_last_failure"] = "Logs endpoint is not reachable. Received status code: 403"
		model["last_failure"] = "2024-10-14T10:49:09.000Z"

		assert.Equal(t, result, model)
	}

	model := new(logsroutingv1.WriteStatus)
	model.Status = core.StringPtr("success")
	model.ReasonForLastFailure = core.StringPtr("Logs endpoint is not reachable. Received status code: 403")
	model.LastFailure = CreateMockDateTime("2024-10-14T10:49:09.000Z")

	result, err := logsrouting.DataSourceIbmLogsRouterTenantsWriteStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
