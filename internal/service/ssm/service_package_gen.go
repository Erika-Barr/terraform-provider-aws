// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	itypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) EphemeralResources(ctx context.Context) []*itypes.ServicePackageEphemeralResource {
	return []*itypes.ServicePackageEphemeralResource{
		{
			Factory:  newEphemeralParameter,
			TypeName: "aws_ssm_parameter",
			Name:     "Parameter",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*itypes.ServicePackageFrameworkDataSource {
	return []*itypes.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourcePatchBaselines,
			TypeName: "aws_ssm_patch_baselines",
			Name:     "Patch Baselines",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:          false,
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*itypes.ServicePackageFrameworkResource {
	return []*itypes.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*itypes.ServicePackageSDKDataSource {
	return []*itypes.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceDocument,
			TypeName: "aws_ssm_document",
			Name:     "Document",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceInstances,
			TypeName: "aws_ssm_instances",
			Name:     "Instances",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceMaintenanceWindows,
			TypeName: "aws_ssm_maintenance_windows",
			Name:     "Maintenance Windows",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceParameter,
			TypeName: "aws_ssm_parameter",
			Name:     "Parameter",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceParametersByPath,
			TypeName: "aws_ssm_parameters_by_path",
			Name:     "Parameters By Path",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourcePatchBaseline,
			TypeName: "aws_ssm_patch_baseline",
			Name:     "Patch Baseline",
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
			Factory:  resourceActivation,
			TypeName: "aws_ssm_activation",
			Name:     "Activation",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceAssociation,
			TypeName: "aws_ssm_association",
			Name:     "Association",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Association",
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceDefaultPatchBaseline,
			TypeName: "aws_ssm_default_patch_baseline",
			Name:     "Default Patch Baseline",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceDocument,
			TypeName: "aws_ssm_document",
			Name:     "Document",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Document",
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceMaintenanceWindow,
			TypeName: "aws_ssm_maintenance_window",
			Name:     "Maintenance Window",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "MaintenanceWindow",
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceMaintenanceWindowTarget,
			TypeName: "aws_ssm_maintenance_window_target",
			Name:     "Maintenance Window Target",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceMaintenanceWindowTask,
			TypeName: "aws_ssm_maintenance_window_task",
			Name:     "Maintenance Window Task",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceParameter,
			TypeName: "aws_ssm_parameter",
			Name:     "Parameter",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Parameter",
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourcePatchBaseline,
			TypeName: "aws_ssm_patch_baseline",
			Name:     "Patch Baseline",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "PatchBaseline",
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourcePatchGroup,
			TypeName: "aws_ssm_patch_group",
			Name:     "Patch Group",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceResourceDataSync,
			TypeName: "aws_ssm_resource_data_sync",
			Name:     "Resource Data Sync",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceServiceSetting,
			TypeName: "aws_ssm_service_setting",
			Name:     "Service Setting",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SSM
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*ssm.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*ssm.Options){
		ssm.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *ssm.Options) {
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

	return ssm.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*ssm.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*ssm.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *ssm.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*ssm.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
