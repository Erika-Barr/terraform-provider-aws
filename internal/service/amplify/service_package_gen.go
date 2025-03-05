// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
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
			Factory:  resourceApp,
			TypeName: "aws_amplify_app",
			Name:     "App",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceBackendEnvironment,
			TypeName: "aws_amplify_backend_environment",
			Name:     "Backend Environment",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceBranch,
			TypeName: "aws_amplify_branch",
			Name:     "Branch",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceDomainAssociation,
			TypeName: "aws_amplify_domain_association",
			Name:     "Domain Association",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceWebhook,
			TypeName: "aws_amplify_webhook",
			Name:     "Webhook",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Amplify
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*amplify.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*amplify.Options){
		amplify.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *amplify.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "amplify",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return amplify.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*amplify.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*amplify.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *amplify.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*amplify.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
