data "merge_merge" "json_single" {
  input {
    format = "json"
    data = jsonencode({
      "hello" = "world"
    })
  }

  output_format = "json"
}
