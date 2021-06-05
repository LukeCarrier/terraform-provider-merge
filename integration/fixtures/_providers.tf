terraform {
  required_version = "= 0.15.4"

  required_providers {
    merge = {
      source = "LukeCarrier/merge"
    }
  }
}

provider "merge" {}
