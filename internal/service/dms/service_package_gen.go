// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package dms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
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
			Factory:  dataSourceCertificate,
			TypeName: "aws_dms_certificate",
			Name:     "Certificate",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrCertificateARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  dataSourceEndpoint,
			TypeName: "aws_dms_endpoint",
			Name:     "Endpoint",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: "endpoint_arn",
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  dataSourceReplicationInstance,
			TypeName: "aws_dms_replication_instance",
			Name:     "Replication Instance",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: "replication_instance_arn",
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  dataSourceReplicationSubnetGroup,
			TypeName: "aws_dms_replication_subnet_group",
			Name:     "Replication Subnet Group",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: "replication_subnet_group_arn",
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  dataSourceReplicationTask,
			TypeName: "aws_dms_replication_task",
			Name:     "Replication Task",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: "replication_task_arn",
			},
			IsRegionOverrideEnabled: false,
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*itypes.ServicePackageSDKResource {
	return []*itypes.ServicePackageSDKResource{
		{
			Factory:  resourceCertificate,
			TypeName: "aws_dms_certificate",
			Name:     "Certificate",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrCertificateARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceEndpoint,
			TypeName: "aws_dms_endpoint",
			Name:     "Endpoint",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: "endpoint_arn",
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceEventSubscription,
			TypeName: "aws_dms_event_subscription",
			Name:     "Event Subscription",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceReplicationConfig,
			TypeName: "aws_dms_replication_config",
			Name:     "Replication Config",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceReplicationInstance,
			TypeName: "aws_dms_replication_instance",
			Name:     "Replication Instance",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: "replication_instance_arn",
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceReplicationSubnetGroup,
			TypeName: "aws_dms_replication_subnet_group",
			Name:     "Replication Subnet Group",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: "replication_subnet_group_arn",
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceReplicationTask,
			TypeName: "aws_dms_replication_task",
			Name:     "Replication Task",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: "replication_task_arn",
			},
			IsRegionOverrideEnabled: false,
		},
		{
			Factory:  resourceS3Endpoint,
			TypeName: "aws_dms_s3_endpoint",
			Name:     "S3 Endpoint",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: "endpoint_arn",
			},
			IsRegionOverrideEnabled: false,
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DMS
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*databasemigrationservice.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*databasemigrationservice.Options){
		databasemigrationservice.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *databasemigrationservice.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "databasemigrationservice",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return databasemigrationservice.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*databasemigrationservice.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*databasemigrationservice.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *databasemigrationservice.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*databasemigrationservice.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
