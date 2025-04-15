// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package securityhub

import (
	"context"
	"unique"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	inttypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*inttypes.ServicePackageFrameworkDataSource {
	return []*inttypes.ServicePackageFrameworkDataSource{
		{
			Factory:  newStandardsControlAssociationsDataSource,
			TypeName: "aws_securityhub_standards_control_associations",
			Name:     "Standards Control Associations",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*inttypes.ServicePackageFrameworkResource {
	return []*inttypes.ServicePackageFrameworkResource{
		{
			Factory:  newAutomationRuleResource,
			TypeName: "aws_securityhub_automation_rule",
			Name:     "Automation Rule",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  newStandardsControlAssociationResource,
			TypeName: "aws_securityhub_standards_control_association",
			Name:     "Standards Control Association",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*inttypes.ServicePackageSDKDataSource {
	return []*inttypes.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*inttypes.ServicePackageSDKResource {
	return []*inttypes.ServicePackageSDKResource{
		{
			Factory:  resourceAccount,
			TypeName: "aws_securityhub_account",
			Name:     "Account",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			}),
		},
		{
			Factory:  resourceActionTarget,
			TypeName: "aws_securityhub_action_target",
			Name:     "Action Target",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			}),
		},
		{
			Factory:  resourceConfigurationPolicy,
			TypeName: "aws_securityhub_configuration_policy",
			Name:     "Configuration Policy",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			}),
		},
		{
			Factory:  resourceConfigurationPolicyAssociation,
			TypeName: "aws_securityhub_configuration_policy_association",
			Name:     "Configuration Policy Association",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			}),
		},
		{
			Factory:  resourceFindingAggregator,
			TypeName: "aws_securityhub_finding_aggregator",
			Name:     "Finding Aggregator",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			}),
		},
		{
			Factory:  resourceInsight,
			TypeName: "aws_securityhub_insight",
			Name:     "Insight",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			}),
		},
		{
			Factory:  resourceInviteAccepter,
			TypeName: "aws_securityhub_invite_accepter",
			Name:     "Invite Accepter",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			}),
		},
		{
			Factory:  resourceMember,
			TypeName: "aws_securityhub_member",
			Name:     "Member",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			}),
		},
		{
			Factory:  resourceOrganizationAdminAccount,
			TypeName: "aws_securityhub_organization_admin_account",
			Name:     "Organization Admin Account",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			}),
		},
		{
			Factory:  resourceOrganizationConfiguration,
			TypeName: "aws_securityhub_organization_configuration",
			Name:     "Organization Configuration",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			}),
		},
		{
			Factory:  resourceProductSubscription,
			TypeName: "aws_securityhub_product_subscription",
			Name:     "Product Subscription",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			}),
		},
		{
			Factory:  resourceStandardsControl,
			TypeName: "aws_securityhub_standards_control",
			Name:     "Standards Control",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			}),
		},
		{
			Factory:  resourceStandardsSubscription,
			TypeName: "aws_securityhub_standards_subscription",
			Name:     "Standards Subscription",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			}),
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SecurityHub
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*securityhub.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*securityhub.Options){
		securityhub.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *securityhub.Options) {
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

	return securityhub.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*securityhub.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*securityhub.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *securityhub.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*securityhub.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
