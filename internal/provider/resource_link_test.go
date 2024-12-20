package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccPritunlLink(t *testing.T) {

	t.Run("creates links without error", func(t *testing.T) {
		linkName := "tfacc-link1"

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("pritunl_link.test", "name", linkName),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { preCheck(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: testPritunlLinkConfig(linkName),
					Check:  check,
				},
				// import test
				importStep("pritunl_link.test"),
			},
		})
	})
}

func testPritunlLinkConfig(name string) string {
	return fmt.Sprintf(`
		resource "pritunl_link" "test" {
			name    = "%[1]s"
		}
	`, name)
}
