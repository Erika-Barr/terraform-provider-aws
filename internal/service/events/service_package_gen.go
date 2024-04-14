// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package events

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	eventbridge_sdkv2 "github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceBus,
			TypeName: "aws_cloudwatch_event_bus",
		},
		{
			Factory:  DataSourceConnection,
			TypeName: "aws_cloudwatch_event_connection",
		},
		{
			Factory:  DataSourceSource,
			TypeName: "aws_cloudwatch_event_source",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAPIDestination,
			TypeName: "aws_cloudwatch_event_api_destination",
			Name:     "API Destination",
		},
		{
			Factory:  resourceArchive,
			TypeName: "aws_cloudwatch_event_archive",
			Name:     "Archive",
		},
		{
			Factory:  ResourceBus,
			TypeName: "aws_cloudwatch_event_bus",
			Name:     "Event Bus",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceBusPolicy,
			TypeName: "aws_cloudwatch_event_bus_policy",
		},
		{
			Factory:  ResourceConnection,
			TypeName: "aws_cloudwatch_event_connection",
		},
		{
			Factory:  ResourceEndpoint,
			TypeName: "aws_cloudwatch_event_endpoint",
			Name:     "Global Endpoint",
		},
		{
			Factory:  ResourcePermission,
			TypeName: "aws_cloudwatch_event_permission",
		},
		{
			Factory:  ResourceRule,
			TypeName: "aws_cloudwatch_event_rule",
			Name:     "Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceTarget,
			TypeName: "aws_cloudwatch_event_target",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Events
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*eventbridge_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return eventbridge_sdkv2.NewFromConfig(cfg, func(o *eventbridge_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
