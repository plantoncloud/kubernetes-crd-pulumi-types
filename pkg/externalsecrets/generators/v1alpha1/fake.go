// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Fake generator is used for testing. It lets you define
// a static set of credentials that is always returned.
type Fake struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// FakeSpec contains the static data.
	Spec FakeSpecPtrOutput `pulumi:"spec"`
}

// NewFake registers a new resource with the given unique name, arguments, and options.
func NewFake(ctx *pulumi.Context,
	name string, args *FakeArgs, opts ...pulumi.ResourceOption) (*Fake, error) {
	if args == nil {
		args = &FakeArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("generators.external-secrets.io/v1alpha1")
	args.Kind = pulumi.StringPtr("Fake")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Fake
	err := ctx.RegisterResource("kubernetes:generators.external-secrets.io/v1alpha1:Fake", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFake gets an existing Fake resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFake(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FakeState, opts ...pulumi.ResourceOption) (*Fake, error) {
	var resource Fake
	err := ctx.ReadResource("kubernetes:generators.external-secrets.io/v1alpha1:Fake", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fake resources.
type fakeState struct {
}

type FakeState struct {
}

func (FakeState) ElementType() reflect.Type {
	return reflect.TypeOf((*fakeState)(nil)).Elem()
}

type fakeArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// FakeSpec contains the static data.
	Spec *FakeSpec `pulumi:"spec"`
}

// The set of arguments for constructing a Fake resource.
type FakeArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// FakeSpec contains the static data.
	Spec FakeSpecPtrInput
}

func (FakeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fakeArgs)(nil)).Elem()
}

type FakeInput interface {
	pulumi.Input

	ToFakeOutput() FakeOutput
	ToFakeOutputWithContext(ctx context.Context) FakeOutput
}

func (*Fake) ElementType() reflect.Type {
	return reflect.TypeOf((**Fake)(nil)).Elem()
}

func (i *Fake) ToFakeOutput() FakeOutput {
	return i.ToFakeOutputWithContext(context.Background())
}

func (i *Fake) ToFakeOutputWithContext(ctx context.Context) FakeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FakeOutput)
}

type FakeOutput struct{ *pulumi.OutputState }

func (FakeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fake)(nil)).Elem()
}

func (o FakeOutput) ToFakeOutput() FakeOutput {
	return o
}

func (o FakeOutput) ToFakeOutputWithContext(ctx context.Context) FakeOutput {
	return o
}

func (o FakeOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fake) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o FakeOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fake) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o FakeOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *Fake) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// FakeSpec contains the static data.
func (o FakeOutput) Spec() FakeSpecPtrOutput {
	return o.ApplyT(func(v *Fake) FakeSpecPtrOutput { return v.Spec }).(FakeSpecPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FakeInput)(nil)).Elem(), &Fake{})
	pulumi.RegisterOutputType(FakeOutput{})
}
