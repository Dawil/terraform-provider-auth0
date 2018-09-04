package main

import (
    "github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
    return &schema.Provider{
        DataSourcesMap: map[string]*schema.Resource{
            "auth0_client": dataSourceAuth0Client(),
        },
    }
}
