package proxmox

import (
	"net/http"
	// …
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	authHeader string
}

func NewProxmoxClient(cfg ProxmoxConfig) *Client {
	// build baseURL = "https://"+cfg.Host+"/api2/json"
	// build authHeader = "PVEAPIToken="+cfg.User+"!"+cfg.TokenName+"="+cfg.TokenValue
	// configure HTTP client (TLS skip if needed)
}

func (c *Client) CreateBridge(node, bridgeName string) error {
	// TODO: POST /nodes/{node}/network
}

func (c *Client) EnsureVM(node string, vmid int, vmDefaults ProxmoxConfig) (int, error) {
	// TODO: GET then POST if 404
}

func (c *Client) AttachNIC(vmid int, bridge string, index int) error {
	// TODO: PUT /nodes/{node}/qemu/{vmid}/config
}
