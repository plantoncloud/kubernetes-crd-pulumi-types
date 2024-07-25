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

type Postgresql struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringOutput        `pulumi:"apiVersion"`
	Kind       pulumi.StringOutput        `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	Spec       PostgresqlSpecOutput       `pulumi:"spec"`
	Status     pulumi.StringMapOutput     `pulumi:"status"`
}

// NewPostgresql registers a new resource with the given unique name, arguments, and options.
func NewPostgresql(ctx *pulumi.Context,
	name string, args *PostgresqlArgs, opts ...pulumi.ResourceOption) (*Postgresql, error) {
	if args == nil {
		args = &PostgresqlArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("acid.zalan.do/v1")
	args.Kind = pulumi.StringPtr("postgresql")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Postgresql
	err := ctx.RegisterResource("kubernetes:acid.zalan.do/v1:postgresql", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPostgresql gets an existing Postgresql resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPostgresql(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PostgresqlState, opts ...pulumi.ResourceOption) (*Postgresql, error) {
	var resource Postgresql
	err := ctx.ReadResource("kubernetes:acid.zalan.do/v1:postgresql", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Postgresql resources.
type postgresqlState struct {
}

type PostgresqlState struct {
}

func (PostgresqlState) ElementType() reflect.Type {
	return reflect.TypeOf((*postgresqlState)(nil)).Elem()
}

type postgresqlArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	Spec       *PostgresqlSpec    `pulumi:"spec"`
	Status     map[string]string  `pulumi:"status"`
}

// The set of arguments for constructing a Postgresql resource.
type PostgresqlArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	Spec       PostgresqlSpecPtrInput
	Status     pulumi.StringMapInput
}

func (PostgresqlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*postgresqlArgs)(nil)).Elem()
}

type PostgresqlInput interface {
	pulumi.Input

	ToPostgresqlOutput() PostgresqlOutput
	ToPostgresqlOutputWithContext(ctx context.Context) PostgresqlOutput
}

func (*Postgresql) ElementType() reflect.Type {
	return reflect.TypeOf((**Postgresql)(nil)).Elem()
}

func (i *Postgresql) ToPostgresqlOutput() PostgresqlOutput {
	return i.ToPostgresqlOutputWithContext(context.Background())
}

func (i *Postgresql) ToPostgresqlOutputWithContext(ctx context.Context) PostgresqlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresqlOutput)
}

type PostgresqlOutput struct{ *pulumi.OutputState }

func (PostgresqlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Postgresql)(nil)).Elem()
}

func (o PostgresqlOutput) ToPostgresqlOutput() PostgresqlOutput {
	return o
}

func (o PostgresqlOutput) ToPostgresqlOutputWithContext(ctx context.Context) PostgresqlOutput {
	return o
}

func (o PostgresqlOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Postgresql) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

func (o PostgresqlOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Postgresql) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o PostgresqlOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *Postgresql) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

func (o PostgresqlOutput) Spec() PostgresqlSpecOutput {
	return o.ApplyT(func(v *Postgresql) PostgresqlSpecOutput { return v.Spec }).(PostgresqlSpecOutput)
}

func (o PostgresqlOutput) Status() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Postgresql) pulumi.StringMapOutput { return v.Status }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PostgresqlInput)(nil)).Elem(), &Postgresql{})
	pulumi.RegisterOutputType(PostgresqlOutput{})
}
