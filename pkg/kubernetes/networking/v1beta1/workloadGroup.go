// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkloadGroup struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Describes a collection of workload instances. See more details at: https://istio.io/docs/reference/config/networking/workload-group.html
	Spec   WorkloadGroupSpecPtrOutput `pulumi:"spec"`
	Status pulumi.MapOutput           `pulumi:"status"`
}

// NewWorkloadGroup registers a new resource with the given unique name, arguments, and options.
func NewWorkloadGroup(ctx *pulumi.Context,
	name string, args *WorkloadGroupArgs, opts ...pulumi.ResourceOption) (*WorkloadGroup, error) {
	if args == nil {
		args = &WorkloadGroupArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("networking.istio.io/v1beta1")
	args.Kind = pulumi.StringPtr("WorkloadGroup")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkloadGroup
	err := ctx.RegisterResource("kubernetes:networking.istio.io/v1beta1:WorkloadGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadGroup gets an existing WorkloadGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadGroupState, opts ...pulumi.ResourceOption) (*WorkloadGroup, error) {
	var resource WorkloadGroup
	err := ctx.ReadResource("kubernetes:networking.istio.io/v1beta1:WorkloadGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadGroup resources.
type workloadGroupState struct {
}

type WorkloadGroupState struct {
}

func (WorkloadGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadGroupState)(nil)).Elem()
}

type workloadGroupArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// Describes a collection of workload instances. See more details at: https://istio.io/docs/reference/config/networking/workload-group.html
	Spec   *WorkloadGroupSpec     `pulumi:"spec"`
	Status map[string]interface{} `pulumi:"status"`
}

// The set of arguments for constructing a WorkloadGroup resource.
type WorkloadGroupArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// Describes a collection of workload instances. See more details at: https://istio.io/docs/reference/config/networking/workload-group.html
	Spec   WorkloadGroupSpecPtrInput
	Status pulumi.MapInput
}

func (WorkloadGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadGroupArgs)(nil)).Elem()
}

type WorkloadGroupInput interface {
	pulumi.Input

	ToWorkloadGroupOutput() WorkloadGroupOutput
	ToWorkloadGroupOutputWithContext(ctx context.Context) WorkloadGroupOutput
}

func (*WorkloadGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadGroup)(nil)).Elem()
}

func (i *WorkloadGroup) ToWorkloadGroupOutput() WorkloadGroupOutput {
	return i.ToWorkloadGroupOutputWithContext(context.Background())
}

func (i *WorkloadGroup) ToWorkloadGroupOutputWithContext(ctx context.Context) WorkloadGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadGroupOutput)
}

type WorkloadGroupOutput struct{ *pulumi.OutputState }

func (WorkloadGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadGroup)(nil)).Elem()
}

func (o WorkloadGroupOutput) ToWorkloadGroupOutput() WorkloadGroupOutput {
	return o
}

func (o WorkloadGroupOutput) ToWorkloadGroupOutputWithContext(ctx context.Context) WorkloadGroupOutput {
	return o
}

func (o WorkloadGroupOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o WorkloadGroupOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WorkloadGroupOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *WorkloadGroup) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Describes a collection of workload instances. See more details at: https://istio.io/docs/reference/config/networking/workload-group.html
func (o WorkloadGroupOutput) Spec() WorkloadGroupSpecPtrOutput {
	return o.ApplyT(func(v *WorkloadGroup) WorkloadGroupSpecPtrOutput { return v.Spec }).(WorkloadGroupSpecPtrOutput)
}

func (o WorkloadGroupOutput) Status() pulumi.MapOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.MapOutput { return v.Status }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadGroupInput)(nil)).Elem(), &WorkloadGroup{})
	pulumi.RegisterOutputType(WorkloadGroupOutput{})
}
