// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
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
			Factory:                 newRecordsDataSource,
			TypeName:                "aws_route53_records",
			Name:                    "Records",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 newZonesDataSource,
			TypeName:                "aws_route53_zones",
			Name:                    "Zones",
			IsRegionOverrideEnabled: false,
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*itypes.ServicePackageFrameworkResource {
	return []*itypes.ServicePackageFrameworkResource{
		{
			Factory:                 newCIDRCollectionResource,
			TypeName:                "aws_route53_cidr_collection",
			Name:                    "CIDR Collection",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 newCIDRLocationResource,
			TypeName:                "aws_route53_cidr_location",
			Name:                    "CIDR Location",
			IsRegionOverrideEnabled: false,
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*itypes.ServicePackageSDKDataSource {
	return []*itypes.ServicePackageSDKDataSource{
		{
			Factory:                 dataSourceDelegationSet,
			TypeName:                "aws_route53_delegation_set",
			Name:                    "Reusable Delegation Set",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 dataSourceTrafficPolicyDocument,
			TypeName:                "aws_route53_traffic_policy_document",
			Name:                    "Traffic Policy Document",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 dataSourceZone,
			TypeName:                "aws_route53_zone",
			Name:                    "Hosted Zone",
			IsRegionOverrideEnabled: false,
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*itypes.ServicePackageSDKResource {
	return []*itypes.ServicePackageSDKResource{
		{
			Factory:                 resourceDelegationSet,
			TypeName:                "aws_route53_delegation_set",
			Name:                    "Reusable Delegation Set",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceHealthCheck,
			TypeName: "aws_route53_health_check",
			Name:     "Health Check",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "healthcheck",
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 resourceHostedZoneDNSSEC,
			TypeName:                "aws_route53_hosted_zone_dnssec",
			Name:                    "Hosted Zone DNSSEC",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 resourceKeySigningKey,
			TypeName:                "aws_route53_key_signing_key",
			Name:                    "Key Signing Key",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 resourceQueryLog,
			TypeName:                "aws_route53_query_log",
			Name:                    "Query Logging Config",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 resourceRecord,
			TypeName:                "aws_route53_record",
			Name:                    "Record",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 resourceTrafficPolicy,
			TypeName:                "aws_route53_traffic_policy",
			Name:                    "Traffic Policy",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 resourceTrafficPolicyInstance,
			TypeName:                "aws_route53_traffic_policy_instance",
			Name:                    "Traffic Policy Instance",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 resourceVPCAssociationAuthorization,
			TypeName:                "aws_route53_vpc_association_authorization",
			Name:                    "VPC Association Authorization",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceZone,
			TypeName: "aws_route53_zone",
			Name:     "Hosted Zone",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: "zone_id",
				ResourceType:        "hostedzone",
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 resourceZoneAssociation,
			TypeName:                "aws_route53_zone_association",
			Name:                    "Zone Association",
			IsRegionOverrideEnabled: false,
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Route53
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*route53.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*route53.Options){
		route53.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *route53.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "route53",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		func(o *route53.Options) {
			switch partition := config["partition"].(string); partition {
			case endpoints.AwsPartitionID:
				if region := endpoints.UsEast1RegionID; o.Region != region {
					tflog.Info(ctx, "overriding effective AWS API region", map[string]any{
						"service":         "route53",
						"original_region": o.Region,
						"override_region": region,
					})
					o.Region = region
				}
			case endpoints.AwsCnPartitionID:
				if region := endpoints.CnNorthwest1RegionID; o.Region != region {
					tflog.Info(ctx, "overriding effective AWS API region", map[string]any{
						"service":         "route53",
						"original_region": o.Region,
						"override_region": region,
					})
					o.Region = region
				}
			case endpoints.AwsUsGovPartitionID:
				if region := endpoints.UsGovWest1RegionID; o.Region != region {
					tflog.Info(ctx, "overriding effective AWS API region", map[string]any{
						"service":         "route53",
						"original_region": o.Region,
						"override_region": region,
					})
					o.Region = region
				}
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return route53.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*route53.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*route53.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *route53.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*route53.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
