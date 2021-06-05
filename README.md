# Merge Terraform provider

Deep merge your JSON and YAML objects. Builds on [cloudposse/utils](https://github.com/cloudposse/terraform-provider-utils), allowing mixing of JSON and YAML inputs.

---

## Usage

Initialise the provider:

```hcl
terraform {
  required_providers {
    merge = {
      source = "LukeCarrier/merge"
    }
  }
}

provider "merge" {}
```

Then use the data source:

```hcl
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

output "json_multi" {
  value = data.merge_merge.json_multi.output
}
```

## Hacking

Get the dependencies:

```console
go mod download
go build
```

Run the unit tests:

```console
go test ./...
```

To test a Terraform state, you'll need to create a `.terraformrc` and set `TF_CLI_CONFIG_FILE` to its path. See `/integration/fixtures/example.terraformrc` for an example.

## Signing

Generate a new key (use [Keybase](https://keybase.io/), GnuPG's CLI is a dumpster fire), then import it into a temporary local keyring:

```console
export GNUPGHOME="$(mktemp -d)"
keybase pgp export -q email@addr -s \
    | gpg --import --allow-secret-key-import --pinentry-mode=loopback
```

Export the private key, and paste the result into a new GitHub secret named `GPG_PRIVATE_KEY`:

```console
gpg --armor --export-secret-key email@addr
```

Add the passphrase to a secret named `PASSPHRASE`.

Finally, export the public key and paste the result into the [Terraform registry](https://registry.terraform.io/settings/gpg-keys):

```console
gpg --armor --export email@addr
```
