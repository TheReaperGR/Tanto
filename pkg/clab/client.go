package clab

import "net/http"

type Client struct {
    baseURL string
}

func NewClabClient(cfg ClabConfig) *Client {
    // return &Client{"http://"+cfg.APIHost+":"+port+"/v1"}
}

func (c *Client) Deploy(topoFilePath string) (string, error) {
    // read file → POST c.baseURL+"/deploy"
    // parse the JSON response for lab_id
}

func (c *Client) Destroy(labID string) error {
    // DELETE c.baseURL+"/destroy/"+labID
}
