package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccPritunlLocation(t *testing.T) {

	t.Run("creates locations without error", func(t *testing.T) {
		linkID := "tfacc-link1"
		locationName := "tfacc-location1"

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("pritunl_location.test", "name", locationName),
			resource.TestCheckResourceAttr("pritunl_location.test", "link_id", linkID),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { preCheck(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: testPritunlLocationConfig(locationName),
					Check:  check,
				},
				// import test
				importStep("pritunl_location.test"),
			},
		})
	})
}

func testPritunlLocationConfig(name string) string {
	return fmt.Sprintf(`
		resource "pritunl_location" "test" {
			name    = "%[1]s"
		}
	`, name)
}
