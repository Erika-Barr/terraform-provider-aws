// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/connect"
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
			Factory:  dataSourceBotAssociation,
			TypeName: "aws_connect_bot_association",
			Name:     "Bot Association",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceContactFlow,
			TypeName: "aws_connect_contact_flow",
			Name:     "Contact Flow",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceContactFlowModule,
			TypeName: "aws_connect_contact_flow_module",
			Name:     "Contact Flow Module",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceHoursOfOperation,
			TypeName: "aws_connect_hours_of_operation",
			Name:     "Hours Of Operation",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceInstance,
			TypeName: "aws_connect_instance",
			Name:     "Instance",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceInstanceStorageConfig,
			TypeName: "aws_connect_instance_storage_config",
			Name:     "Instance Storage Config",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceLambdaFunctionAssociation,
			TypeName: "aws_connect_lambda_function_association",
			Name:     "Lambda Function Association",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourcePrompt,
			TypeName: "aws_connect_prompt",
			Name:     "Prompt",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceQueue,
			TypeName: "aws_connect_queue",
			Name:     "Queue",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceQuickConnect,
			TypeName: "aws_connect_quick_connect",
			Name:     "Quick Connect",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceRoutingProfile,
			TypeName: "aws_connect_routing_profile",
			Name:     "Routing Profile",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceSecurityProfile,
			TypeName: "aws_connect_security_profile",
			Name:     "Security Profile",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  DataSourceUser,
			TypeName: "aws_connect_user",
			Name:     "User",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceUserHierarchyGroup,
			TypeName: "aws_connect_user_hierarchy_group",
			Name:     "User Hierarchy Group",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceUserHierarchyStructure,
			TypeName: "aws_connect_user_hierarchy_structure",
			Name:     "User Hierarchy Structure",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  dataSourceVocabulary,
			TypeName: "aws_connect_vocabulary",
			Name:     "Vocabulary",
			Tags:     &itypes.ServicePackageResourceTags{},
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
			Factory:  resourceBotAssociation,
			TypeName: "aws_connect_bot_association",
			Name:     "Bot Association",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceContactFlow,
			TypeName: "aws_connect_contact_flow",
			Name:     "Contact Flow",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceContactFlowModule,
			TypeName: "aws_connect_contact_flow_module",
			Name:     "Contact Flow Module",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceHoursOfOperation,
			TypeName: "aws_connect_hours_of_operation",
			Name:     "Hours Of Operation",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceInstance,
			TypeName: "aws_connect_instance",
			Name:     "Instance",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceInstanceStorageConfig,
			TypeName: "aws_connect_instance_storage_config",
			Name:     "Instance Storage Config",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceLambdaFunctionAssociation,
			TypeName: "aws_connect_lambda_function_association",
			Name:     "Lambda Function Association",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourcePhoneNumber,
			TypeName: "aws_connect_phone_number",
			Name:     "Phone Number",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceQueue,
			TypeName: "aws_connect_queue",
			Name:     "Queue",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceQuickConnect,
			TypeName: "aws_connect_quick_connect",
			Name:     "Quick Connect",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceRoutingProfile,
			TypeName: "aws_connect_routing_profile",
			Name:     "Routing Profile",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceSecurityProfile,
			TypeName: "aws_connect_security_profile",
			Name:     "Security Profile",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceUser,
			TypeName: "aws_connect_user",
			Name:     "User",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceUserHierarchyGroup,
			TypeName: "aws_connect_user_hierarchy_group",
			Name:     "User Hierarchy Group",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceUserHierarchyStructure,
			TypeName: "aws_connect_user_hierarchy_structure",
			Name:     "User Hierarchy Structure",
			Region: &itypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceVocabulary,
			TypeName: "aws_connect_vocabulary",
			Name:     "Vocabulary",
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
	return names.Connect
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*connect.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*connect.Options){
		connect.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *connect.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "connect",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return connect.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*connect.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*connect.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *connect.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*connect.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
