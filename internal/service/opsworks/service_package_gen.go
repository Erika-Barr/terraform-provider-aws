// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opsworks"
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
			Factory:                 resourceApplication,
			TypeName:                "aws_opsworks_application",
			Name:                    "Application",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceCustomLayer,
			TypeName: "aws_opsworks_custom_layer",
			Name:     "Custom Layer",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceECSClusterLayer,
			TypeName: "aws_opsworks_ecs_cluster_layer",
			Name:     "ECS Cluster Layer",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceGangliaLayer,
			TypeName: "aws_opsworks_ganglia_layer",
			Name:     "Ganglia Layer",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceHAProxyLayer,
			TypeName: "aws_opsworks_haproxy_layer",
			Name:     "HAProxy Layer",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 resourceInstance,
			TypeName:                "aws_opsworks_instance",
			Name:                    "Instance",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceJavaAppLayer,
			TypeName: "aws_opsworks_java_app_layer",
			Name:     "Java App Layer",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceMemcachedLayer,
			TypeName: "aws_opsworks_memcached_layer",
			Name:     "Memcached Layer",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceMySQLLayer,
			TypeName: "aws_opsworks_mysql_layer",
			Name:     "MySQL Layer",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceNodejsAppLayer,
			TypeName: "aws_opsworks_nodejs_app_layer",
			Name:     "NodeJS App Layer",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 resourcePermission,
			TypeName:                "aws_opsworks_permission",
			Name:                    "Permission",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourcePHPAppLayer,
			TypeName: "aws_opsworks_php_app_layer",
			Name:     "PHP App Layer",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceRailsAppLayer,
			TypeName: "aws_opsworks_rails_app_layer",
			Name:     "Rails App Layer",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 resourceRDSDBInstance,
			TypeName:                "aws_opsworks_rds_db_instance",
			Name:                    "RDS DB Instance",
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 resourceStack,
			TypeName:                "aws_opsworks_stack",
			Name:                    "Stack",
			Tags:                    &itypes.ServicePackageResourceTags{},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceStaticWebLayer,
			TypeName: "aws_opsworks_static_web_layer",
			Name:     "Static Web Layer",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:                 resourceUserProfile,
			TypeName:                "aws_opsworks_user_profile",
			Name:                    "Profile",
			IsRegionOverrideEnabled: false,
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.OpsWorks
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*opsworks.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*opsworks.Options){
		opsworks.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *opsworks.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "opsworks",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return opsworks.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*opsworks.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*opsworks.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *opsworks.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*opsworks.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
