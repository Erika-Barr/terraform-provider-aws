// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
	"github.com/hashicorp/aws-sdk-go-base/v2/endpoints"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	itypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*itypes.ServicePackageFrameworkDataSource {
	return []*itypes.ServicePackageFrameworkDataSource{
		{
			Factory:  newAcceleratorDataSource,
			TypeName: "aws_globalaccelerator_accelerator",
			Name:     "Accelerator",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          true,
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*itypes.ServicePackageFrameworkResource {
	return []*itypes.ServicePackageFrameworkResource{
		{
			Factory:  newCrossAccountAttachmentResource,
			TypeName: "aws_globalaccelerator_cross_account_attachment",
			Name:     "Cross-account Attachment",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          true,
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*itypes.ServicePackageSDKDataSource {
	return []*itypes.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceCustomRoutingAccelerator,
			TypeName: "aws_globalaccelerator_custom_routing_accelerator",
			Name:     "Custom Routing Accelerator",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          true,
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*itypes.ServicePackageSDKResource {
	return []*itypes.ServicePackageSDKResource{
		{
			Factory:  resourceAccelerator,
			TypeName: "aws_globalaccelerator_accelerator",
			Name:     "Accelerator",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          true,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceCustomRoutingAccelerator,
			TypeName: "aws_globalaccelerator_custom_routing_accelerator",
			Name:     "Custom Routing Accelerator",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          true,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceCustomRoutingEndpointGroup,
			TypeName: "aws_globalaccelerator_custom_routing_endpoint_group",
			Name:     "Custom Routing Endpoint Group",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          true,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceCustomRoutingListener,
			TypeName: "aws_globalaccelerator_custom_routing_listener",
			Name:     "Custom Routing Listener",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          true,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceEndpointGroup,
			TypeName: "aws_globalaccelerator_endpoint_group",
			Name:     "Endpoint Group",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          true,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceListener,
			TypeName: "aws_globalaccelerator_listener",
			Name:     "Listener",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          true,
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.GlobalAccelerator
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*globalaccelerator.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*globalaccelerator.Options){
		globalaccelerator.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *globalaccelerator.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "globalaccelerator",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		func(o *globalaccelerator.Options) {
			switch partition := config["partition"].(string); partition {
			case endpoints.AwsPartitionID:
				if region := endpoints.UsWest2RegionID; o.Region != region {
					tflog.Info(ctx, "overriding effective AWS API region", map[string]any{
						"service":         "globalaccelerator",
						"original_region": o.Region,
						"override_region": region,
					})
					o.Region = region
				}
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return globalaccelerator.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*globalaccelerator.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*globalaccelerator.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *globalaccelerator.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*globalaccelerator.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
