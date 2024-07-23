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

type EnvoyFilter struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Customizing Envoy configuration generated by Istio. See more details at: https://istio.io/docs/reference/config/networking/envoy-filter.html
	Spec   EnvoyFilterSpecPtrOutput `pulumi:"spec"`
	Status pulumi.MapOutput         `pulumi:"status"`
}

// NewEnvoyFilter registers a new resource with the given unique name, arguments, and options.
func NewEnvoyFilter(ctx *pulumi.Context,
	name string, args *EnvoyFilterArgs, opts ...pulumi.ResourceOption) (*EnvoyFilter, error) {
	if args == nil {
		args = &EnvoyFilterArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("networking.istio.io/v1alpha3")
	args.Kind = pulumi.StringPtr("EnvoyFilter")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EnvoyFilter
	err := ctx.RegisterResource("kubernetes:networking.istio.io/v1alpha3:EnvoyFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvoyFilter gets an existing EnvoyFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvoyFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvoyFilterState, opts ...pulumi.ResourceOption) (*EnvoyFilter, error) {
	var resource EnvoyFilter
	err := ctx.ReadResource("kubernetes:networking.istio.io/v1alpha3:EnvoyFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvoyFilter resources.
type envoyFilterState struct {
}

type EnvoyFilterState struct {
}

func (EnvoyFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*envoyFilterState)(nil)).Elem()
}

type envoyFilterArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// Customizing Envoy configuration generated by Istio. See more details at: https://istio.io/docs/reference/config/networking/envoy-filter.html
	Spec   *EnvoyFilterSpec       `pulumi:"spec"`
	Status map[string]interface{} `pulumi:"status"`
}

// The set of arguments for constructing a EnvoyFilter resource.
type EnvoyFilterArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// Customizing Envoy configuration generated by Istio. See more details at: https://istio.io/docs/reference/config/networking/envoy-filter.html
	Spec   EnvoyFilterSpecPtrInput
	Status pulumi.MapInput
}

func (EnvoyFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*envoyFilterArgs)(nil)).Elem()
}

type EnvoyFilterInput interface {
	pulumi.Input

	ToEnvoyFilterOutput() EnvoyFilterOutput
	ToEnvoyFilterOutputWithContext(ctx context.Context) EnvoyFilterOutput
}

func (*EnvoyFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvoyFilter)(nil)).Elem()
}

func (i *EnvoyFilter) ToEnvoyFilterOutput() EnvoyFilterOutput {
	return i.ToEnvoyFilterOutputWithContext(context.Background())
}

func (i *EnvoyFilter) ToEnvoyFilterOutputWithContext(ctx context.Context) EnvoyFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvoyFilterOutput)
}

type EnvoyFilterOutput struct{ *pulumi.OutputState }

func (EnvoyFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvoyFilter)(nil)).Elem()
}

func (o EnvoyFilterOutput) ToEnvoyFilterOutput() EnvoyFilterOutput {
	return o
}

func (o EnvoyFilterOutput) ToEnvoyFilterOutputWithContext(ctx context.Context) EnvoyFilterOutput {
	return o
}

func (o EnvoyFilterOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvoyFilter) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o EnvoyFilterOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvoyFilter) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o EnvoyFilterOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *EnvoyFilter) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Customizing Envoy configuration generated by Istio. See more details at: https://istio.io/docs/reference/config/networking/envoy-filter.html
func (o EnvoyFilterOutput) Spec() EnvoyFilterSpecPtrOutput {
	return o.ApplyT(func(v *EnvoyFilter) EnvoyFilterSpecPtrOutput { return v.Spec }).(EnvoyFilterSpecPtrOutput)
}

func (o EnvoyFilterOutput) Status() pulumi.MapOutput {
	return o.ApplyT(func(v *EnvoyFilter) pulumi.MapOutput { return v.Status }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvoyFilterInput)(nil)).Elem(), &EnvoyFilter{})
	pulumi.RegisterOutputType(EnvoyFilterOutput{})
}
