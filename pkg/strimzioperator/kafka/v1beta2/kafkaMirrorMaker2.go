// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KafkaMirrorMaker2 struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// The specification of the Kafka MirrorMaker 2 cluster.
	Spec KafkaMirrorMaker2SpecPtrOutput `pulumi:"spec"`
	// The status of the Kafka MirrorMaker 2 cluster.
	Status KafkaMirrorMaker2StatusPtrOutput `pulumi:"status"`
}

// NewKafkaMirrorMaker2 registers a new resource with the given unique name, arguments, and options.
func NewKafkaMirrorMaker2(ctx *pulumi.Context,
	name string, args *KafkaMirrorMaker2Args, opts ...pulumi.ResourceOption) (*KafkaMirrorMaker2, error) {
	if args == nil {
		args = &KafkaMirrorMaker2Args{}
	}

	args.ApiVersion = pulumi.StringPtr("kafka.strimzi.io/v1beta2")
	args.Kind = pulumi.StringPtr("KafkaMirrorMaker2")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource KafkaMirrorMaker2
	err := ctx.RegisterResource("kubernetes:kafka.strimzi.io/v1beta2:KafkaMirrorMaker2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKafkaMirrorMaker2 gets an existing KafkaMirrorMaker2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKafkaMirrorMaker2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KafkaMirrorMaker2State, opts ...pulumi.ResourceOption) (*KafkaMirrorMaker2, error) {
	var resource KafkaMirrorMaker2
	err := ctx.ReadResource("kubernetes:kafka.strimzi.io/v1beta2:KafkaMirrorMaker2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KafkaMirrorMaker2 resources.
type kafkaMirrorMaker2State struct {
}

type KafkaMirrorMaker2State struct {
}

func (KafkaMirrorMaker2State) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaMirrorMaker2State)(nil)).Elem()
}

type kafkaMirrorMaker2Args struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// The specification of the Kafka MirrorMaker 2 cluster.
	Spec *KafkaMirrorMaker2Spec `pulumi:"spec"`
	// The status of the Kafka MirrorMaker 2 cluster.
	Status *KafkaMirrorMaker2Status `pulumi:"status"`
}

// The set of arguments for constructing a KafkaMirrorMaker2 resource.
type KafkaMirrorMaker2Args struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// The specification of the Kafka MirrorMaker 2 cluster.
	Spec KafkaMirrorMaker2SpecPtrInput
	// The status of the Kafka MirrorMaker 2 cluster.
	Status KafkaMirrorMaker2StatusPtrInput
}

func (KafkaMirrorMaker2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaMirrorMaker2Args)(nil)).Elem()
}

type KafkaMirrorMaker2Input interface {
	pulumi.Input

	ToKafkaMirrorMaker2Output() KafkaMirrorMaker2Output
	ToKafkaMirrorMaker2OutputWithContext(ctx context.Context) KafkaMirrorMaker2Output
}

func (*KafkaMirrorMaker2) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaMirrorMaker2)(nil)).Elem()
}

func (i *KafkaMirrorMaker2) ToKafkaMirrorMaker2Output() KafkaMirrorMaker2Output {
	return i.ToKafkaMirrorMaker2OutputWithContext(context.Background())
}

func (i *KafkaMirrorMaker2) ToKafkaMirrorMaker2OutputWithContext(ctx context.Context) KafkaMirrorMaker2Output {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaMirrorMaker2Output)
}

type KafkaMirrorMaker2Output struct{ *pulumi.OutputState }

func (KafkaMirrorMaker2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaMirrorMaker2)(nil)).Elem()
}

func (o KafkaMirrorMaker2Output) ToKafkaMirrorMaker2Output() KafkaMirrorMaker2Output {
	return o
}

func (o KafkaMirrorMaker2Output) ToKafkaMirrorMaker2OutputWithContext(ctx context.Context) KafkaMirrorMaker2Output {
	return o
}

func (o KafkaMirrorMaker2Output) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KafkaMirrorMaker2) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o KafkaMirrorMaker2Output) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KafkaMirrorMaker2) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o KafkaMirrorMaker2Output) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *KafkaMirrorMaker2) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// The specification of the Kafka MirrorMaker 2 cluster.
func (o KafkaMirrorMaker2Output) Spec() KafkaMirrorMaker2SpecPtrOutput {
	return o.ApplyT(func(v *KafkaMirrorMaker2) KafkaMirrorMaker2SpecPtrOutput { return v.Spec }).(KafkaMirrorMaker2SpecPtrOutput)
}

// The status of the Kafka MirrorMaker 2 cluster.
func (o KafkaMirrorMaker2Output) Status() KafkaMirrorMaker2StatusPtrOutput {
	return o.ApplyT(func(v *KafkaMirrorMaker2) KafkaMirrorMaker2StatusPtrOutput { return v.Status }).(KafkaMirrorMaker2StatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaMirrorMaker2Input)(nil)).Elem(), &KafkaMirrorMaker2{})
	pulumi.RegisterOutputType(KafkaMirrorMaker2Output{})
}
