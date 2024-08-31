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

// Elasticsearch represents an Elasticsearch resource in a Kubernetes cluster.
type Elasticsearch struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// ElasticsearchSpec holds the specification of an Elasticsearch cluster.
	Spec ElasticsearchSpecPtrOutput `pulumi:"spec"`
	// ElasticsearchStatus defines the observed state of Elasticsearch
	Status ElasticsearchStatusPtrOutput `pulumi:"status"`
}

// NewElasticsearch registers a new resource with the given unique name, arguments, and options.
func NewElasticsearch(ctx *pulumi.Context,
	name string, args *ElasticsearchArgs, opts ...pulumi.ResourceOption) (*Elasticsearch, error) {
	if args == nil {
		args = &ElasticsearchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("elasticsearch.k8s.elastic.co/v1beta1")
	args.Kind = pulumi.StringPtr("Elasticsearch")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Elasticsearch
	err := ctx.RegisterResource("kubernetes:elasticsearch.k8s.elastic.co/v1beta1:Elasticsearch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElasticsearch gets an existing Elasticsearch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticsearch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticsearchState, opts ...pulumi.ResourceOption) (*Elasticsearch, error) {
	var resource Elasticsearch
	err := ctx.ReadResource("kubernetes:elasticsearch.k8s.elastic.co/v1beta1:Elasticsearch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Elasticsearch resources.
type elasticsearchState struct {
}

type ElasticsearchState struct {
}

func (ElasticsearchState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticsearchState)(nil)).Elem()
}

type elasticsearchArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// ElasticsearchSpec holds the specification of an Elasticsearch cluster.
	Spec *ElasticsearchSpec `pulumi:"spec"`
	// ElasticsearchStatus defines the observed state of Elasticsearch
	Status *ElasticsearchStatus `pulumi:"status"`
}

// The set of arguments for constructing a Elasticsearch resource.
type ElasticsearchArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// ElasticsearchSpec holds the specification of an Elasticsearch cluster.
	Spec ElasticsearchSpecPtrInput
	// ElasticsearchStatus defines the observed state of Elasticsearch
	Status ElasticsearchStatusPtrInput
}

func (ElasticsearchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticsearchArgs)(nil)).Elem()
}

type ElasticsearchInput interface {
	pulumi.Input

	ToElasticsearchOutput() ElasticsearchOutput
	ToElasticsearchOutputWithContext(ctx context.Context) ElasticsearchOutput
}

func (*Elasticsearch) ElementType() reflect.Type {
	return reflect.TypeOf((**Elasticsearch)(nil)).Elem()
}

func (i *Elasticsearch) ToElasticsearchOutput() ElasticsearchOutput {
	return i.ToElasticsearchOutputWithContext(context.Background())
}

func (i *Elasticsearch) ToElasticsearchOutputWithContext(ctx context.Context) ElasticsearchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticsearchOutput)
}

type ElasticsearchOutput struct{ *pulumi.OutputState }

func (ElasticsearchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Elasticsearch)(nil)).Elem()
}

func (o ElasticsearchOutput) ToElasticsearchOutput() ElasticsearchOutput {
	return o
}

func (o ElasticsearchOutput) ToElasticsearchOutputWithContext(ctx context.Context) ElasticsearchOutput {
	return o
}

func (o ElasticsearchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Elasticsearch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o ElasticsearchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Elasticsearch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ElasticsearchOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *Elasticsearch) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// ElasticsearchSpec holds the specification of an Elasticsearch cluster.
func (o ElasticsearchOutput) Spec() ElasticsearchSpecPtrOutput {
	return o.ApplyT(func(v *Elasticsearch) ElasticsearchSpecPtrOutput { return v.Spec }).(ElasticsearchSpecPtrOutput)
}

// ElasticsearchStatus defines the observed state of Elasticsearch
func (o ElasticsearchOutput) Status() ElasticsearchStatusPtrOutput {
	return o.ApplyT(func(v *Elasticsearch) ElasticsearchStatusPtrOutput { return v.Status }).(ElasticsearchStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticsearchInput)(nil)).Elem(), &Elasticsearch{})
	pulumi.RegisterOutputType(ElasticsearchOutput{})
}
