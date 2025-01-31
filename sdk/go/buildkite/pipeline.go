// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package buildkite

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: pipeline
//
// This resource allows you to create and manage pipelines for repositories.
//
// Buildkite Documentation: https://buildkite.com/docs/pipelines
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"io/ioutil"
//
// 	"github.com/grapl-security/pulumi-buildkite/sdk/go/buildkite"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := buildkite.NewPipeline(ctx, "repo2", &buildkite.PipelineArgs{
// 			Repository: pulumi.String("git@github.com:org/repo2"),
// 			Steps:      readFileOrPanic("./steps.yml"),
// 			Teams: PipelineTeamArray{
// 				&PipelineTeamArgs{
// 					Slug:        pulumi.String("everyone"),
// 					AccessLevel: pulumi.String("READ_ONLY"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### With GitHub Provider Settings
//
// ```go
// package main
//
// import (
// 	"io/ioutil"
//
// 	"github.com/grapl-security/pulumi-buildkite/sdk/go/buildkite"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := buildkite.NewPipeline(ctx, "repo2-deploy", &buildkite.PipelineArgs{
// 			Repository: pulumi.String("git@github.com:org/repo2"),
// 			Steps:      readFileOrPanic("./deploy-steps.yml"),
// 			ProviderSettings: &PipelineProviderSettingsArgs{
// 				TriggerMode: pulumi.String("none"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = buildkite.NewPipeline(ctx, "repo2-release", &buildkite.PipelineArgs{
// 			Repository: pulumi.String("git@github.com:org/repo2"),
// 			Steps:      readFileOrPanic("./release-steps.yml"),
// 			ProviderSettings: &PipelineProviderSettingsArgs{
// 				BuildBranches:     pulumi.Bool(false),
// 				BuildTags:         pulumi.Bool(true),
// 				BuildPullRequests: pulumi.Bool(false),
// 				TriggerMode:       pulumi.String("code"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Pipelines can be imported using the `GraphQL ID` (not UUID), e.g.
//
// ```sh
//  $ pulumi import buildkite:index/pipeline:Pipeline fleet UGlwZWxpbmUtLS00MzVjYWQ1OC1lODFkLTQ1YWYtODYzNy1iMWNmODA3MDIzOGQ=
// ```
type Pipeline struct {
	pulumi.CustomResourceState

	AllowRebuilds pulumi.BoolOutput `pulumi:"allowRebuilds"`
	// The pipeline's last build status so you can display build status badge.
	BadgeUrl pulumi.StringOutput `pulumi:"badgeUrl"`
	// Limit which branches and tags cause new builds to be created, either via a code push or via the Builds REST API.
	BranchConfiguration pulumi.StringOutput `pulumi:"branchConfiguration"`
	// A boolean to enable automatically cancelling any running builds on the same branch when a new build is created.
	CancelIntermediateBuilds pulumi.BoolOutput `pulumi:"cancelIntermediateBuilds"`
	// Limit which branches build cancelling applies to, for example !master will ensure that the master branch won't have it's builds automatically cancelled.
	CancelIntermediateBuildsBranchFilter pulumi.StringOutput `pulumi:"cancelIntermediateBuildsBranchFilter"`
	// The GraphQL ID of the cluster you want to use for the pipeline.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The default branch to prefill when new builds are created or triggered, usually main or master but can be anything.
	DefaultBranch pulumi.StringOutput `pulumi:"defaultBranch"`
	// A description of the pipeline.
	Description pulumi.StringOutput `pulumi:"description"`
	// The name of the pipeline.
	Name pulumi.StringOutput `pulumi:"name"`
	// Source control provider settings for the pipeline. See Provider Settings Configuration below for details.
	ProviderSettings PipelineProviderSettingsOutput `pulumi:"providerSettings"`
	// The git URL of the repository.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// A boolean to enable automatically skipping any unstarted builds on the same branch when a new build is created.
	SkipIntermediateBuilds pulumi.BoolOutput `pulumi:"skipIntermediateBuilds"`
	// Limit which branches build skipping applies to, for example `!master` will ensure that the master branch won't have it's builds automatically skipped.
	SkipIntermediateBuildsBranchFilter pulumi.StringOutput `pulumi:"skipIntermediateBuildsBranchFilter"`
	// The buildkite slug of the team.
	Slug pulumi.StringOutput `pulumi:"slug"`
	// The string YAML steps to run the pipeline.
	Steps pulumi.StringOutput      `pulumi:"steps"`
	Tags  pulumi.StringArrayOutput `pulumi:"tags"`
	// Set team access for the pipeline. Can be specified multiple times for each team. See Teams Configuration below for details.
	Teams PipelineTeamArrayOutput `pulumi:"teams"`
	// The Buildkite webhook URL to configure on the repository to trigger builds on this pipeline.
	WebhookUrl pulumi.StringOutput `pulumi:"webhookUrl"`
}

