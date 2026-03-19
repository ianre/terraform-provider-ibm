// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logsrouting_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logsrouting"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-router-go-sdk/logsroutingv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsRouterTenantBasic(t *testing.T) {
	var conf logsroutingv1.Tenant
	ibmApiVersion := fmt.Sprintf("tf_ibm_api_version_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	ibmApiVersionUpdate := fmt.Sprintf("tf_ibm_api_version_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsRouterTenantDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsRouterTenantConfigBasic(ibmApiVersion, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsRouterTenantExists("ibm_logs-router_tenant.logs_router_tenant_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs-router_tenant.logs_router_tenant_instance", "ibm_api_version", ibmApiVersion),
					resource.TestCheckResourceAttr("ibm_logs-router_tenant.logs_router_tenant_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsRouterTenantConfigBasic(ibmApiVersionUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs-router_tenant.logs_router_tenant_instance", "ibm_api_version", ibmApiVersionUpdate),
					resource.TestCheckResourceAttr("ibm_logs-router_tenant.logs_router_tenant_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs-router_tenant.logs_router_tenant_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsRouterTenantConfigBasic(ibmApiVersion string, name string) string {
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
	`, ibmApiVersion, name)
}

func testAccCheckIbmLogsRouterTenantExists(n string, obj logsroutingv1.Tenant) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsRoutingClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsRoutingV1()
		if err != nil {
			return err
		}

		getTenantDetailOptions := &logsroutingv1.GetTenantDetailOptions{}

		getTenantDetailOptions.SetTenantID(rs.Primary.ID)

		tenant, _, err := logsRoutingClient.GetTenantDetail(getTenantDetailOptions)
		if err != nil {
			return err
		}

		obj = *tenant
		return nil
	}
}

func testAccCheckIbmLogsRouterTenantDestroy(s *terraform.State) error {
	logsRoutingClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsRoutingV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs-router_tenant" {
			continue
		}

		getTenantDetailOptions := &logsroutingv1.GetTenantDetailOptions{}

		getTenantDetailOptions.SetTenantID(rs.Primary.ID)

		// Try to find the key
		_, response, err := logsRoutingClient.GetTenantDetail(getTenantDetailOptions)

		if err == nil {
			return fmt.Errorf("logs-router_tenant still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs-router_tenant (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmLogsRouterTenantTargetToMap(t *testing.T) {
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

	result, err := logsrouting.ResourceIbmLogsRouterTenantTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRouterTenantTargetParametersTypeLogsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["host"] = "www.example.com"
		model["port"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(logsroutingv1.TargetParametersTypeLogs)
	model.Host = core.StringPtr("www.example.com")
	model.Port = core.Int64Ptr(int64(1))

	result, err := logsrouting.ResourceIbmLogsRouterTenantTargetParametersTypeLogsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRouterTenantTargetTypeLogsToMap(t *testing.T) {
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

	result, err := logsrouting.ResourceIbmLogsRouterTenantTargetTypeLogsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRouterTenantWriteStatusToMap(t *testing.T) {
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

	result, err := logsrouting.ResourceIbmLogsRouterTenantWriteStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRouterTenantMapToTargetPrototype(t *testing.T) {
	checkResult := func(result logsroutingv1.TargetPrototypeIntf) {
		targetParametersTypeLogsPrototypeModel := new(logsroutingv1.TargetParametersTypeLogsPrototype)
		targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
		targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(1))

		model := new(logsroutingv1.TargetPrototype)
		model.Name = core.StringPtr("my-log-sink")
		model.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
		model.Parameters = targetParametersTypeLogsPrototypeModel

		assert.Equal(t, result, model)
	}

	targetParametersTypeLogsPrototypeModel := make(map[string]interface{})
	targetParametersTypeLogsPrototypeModel["host"] = "www.example.com"
	targetParametersTypeLogsPrototypeModel["port"] = int(1)

	model := make(map[string]interface{})
	model["name"] = "my-log-sink"
	model["log_sink_crn"] = "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::"
	model["parameters"] = []interface{}{targetParametersTypeLogsPrototypeModel}

	result, err := logsrouting.ResourceIbmLogsRouterTenantMapToTargetPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRouterTenantMapToTargetParametersTypeLogsPrototype(t *testing.T) {
	checkResult := func(result *logsroutingv1.TargetParametersTypeLogsPrototype) {
		model := new(logsroutingv1.TargetParametersTypeLogsPrototype)
		model.Host = core.StringPtr("www.example.com")
		model.Port = core.Int64Ptr(int64(1))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["host"] = "www.example.com"
	model["port"] = int(1)

	result, err := logsrouting.ResourceIbmLogsRouterTenantMapToTargetParametersTypeLogsPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRouterTenantMapToTargetPrototypeTargetTypeLogsPrototype(t *testing.T) {
	checkResult := func(result *logsroutingv1.TargetPrototypeTargetTypeLogsPrototype) {
		targetParametersTypeLogsPrototypeModel := new(logsroutingv1.TargetParametersTypeLogsPrototype)
		targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
		targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(8080))

		model := new(logsroutingv1.TargetPrototypeTargetTypeLogsPrototype)
		model.Name = core.StringPtr("my-log-sink")
		model.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
		model.Parameters = targetParametersTypeLogsPrototypeModel

		assert.Equal(t, result, model)
	}

	targetParametersTypeLogsPrototypeModel := make(map[string]interface{})
	targetParametersTypeLogsPrototypeModel["host"] = "www.example.com"
	targetParametersTypeLogsPrototypeModel["port"] = int(8080)

	model := make(map[string]interface{})
	model["name"] = "my-log-sink"
	model["log_sink_crn"] = "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::"
	model["parameters"] = []interface{}{targetParametersTypeLogsPrototypeModel}

	result, err := logsrouting.ResourceIbmLogsRouterTenantMapToTargetPrototypeTargetTypeLogsPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}
