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

type KafkaUser struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// The specification of the user.
	Spec KafkaUserSpecPtrOutput `pulumi:"spec"`
	// The status of the Kafka User.
	Status KafkaUserStatusPtrOutput `pulumi:"status"`
}

// NewKafkaUser registers a new resource with the given unique name, arguments, and options.
func NewKafkaUser(ctx *pulumi.Context,
	name string, args *KafkaUserArgs, opts ...pulumi.ResourceOption) (*KafkaUser, error) {
	if args == nil {
		args = &KafkaUserArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("kafka.strimzi.io/v1beta2")
	args.Kind = pulumi.StringPtr("KafkaUser")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource KafkaUser
	err := ctx.RegisterResource("kubernetes:kafka.strimzi.io/v1beta2:KafkaUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKafkaUser gets an existing KafkaUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKafkaUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KafkaUserState, opts ...pulumi.ResourceOption) (*KafkaUser, error) {
	var resource KafkaUser
	err := ctx.ReadResource("kubernetes:kafka.strimzi.io/v1beta2:KafkaUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KafkaUser resources.
type kafkaUserState struct {
}

type KafkaUserState struct {
}

func (KafkaUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaUserState)(nil)).Elem()
}

type kafkaUserArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// The specification of the user.
	Spec *KafkaUserSpec `pulumi:"spec"`
	// The status of the Kafka User.
	Status *KafkaUserStatus `pulumi:"status"`
}

// The set of arguments for constructing a KafkaUser resource.
type KafkaUserArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// The specification of the user.
	Spec KafkaUserSpecPtrInput
	// The status of the Kafka User.
	Status KafkaUserStatusPtrInput
}

func (KafkaUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaUserArgs)(nil)).Elem()
}

type KafkaUserInput interface {
	pulumi.Input

	ToKafkaUserOutput() KafkaUserOutput
	ToKafkaUserOutputWithContext(ctx context.Context) KafkaUserOutput
}

func (*KafkaUser) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaUser)(nil)).Elem()
}

func (i *KafkaUser) ToKafkaUserOutput() KafkaUserOutput {
	return i.ToKafkaUserOutputWithContext(context.Background())
}

func (i *KafkaUser) ToKafkaUserOutputWithContext(ctx context.Context) KafkaUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaUserOutput)
}

type KafkaUserOutput struct{ *pulumi.OutputState }

func (KafkaUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaUser)(nil)).Elem()
}

func (o KafkaUserOutput) ToKafkaUserOutput() KafkaUserOutput {
	return o
}

func (o KafkaUserOutput) ToKafkaUserOutputWithContext(ctx context.Context) KafkaUserOutput {
	return o
}

func (o KafkaUserOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KafkaUser) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o KafkaUserOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KafkaUser) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o KafkaUserOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *KafkaUser) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// The specification of the user.
func (o KafkaUserOutput) Spec() KafkaUserSpecPtrOutput {
	return o.ApplyT(func(v *KafkaUser) KafkaUserSpecPtrOutput { return v.Spec }).(KafkaUserSpecPtrOutput)
}

// The status of the Kafka User.
func (o KafkaUserOutput) Status() KafkaUserStatusPtrOutput {
	return o.ApplyT(func(v *KafkaUser) KafkaUserStatusPtrOutput { return v.Status }).(KafkaUserStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaUserInput)(nil)).Elem(), &KafkaUser{})
	pulumi.RegisterOutputType(KafkaUserOutput{})
}