// NewPipeline registers a new resource with the given unique name, arguments, and options.
func NewPipeline(ctx *pulumi.Context,
	name string, args *PipelineArgs, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.Steps == nil {
		return nil, errors.New("invalid value for required argument 'Steps'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Pipeline
	err := ctx.RegisterResource("buildkite:index/pipeline:Pipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipeline gets an existing Pipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineState, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	var resource Pipeline
	err := ctx.ReadResource("buildkite:index/pipeline:Pipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pipeline resources.
type pipelineState struct {
	AllowRebuilds *bool `pulumi:"allowRebuilds"`
	// The pipeline's last build status so you can display build status badge.
	BadgeUrl *string `pulumi:"badgeUrl"`
	// Limit which branches and tags cause new builds to be created, either via a code push or via the Builds REST API.
	BranchConfiguration *string `pulumi:"branchConfiguration"`
	// A boolean to enable automatically cancelling any running builds on the same branch when a new build is created.
	CancelIntermediateBuilds *bool `pulumi:"cancelIntermediateBuilds"`
	// Limit which branches build cancelling applies to, for example !master will ensure that the master branch won't have it's builds automatically cancelled.
	CancelIntermediateBuildsBranchFilter *string `pulumi:"cancelIntermediateBuildsBranchFilter"`
	// The GraphQL ID of the cluster you want to use for the pipeline.
	ClusterId *string `pulumi:"clusterId"`
	// The default branch to prefill when new builds are created or triggered, usually main or master but can be anything.
	DefaultBranch *string `pulumi:"defaultBranch"`
	// A description of the pipeline.
	Description *string `pulumi:"description"`
	// The name of the pipeline.
	Name *string `pulumi:"name"`
	// Source control provider settings for the pipeline. See Provider Settings Configuration below for details.
	ProviderSettings *PipelineProviderSettings `pulumi:"providerSettings"`
	// The git URL of the repository.
	Repository *string `pulumi:"repository"`
	// A boolean to enable automatically skipping any unstarted builds on the same branch when a new build is created.
	SkipIntermediateBuilds *bool `pulumi:"skipIntermediateBuilds"`
	// Limit which branches build skipping applies to, for example `!master` will ensure that the master branch won't have it's builds automatically skipped.
	SkipIntermediateBuildsBranchFilter *string `pulumi:"skipIntermediateBuildsBranchFilter"`
	// The buildkite slug of the team.
	Slug *string `pulumi:"slug"`
	// The string YAML steps to run the pipeline.
	Steps *string  `pulumi:"steps"`
	Tags  []string `pulumi:"tags"`
	// Set team access for the pipeline. Can be specified multiple times for each team. See Teams Configuration below for details.
	Teams []PipelineTeam `pulumi:"teams"`
	// The Buildkite webhook URL to configure on the repository to trigger builds on this pipeline.
	WebhookUrl *string `pulumi:"webhookUrl"`
}

type PipelineState struct {
	AllowRebuilds pulumi.BoolPtrInput
	// The pipeline's last build status so you can display build status badge.
	BadgeUrl pulumi.StringPtrInput
	// Limit which branches and tags cause new builds to be created, either via a code push or via the Builds REST API.
	BranchConfiguration pulumi.StringPtrInput
	// A boolean to enable automatically cancelling any running builds on the same branch when a new build is created.
	CancelIntermediateBuilds pulumi.BoolPtrInput
	// Limit which branches build cancelling applies to, for example !master will ensure that the master branch won't have it's builds automatically cancelled.
	CancelIntermediateBuildsBranchFilter pulumi.StringPtrInput
	// The GraphQL ID of the cluster you want to use for the pipeline.
	ClusterId pulumi.StringPtrInput
	// The default branch to prefill when new builds are created or triggered, usually main or master but can be anything.
	DefaultBranch pulumi.StringPtrInput
	// A description of the pipeline.
	Description pulumi.StringPtrInput
	// The name of the pipeline.
	Name pulumi.StringPtrInput
	// Source control provider settings for the pipeline. See Provider Settings Configuration below for details.
	ProviderSettings PipelineProviderSettingsPtrInput
	// The git URL of the repository.
	Repository pulumi.StringPtrInput
	// A boolean to enable automatically skipping any unstarted builds on the same branch when a new build is created.
	SkipIntermediateBuilds pulumi.BoolPtrInput
	// Limit which branches build skipping applies to, for example `!master` will ensure that the master branch won't have it's builds automatically skipped.
	SkipIntermediateBuildsBranchFilter pulumi.StringPtrInput
	// The buildkite slug of the team.
	Slug pulumi.StringPtrInput
	// The string YAML steps to run the pipeline.
	Steps pulumi.StringPtrInput
	Tags  pulumi.StringArrayInput
	// Set team access for the pipeline. Can be specified multiple times for each team. See Teams Configuration below for details.
	Teams PipelineTeamArrayInput
	// The Buildkite webhook URL to configure on the repository to trigger builds on this pipeline.
	WebhookUrl pulumi.StringPtrInput
}

func (PipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineState)(nil)).Elem()
}

type pipelineArgs struct {
	AllowRebuilds *bool `pulumi:"allowRebuilds"`
	// Limit which branches and tags cause new builds to be created, either via a code push or via the Builds REST API.
	BranchConfiguration *string `pulumi:"branchConfiguration"`
	// A boolean to enable automatically cancelling any running builds on the same branch when a new build is created.
	CancelIntermediateBuilds *bool `pulumi:"cancelIntermediateBuilds"`
	// Limit which branches build cancelling applies to, for example !master will ensure that the master branch won't have it's builds automatically cancelled.
	CancelIntermediateBuildsBranchFilter *string `pulumi:"cancelIntermediateBuildsBranchFilter"`
	// The GraphQL ID of the cluster you want to use for the pipeline.
	ClusterId *string `pulumi:"clusterId"`
	// The default branch to prefill when new builds are created or triggered, usually main or master but can be anything.
	DefaultBranch *string `pulumi:"defaultBranch"`
	// A description of the pipeline.
	Description *string `pulumi:"description"`
	// The name of the pipeline.
	Name *string `pulumi:"name"`
	// Source control provider settings for the pipeline. See Provider Settings Configuration below for details.
	ProviderSettings *PipelineProviderSettings `pulumi:"providerSettings"`
	// The git URL of the repository.
	Repository string `pulumi:"repository"`
	// A boolean to enable automatically skipping any unstarted builds on the same branch when a new build is created.
	SkipIntermediateBuilds *bool `pulumi:"skipIntermediateBuilds"`
	// Limit which branches build skipping applies to, for example `!master` will ensure that the master branch won't have it's builds automatically skipped.
	SkipIntermediateBuildsBranchFilter *string `pulumi:"skipIntermediateBuildsBranchFilter"`
	// The string YAML steps to run the pipeline.
	Steps string   `pulumi:"steps"`
	Tags  []string `pulumi:"tags"`
	// Set team access for the pipeline. Can be specified multiple times for each team. See Teams Configuration below for details.
	Teams []PipelineTeam `pulumi:"teams"`
}

// The set of arguments for constructing a Pipeline resource.
type PipelineArgs struct {
	AllowRebuilds pulumi.BoolPtrInput
	// Limit which branches and tags cause new builds to be created, either via a code push or via the Builds REST API.
	BranchConfiguration pulumi.StringPtrInput
	// A boolean to enable automatically cancelling any running builds on the same branch when a new build is created.
	CancelIntermediateBuilds pulumi.BoolPtrInput
	// Limit which branches build cancelling applies to, for example !master will ensure that the master branch won't have it's builds automatically cancelled.
	CancelIntermediateBuildsBranchFilter pulumi.StringPtrInput
	// The GraphQL ID of the cluster you want to use for the pipeline.
	ClusterId pulumi.StringPtrInput
	// The default branch to prefill when new builds are created or triggered, usually main or master but can be anything.
	DefaultBranch pulumi.StringPtrInput
	// A description of the pipeline.
	Description pulumi.StringPtrInput
	// The name of the pipeline.
	Name pulumi.StringPtrInput
	// Source control provider settings for the pipeline. See Provider Settings Configuration below for details.
	ProviderSettings PipelineProviderSettingsPtrInput
	// The git URL of the repository.
	Repository pulumi.StringInput
	// A boolean to enable automatically skipping any unstarted builds on the same branch when a new build is created.
	SkipIntermediateBuilds pulumi.BoolPtrInput
	// Limit which branches build skipping applies to, for example `!master` will ensure that the master branch won't have it's builds automatically skipped.
	SkipIntermediateBuildsBranchFilter pulumi.StringPtrInput
	// The string YAML steps to run the pipeline.
	Steps pulumi.StringInput
	Tags  pulumi.StringArrayInput
	// Set team access for the pipeline. Can be specified multiple times for each team. See Teams Configuration below for details.
	Teams PipelineTeamArrayInput
}

func (PipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineArgs)(nil)).Elem()
}

type PipelineInput interface {
	pulumi.Input

	ToPipelineOutput() PipelineOutput
	ToPipelineOutputWithContext(ctx context.Context) PipelineOutput
}

func (*Pipeline) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil)).Elem()
}

