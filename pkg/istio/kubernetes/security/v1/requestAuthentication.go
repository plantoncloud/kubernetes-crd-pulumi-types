// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RequestAuthentication struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Request authentication configuration for workloads. See more details at: https://istio.io/docs/reference/config/security/request_authentication.html
	Spec   RequestAuthenticationSpecPtrOutput `pulumi:"spec"`
	Status pulumi.MapOutput                   `pulumi:"status"`
}

// NewRequestAuthentication registers a new resource with the given unique name, arguments, and options.
func NewRequestAuthentication(ctx *pulumi.Context,
	name string, args *RequestAuthenticationArgs, opts ...pulumi.ResourceOption) (*RequestAuthentication, error) {
	if args == nil {
		args = &RequestAuthenticationArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("security.istio.io/v1")
	args.Kind = pulumi.StringPtr("RequestAuthentication")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RequestAuthentication
	err := ctx.RegisterResource("kubernetes:security.istio.io/v1:RequestAuthentication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRequestAuthentication gets an existing RequestAuthentication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRequestAuthentication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RequestAuthenticationState, opts ...pulumi.ResourceOption) (*RequestAuthentication, error) {
	var resource RequestAuthentication
	err := ctx.ReadResource("kubernetes:security.istio.io/v1:RequestAuthentication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RequestAuthentication resources.
type requestAuthenticationState struct {
}

type RequestAuthenticationState struct {
}

func (RequestAuthenticationState) ElementType() reflect.Type {
	return reflect.TypeOf((*requestAuthenticationState)(nil)).Elem()
}

type requestAuthenticationArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// Request authentication configuration for workloads. See more details at: https://istio.io/docs/reference/config/security/request_authentication.html
	Spec   *RequestAuthenticationSpec `pulumi:"spec"`
	Status map[string]interface{}     `pulumi:"status"`
}

// The set of arguments for constructing a RequestAuthentication resource.
type RequestAuthenticationArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// Request authentication configuration for workloads. See more details at: https://istio.io/docs/reference/config/security/request_authentication.html
	Spec   RequestAuthenticationSpecPtrInput
	Status pulumi.MapInput
}

func (RequestAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*requestAuthenticationArgs)(nil)).Elem()
}

type RequestAuthenticationInput interface {
	pulumi.Input

	ToRequestAuthenticationOutput() RequestAuthenticationOutput
	ToRequestAuthenticationOutputWithContext(ctx context.Context) RequestAuthenticationOutput
}

func (*RequestAuthentication) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestAuthentication)(nil)).Elem()
}

func (i *RequestAuthentication) ToRequestAuthenticationOutput() RequestAuthenticationOutput {
	return i.ToRequestAuthenticationOutputWithContext(context.Background())
}

func (i *RequestAuthentication) ToRequestAuthenticationOutputWithContext(ctx context.Context) RequestAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestAuthenticationOutput)
}

type RequestAuthenticationOutput struct{ *pulumi.OutputState }

func (RequestAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestAuthentication)(nil)).Elem()
}

func (o RequestAuthenticationOutput) ToRequestAuthenticationOutput() RequestAuthenticationOutput {
	return o
}

func (o RequestAuthenticationOutput) ToRequestAuthenticationOutputWithContext(ctx context.Context) RequestAuthenticationOutput {
	return o
}

func (o RequestAuthenticationOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestAuthentication) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o RequestAuthenticationOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestAuthentication) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o RequestAuthenticationOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *RequestAuthentication) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Request authentication configuration for workloads. See more details at: https://istio.io/docs/reference/config/security/request_authentication.html
func (o RequestAuthenticationOutput) Spec() RequestAuthenticationSpecPtrOutput {
	return o.ApplyT(func(v *RequestAuthentication) RequestAuthenticationSpecPtrOutput { return v.Spec }).(RequestAuthenticationSpecPtrOutput)
}

func (o RequestAuthenticationOutput) Status() pulumi.MapOutput {
	return o.ApplyT(func(v *RequestAuthentication) pulumi.MapOutput { return v.Status }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RequestAuthenticationInput)(nil)).Elem(), &RequestAuthentication{})
	pulumi.RegisterOutputType(RequestAuthenticationOutput{})
}
