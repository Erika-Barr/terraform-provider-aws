// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package account

import (
	"context"
	"unique"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	inttypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*inttypes.ServicePackageFrameworkDataSource {
	return []*inttypes.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*inttypes.ServicePackageFrameworkResource {
	return []*inttypes.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*inttypes.ServicePackageSDKDataSource {
	return []*inttypes.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*inttypes.ServicePackageSDKResource {
	return []*inttypes.ServicePackageSDKResource{
		{
			Factory:  resourceAlternateContact,
			TypeName: "aws_account_alternate_contact",
			Name:     "Alternate Contact",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  resourcePrimaryContact,
			TypeName: "aws_account_primary_contact",
			Name:     "Primary Contact",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
		{
			Factory:  resourceRegion,
			TypeName: "aws_account_region",
			Name:     "Region",
			Region: unique.Make(inttypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			}),
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Account
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*account.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*account.Options){
		account.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *account.Options) {
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

	return account.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*account.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*account.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *account.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*account.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
