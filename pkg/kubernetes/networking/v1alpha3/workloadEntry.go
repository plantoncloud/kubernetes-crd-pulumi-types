// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha3

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkloadEntry struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Configuration affecting VMs onboarded into the mesh. See more details at: https://istio.io/docs/reference/config/networking/workload-entry.html
	Spec   WorkloadEntrySpecOutput `pulumi:"spec"`
	Status pulumi.MapOutput        `pulumi:"status"`
}

// NewWorkloadEntry registers a new resource with the given unique name, arguments, and options.
func NewWorkloadEntry(ctx *pulumi.Context,
	name string, args *WorkloadEntryArgs, opts ...pulumi.ResourceOption) (*WorkloadEntry, error) {
	if args == nil {
		args = &WorkloadEntryArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("networking.istio.io/v1alpha3")
	args.Kind = pulumi.StringPtr("WorkloadEntry")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkloadEntry
	err := ctx.RegisterResource("kubernetes:networking.istio.io/v1alpha3:WorkloadEntry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadEntry gets an existing WorkloadEntry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadEntryState, opts ...pulumi.ResourceOption) (*WorkloadEntry, error) {
	var resource WorkloadEntry
	err := ctx.ReadResource("kubernetes:networking.istio.io/v1alpha3:WorkloadEntry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadEntry resources.
type workloadEntryState struct {
}

type WorkloadEntryState struct {
}

func (WorkloadEntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadEntryState)(nil)).Elem()
}

type workloadEntryArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// Configuration affecting VMs onboarded into the mesh. See more details at: https://istio.io/docs/reference/config/networking/workload-entry.html
	Spec   *WorkloadEntrySpec     `pulumi:"spec"`
	Status map[string]interface{} `pulumi:"status"`
}

// The set of arguments for constructing a WorkloadEntry resource.
type WorkloadEntryArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// Configuration affecting VMs onboarded into the mesh. See more details at: https://istio.io/docs/reference/config/networking/workload-entry.html
	Spec   WorkloadEntrySpecPtrInput
	Status pulumi.MapInput
}

func (WorkloadEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadEntryArgs)(nil)).Elem()
}

type WorkloadEntryInput interface {
	pulumi.Input

	ToWorkloadEntryOutput() WorkloadEntryOutput
	ToWorkloadEntryOutputWithContext(ctx context.Context) WorkloadEntryOutput
}

func (*WorkloadEntry) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadEntry)(nil)).Elem()
}

func (i *WorkloadEntry) ToWorkloadEntryOutput() WorkloadEntryOutput {
	return i.ToWorkloadEntryOutputWithContext(context.Background())
}

func (i *WorkloadEntry) ToWorkloadEntryOutputWithContext(ctx context.Context) WorkloadEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadEntryOutput)
}

type WorkloadEntryOutput struct{ *pulumi.OutputState }

func (WorkloadEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadEntry)(nil)).Elem()
}

func (o WorkloadEntryOutput) ToWorkloadEntryOutput() WorkloadEntryOutput {
	return o
}

func (o WorkloadEntryOutput) ToWorkloadEntryOutputWithContext(ctx context.Context) WorkloadEntryOutput {
	return o
}

func (o WorkloadEntryOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadEntry) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o WorkloadEntryOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadEntry) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WorkloadEntryOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *WorkloadEntry) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Configuration affecting VMs onboarded into the mesh. See more details at: https://istio.io/docs/reference/config/networking/workload-entry.html
func (o WorkloadEntryOutput) Spec() WorkloadEntrySpecOutput {
	return o.ApplyT(func(v *WorkloadEntry) WorkloadEntrySpecOutput { return v.Spec }).(WorkloadEntrySpecOutput)
}

func (o WorkloadEntryOutput) Status() pulumi.MapOutput {
	return o.ApplyT(func(v *WorkloadEntry) pulumi.MapOutput { return v.Status }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadEntryInput)(nil)).Elem(), &WorkloadEntry{})
	pulumi.RegisterOutputType(WorkloadEntryOutput{})
}
