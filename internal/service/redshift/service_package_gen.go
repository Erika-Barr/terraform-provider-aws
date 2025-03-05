// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	itypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*itypes.ServicePackageFrameworkDataSource {
	return []*itypes.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourceDataShares,
			TypeName: "aws_redshift_data_shares",
			Name:     "Data Shares",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newDataSourceProducerDataShares,
			TypeName: "aws_redshift_producer_data_shares",
			Name:     "Producer Data Shares",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*itypes.ServicePackageFrameworkResource {
	return []*itypes.ServicePackageFrameworkResource{
		{
			Factory:  newResourceDataShareAuthorization,
			TypeName: "aws_redshift_data_share_authorization",
			Name:     "Data Share Authorization",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newResourceDataShareConsumerAssociation,
			TypeName: "aws_redshift_data_share_consumer_association",
			Name:     "Data Share Consumer Association",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newResourceLogging,
			TypeName: "aws_redshift_logging",
			Name:     "Logging",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  newResourceSnapshotCopy,
			TypeName: "aws_redshift_snapshot_copy",
			Name:     "Snapshot Copy",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*itypes.ServicePackageSDKDataSource {
	return []*itypes.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceCluster,
			TypeName: "aws_redshift_cluster",
			Name:     "Cluster",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourceClusterCredentials,
			TypeName: "aws_redshift_cluster_credentials",
			Name:     "Cluster Credentials",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourceOrderableCluster,
			TypeName: "aws_redshift_orderable_cluster",
			Name:     "Orderable Cluster",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourceServiceAccount,
			TypeName: "aws_redshift_service_account",
			Name:     "Service Account",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  dataSourceSubnetGroup,
			TypeName: "aws_redshift_subnet_group",
			Name:     "Subnet Group",
			Tags:     &itypes.ServicePackageResourceTags{},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*itypes.ServicePackageSDKResource {
	return []*itypes.ServicePackageSDKResource{
		{
			Factory:  resourceAuthenticationProfile,
			TypeName: "aws_redshift_authentication_profile",
			Name:     "Authentication Profile",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceCluster,
			TypeName: "aws_redshift_cluster",
			Name:     "Cluster",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceClusterIAMRoles,
			TypeName: "aws_redshift_cluster_iam_roles",
			Name:     "Cluster IAM Roles",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceClusterSnapshot,
			TypeName: "aws_redshift_cluster_snapshot",
			Name:     "Cluster Snapshot",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceEndpointAccess,
			TypeName: "aws_redshift_endpoint_access",
			Name:     "Endpoint Access",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceEndpointAuthorization,
			TypeName: "aws_redshift_endpoint_authorization",
			Name:     "Endpoint Authorization",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceEventSubscription,
			TypeName: "aws_redshift_event_subscription",
			Name:     "Event Subscription",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceHSMClientCertificate,
			TypeName: "aws_redshift_hsm_client_certificate",
			Name:     "HSM Client Certificate",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceHSMConfiguration,
			TypeName: "aws_redshift_hsm_configuration",
			Name:     "HSM Configuration",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceParameterGroup,
			TypeName: "aws_redshift_parameter_group",
			Name:     "Parameter Group",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourcePartner,
			TypeName: "aws_redshift_partner",
			Name:     "Partner",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceResourcePolicy,
			TypeName: "aws_redshift_resource_policy",
			Name:     "Resource Policy",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceScheduledAction,
			TypeName: "aws_redshift_scheduled_action",
			Name:     "Scheduled Action",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceSnapshotCopyGrant,
			TypeName: "aws_redshift_snapshot_copy_grant",
			Name:     "Snapshot Copy Grant",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceSnapshotSchedule,
			TypeName: "aws_redshift_snapshot_schedule",
			Name:     "Snapshot Schedule",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceSnapshotScheduleAssociation,
			TypeName: "aws_redshift_snapshot_schedule_association",
			Name:     "Snapshot Schedule Association",
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceSubnetGroup,
			TypeName: "aws_redshift_subnet_group",
			Name:     "Subnet Group",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
		{
			Factory:  resourceUsageLimit,
			TypeName: "aws_redshift_usage_limit",
			Name:     "Usage Limit",
			Tags: &itypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &itypes.ServicePackageResourceRegion{
				IsOverrideEnabled: false,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Redshift
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*redshift.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*redshift.Options){
		redshift.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *redshift.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "redshift",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return redshift.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*redshift.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*redshift.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *redshift.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*redshift.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
