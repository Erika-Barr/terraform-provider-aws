// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package ssoadmin

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	itypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*itypes.ServicePackageFrameworkDataSource {
	return []*itypes.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourceApplication,
			TypeName: "aws_ssoadmin_application",
			Name:     "Application",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newDataSourceApplicationAssignments,
			TypeName: "aws_ssoadmin_application_assignments",
			Name:     "Application Assignments",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newDataSourceApplicationProviders,
			TypeName: "aws_ssoadmin_application_providers",
			Name:     "Application Providers",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newPermissionSetsDataSource,
			TypeName: "aws_ssoadmin_permission_sets",
			Name:     "Permission Sets",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newDataSourcePrincipalApplicationAssignments,
			TypeName: "aws_ssoadmin_principal_application_assignments",
			Name:     "Principal Application Assignments",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*itypes.ServicePackageFrameworkResource {
	return []*itypes.ServicePackageFrameworkResource{
		{
			Factory:  newResourceApplication,
			TypeName: "aws_ssoadmin_application",
			Name:     "Application",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newResourceApplicationAccessScope,
			TypeName: "aws_ssoadmin_application_access_scope",
			Name:     "Application Access Scope",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newResourceApplicationAssignment,
			TypeName: "aws_ssoadmin_application_assignment",
			Name:     "Application Assignment",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newResourceApplicationAssignmentConfiguration,
			TypeName: "aws_ssoadmin_application_assignment_configuration",
			Name:     "Application Assignment Configuration",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newResourceTrustedTokenIssuer,
			TypeName: "aws_ssoadmin_trusted_token_issuer",
			Name:     "Trusted Token Issuer",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*itypes.ServicePackageSDKDataSource {
	return []*itypes.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceInstances,
			TypeName: "aws_ssoadmin_instances",
			Name:     "Instances",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  DataSourcePermissionSet,
			TypeName: "aws_ssoadmin_permission_set",
			Name:     "Permission Set",
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
			Factory:  ResourceAccountAssignment,
			TypeName: "aws_ssoadmin_account_assignment",
			Name:     "Account Assignment",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceCustomerManagedPolicyAttachment,
			TypeName: "aws_ssoadmin_customer_managed_policy_attachment",
			Name:     "Customer Managed Policy Attachment",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceAccessControlAttributes,
			TypeName: "aws_ssoadmin_instance_access_control_attributes",
			Name:     "Instance Access Control Attributes",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceManagedPolicyAttachment,
			TypeName: "aws_ssoadmin_managed_policy_attachment",
			Name:     "Managed Policy Attachment",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourcePermissionSet,
			TypeName: "aws_ssoadmin_permission_set",
			Name:     "Permission Set",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourcePermissionSetInlinePolicy,
			TypeName: "aws_ssoadmin_permission_set_inline_policy",
			Name:     "Permission Set Inline Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourcePermissionsBoundaryAttachment,
			TypeName: "aws_ssoadmin_permissions_boundary_attachment",
			Name:     "Permissions Boundary Attachment",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SSOAdmin
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*ssoadmin.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*ssoadmin.Options){
		ssoadmin.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *ssoadmin.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "ssoadmin",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return ssoadmin.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*ssoadmin.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*ssoadmin.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *ssoadmin.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*ssoadmin.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
