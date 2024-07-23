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

type DestinationRule struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Configuration affecting load balancing, outlier detection, etc. See more details at: https://istio.io/docs/reference/config/networking/destination-rule.html
	Spec   DestinationRuleSpecPtrOutput `pulumi:"spec"`
	Status pulumi.MapOutput             `pulumi:"status"`
}

// NewDestinationRule registers a new resource with the given unique name, arguments, and options.
func NewDestinationRule(ctx *pulumi.Context,
	name string, args *DestinationRuleArgs, opts ...pulumi.ResourceOption) (*DestinationRule, error) {
	if args == nil {
		args = &DestinationRuleArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("networking.istio.io/v1alpha3")
	args.Kind = pulumi.StringPtr("DestinationRule")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DestinationRule
	err := ctx.RegisterResource("kubernetes:networking.istio.io/v1alpha3:DestinationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationRule gets an existing DestinationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationRuleState, opts ...pulumi.ResourceOption) (*DestinationRule, error) {
	var resource DestinationRule
	err := ctx.ReadResource("kubernetes:networking.istio.io/v1alpha3:DestinationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationRule resources.
type destinationRuleState struct {
}

type DestinationRuleState struct {
}

func (DestinationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationRuleState)(nil)).Elem()
}

type destinationRuleArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// Configuration affecting load balancing, outlier detection, etc. See more details at: https://istio.io/docs/reference/config/networking/destination-rule.html
	Spec   *DestinationRuleSpec   `pulumi:"spec"`
	Status map[string]interface{} `pulumi:"status"`
}

// The set of arguments for constructing a DestinationRule resource.
type DestinationRuleArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// Configuration affecting load balancing, outlier detection, etc. See more details at: https://istio.io/docs/reference/config/networking/destination-rule.html
	Spec   DestinationRuleSpecPtrInput
	Status pulumi.MapInput
}

func (DestinationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationRuleArgs)(nil)).Elem()
}

type DestinationRuleInput interface {
	pulumi.Input

	ToDestinationRuleOutput() DestinationRuleOutput
	ToDestinationRuleOutputWithContext(ctx context.Context) DestinationRuleOutput
}

func (*DestinationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationRule)(nil)).Elem()
}

func (i *DestinationRule) ToDestinationRuleOutput() DestinationRuleOutput {
	return i.ToDestinationRuleOutputWithContext(context.Background())
}

func (i *DestinationRule) ToDestinationRuleOutputWithContext(ctx context.Context) DestinationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationRuleOutput)
}

type DestinationRuleOutput struct{ *pulumi.OutputState }

func (DestinationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationRule)(nil)).Elem()
}

func (o DestinationRuleOutput) ToDestinationRuleOutput() DestinationRuleOutput {
	return o
}

func (o DestinationRuleOutput) ToDestinationRuleOutputWithContext(ctx context.Context) DestinationRuleOutput {
	return o
}

func (o DestinationRuleOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationRule) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o DestinationRuleOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationRule) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o DestinationRuleOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *DestinationRule) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Configuration affecting load balancing, outlier detection, etc. See more details at: https://istio.io/docs/reference/config/networking/destination-rule.html
func (o DestinationRuleOutput) Spec() DestinationRuleSpecPtrOutput {
	return o.ApplyT(func(v *DestinationRule) DestinationRuleSpecPtrOutput { return v.Spec }).(DestinationRuleSpecPtrOutput)
}

func (o DestinationRuleOutput) Status() pulumi.MapOutput {
	return o.ApplyT(func(v *DestinationRule) pulumi.MapOutput { return v.Status }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationRuleInput)(nil)).Elem(), &DestinationRule{})
	pulumi.RegisterOutputType(DestinationRuleOutput{})
}
