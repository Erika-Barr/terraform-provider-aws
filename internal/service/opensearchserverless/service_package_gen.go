// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package opensearchserverless

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	itypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*itypes.ServicePackageFrameworkDataSource {
	return []*itypes.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourceAccessPolicy,
			TypeName: "aws_opensearchserverless_access_policy",
			Name:     "Access Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newDataSourceCollection,
			TypeName: "aws_opensearchserverless_collection",
			Name:     "Collection",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newDataSourceLifecyclePolicy,
			TypeName: "aws_opensearchserverless_lifecycle_policy",
			Name:     "Lifecycle Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newDataSourceSecurityConfig,
			TypeName: "aws_opensearchserverless_security_config",
			Name:     "Security Config",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*itypes.ServicePackageFrameworkResource {
	return []*itypes.ServicePackageFrameworkResource{
		{
			Factory:  newResourceAccessPolicy,
			TypeName: "aws_opensearchserverless_access_policy",
			Name:     "Access Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newResourceCollection,
			TypeName: "aws_opensearchserverless_collection",
			Name:     "Collection",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newResourceLifecyclePolicy,
			TypeName: "aws_opensearchserverless_lifecycle_policy",
			Name:     "Lifecycle Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newResourceSecurityConfig,
			TypeName: "aws_opensearchserverless_security_config",
			Name:     "Security Config",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newResourceSecurityPolicy,
			TypeName: "aws_opensearchserverless_security_policy",
			Name:     "Security Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newVPCEndpointResource,
			TypeName: "aws_opensearchserverless_vpc_endpoint",
			Name:     "VPC Endpoint",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*itypes.ServicePackageSDKDataSource {
	return []*itypes.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceSecurityPolicy,
			TypeName: "aws_opensearchserverless_security_policy",
			Name:     "Security Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceVPCEndpoint,
			TypeName: "aws_opensearchserverless_vpc_endpoint",
			Name:     "VPC Endpoint",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*itypes.ServicePackageSDKResource {
	return []*itypes.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.OpenSearchServerless
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*opensearchserverless.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*opensearchserverless.Options){
		opensearchserverless.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *opensearchserverless.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "opensearchserverless",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return opensearchserverless.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*opensearchserverless.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*opensearchserverless.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *opensearchserverless.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*opensearchserverless.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
