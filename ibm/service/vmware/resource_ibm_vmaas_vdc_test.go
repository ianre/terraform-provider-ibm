// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vmware_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vmware"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vmware-go-sdk/vmwarev1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmVmaasVdcBasic(t *testing.T) {
	var conf vmwarev1.VDC

	ds_id := acc.Vmaas_Directorsite_id
	ds_pvdc_id := acc.Vmaas_Directorsite_pvdc_id
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckVMwareService(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmVmaasVdcDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmVmaasVdcConfigBasic(ds_id, ds_pvdc_id, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmVmaasVdcExists("ibm_vmaas_vdc.vmaas_vdc_instance", conf),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "name", name),
				),
			},
		},
	})
}

func TestAccIbmVmaasVdcAllArgs(t *testing.T) {
	var conf vmwarev1.VDC

	ds_id := acc.Vmaas_Directorsite_id
	ds_pvdc_id := acc.Vmaas_Directorsite_pvdc_id

	cpu := fmt.Sprintf("%d", acctest.RandIntRange(0, 2000))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	ram := fmt.Sprintf("%d", acctest.RandIntRange(0, 40960))
	fastProvisioningEnabled := "false"
	rhelByol := "false"
	windowsByol := "false"

	cpuUpdate := fmt.Sprintf("%d", acctest.RandIntRange(0, 2000))
	nameUpdate := name
	ramUpdate := fmt.Sprintf("%d", acctest.RandIntRange(0, 40960))
	fastProvisioningEnabledUpdate := fastProvisioningEnabled
	rhelByolUpdate := rhelByol
	windowsByolUpdate := windowsByol

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckVMwareService(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmVmaasVdcDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmVmaasVdcConfig(cpu, ds_id, ds_pvdc_id, name, ram, fastProvisioningEnabled, rhelByol, windowsByol),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmVmaasVdcExists("ibm_vmaas_vdc.vmaas_vdc_instance", conf),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "cpu", cpu),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "ram", ram),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "fast_provisioning_enabled", fastProvisioningEnabled),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "rhel_byol", rhelByol),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "windows_byol", windowsByol),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmVmaasVdcConfig(cpuUpdate, ds_id, ds_pvdc_id, nameUpdate, ramUpdate, fastProvisioningEnabledUpdate, rhelByolUpdate, windowsByolUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "cpu", cpuUpdate),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "ram", ramUpdate),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "fast_provisioning_enabled", fastProvisioningEnabledUpdate),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "rhel_byol", rhelByolUpdate),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "windows_byol", windowsByolUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_vmaas_vdc.vmaas_vdc_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmVmaasVdcConfigBasic(ds_id string, ds_pvdc_id string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_vmaas_vdc" "vmaas_vdc_instance" {
			director_site {
				id = "%s"
				pvdc {
					id = "%s"
					provider_type {
						name = "on_demand"
					}
				}
			}
			name = "%s"
		}
	`, ds_id, ds_pvdc_id, name)
}

func testAccCheckIbmVmaasVdcConfig(cpu string, ds_id string, ds_pvdc_id string, name string, ram string, fastProvisioningEnabled string, rhelByol string, windowsByol string) string {
	return fmt.Sprintf(`

		resource "ibm_vmaas_vdc" "vmaas_vdc_instance" {
			cpu = %s
			director_site {
				id = "%s"
				pvdc {
					id = "%s"
					provider_type {
						name = "on_demand"
					}
				}
			}
			name = "%s"
			ram = %s
			fast_provisioning_enabled = %s
			rhel_byol = %s
			windows_byol = %s
		}
	`, cpu, ds_id, ds_pvdc_id, name, ram, fastProvisioningEnabled, rhelByol, windowsByol)
}

func testAccCheckIbmVmaasVdcExists(n string, obj vmwarev1.VDC) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vmwareClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VmwareV1()
		if err != nil {
			return err
		}

		getVdcOptions := &vmwarev1.GetVdcOptions{}

		getVdcOptions.SetID(rs.Primary.ID)

		vDC, _, err := vmwareClient.GetVdc(getVdcOptions)
		if err != nil {
			return err
		}

		obj = *vDC
		return nil
	}
}

func testAccCheckIbmVmaasVdcDestroy(s *terraform.State) error {
	vmwareClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VmwareV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_vmaas_vdc" {
			continue
		}

		getVdcOptions := &vmwarev1.GetVdcOptions{}

		getVdcOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := vmwareClient.GetVdc(getVdcOptions)

		if err == nil {
			return fmt.Errorf("vmaas_vdc still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for vmaas_vdc (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmVmaasVdcVDCDirectorSiteToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		vdcProviderTypeModel := make(map[string]interface{})
		vdcProviderTypeModel["name"] = "paygo"

		directorSitePvdcModel := make(map[string]interface{})
		directorSitePvdcModel["id"] = "pvdc_id"
		directorSitePvdcModel["provider_type"] = []map[string]interface{}{vdcProviderTypeModel}

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["pvdc"] = []map[string]interface{}{directorSitePvdcModel}
		model["url"] = "testString"

		assert.Equal(t, result, model)
	}

	vdcProviderTypeModel := new(vmwarev1.VDCProviderType)
	vdcProviderTypeModel.Name = core.StringPtr("paygo")

	directorSitePvdcModel := new(vmwarev1.DirectorSitePVDC)
	directorSitePvdcModel.ID = core.StringPtr("pvdc_id")
	directorSitePvdcModel.ProviderType = vdcProviderTypeModel

	model := new(vmwarev1.VDCDirectorSite)
	model.ID = core.StringPtr("testString")
	model.Pvdc = directorSitePvdcModel
	model.URL = core.StringPtr("testString")

	result, err := vmware.ResourceIbmVmaasVdcVDCDirectorSiteToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmVmaasVdcDirectorSitePVDCToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		vdcProviderTypeModel := make(map[string]interface{})
		vdcProviderTypeModel["name"] = "paygo"

		model := make(map[string]interface{})
		model["id"] = "pvdc_id"
		model["provider_type"] = []map[string]interface{}{vdcProviderTypeModel}

		assert.Equal(t, result, model)
	}

	vdcProviderTypeModel := new(vmwarev1.VDCProviderType)
	vdcProviderTypeModel.Name = core.StringPtr("paygo")

	model := new(vmwarev1.DirectorSitePVDC)
	model.ID = core.StringPtr("pvdc_id")
	model.ProviderType = vdcProviderTypeModel

	result, err := vmware.ResourceIbmVmaasVdcDirectorSitePVDCToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmVmaasVdcVDCProviderTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "paygo"

		assert.Equal(t, result, model)
	}

	model := new(vmwarev1.VDCProviderType)
	model.Name = core.StringPtr("paygo")

	result, err := vmware.ResourceIbmVmaasVdcVDCProviderTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmVmaasVdcEdgeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		transitGatewayConnectionModel := make(map[string]interface{})
		transitGatewayConnectionModel["name"] = "testString"
		transitGatewayConnectionModel["transit_gateway_connection_name"] = "testString"
		transitGatewayConnectionModel["status"] = "pending"
		transitGatewayConnectionModel["local_gateway_ip"] = "testString"
		transitGatewayConnectionModel["remote_gateway_ip"] = "testString"
		transitGatewayConnectionModel["local_tunnel_ip"] = "testString"
		transitGatewayConnectionModel["remote_tunnel_ip"] = "testString"
		transitGatewayConnectionModel["local_bgp_asn"] = int(1)
		transitGatewayConnectionModel["remote_bgp_asn"] = int(1)
		transitGatewayConnectionModel["network_account_id"] = "testString"
		transitGatewayConnectionModel["network_type"] = "testString"
		transitGatewayConnectionModel["base_network_type"] = "testString"
		transitGatewayConnectionModel["zone"] = "testString"

		transitGatewayModel := make(map[string]interface{})
		transitGatewayModel["id"] = "testString"
		transitGatewayModel["connections"] = []map[string]interface{}{transitGatewayConnectionModel}
		transitGatewayModel["status"] = "pending"
		transitGatewayModel["region"] = "testString"

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["public_ips"] = []string{"testString"}
		model["private_ips"] = []string{"testString"}
		model["private_only"] = true
		model["size"] = "medium"
		model["status"] = "creating"
		model["transit_gateways"] = []map[string]interface{}{transitGatewayModel}
		model["type"] = "performance"
		model["version"] = "testString"

		assert.Equal(t, result, model)
	}

	transitGatewayConnectionModel := new(vmwarev1.TransitGatewayConnection)
	transitGatewayConnectionModel.Name = core.StringPtr("testString")
	transitGatewayConnectionModel.TransitGatewayConnectionName = core.StringPtr("testString")
	transitGatewayConnectionModel.Status = core.StringPtr("pending")
	transitGatewayConnectionModel.LocalGatewayIp = core.StringPtr("testString")
	transitGatewayConnectionModel.RemoteGatewayIp = core.StringPtr("testString")
	transitGatewayConnectionModel.LocalTunnelIp = core.StringPtr("testString")
	transitGatewayConnectionModel.RemoteTunnelIp = core.StringPtr("testString")
	transitGatewayConnectionModel.LocalBgpAsn = core.Int64Ptr(int64(1))
	transitGatewayConnectionModel.RemoteBgpAsn = core.Int64Ptr(int64(1))
	transitGatewayConnectionModel.NetworkAccountID = core.StringPtr("testString")
	transitGatewayConnectionModel.NetworkType = core.StringPtr("testString")
	transitGatewayConnectionModel.BaseNetworkType = core.StringPtr("testString")
	transitGatewayConnectionModel.Zone = core.StringPtr("testString")

	transitGatewayModel := new(vmwarev1.TransitGateway)
	transitGatewayModel.ID = core.StringPtr("testString")
	transitGatewayModel.Connections = []vmwarev1.TransitGatewayConnection{*transitGatewayConnectionModel}
	transitGatewayModel.Status = core.StringPtr("pending")
	transitGatewayModel.Region = core.StringPtr("testString")

	model := new(vmwarev1.Edge)
	model.ID = core.StringPtr("testString")
	model.PublicIps = []string{"testString"}
	model.PrivateIps = []string{"testString"}
	model.PrivateOnly = core.BoolPtr(true)
	model.Size = core.StringPtr("medium")
	model.Status = core.StringPtr("creating")
	model.TransitGateways = []vmwarev1.TransitGateway{*transitGatewayModel}
	model.Type = core.StringPtr("performance")
	model.Version = core.StringPtr("testString")

	result, err := vmware.ResourceIbmVmaasVdcEdgeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmVmaasVdcTransitGatewayToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		transitGatewayConnectionModel := make(map[string]interface{})
		transitGatewayConnectionModel["name"] = "testString"
		transitGatewayConnectionModel["transit_gateway_connection_name"] = "testString"
		transitGatewayConnectionModel["status"] = "pending"
		transitGatewayConnectionModel["local_gateway_ip"] = "testString"
		transitGatewayConnectionModel["remote_gateway_ip"] = "testString"
		transitGatewayConnectionModel["local_tunnel_ip"] = "testString"
		transitGatewayConnectionModel["remote_tunnel_ip"] = "testString"
		transitGatewayConnectionModel["local_bgp_asn"] = int(1)
		transitGatewayConnectionModel["remote_bgp_asn"] = int(1)
		transitGatewayConnectionModel["network_account_id"] = "testString"
		transitGatewayConnectionModel["network_type"] = "testString"
		transitGatewayConnectionModel["base_network_type"] = "testString"
		transitGatewayConnectionModel["zone"] = "testString"

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["connections"] = []map[string]interface{}{transitGatewayConnectionModel}
		model["status"] = "pending"
		model["region"] = "testString"

		assert.Equal(t, result, model)
	}

	transitGatewayConnectionModel := new(vmwarev1.TransitGatewayConnection)
	transitGatewayConnectionModel.Name = core.StringPtr("testString")
	transitGatewayConnectionModel.TransitGatewayConnectionName = core.StringPtr("testString")
	transitGatewayConnectionModel.Status = core.StringPtr("pending")
	transitGatewayConnectionModel.LocalGatewayIp = core.StringPtr("testString")
	transitGatewayConnectionModel.RemoteGatewayIp = core.StringPtr("testString")
	transitGatewayConnectionModel.LocalTunnelIp = core.StringPtr("testString")
	transitGatewayConnectionModel.RemoteTunnelIp = core.StringPtr("testString")
	transitGatewayConnectionModel.LocalBgpAsn = core.Int64Ptr(int64(1))
	transitGatewayConnectionModel.RemoteBgpAsn = core.Int64Ptr(int64(1))
	transitGatewayConnectionModel.NetworkAccountID = core.StringPtr("testString")
	transitGatewayConnectionModel.NetworkType = core.StringPtr("testString")
	transitGatewayConnectionModel.BaseNetworkType = core.StringPtr("testString")
	transitGatewayConnectionModel.Zone = core.StringPtr("testString")

	model := new(vmwarev1.TransitGateway)
	model.ID = core.StringPtr("testString")
	model.Connections = []vmwarev1.TransitGatewayConnection{*transitGatewayConnectionModel}
	model.Status = core.StringPtr("pending")
	model.Region = core.StringPtr("testString")

	result, err := vmware.ResourceIbmVmaasVdcTransitGatewayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmVmaasVdcTransitGatewayConnectionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["transit_gateway_connection_name"] = "testString"
		model["status"] = "pending"
		model["local_gateway_ip"] = "testString"
		model["remote_gateway_ip"] = "testString"
		model["local_tunnel_ip"] = "testString"
		model["remote_tunnel_ip"] = "testString"
		model["local_bgp_asn"] = int(1)
		model["remote_bgp_asn"] = int(1)
		model["network_account_id"] = "testString"
		model["network_type"] = "testString"
		model["base_network_type"] = "testString"
		model["zone"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(vmwarev1.TransitGatewayConnection)
	model.Name = core.StringPtr("testString")
	model.TransitGatewayConnectionName = core.StringPtr("testString")
	model.Status = core.StringPtr("pending")
	model.LocalGatewayIp = core.StringPtr("testString")
	model.RemoteGatewayIp = core.StringPtr("testString")
	model.LocalTunnelIp = core.StringPtr("testString")
	model.RemoteTunnelIp = core.StringPtr("testString")
	model.LocalBgpAsn = core.Int64Ptr(int64(1))
	model.RemoteBgpAsn = core.Int64Ptr(int64(1))
	model.NetworkAccountID = core.StringPtr("testString")
	model.NetworkType = core.StringPtr("testString")
	model.BaseNetworkType = core.StringPtr("testString")
	model.Zone = core.StringPtr("testString")

	result, err := vmware.ResourceIbmVmaasVdcTransitGatewayConnectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmVmaasVdcStatusReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "insufficent_cpu"
		model["message"] = "testString"
		model["more_info"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(vmwarev1.StatusReason)
	model.Code = core.StringPtr("insufficent_cpu")
	model.Message = core.StringPtr("testString")
	model.MoreInfo = core.StringPtr("testString")

	result, err := vmware.ResourceIbmVmaasVdcStatusReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmVmaasVdcMapToVDCDirectorSitePrototype(t *testing.T) {
	checkResult := func(result *vmwarev1.VDCDirectorSitePrototype) {
		vdcProviderTypeModel := new(vmwarev1.VDCProviderType)
		vdcProviderTypeModel.Name = core.StringPtr("paygo")

		directorSitePvdcModel := new(vmwarev1.DirectorSitePVDC)
		directorSitePvdcModel.ID = core.StringPtr("pvdc_id")
		directorSitePvdcModel.ProviderType = vdcProviderTypeModel

		model := new(vmwarev1.VDCDirectorSitePrototype)
		model.ID = core.StringPtr("site_id")
		model.Pvdc = directorSitePvdcModel

		assert.Equal(t, result, model)
	}

	vdcProviderTypeModel := make(map[string]interface{})
	vdcProviderTypeModel["name"] = "paygo"

	directorSitePvdcModel := make(map[string]interface{})
	directorSitePvdcModel["id"] = "pvdc_id"
	directorSitePvdcModel["provider_type"] = []interface{}{vdcProviderTypeModel}

	model := make(map[string]interface{})
	model["id"] = "site_id"
	model["pvdc"] = []interface{}{directorSitePvdcModel}

	result, err := vmware.ResourceIbmVmaasVdcMapToVDCDirectorSitePrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmVmaasVdcMapToDirectorSitePVDC(t *testing.T) {
	checkResult := func(result *vmwarev1.DirectorSitePVDC) {
		vdcProviderTypeModel := new(vmwarev1.VDCProviderType)
		vdcProviderTypeModel.Name = core.StringPtr("paygo")

		model := new(vmwarev1.DirectorSitePVDC)
		model.ID = core.StringPtr("pvdc_id")
		model.ProviderType = vdcProviderTypeModel

		assert.Equal(t, result, model)
	}

	vdcProviderTypeModel := make(map[string]interface{})
	vdcProviderTypeModel["name"] = "paygo"

	model := make(map[string]interface{})
	model["id"] = "pvdc_id"
	model["provider_type"] = []interface{}{vdcProviderTypeModel}

	result, err := vmware.ResourceIbmVmaasVdcMapToDirectorSitePVDC(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmVmaasVdcMapToVDCProviderType(t *testing.T) {
	checkResult := func(result *vmwarev1.VDCProviderType) {
		model := new(vmwarev1.VDCProviderType)
		model.Name = core.StringPtr("paygo")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "paygo"

	result, err := vmware.ResourceIbmVmaasVdcMapToVDCProviderType(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmVmaasVdcMapToVDCEdgePrototype(t *testing.T) {
	checkResult := func(result *vmwarev1.VDCEdgePrototype) {
		model := new(vmwarev1.VDCEdgePrototype)
		model.Size = core.StringPtr("medium")
		model.Type = core.StringPtr("performance")
		model.PrivateOnly = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["size"] = "medium"
	model["type"] = "performance"
	model["private_only"] = true

	result, err := vmware.ResourceIbmVmaasVdcMapToVDCEdgePrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmVmaasVdcMapToResourceGroupIdentity(t *testing.T) {
	checkResult := func(result *vmwarev1.ResourceGroupIdentity) {
		model := new(vmwarev1.ResourceGroupIdentity)
		model.ID = core.StringPtr("some_resourcegroupid")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "some_resourcegroupid"

	result, err := vmware.ResourceIbmVmaasVdcMapToResourceGroupIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}
