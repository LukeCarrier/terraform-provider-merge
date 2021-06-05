data "merge_merge" "yaml_multi" {
  input {
    format = "yaml"
    data = yamlencode({
      "hello" = "world"
    })
  }

  input {
    format = "yaml"
    data = yamlencode({
      "hello" = "galaxy"
    })
  }

  output_format = "yaml"
}
