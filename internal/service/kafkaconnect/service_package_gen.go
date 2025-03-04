// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package kafkaconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafkaconnect"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	itypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*itypes.ServicePackageFrameworkDataSource {
	return []*itypes.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*itypes.ServicePackageFrameworkResource {
	return []*itypes.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*itypes.ServicePackageSDKDataSource {
	return []*itypes.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceConnector,
			TypeName: "aws_mskconnect_connector",
			Name:     "Connector",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  dataSourceCustomPlugin,
			TypeName: "aws_mskconnect_custom_plugin",
			Name:     "Custom Plugin",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  dataSourceWorkerConfiguration,
			TypeName: "aws_mskconnect_worker_configuration",
			Name:     "Worker Configuration",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*itypes.ServicePackageSDKResource {
	return []*itypes.ServicePackageSDKResource{
		{
			Factory:  resourceConnector,
			TypeName: "aws_mskconnect_connector",
			Name:     "Connector",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceCustomPlugin,
			TypeName: "aws_mskconnect_custom_plugin",
			Name:     "Custom Plugin",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceWorkerConfiguration,
			TypeName: "aws_mskconnect_worker_configuration",
			Name:     "Worker Configuration",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.KafkaConnect
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*kafkaconnect.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*kafkaconnect.Options){
		kafkaconnect.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *kafkaconnect.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "kafkaconnect",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return kafkaconnect.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*kafkaconnect.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*kafkaconnect.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *kafkaconnect.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*kafkaconnect.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
