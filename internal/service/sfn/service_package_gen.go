// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package sfn

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
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
			Factory:  dataSourceActivity,
			TypeName: "aws_sfn_activity",
			Name:     "Activity",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceAlias,
			TypeName: "aws_sfn_alias",
			Name:     "Alias",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceStateMachine,
			TypeName: "aws_sfn_state_machine",
			Name:     "State Machine",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceStateMachineVersions,
			TypeName: "aws_sfn_state_machine_versions",
			Name:     "State Machine Versions",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*itypes.ServicePackageSDKResource {
	return []*itypes.ServicePackageSDKResource{
		{
			Factory:  resourceActivity,
			TypeName: "aws_sfn_activity",
			Name:     "Activity",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceAlias,
			TypeName: "aws_sfn_alias",
			Name:     "Alias",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceStateMachine,
			TypeName: "aws_sfn_state_machine",
			Name:     "State Machine",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SFN
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*sfn.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*sfn.Options){
		sfn.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *sfn.Options) {
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

	return sfn.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*sfn.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*sfn.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *sfn.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*sfn.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
