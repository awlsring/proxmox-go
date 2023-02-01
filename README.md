# Proxmox Go Client

This is a Go client library for the Proxmox API. It is automatically generated from my [Proxmox Smithy Model](https://github.com/awlsring/ProxmoxModel).

This library is not yet complete, but it is functional. Supported operations can be found in `docs` or in [the model package.](https://github.com/awlsring/ProxmoxModel)

## Use

You can create a client for your proxmox instance with the following code.

```golang
import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/awlsring/proxmox-go/proxmox"
)

func main() {
	endpoint := os.Getenv("PROXMOX_ENDPOINT")
	token := os.Getenv("PROXMOX_TOKEN")
	cfg := proxmox.NewConfiguration()
	cfg.Servers[0] = proxmox.ServerConfiguration{
		URL: fmt.Sprintf("%s/api2/json", endpoint),
	}
	cfg.HTTPClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 10,
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		},
	}
	
	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("PVEAPIToken=%s", token))

	// this is a client that can call methods on the API
	client := proxmox.NewAPIClient(cfg)
}
```

The `client` in this example will have methods for all of the operations supported in the model.