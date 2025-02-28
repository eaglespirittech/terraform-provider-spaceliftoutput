---
page_title: "spaceliftoutput_stack_output Data Source - terraform-provider-spaceliftoutput"
subcategory: ""
description: |-
  Retrieves a single output from a Spacelift stack.
---

# spaceliftoutput_stack_output (Data Source)

This data source allows you to retrieve a specific output from a Spacelift stack. It fetches the output from the Spacelift API and makes it available in your Terraform configuration.

## Example Usage

```terraform
data "spaceliftoutput_stack_output" "example" {
  stack_id    = "your-stack-id"
  output_name = "output_name"
}

output "output_value" {
  value = data.spaceliftoutput_stack_output.example.value
}
```

## Schema

### Required

- **stack_id** (String) - The ID of the Spacelift stack.
- **output_name** (String) - The name of the output to retrieve.

### Read-Only

- **id** (String) - The ID of the data source. This is a combination of the stack_id and output_name.
- **value** (String) - The value of the specified output.
- **last_check** (String) - The timestamp of the last check. 