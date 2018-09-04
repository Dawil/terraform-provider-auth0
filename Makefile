build:
	go build -o terraform-provider-auth0
	cp terraform-provider-auth0 .terraform/plugins/linux_amd64/.
