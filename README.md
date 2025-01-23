# nova-api-client-go

A simple client for the [Wandelbots Nova](https://www.wandelbots.com/) API.

## basic usage

```bash
go get github.com/wandelbotsgmbh/nova-api-client-go/v25
```

Example:
```golang
func ListControllers(host string, cell string) ([]v25.ControllerInstance, error) {
	client, err := v25.NewClientWithResponses(host)
	if err != nil {
		return nil, err
	}

	resp, err := client.ListControllersWithResponse(context.TODO(), cell)
	if err != nil {
		return []v25.ControllerInstance{}, err
	}

	if !StatusSuccessfull(resp.StatusCode()) {
		return []v25.ControllerInstance{}, fmt.Errorf("failed to list controllers %s", resp.Status())
	}

	return resp.JSON200.Instances, nil
}
```
