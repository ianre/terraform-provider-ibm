// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logsrouting_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logsrouting"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-router-go-sdk/ibmcloudlogsroutingv0"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsRouterTenantBasic(t *testing.T) {
	var conf ibmcloudlogsroutingv0.Tenant
	name := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsRouterTenantDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsRouterTenantConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsRouterTenantExists("ibm_logs_router_tenant.logs_router_tenant_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_router_tenant.logs_router_tenant_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsRouterTenantConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_router_tenant.logs_router_tenant_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_router_tenant.logs_router_tenant",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsRouterTenantConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_router_tenant" "logs_router_tenant_instance" {
			name = "%s"
			targets {
				name = "my-log-sink"
				log_sink_crn = "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::"
				parameters {
					host = "www.example.com"
					port = 8080
				}
			}
		}
	`, name)
}

func testAccCheckIbmLogsRouterTenantExists(n string, obj ibmcloudlogsroutingv0.Tenant) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		ibmCloudLogsRoutingClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsRoutingV0()
		if err != nil {
			return err
		}

		getTenantDetailOptions := &ibmcloudlogsroutingv0.GetTenantDetailOptions{}
		tenantId := strfmt.UUID(rs.Primary.ID)
		getTenantDetailOptions.SetTenantID(&tenantId)

		tenant, _, err := ibmCloudLogsRoutingClient.GetTenantDetail(getTenantDetailOptions)
		if err != nil {
			return err
		}

		obj = *tenant
		return nil
	}
}

func testAccCheckIbmLogsRouterTenantDestroy(s *terraform.State) error {
	ibmCloudLogsRoutingClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsRoutingV0()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_router_tenant" {
			continue
		}

		getTenantDetailOptions := &ibmcloudlogsroutingv0.GetTenantDetailOptions{}
		tenantId := strfmt.UUID(rs.Primary.ID)
		getTenantDetailOptions.SetTenantID(&tenantId)

		// Try to find the key
		_, response, err := ibmCloudLogsRoutingClient.GetTenantDetail(getTenantDetailOptions)

		if err == nil {
			return fmt.Errorf("logs_router_tenant still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_router_tenant (%s) has been destroyed: %s", rs.Primary.ID, err)
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
		model["created_at"] = "2024-06-20T18:30:00.143156 + 0000 UTC"
		model["updated_at"] = "2024-06-20T18:30:00.143156 + 0000 UTC"
		model["log_sink_crn"] = "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::"
		model["parameters"] = []map[string]interface{}{targetParametersTypeLogsModel}

		assert.Equal(t, result, model)
	}

	targetParametersTypeLogsModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogs)
	targetParametersTypeLogsModel.Host = core.StringPtr("www.example.com")
	targetParametersTypeLogsModel.Port = core.Int64Ptr(int64(8080))

	model := new(ibmcloudlogsroutingv0.Target)
	model.ID = CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
	model.Name = core.StringPtr("my-log-sink")
	model.Etag = core.StringPtr("822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049")
	model.Type = core.StringPtr("logs")
	model.CreatedAt = core.StringPtr("2024-06-20T18:30:00.143156 + 0000 UTC")
	model.UpdatedAt = core.StringPtr("2024-06-20T18:30:00.143156 + 0000 UTC")
	model.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
	model.Parameters = targetParametersTypeLogsModel

	result, err := logsrouting.ResourceIbmLogsRouterTenantTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

/*
func TestResourceIbmLogsRouterTenantTargetParametersTypeLogsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["host"] = "www.example.com"
		model["port"] = int(8080)

		assert.Equal(t, result, model)
	}

	model := new(ibmcloudlogsroutingv0.TargetParametersTypeLogs)
	model.Host = core.StringPtr("www.example.com")
	model.Port = core.Int64Ptr(int64(8080))

	result, err := ibmcloudlogsrouting.ResourceIbmLogsRouterTenantTargetParametersTypeLogsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
*/
/*
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
		model["created_at"] = "2024-06-20T18:30:00.143156 + 0000 UTC"
		model["updated_at"] = "2024-06-20T18:30:00.143156 + 0000 UTC"
		model["log_sink_crn"] = "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::"
		model["parameters"] = []map[string]interface{}{targetParametersTypeLogsModel}

		assert.Equal(t, result, model)
	}

	targetParametersTypeLogsModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogs)
	targetParametersTypeLogsModel.Host = core.StringPtr("www.example.com")
	targetParametersTypeLogsModel.Port = core.Int64Ptr(int64(8080))

	model := new(ibmcloudlogsroutingv0.TargetTypeLogs)
	model.ID = CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
	model.Name = core.StringPtr("my-log-sink")
	model.Etag = core.StringPtr("822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049")
	model.Type = core.StringPtr("logs")
	model.CreatedAt = core.StringPtr("2024-06-20T18:30:00.143156 + 0000 UTC")
	model.UpdatedAt = core.StringPtr("2024-06-20T18:30:00.143156 + 0000 UTC")
	model.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
	model.Parameters = targetParametersTypeLogsModel

	result, err := ibmcloudlogsrouting.ResourceIbmLogsRouterTenantTargetTypeLogsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
*/

/*
func TestResourceIbmLogsRouterTenantWriteStatusToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["status"] = "success"
		model["reason_for_last_failure"] = "Logs endpoint is not reachable. Received status code: 403"
		model["last_failure"] = "2024-10-14T10:49:09 + 0000 UTC"

		assert.Equal(t, result, model)
	}

	model := new(ibmcloudlogsroutingv0.WriteStatus)
	model.Status = core.StringPtr("success")
	model.ReasonForLastFailure = core.StringPtr("Logs endpoint is not reachable. Received status code: 403")
	model.LastFailure = core.StringPtr("2024-10-14T10:49:09 + 0000 UTC")

	result, err := ibmcloudlogsrouting.ResourceIbmLogsRouterTenantWriteStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
*/

func TestResourceIbmLogsRouterTenantMapToTargetPrototype(t *testing.T) {
	checkResult := func(result ibmcloudlogsroutingv0.TargetPrototypeIntf) {
		targetParametersTypeLogsPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogsPrototype)
		targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
		targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(8080))

		model := new(ibmcloudlogsroutingv0.TargetPrototype)
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

	result, err := logsrouting.ResourceIbmLogsRouterTenantMapToTargetPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

/*
func TestResourceIbmLogsRouterTenantMapToTargetParametersTypeLogsPrototype(t *testing.T) {
	checkResult := func(result *ibmcloudlogsroutingv0.TargetParametersTypeLogsPrototype) {
		model := new(ibmcloudlogsroutingv0.TargetParametersTypeLogsPrototype)
		model.Host = core.StringPtr("www.example.com")
		model.Port = core.Int64Ptr(int64(8080))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["host"] = "www.example.com"
	model["port"] = int(8080)

	result, err := ibmcloudlogsrouting.ResourceIbmLogsRouterTenantMapToTargetParametersTypeLogsPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}
*/
/*
func TestResourceIbmLogsRouterTenantMapToTargetPrototypeTargetTypeLogsPrototype(t *testing.T) {
	checkResult := func(result *ibmcloudlogsroutingv0.TargetPrototypeTargetTypeLogsPrototype) {
		targetParametersTypeLogsPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogsPrototype)
		targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
		targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(8080))

		model := new(ibmcloudlogsroutingv0.TargetPrototypeTargetTypeLogsPrototype)
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

	result, err := ibmcloudlogsrouting.ResourceIbmLogsRouterTenantMapToTargetPrototypeTargetTypeLogsPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}
*/