func (i *Pipeline) ToPipelineOutput() PipelineOutput {
	return i.ToPipelineOutputWithContext(context.Background())
}

func (i *Pipeline) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineOutput)
}

// PipelineArrayInput is an input type that accepts PipelineArray and PipelineArrayOutput values.
// You can construct a concrete instance of `PipelineArrayInput` via:
//
//          PipelineArray{ PipelineArgs{...} }
type PipelineArrayInput interface {
	pulumi.Input

	ToPipelineArrayOutput() PipelineArrayOutput
	ToPipelineArrayOutputWithContext(context.Context) PipelineArrayOutput
}

type PipelineArray []PipelineInput

func (PipelineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pipeline)(nil)).Elem()
}

func (i PipelineArray) ToPipelineArrayOutput() PipelineArrayOutput {
	return i.ToPipelineArrayOutputWithContext(context.Background())
}

func (i PipelineArray) ToPipelineArrayOutputWithContext(ctx context.Context) PipelineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineArrayOutput)
}

// PipelineMapInput is an input type that accepts PipelineMap and PipelineMapOutput values.
// You can construct a concrete instance of `PipelineMapInput` via:
//
//          PipelineMap{ "key": PipelineArgs{...} }
type PipelineMapInput interface {
	pulumi.Input

	ToPipelineMapOutput() PipelineMapOutput
	ToPipelineMapOutputWithContext(context.Context) PipelineMapOutput
}

