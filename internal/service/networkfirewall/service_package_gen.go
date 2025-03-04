// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package networkfirewall

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
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
	return []*itypes.ServicePackageFrameworkResource{
		{
			Factory:  newTLSInspectionConfigurationResource,
			TypeName: "aws_networkfirewall_tls_inspection_configuration",
			Name:     "TLS Inspection Configuration",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*itypes.ServicePackageSDKDataSource {
	return []*itypes.ServicePackageSDKDataSource{
		{
			Factory:                 dataSourceFirewall,
			TypeName:                "aws_networkfirewall_firewall",
			Name:                    "Firewall",
			Tags:                    &itypes.ServicePackageResourceTags{},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 dataSourceFirewallPolicy,
			TypeName:                "aws_networkfirewall_firewall_policy",
			Name:                    "Firewall Policy",
			Tags:                    &itypes.ServicePackageResourceTags{},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 dataSourceResourcePolicy,
			TypeName:                "aws_networkfirewall_resource_policy",
			Name:                    "Resource Policy",
			IsRegionOverrideEnabled: false,
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*itypes.ServicePackageSDKResource {
	return []*itypes.ServicePackageSDKResource{
		{
			Factory:  resourceFirewall,
			TypeName: "aws_networkfirewall_firewall",
			Name:     "Firewall",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceFirewallPolicy,
			TypeName: "aws_networkfirewall_firewall_policy",
			Name:     "Firewall Policy",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 resourceLoggingConfiguration,
			TypeName:                "aws_networkfirewall_logging_configuration",
			Name:                    "Logging Configuration",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 resourceResourcePolicy,
			TypeName:                "aws_networkfirewall_resource_policy",
			Name:                    "Resource Policy",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceRuleGroup,
			TypeName: "aws_networkfirewall_rule_group",
			Name:     "Rule Group",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
			IsRegionOverrideEnabled: false,
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.NetworkFirewall
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*networkfirewall.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*networkfirewall.Options){
		networkfirewall.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *networkfirewall.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "networkfirewall",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return networkfirewall.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*networkfirewall.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*networkfirewall.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *networkfirewall.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*networkfirewall.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
