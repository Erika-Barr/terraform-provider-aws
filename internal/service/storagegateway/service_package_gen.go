// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
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
	return []*inttypes.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceLocalDisk,
			TypeName: "aws_storagegateway_local_disk",
			Name:     "Local Disk",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*inttypes.ServicePackageSDKResource {
	return []*inttypes.ServicePackageSDKResource{
		{
			Factory:  resourceCache,
			TypeName: "aws_storagegateway_cache",
			Name:     "Cache",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceCachediSCSIVolume,
			TypeName: "aws_storagegateway_cached_iscsi_volume",
			Name:     "Cached iSCSI Volume",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceFileSystemAssociation,
			TypeName: "aws_storagegateway_file_system_association",
			Name:     "File System Association",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceGateway,
			TypeName: "aws_storagegateway_gateway",
			Name:     "Gateway",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceNFSFileShare,
			TypeName: "aws_storagegateway_nfs_file_share",
			Name:     "NFS File Share",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceSMBFileShare,
			TypeName: "aws_storagegateway_smb_file_share",
			Name:     "SMB File Share",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceStorediSCSIVolume,
			TypeName: "aws_storagegateway_stored_iscsi_volume",
			Name:     "Stored iSCSI Volume",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceTapePool,
			TypeName: "aws_storagegateway_tape_pool",
			Name:     "Tape Pool",
			Tags: &inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  resourceUploadBuffer,
			TypeName: "aws_storagegateway_upload_buffer",
			Name:     "Upload Buffer",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
		{
			Factory:  ResourceWorkingStorage,
			TypeName: "aws_storagegateway_working_storage",
			Name:     "Working Storage",
			Region: &inttypes.ServicePackageResourceRegion{
				IsGlobal:                      false,
				IsOverrideEnabled:             true,
				IsValidateOverrideInPartition: true,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.StorageGateway
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*storagegateway.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*storagegateway.Options){
		storagegateway.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *storagegateway.Options) {
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

	return storagegateway.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*storagegateway.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*storagegateway.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *storagegateway.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*storagegateway.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