type PipelineMap map[string]PipelineInput

func (PipelineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pipeline)(nil)).Elem()
}

func (i PipelineMap) ToPipelineMapOutput() PipelineMapOutput {
	return i.ToPipelineMapOutputWithContext(context.Background())
}

func (i PipelineMap) ToPipelineMapOutputWithContext(ctx context.Context) PipelineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineMapOutput)
}

type PipelineOutput struct{ *pulumi.OutputState }

func (PipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil)).Elem()
}

func (o PipelineOutput) ToPipelineOutput() PipelineOutput {
	return o
}

func (o PipelineOutput) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return o
}

func (o PipelineOutput) AllowRebuilds() pulumi.BoolOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.BoolOutput { return v.AllowRebuilds }).(pulumi.BoolOutput)
}

// The pipeline's last build status so you can display build status badge.
func (o PipelineOutput) BadgeUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.BadgeUrl }).(pulumi.StringOutput)
}

// Limit which branches and tags cause new builds to be created, either via a code push or via the Builds REST API.
func (o PipelineOutput) BranchConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.BranchConfiguration }).(pulumi.StringOutput)
}

// A boolean to enable automatically cancelling any running builds on the same branch when a new build is created.
func (o PipelineOutput) CancelIntermediateBuilds() pulumi.BoolOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.BoolOutput { return v.CancelIntermediateBuilds }).(pulumi.BoolOutput)
}

