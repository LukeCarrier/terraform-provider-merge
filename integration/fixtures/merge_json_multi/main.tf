data "merge_merge" "json_multi" {
  input {
    format = "json"
    data = jsonencode({
      "hello" = "world"
    })
  }

  input {
    format = "json"
    data = jsonencode({
      "hello" = "galaxy"
    })
  }

  output_format = "json"
}
