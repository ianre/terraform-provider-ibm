// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.112.0-f88e9264-20260220-115155
 */

package logsrouting_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logsrouting"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-router-go-sdk/logsroutingv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsRouterTargetsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsRouterTargetsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs-router_targets.logs_router_targets_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs-router_targets.logs_router_targets_instance", "ibm_api_version"),
					resource.TestCheckResourceAttrSet("data.ibm_logs-router_targets.logs_router_targets_instance", "tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsRouterTargetsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_logs-router_targets" "logs_router_targets_instance" {
			IBM-API-Version = "IBM-API-Version"
			tenant_id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
			name = "my-log-sink"
		}
	`)
}

func TestDataSourceIbmLogsRouterTargetsTargetToMap(t *testing.T) {
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

	result, err := logsrouting.DataSourceIbmLogsRouterTargetsTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsRouterTargetsTargetParametersTypeLogsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["host"] = "www.example.com"
		model["port"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(logsroutingv1.TargetParametersTypeLogs)
	model.Host = core.StringPtr("www.example.com")
	model.Port = core.Int64Ptr(int64(1))

	result, err := logsrouting.DataSourceIbmLogsRouterTargetsTargetParametersTypeLogsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsRouterTargetsTargetTypeLogsToMap(t *testing.T) {
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

	result, err := logsrouting.DataSourceIbmLogsRouterTargetsTargetTypeLogsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
