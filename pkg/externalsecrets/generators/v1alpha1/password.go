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

// Password generates a random password based on the
// configuration parameters in spec.
// You can specify the length, characterset and other attributes.
type Password struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// PasswordSpec controls the behavior of the password generator.
	Spec PasswordSpecPtrOutput `pulumi:"spec"`
}

// NewPassword registers a new resource with the given unique name, arguments, and options.
func NewPassword(ctx *pulumi.Context,
	name string, args *PasswordArgs, opts ...pulumi.ResourceOption) (*Password, error) {
	if args == nil {
		args = &PasswordArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("generators.external-secrets.io/v1alpha1")
	args.Kind = pulumi.StringPtr("Password")
	if args.Spec != nil {
		args.Spec = args.Spec.ToPasswordSpecPtrOutput().ApplyT(func(v *PasswordSpec) *PasswordSpec { return v.Defaults() }).(PasswordSpecPtrOutput)
	}
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Password
	err := ctx.RegisterResource("kubernetes:generators.external-secrets.io/v1alpha1:Password", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPassword gets an existing Password resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPassword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PasswordState, opts ...pulumi.ResourceOption) (*Password, error) {
	var resource Password
	err := ctx.ReadResource("kubernetes:generators.external-secrets.io/v1alpha1:Password", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Password resources.
type passwordState struct {
}

type PasswordState struct {
}

func (PasswordState) ElementType() reflect.Type {
	return reflect.TypeOf((*passwordState)(nil)).Elem()
}

type passwordArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// PasswordSpec controls the behavior of the password generator.
	Spec *PasswordSpec `pulumi:"spec"`
}

// The set of arguments for constructing a Password resource.
type PasswordArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// PasswordSpec controls the behavior of the password generator.
	Spec PasswordSpecPtrInput
}

func (PasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*passwordArgs)(nil)).Elem()
}

type PasswordInput interface {
	pulumi.Input

	ToPasswordOutput() PasswordOutput
	ToPasswordOutputWithContext(ctx context.Context) PasswordOutput
}

func (*Password) ElementType() reflect.Type {
	return reflect.TypeOf((**Password)(nil)).Elem()
}

func (i *Password) ToPasswordOutput() PasswordOutput {
	return i.ToPasswordOutputWithContext(context.Background())
}

func (i *Password) ToPasswordOutputWithContext(ctx context.Context) PasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordOutput)
}

type PasswordOutput struct{ *pulumi.OutputState }

func (PasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Password)(nil)).Elem()
}

func (o PasswordOutput) ToPasswordOutput() PasswordOutput {
	return o
}

func (o PasswordOutput) ToPasswordOutputWithContext(ctx context.Context) PasswordOutput {
	return o
}

func (o PasswordOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Password) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o PasswordOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Password) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o PasswordOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *Password) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// PasswordSpec controls the behavior of the password generator.
func (o PasswordOutput) Spec() PasswordSpecPtrOutput {
	return o.ApplyT(func(v *Password) PasswordSpecPtrOutput { return v.Spec }).(PasswordSpecPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordInput)(nil)).Elem(), &Password{})
	pulumi.RegisterOutputType(PasswordOutput{})
}
