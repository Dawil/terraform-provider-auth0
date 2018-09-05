# Terraform Provider for Auth0

NB: super basic at present. This is the thin end of the wedge.

To help version control your auth0 configurations.

## Usage

Only a terraform data, not resource. Support:

* client_id
* name

You will need an Auth0 Application that is authorized to use the Auth0 Management API.

```
provider "auth0" {
  domain = "<OMITTED>.au.auth0.com"
  client_id = "<OMITTED>"
  client_secret = "<OMITTED>"
}

data "auth0_client" "foo" {
  client_id = "<OMITTED>"
}

output "name" {
  value = "${data.auth0_client.foo.name}"
}
```

## Developing

I create a `main.tf` file for testing purposes and run

```
make build && terraform init && terraform apply || cat foo.txt
```

`foo.txt` is a log file that Icreate by adding:

```
 import (
     "github.com/hashicorp/terraform/plugin"
     "github.com/hashicorp/terraform/terraform"
+    "log"
+    "os"
+)
+
+var (
+    outfile, _ = os.Create("foo.txt")
+    l = log.New(outfile, "", 0)
 )
```
 
To `main.go`. I can then log using `l.println()`. This is because error reporting from terraform providers isn't great.
