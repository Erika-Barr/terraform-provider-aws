// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
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
			Factory:  dataSourceConstraint,
			TypeName: "aws_servicecatalog_constraint",
			Name:     "Constraint",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourceLaunchPaths,
			TypeName: "aws_servicecatalog_launch_paths",
			Name:     "Launch Paths",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourcePortfolio,
			TypeName: "aws_servicecatalog_portfolio",
			Name:     "Portfolio",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourcePortfolioConstraints,
			TypeName: "aws_servicecatalog_portfolio_constraints",
			Name:     "Portfolio Constraints",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourceProduct,
			TypeName: "aws_servicecatalog_product",
			Name:     "Product",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourceProvisioningArtifacts,
			TypeName: "aws_servicecatalog_provisioning_artifacts",
			Name:     "Provisioning Artifacts",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*itypes.ServicePackageSDKResource {
	return []*itypes.ServicePackageSDKResource{
		{
			Factory:  resourceBudgetResourceAssociation,
			TypeName: "aws_servicecatalog_budget_resource_association",
			Name:     "Budget Resource Association",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceConstraint,
			TypeName: "aws_servicecatalog_constraint",
			Name:     "Constraint",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceOrganizationsAccess,
			TypeName: "aws_servicecatalog_organizations_access",
			Name:     "Organizations Access",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourcePortfolio,
			TypeName: "aws_servicecatalog_portfolio",
			Name:     "Portfolio",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourcePortfolioShare,
			TypeName: "aws_servicecatalog_portfolio_share",
			Name:     "Portfolio Share",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourcePrincipalPortfolioAssociation,
			TypeName: "aws_servicecatalog_principal_portfolio_association",
			Name:     "Principal Portfolio Association",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceProduct,
			TypeName: "aws_servicecatalog_product",
			Name:     "Product",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceProductPortfolioAssociation,
			TypeName: "aws_servicecatalog_product_portfolio_association",
			Name:     "Product Portfolio Association",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceProvisionedProduct,
			TypeName: "aws_servicecatalog_provisioned_product",
			Name:     "Provisioned Product",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceProvisioningArtifact,
			TypeName: "aws_servicecatalog_provisioning_artifact",
			Name:     "Provisioning Artifact",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceServiceAction,
			TypeName: "aws_servicecatalog_service_action",
			Name:     "Service Action",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceTagOption,
			TypeName: "aws_servicecatalog_tag_option",
			Name:     "Tag Option",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceTagOptionResourceAssociation,
			TypeName: "aws_servicecatalog_tag_option_resource_association",
			Name:     "Tag Option Resource Association",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ServiceCatalog
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*servicecatalog.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*servicecatalog.Options){
		servicecatalog.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *servicecatalog.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "servicecatalog",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return servicecatalog.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*servicecatalog.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*servicecatalog.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *servicecatalog.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*servicecatalog.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
