---
page_title: "spaceliftoutput_stack_outputs Data Source - terraform-provider-spaceliftoutput"
subcategory: ""
description: |-
  Retrieves all outputs from a Spacelift stack.
---

# spaceliftoutput_stack_outputs (Data Source)

This data source allows you to retrieve all outputs from a Spacelift stack. It fetches the outputs from the Spacelift API and makes them available in your Terraform configuration.

## Example Usage

```terraform
data "spaceliftoutput_stack_outputs" "example" {
  stack_id = "your-stack-id"
}

output "all_stack_outputs" {
  value = data.spaceliftoutput_stack_outputs.example.outputs
}

# Example of using a specific output
output "specific_output" {
  value = data.spaceliftoutput_stack_outputs.example.outputs["output_name"]
}
```

## Schema

### Required

- **stack_id** (String) - The ID of the Spacelift stack.

### Read-Only

- **id** (String) - The ID of the data source. This is the same as the stack_id.
- **outputs** (Map of String) - The outputs of the Spacelift stack. The keys are the output names and the values are the output values.
- **last_check** (String) - The timestamp of the last check. 