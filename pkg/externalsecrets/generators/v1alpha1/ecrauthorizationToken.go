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

// ECRAuthorizationTokenSpec uses the GetAuthorizationToken API to retrieve an
// authorization token.
// The authorization token is valid for 12 hours.
// The authorizationToken returned is a base64 encoded string that can be decoded
// and used in a docker login command to authenticate to a registry.
// For more information, see Registry authentication (https://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html#registry_auth) in the Amazon Elastic Container Registry User Guide.
type ECRAuthorizationToken struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput             `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput             `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput         `pulumi:"metadata"`
	Spec       ECRAuthorizationTokenSpecPtrOutput `pulumi:"spec"`
}

// NewECRAuthorizationToken registers a new resource with the given unique name, arguments, and options.
func NewECRAuthorizationToken(ctx *pulumi.Context,
	name string, args *ECRAuthorizationTokenArgs, opts ...pulumi.ResourceOption) (*ECRAuthorizationToken, error) {
	if args == nil {
		args = &ECRAuthorizationTokenArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("generators.external-secrets.io/v1alpha1")
	args.Kind = pulumi.StringPtr("ECRAuthorizationToken")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ECRAuthorizationToken
	err := ctx.RegisterResource("kubernetes:generators.external-secrets.io/v1alpha1:ECRAuthorizationToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetECRAuthorizationToken gets an existing ECRAuthorizationToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetECRAuthorizationToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ECRAuthorizationTokenState, opts ...pulumi.ResourceOption) (*ECRAuthorizationToken, error) {
	var resource ECRAuthorizationToken
	err := ctx.ReadResource("kubernetes:generators.external-secrets.io/v1alpha1:ECRAuthorizationToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ECRAuthorizationToken resources.
type ecrauthorizationTokenState struct {
}

type ECRAuthorizationTokenState struct {
}

func (ECRAuthorizationTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*ecrauthorizationTokenState)(nil)).Elem()
}

type ecrauthorizationTokenArgs struct {
	ApiVersion *string                    `pulumi:"apiVersion"`
	Kind       *string                    `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta         `pulumi:"metadata"`
	Spec       *ECRAuthorizationTokenSpec `pulumi:"spec"`
}

// The set of arguments for constructing a ECRAuthorizationToken resource.
type ECRAuthorizationTokenArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	Spec       ECRAuthorizationTokenSpecPtrInput
}

func (ECRAuthorizationTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ecrauthorizationTokenArgs)(nil)).Elem()
}

type ECRAuthorizationTokenInput interface {
	pulumi.Input

	ToECRAuthorizationTokenOutput() ECRAuthorizationTokenOutput
	ToECRAuthorizationTokenOutputWithContext(ctx context.Context) ECRAuthorizationTokenOutput
}

func (*ECRAuthorizationToken) ElementType() reflect.Type {
	return reflect.TypeOf((**ECRAuthorizationToken)(nil)).Elem()
}

func (i *ECRAuthorizationToken) ToECRAuthorizationTokenOutput() ECRAuthorizationTokenOutput {
	return i.ToECRAuthorizationTokenOutputWithContext(context.Background())
}

func (i *ECRAuthorizationToken) ToECRAuthorizationTokenOutputWithContext(ctx context.Context) ECRAuthorizationTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECRAuthorizationTokenOutput)
}

type ECRAuthorizationTokenOutput struct{ *pulumi.OutputState }

func (ECRAuthorizationTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ECRAuthorizationToken)(nil)).Elem()
}

func (o ECRAuthorizationTokenOutput) ToECRAuthorizationTokenOutput() ECRAuthorizationTokenOutput {
	return o
}

func (o ECRAuthorizationTokenOutput) ToECRAuthorizationTokenOutputWithContext(ctx context.Context) ECRAuthorizationTokenOutput {
	return o
}

func (o ECRAuthorizationTokenOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ECRAuthorizationToken) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o ECRAuthorizationTokenOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ECRAuthorizationToken) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ECRAuthorizationTokenOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *ECRAuthorizationToken) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

func (o ECRAuthorizationTokenOutput) Spec() ECRAuthorizationTokenSpecPtrOutput {
	return o.ApplyT(func(v *ECRAuthorizationToken) ECRAuthorizationTokenSpecPtrOutput { return v.Spec }).(ECRAuthorizationTokenSpecPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ECRAuthorizationTokenInput)(nil)).Elem(), &ECRAuthorizationToken{})
	pulumi.RegisterOutputType(ECRAuthorizationTokenOutput{})
}
