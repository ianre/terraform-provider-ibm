// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoverySearchObjectsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoverySearchObjectsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_search_objects.baas_search_objects_instance", "id"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_search_objects.baas_search_objects_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_search_objects.baas_search_objects_instance", "objects.#"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoverySearchObjectsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_search_objects" "baas_search_objects_instance" {
			x_ibm_tenant_id = "%s"
		}
	`, tenantId)
}
