package bootstrap

import (
	"context"

	"github.com/hashicorp/consul/api"
	v1 "github.com/solo-io/gloo/projects/controller/pkg/api/v1"
	"github.com/solo-io/gloo/projects/controller/pkg/bootstrap/clients"
)

// Deprecated. Use bootstrap/clients
func ConsulClientForSettings(ctx context.Context, settings *v1.Settings) (*api.Client, error) {
	return clients.ConsulClientForSettings(ctx, settings)
}
