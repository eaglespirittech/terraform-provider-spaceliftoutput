data "spaceliftoutput_stack_output" "example" {
  stack_id    = "your-stack-id"
  output_name = "output_name"
}

output "output_value" {
  value = data.spaceliftoutput_stack_output.example.value
} 