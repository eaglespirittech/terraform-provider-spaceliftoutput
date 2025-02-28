package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccStackOutputsDataSource(t *testing.T) {
	t.Skip("Skipping acceptance test")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccStackOutputsDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.spaceliftoutput_stack_outputs.test", "stack_id", "test-stack"),
					resource.TestCheckResourceAttrSet("data.spaceliftoutput_stack_outputs.test", "last_check"),
				),
			},
		},
	})
}

const testAccStackOutputsDataSourceConfig = `
provider "spaceliftoutput" {}

data "spaceliftoutput_stack_outputs" "test" {
  stack_id = "test-stack"
}
`
