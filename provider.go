package main

import (
    "github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
    return &schema.Provider{
        Schema: map[string]*schema.Schema{
            "domain": &schema.Schema{
                Type: schema.TypeString,
                Required: true,
            },
            "client_id": &schema.Schema{
                Type: schema.TypeString,
                Required: true,
            },
            "client_secret": &schema.Schema{
                Type: schema.TypeString,
                Required: true,
            },
        },
        DataSourcesMap: map[string]*schema.Resource{
            "auth0_client": dataSourceAuth0Client(),
        },
        ConfigureFunc: providerConfigure,
    }
}

type Config struct {
    Domain string
    ClientId string
    ClientSecret string
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
    config := Config{
        Domain: d.Get("domain").(string),
        ClientId: d.Get("client_id").(string),
        ClientSecret: d.Get("client_secret").(string),
    }
    return &config, nil
}
