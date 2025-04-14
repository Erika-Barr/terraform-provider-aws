// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	inttypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*inttypes.ServicePackageFrameworkDataSource {
	return []*inttypes.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourceContributorManagedInsightRules,
			TypeName: "aws_cloudwatch_contributor_managed_insight_rules",
			Name:     "Contributor Managed Insight Rules",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*inttypes.ServicePackageFrameworkResource {
	return []*inttypes.ServicePackageFrameworkResource{
		{
			Factory:  newResourceContributorInsightRule,
			TypeName: "aws_cloudwatch_contributor_insight_rule",
			Name:     "Contributor Insight Rule",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrResourceARN,
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newResourceContributorManagedInsightRule,
			TypeName: "aws_cloudwatch_contributor_managed_insight_rule",
			Name:     "Contributor Managed Insight Rule",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*inttypes.ServicePackageSDKDataSource {
	return []*inttypes.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*inttypes.ServicePackageSDKResource {
	return []*inttypes.ServicePackageSDKResource{
		{
			Factory:  resourceCompositeAlarm,
			TypeName: "aws_cloudwatch_composite_alarm",
			Name:     "Composite Alarm",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceDashboard,
			TypeName: "aws_cloudwatch_dashboard",
			Name:     "Dashboard",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceMetricAlarm,
			TypeName: "aws_cloudwatch_metric_alarm",
			Name:     "Metric Alarm",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceMetricStream,
			TypeName: "aws_cloudwatch_metric_stream",
			Name:     "Metric Stream",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CloudWatch
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*cloudwatch.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*cloudwatch.Options){
		cloudwatch.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *cloudwatch.Options) {
			if region := config[names.AttrRegion].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         p.ServicePackageName(),
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return cloudwatch.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*cloudwatch.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*cloudwatch.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *cloudwatch.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*cloudwatch.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
