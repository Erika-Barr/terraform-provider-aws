// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package guardduty

import (
	"context"
	"unique"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourceFindingIds,
			TypeName: "aws_guardduty_finding_ids",
			Name:     "Finding Ids",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newResourceMalwareProtectionPlan,
			TypeName: "aws_guardduty_malware_protection_plan",
			Name:     "Malware Protection Plan",
			Tags: unique.Make(types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
		},
		{
			Factory:  newMemberDetectorFeatureResource,
			TypeName: "aws_guardduty_member_detector_feature",
			Name:     "Member Detector Feature",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceDetector,
			TypeName: "aws_guardduty_detector",
			Name:     "Detector",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDetector,
			TypeName: "aws_guardduty_detector",
			Name:     "Detector",
			Tags: unique.Make(types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
		},
		{
			Factory:  ResourceDetectorFeature,
			TypeName: "aws_guardduty_detector_feature",
			Name:     "Detector Feature",
		},
		{
			Factory:  ResourceFilter,
			TypeName: "aws_guardduty_filter",
			Name:     "Filter",
			Tags: unique.Make(types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
		},
		{
			Factory:  resourceInviteAccepter,
			TypeName: "aws_guardduty_invite_accepter",
			Name:     "Invite Accepter",
		},
		{
			Factory:  ResourceIPSet,
			TypeName: "aws_guardduty_ipset",
			Name:     "IP Set",
			Tags: unique.Make(types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
		},
		{
			Factory:  ResourceMember,
			TypeName: "aws_guardduty_member",
			Name:     "Member",
		},
		{
			Factory:  ResourceOrganizationAdminAccount,
			TypeName: "aws_guardduty_organization_admin_account",
			Name:     "Organization Admin Account",
		},
		{
			Factory:  ResourceOrganizationConfiguration,
			TypeName: "aws_guardduty_organization_configuration",
			Name:     "Organization Configuration",
		},
		{
			Factory:  ResourceOrganizationConfigurationFeature,
			TypeName: "aws_guardduty_organization_configuration_feature",
			Name:     "Organization Configuration Feature",
		},
		{
			Factory:  ResourcePublishingDestination,
			TypeName: "aws_guardduty_publishing_destination",
			Name:     "Publishing Destination",
		},
		{
			Factory:  ResourceThreatIntelSet,
			TypeName: "aws_guardduty_threatintelset",
			Name:     "Threat Intel Set",
			Tags: unique.Make(types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.GuardDuty
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*guardduty.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*guardduty.Options){
		guardduty.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return guardduty.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*guardduty.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*guardduty.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *guardduty.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*guardduty.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
