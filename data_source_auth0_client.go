package main

import (
//    "errors"
    "io/ioutil"
    "fmt"
    "bytes"
    "encoding/json"
    "net/http"
    "github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAuth0Client() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceAuth0ClientRead,
        Schema: map[string]*schema.Schema{
            "client_id": {
                Type: schema.TypeString,
                Required: true,
            },
            "domain": {
                Type: schema.TypeString,
                Required: true,
            },
            "client_secret": {
                Type: schema.TypeString,
                Required: true,
            },
            "name": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func dataSourceAuth0ClientRead(d *schema.ResourceData, meta interface{}) error {
    clientId := d.Get("client_id").(string)
    d.SetId(clientId)
    domain := d.Get("domain").(string)
    clientSecret := d.Get("client_secret").(string)

    body := map[string]string{
        "grant_type": "client_credentials",
        "client_id": clientId,
        "client_secret": clientSecret,
        "audience": fmt.Sprintf(
            "https://%s/api/v2/", domain,
        ),
    }
    buffer := new(bytes.Buffer)

    json.NewEncoder(buffer).Encode(body)
    resp, _ := http.Post(
        fmt.Sprintf(
            "https://%s/oauth/token",
            domain,
        ),
        "application/json",
        buffer,
    )
    b, _ := ioutil.ReadAll(resp.Body)
    var respbody map[string]interface{}
    json.Unmarshal(b, &respbody)

    // get access token
    accessToken := respbody["access_token"].(string)

    // get client and set attributes
    client := &http.Client{}
    req, _ := http.NewRequest(
        "GET",
        fmt.Sprintf(
            "https://%s/api/v2/clients/%s",
            domain,
            clientId,
        ),
        nil,
    )
    req.Header.Set(
        "Content-Type",
        "application/json",
    )
    req.Header.Set(
        "Authorization",
        fmt.Sprintf(
            "Bearer %s",
            accessToken,
        ),
    )
    res2, _ := client.Do(req)
    b2, _ := ioutil.ReadAll(res2.Body)
    //out := new(bytes.Buffer)
    //out.ReadFrom(res2.Body)
    var respbody2 map[string]interface{}
    json.Unmarshal(b2, &respbody2)

    d.Set("name", respbody2["name"].(string))

    //return errors.New(out.String())
    return nil
}
