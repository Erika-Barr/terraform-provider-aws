// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package scheduler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
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
	return []*itypes.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*itypes.ServicePackageSDKResource {
	return []*itypes.ServicePackageSDKResource{
		{
			Factory:  resourceSchedule,
			TypeName: "aws_scheduler_schedule",
			Name:     "Schedule",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceScheduleGroup,
			TypeName: "aws_scheduler_schedule_group",
			Name:     "Schedule Group",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
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
	return names.Scheduler
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*scheduler.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*scheduler.Options){
		scheduler.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *scheduler.Options) {
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

	return scheduler.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*scheduler.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*scheduler.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *scheduler.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*scheduler.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
