// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package {{ .ProviderPackage }}

import (
	"context"

{{ if .GenerateClient }}
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/{{ .GoV2Package }}"
	"github.com/hashicorp/terraform-plugin-log/tflog"
{{- end }}
{{- if gt (len .EndpointRegionOverrides) 0 }}
	"github.com/hashicorp/aws-sdk-go-base/v2/endpoints"
{{- end }}
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	itypes "github.com/hashicorp/terraform-provider-aws/internal/types"
{{- if ne .ProviderPackage "meta" }}
	"github.com/hashicorp/terraform-provider-aws/names"
{{- end }}
)

type servicePackage struct {}

{{- if .EphemeralResources }}
func (p *servicePackage) EphemeralResources(ctx context.Context) []*itypes.ServicePackageEphemeralResource {
	return []*itypes.ServicePackageEphemeralResource {
{{- range $key, $value := .EphemeralResources }}
		{
			Factory:                 {{ $value.FactoryName }},
			TypeName:                "{{ $key }}",
			Name:                    "{{ $value.Name }}",
			Region: &itypes.ServicePackageResourceRegion {
				IsGlobal:          {{ or $.IsGlobal $value.IsGlobal }},
				IsOverrideEnabled: {{ $value.RegionOverrideEnabled }},
	{{- if $value.RegionOverrideEnabled }}
				IsValidateOverrideInPartition: {{ $value.ValidateRegionOverrideInPartition }},
	{{- end }}
			},
		},
{{- end }}
	}
}
{{- end }}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*itypes.ServicePackageFrameworkDataSource {
	return []*itypes.ServicePackageFrameworkDataSource {
{{- range $key, $value := .FrameworkDataSources }}
		{
			Factory:  {{ $value.FactoryName }},
			TypeName: "{{ $key }}",
			Name:     "{{ $value.Name }}",
			{{- if .TransparentTagging }}
			Tags: &itypes.ServicePackageResourceTags {
				{{- if ne .TagsIdentifierAttribute "" }}
				IdentifierAttribute: {{ .TagsIdentifierAttribute }},
				{{- end }}
				{{- if ne .TagsResourceType "" }}
				ResourceType: "{{ .TagsResourceType }}",
				{{- end }}
			},
			{{- end }}
			Region: &itypes.ServicePackageResourceRegion {
				IsGlobal:          {{ or $.IsGlobal $value.IsGlobal }},
				IsOverrideEnabled: {{ $value.RegionOverrideEnabled }},
	{{- if $value.RegionOverrideEnabled }}
				IsValidateOverrideInPartition: {{ $value.ValidateRegionOverrideInPartition }},
	{{- end }}
			},
		},
{{- end }}
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*itypes.ServicePackageFrameworkResource {
	return []*itypes.ServicePackageFrameworkResource {
{{- range $key, $value := .FrameworkResources }}
		{
			Factory:  {{ $value.FactoryName }},
			TypeName: "{{ $key }}",
			Name:     "{{ $value.Name }}",
			{{- if .TransparentTagging }}
			Tags: &itypes.ServicePackageResourceTags {
				{{- if ne .TagsIdentifierAttribute "" }}
				IdentifierAttribute: {{ .TagsIdentifierAttribute }},
				{{- end }}
				{{- if ne .TagsResourceType "" }}
				ResourceType: "{{ .TagsResourceType }}",
				{{- end }}
			},
			{{- end }}
			Region: &itypes.ServicePackageResourceRegion {
				IsGlobal:          {{ or $.IsGlobal $value.IsGlobal }},
				IsOverrideEnabled: {{ $value.RegionOverrideEnabled }},
	{{- if $value.RegionOverrideEnabled }}
				IsValidateOverrideInPartition: {{ $value.ValidateRegionOverrideInPartition }},
	{{- end }}
			},
		},
{{- end }}
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*itypes.ServicePackageSDKDataSource {
	return []*itypes.ServicePackageSDKDataSource {
{{- range $key, $value := .SDKDataSources }}
		{
			Factory:  {{ $value.FactoryName }},
			TypeName: "{{ $key }}",
			Name:     "{{ $value.Name }}",
			{{- if $value.TransparentTagging }}
			Tags: &itypes.ServicePackageResourceTags {
				{{- if ne $value.TagsIdentifierAttribute "" }}
				IdentifierAttribute: {{ $value.TagsIdentifierAttribute }},
				{{- end }}
				{{- if ne .TagsResourceType "" }}
				ResourceType: "{{ .TagsResourceType }}",
				{{- end }}
			},
			{{- end }}
			Region: &itypes.ServicePackageResourceRegion {
				IsGlobal:          {{ or $.IsGlobal $value.IsGlobal }},
				IsOverrideEnabled: {{ $value.RegionOverrideEnabled }},
	{{- if $value.RegionOverrideEnabled }}
				IsValidateOverrideInPartition: {{ $value.ValidateRegionOverrideInPartition }},
	{{- end }}
			},
		},
{{- end }}
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*itypes.ServicePackageSDKResource {
	return []*itypes.ServicePackageSDKResource {
{{- range $key, $value := .SDKResources }}
		{
			Factory:  {{ $value.FactoryName }},
			TypeName: "{{ $key }}",
			Name:     "{{ $value.Name }}",
			{{- if $value.TransparentTagging }}
			Tags: &itypes.ServicePackageResourceTags {
				{{- if ne $value.TagsIdentifierAttribute "" }}
				IdentifierAttribute: {{ $value.TagsIdentifierAttribute }},
				{{- end }}
				{{- if ne .TagsResourceType "" }}
				ResourceType: "{{ .TagsResourceType }}",
				{{- end }}
			},
			{{- end }}
			Region: &itypes.ServicePackageResourceRegion {
				IsGlobal:          {{ or $.IsGlobal $value.IsGlobal }},
				IsOverrideEnabled: {{ $value.RegionOverrideEnabled }},
	{{- if $value.RegionOverrideEnabled }}
				IsValidateOverrideInPartition: {{ $value.ValidateRegionOverrideInPartition }},
	{{- end }}
			},
		},
{{- end }}
	}
}

func (p *servicePackage) ServicePackageName() string {
{{- if eq .ProviderPackage "meta" }}
	return "{{ .ProviderPackage }}"
{{- else }}
	return names.{{ .ProviderNameUpper }}
{{- end }}
}

{{- if .GenerateClient }}
// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*{{ .GoV2Package }}.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*{{ .GoV2Package }}.Options){
		{{ .GoV2Package }}.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *{{ .GoV2Package }}.Options) {
			if region := config[names.AttrRegion].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         p.ServicePackageName(),
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
{{- if gt (len .EndpointRegionOverrides) 0 }}
		func(o *{{ .GoV2Package }}.Options) {
			switch partition := config["partition"].(string); partition {
	{{- range $k, $v := .EndpointRegionOverrides }}
			case endpoints.{{ $k | Camel }}PartitionID:
				if region := endpoints.{{ $v | Camel }}RegionID; o.Region != region {
					tflog.Info(ctx, "overriding effective AWS API region", map[string]any{
					    "service":         p.ServicePackageName(),
						"original_region": o.Region,
						"override_region": region,
					})
					o.Region = region
				}
	{{- end }}
			}
		},
{{- end }}
		withExtraOptions(ctx, p, config),
	}

	return {{ .GoV2Package }}.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*{{ .GoV2Package }}.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*{{ .GoV2Package }}.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *{{ .GoV2Package }}.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func (*{{ .GoV2Package }}.Options) {}
}
{{- end }}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