// Limit which branches build cancelling applies to, for example !master will ensure that the master branch won't have it's builds automatically cancelled.
func (o PipelineOutput) CancelIntermediateBuildsBranchFilter() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.CancelIntermediateBuildsBranchFilter }).(pulumi.StringOutput)
}

// The GraphQL ID of the cluster you want to use for the pipeline.
func (o PipelineOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The default branch to prefill when new builds are created or triggered, usually main or master but can be anything.
func (o PipelineOutput) DefaultBranch() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.DefaultBranch }).(pulumi.StringOutput)
}

// A description of the pipeline.
func (o PipelineOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The name of the pipeline.
func (o PipelineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Source control provider settings for the pipeline. See Provider Settings Configuration below for details.
func (o PipelineOutput) ProviderSettings() PipelineProviderSettingsOutput {
	return o.ApplyT(func(v *Pipeline) PipelineProviderSettingsOutput { return v.ProviderSettings }).(PipelineProviderSettingsOutput)
}

// The git URL of the repository.
func (o PipelineOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// A boolean to enable automatically skipping any unstarted builds on the same branch when a new build is created.
func (o PipelineOutput) SkipIntermediateBuilds() pulumi.BoolOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.BoolOutput { return v.SkipIntermediateBuilds }).(pulumi.BoolOutput)
}

// Limit which branches build skipping applies to, for example `!master` will ensure that the master branch won't have it's builds automatically skipped.
func (o PipelineOutput) SkipIntermediateBuildsBranchFilter() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.SkipIntermediateBuildsBranchFilter }).(pulumi.StringOutput)
}

// The buildkite slug of the team.
func (o PipelineOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// The string YAML steps to run the pipeline.
func (o PipelineOutput) Steps() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.Steps }).(pulumi.StringOutput)
}

func (o PipelineOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Set team access for the pipeline. Can be specified multiple times for each team. See Teams Configuration below for details.
func (o PipelineOutput) Teams() PipelineTeamArrayOutput {
	return o.ApplyT(func(v *Pipeline) PipelineTeamArrayOutput { return v.Teams }).(PipelineTeamArrayOutput)
}

// The Buildkite webhook URL to configure on the repository to trigger builds on this pipeline.
func (o PipelineOutput) WebhookUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.WebhookUrl }).(pulumi.StringOutput)
}

type PipelineArrayOutput struct{ *pulumi.OutputState }

func (PipelineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pipeline)(nil)).Elem()
}

func (o PipelineArrayOutput) ToPipelineArrayOutput() PipelineArrayOutput {
	return o
}

func (o PipelineArrayOutput) ToPipelineArrayOutputWithContext(ctx context.Context) PipelineArrayOutput {
	return o
}

func (o PipelineArrayOutput) Index(i pulumi.IntInput) PipelineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Pipeline {
		return vs[0].([]*Pipeline)[vs[1].(int)]
	}).(PipelineOutput)
}

type PipelineMapOutput struct{ *pulumi.OutputState }

func (PipelineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pipeline)(nil)).Elem()
}

func (o PipelineMapOutput) ToPipelineMapOutput() PipelineMapOutput {
	return o
}

func (o PipelineMapOutput) ToPipelineMapOutputWithContext(ctx context.Context) PipelineMapOutput {
	return o
}

func (o PipelineMapOutput) MapIndex(k pulumi.StringInput) PipelineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Pipeline {
		return vs[0].(map[string]*Pipeline)[vs[1].(string)]
	}).(PipelineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineInput)(nil)).Elem(), &Pipeline{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineArrayInput)(nil)).Elem(), PipelineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineMapInput)(nil)).Elem(), PipelineMap{})
	pulumi.RegisterOutputType(PipelineOutput{})
	pulumi.RegisterOutputType(PipelineArrayOutput{})
	pulumi.RegisterOutputType(PipelineMapOutput{})
}
