// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	itypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*itypes.ServicePackageFrameworkDataSource {
	return []*itypes.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourceOriginAccessControl,
			TypeName: "aws_cloudfront_origin_access_control",
			Name:     "Origin Access Control",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*itypes.ServicePackageFrameworkResource {
	return []*itypes.ServicePackageFrameworkResource{
		{
			Factory:  newContinuousDeploymentPolicyResource,
			TypeName: "aws_cloudfront_continuous_deployment_policy",
			Name:     "Continuous Deployment Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newKeyValueStoreResource,
			TypeName: "aws_cloudfront_key_value_store",
			Name:     "Key Value Store",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newVPCOriginResource,
			TypeName: "aws_cloudfront_vpc_origin",
			Name:     "VPC Origin",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*itypes.ServicePackageSDKDataSource {
	return []*itypes.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceCachePolicy,
			TypeName: "aws_cloudfront_cache_policy",
			Name:     "Cache Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourceDistribution,
			TypeName: "aws_cloudfront_distribution",
			Name:     "Distribution",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourceFunction,
			TypeName: "aws_cloudfront_function",
			Name:     "Function",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourceLogDeliveryCanonicalUserID,
			TypeName: "aws_cloudfront_log_delivery_canonical_user_id",
			Name:     "Log Delivery Canonical User ID",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourceOriginAccessIdentities,
			TypeName: "aws_cloudfront_origin_access_identities",
			Name:     "Origin Access Identities",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourceOriginAccessIdentity,
			TypeName: "aws_cloudfront_origin_access_identity",
			Name:     "Origin Access Identity",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourceOriginRequestPolicy,
			TypeName: "aws_cloudfront_origin_request_policy",
			Name:     "Origin Request Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourceRealtimeLogConfig,
			TypeName: "aws_cloudfront_realtime_log_config",
			Name:     "Real-time Log Config",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourceResponseHeadersPolicy,
			TypeName: "aws_cloudfront_response_headers_policy",
			Name:     "Response Headers Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*itypes.ServicePackageSDKResource {
	return []*itypes.ServicePackageSDKResource{
		{
			Factory:  resourceCachePolicy,
			TypeName: "aws_cloudfront_cache_policy",
			Name:     "Cache Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceDistribution,
			TypeName: "aws_cloudfront_distribution",
			Name:     "Distribution",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceFieldLevelEncryptionConfig,
			TypeName: "aws_cloudfront_field_level_encryption_config",
			Name:     "Field-level Encryption Config",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceFieldLevelEncryptionProfile,
			TypeName: "aws_cloudfront_field_level_encryption_profile",
			Name:     "Field-level Encryption Profile",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceFunction,
			TypeName: "aws_cloudfront_function",
			Name:     "Function",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceKeyGroup,
			TypeName: "aws_cloudfront_key_group",
			Name:     "Key Group",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceMonitoringSubscription,
			TypeName: "aws_cloudfront_monitoring_subscription",
			Name:     "Monitoring Subscription",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceOriginAccessControl,
			TypeName: "aws_cloudfront_origin_access_control",
			Name:     "Origin Access Control",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceOriginAccessIdentity,
			TypeName: "aws_cloudfront_origin_access_identity",
			Name:     "Origin Access Identity",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceOriginRequestPolicy,
			TypeName: "aws_cloudfront_origin_request_policy",
			Name:     "Origin Request Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourcePublicKey,
			TypeName: "aws_cloudfront_public_key",
			Name:     "Public Key",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceRealtimeLogConfig,
			TypeName: "aws_cloudfront_realtime_log_config",
			Name:     "Real-time Log Config",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceResponseHeadersPolicy,
			TypeName: "aws_cloudfront_response_headers_policy",
			Name:     "Response Headers Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CloudFront
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*cloudfront.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*cloudfront.Options){
		cloudfront.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *cloudfront.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "cloudfront",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return cloudfront.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*cloudfront.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*cloudfront.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *cloudfront.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*cloudfront.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
