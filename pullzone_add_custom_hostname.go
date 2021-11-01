package bunny

import (
	"context"
	"fmt"
)

// AddCustomHostnameOptions  represents the message that is sent to the
// Add Custom Hostname API Endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_addhostname
type AddCustomHostnameOptions struct {
	// Hostname the hostname to add. (Required)
	Hostname *string `json:"Hostname,omitempty"`
}

// AddCustomHostname adds a custom hostname to the Pull Zone.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_addhostname
func (s *PullZoneService) AddCustomHostname(ctx context.Context, pullZoneID int64, opts *AddCustomHostnameOptions) error {
	req, err := s.client.newPostRequest(fmt.Sprintf("pullzone/%d/addHostname", pullZoneID), opts)
	if err != nil {
		return err
	}

	return s.client.sendRequest(ctx, req, nil)
}
