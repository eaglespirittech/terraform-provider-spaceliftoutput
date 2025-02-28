package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccStackOutputDataSource(t *testing.T) {
	t.Skip("Skipping acceptance test")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccStackOutputDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.spaceliftoutput_stack_output.test", "stack_id", "test-stack"),
					resource.TestCheckResourceAttr("data.spaceliftoutput_stack_output.test", "output_name", "test-output"),
					resource.TestCheckResourceAttrSet("data.spaceliftoutput_stack_output.test", "last_check"),
				),
			},
		},
	})
}

const testAccStackOutputDataSourceConfig = `
provider "spaceliftoutput" {}

data "spaceliftoutput_stack_output" "test" {
  stack_id    = "test-stack"
  output_name = "test-output"
}
`
