// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package servicecatalog

import (
	"context"
	"unique"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceConstraint,
			TypeName: "aws_servicecatalog_constraint",
			Name:     "Constraint",
		},
		{
			Factory:  dataSourceLaunchPaths,
			TypeName: "aws_servicecatalog_launch_paths",
			Name:     "Launch Paths",
		},
		{
			Factory:  dataSourcePortfolio,
			TypeName: "aws_servicecatalog_portfolio",
			Name:     "Portfolio",
			Tags:     unique.Make(types.ServicePackageResourceTags{}),
		},
		{
			Factory:  dataSourcePortfolioConstraints,
			TypeName: "aws_servicecatalog_portfolio_constraints",
			Name:     "Portfolio Constraints",
		},
		{
			Factory:  dataSourceProduct,
			TypeName: "aws_servicecatalog_product",
			Name:     "Product",
			Tags:     unique.Make(types.ServicePackageResourceTags{}),
		},
		{
			Factory:  dataSourceProvisioningArtifacts,
			TypeName: "aws_servicecatalog_provisioning_artifacts",
			Name:     "Provisioning Artifacts",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceBudgetResourceAssociation,
			TypeName: "aws_servicecatalog_budget_resource_association",
			Name:     "Budget Resource Association",
		},
		{
			Factory:  resourceConstraint,
			TypeName: "aws_servicecatalog_constraint",
			Name:     "Constraint",
		},
		{
			Factory:  resourceOrganizationsAccess,
			TypeName: "aws_servicecatalog_organizations_access",
			Name:     "Organizations Access",
		},
		{
			Factory:  resourcePortfolio,
			TypeName: "aws_servicecatalog_portfolio",
			Name:     "Portfolio",
			Tags:     unique.Make(types.ServicePackageResourceTags{}),
		},
		{
			Factory:  resourcePortfolioShare,
			TypeName: "aws_servicecatalog_portfolio_share",
			Name:     "Portfolio Share",
		},
		{
			Factory:  resourcePrincipalPortfolioAssociation,
			TypeName: "aws_servicecatalog_principal_portfolio_association",
			Name:     "Principal Portfolio Association",
		},
		{
			Factory:  resourceProduct,
			TypeName: "aws_servicecatalog_product",
			Name:     "Product",
			Tags:     unique.Make(types.ServicePackageResourceTags{}),
		},
		{
			Factory:  resourceProductPortfolioAssociation,
			TypeName: "aws_servicecatalog_product_portfolio_association",
			Name:     "Product Portfolio Association",
		},
		{
			Factory:  resourceProvisionedProduct,
			TypeName: "aws_servicecatalog_provisioned_product",
			Name:     "Provisioned Product",
			Tags:     unique.Make(types.ServicePackageResourceTags{}),
		},
		{
			Factory:  resourceProvisioningArtifact,
			TypeName: "aws_servicecatalog_provisioning_artifact",
			Name:     "Provisioning Artifact",
		},
		{
			Factory:  resourceServiceAction,
			TypeName: "aws_servicecatalog_service_action",
			Name:     "Service Action",
		},
		{
			Factory:  resourceTagOption,
			TypeName: "aws_servicecatalog_tag_option",
			Name:     "Tag Option",
		},
		{
			Factory:  resourceTagOptionResourceAssociation,
			TypeName: "aws_servicecatalog_tag_option_resource_association",
			Name:     "Tag Option Resource Association",
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
