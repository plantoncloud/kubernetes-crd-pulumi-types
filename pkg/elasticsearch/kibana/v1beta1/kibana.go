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

// Kibana represents a Kibana resource in a Kubernetes cluster.
type Kibana struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// KibanaSpec holds the specification of a Kibana instance.
	Spec KibanaSpecPtrOutput `pulumi:"spec"`
	// KibanaStatus defines the observed state of Kibana
	Status KibanaStatusPtrOutput `pulumi:"status"`
}

// NewKibana registers a new resource with the given unique name, arguments, and options.
func NewKibana(ctx *pulumi.Context,
	name string, args *KibanaArgs, opts ...pulumi.ResourceOption) (*Kibana, error) {
	if args == nil {
		args = &KibanaArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("kibana.k8s.elastic.co/v1beta1")
	args.Kind = pulumi.StringPtr("Kibana")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Kibana
	err := ctx.RegisterResource("kubernetes:kibana.k8s.elastic.co/v1beta1:Kibana", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKibana gets an existing Kibana resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKibana(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KibanaState, opts ...pulumi.ResourceOption) (*Kibana, error) {
	var resource Kibana
	err := ctx.ReadResource("kubernetes:kibana.k8s.elastic.co/v1beta1:Kibana", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Kibana resources.
type kibanaState struct {
}

type KibanaState struct {
}

func (KibanaState) ElementType() reflect.Type {
	return reflect.TypeOf((*kibanaState)(nil)).Elem()
}

type kibanaArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// KibanaSpec holds the specification of a Kibana instance.
	Spec *KibanaSpec `pulumi:"spec"`
	// KibanaStatus defines the observed state of Kibana
	Status *KibanaStatus `pulumi:"status"`
}

// The set of arguments for constructing a Kibana resource.
type KibanaArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// KibanaSpec holds the specification of a Kibana instance.
	Spec KibanaSpecPtrInput
	// KibanaStatus defines the observed state of Kibana
	Status KibanaStatusPtrInput
}

func (KibanaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kibanaArgs)(nil)).Elem()
}

type KibanaInput interface {
	pulumi.Input

	ToKibanaOutput() KibanaOutput
	ToKibanaOutputWithContext(ctx context.Context) KibanaOutput
}

func (*Kibana) ElementType() reflect.Type {
	return reflect.TypeOf((**Kibana)(nil)).Elem()
}

func (i *Kibana) ToKibanaOutput() KibanaOutput {
	return i.ToKibanaOutputWithContext(context.Background())
}

func (i *Kibana) ToKibanaOutputWithContext(ctx context.Context) KibanaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KibanaOutput)
}

type KibanaOutput struct{ *pulumi.OutputState }

func (KibanaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Kibana)(nil)).Elem()
}

func (o KibanaOutput) ToKibanaOutput() KibanaOutput {
	return o
}

func (o KibanaOutput) ToKibanaOutputWithContext(ctx context.Context) KibanaOutput {
	return o
}

func (o KibanaOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Kibana) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o KibanaOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Kibana) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o KibanaOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *Kibana) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// KibanaSpec holds the specification of a Kibana instance.
func (o KibanaOutput) Spec() KibanaSpecPtrOutput {
	return o.ApplyT(func(v *Kibana) KibanaSpecPtrOutput { return v.Spec }).(KibanaSpecPtrOutput)
}

// KibanaStatus defines the observed state of Kibana
func (o KibanaOutput) Status() KibanaStatusPtrOutput {
	return o.ApplyT(func(v *Kibana) KibanaStatusPtrOutput { return v.Status }).(KibanaStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KibanaInput)(nil)).Elem(), &Kibana{})
	pulumi.RegisterOutputType(KibanaOutput{})
}
