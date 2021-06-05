data "merge_merge" "yaml_single" {
  input {
    format = "yaml"
    data = yamlencode({
      "hello" = "world"
    })
  }

  output_format = "yaml"
}
