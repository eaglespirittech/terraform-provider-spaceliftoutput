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