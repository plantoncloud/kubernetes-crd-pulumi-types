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

type Telemetry struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Telemetry configuration for workloads. See more details at: https://istio.io/docs/reference/config/telemetry.html
	Spec   TelemetrySpecPtrOutput `pulumi:"spec"`
	Status pulumi.MapOutput       `pulumi:"status"`
}

// NewTelemetry registers a new resource with the given unique name, arguments, and options.
func NewTelemetry(ctx *pulumi.Context,
	name string, args *TelemetryArgs, opts ...pulumi.ResourceOption) (*Telemetry, error) {
	if args == nil {
		args = &TelemetryArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("telemetry.istio.io/v1alpha1")
	args.Kind = pulumi.StringPtr("Telemetry")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Telemetry
	err := ctx.RegisterResource("kubernetes:telemetry.istio.io/v1alpha1:Telemetry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTelemetry gets an existing Telemetry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTelemetry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TelemetryState, opts ...pulumi.ResourceOption) (*Telemetry, error) {
	var resource Telemetry
	err := ctx.ReadResource("kubernetes:telemetry.istio.io/v1alpha1:Telemetry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Telemetry resources.
type telemetryState struct {
}

type TelemetryState struct {
}

func (TelemetryState) ElementType() reflect.Type {
	return reflect.TypeOf((*telemetryState)(nil)).Elem()
}

type telemetryArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// Telemetry configuration for workloads. See more details at: https://istio.io/docs/reference/config/telemetry.html
	Spec   *TelemetrySpec         `pulumi:"spec"`
	Status map[string]interface{} `pulumi:"status"`
}

// The set of arguments for constructing a Telemetry resource.
type TelemetryArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// Telemetry configuration for workloads. See more details at: https://istio.io/docs/reference/config/telemetry.html
	Spec   TelemetrySpecPtrInput
	Status pulumi.MapInput
}

func (TelemetryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*telemetryArgs)(nil)).Elem()
}

type TelemetryInput interface {
	pulumi.Input

	ToTelemetryOutput() TelemetryOutput
	ToTelemetryOutputWithContext(ctx context.Context) TelemetryOutput
}

func (*Telemetry) ElementType() reflect.Type {
	return reflect.TypeOf((**Telemetry)(nil)).Elem()
}

func (i *Telemetry) ToTelemetryOutput() TelemetryOutput {
	return i.ToTelemetryOutputWithContext(context.Background())
}

func (i *Telemetry) ToTelemetryOutputWithContext(ctx context.Context) TelemetryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TelemetryOutput)
}

type TelemetryOutput struct{ *pulumi.OutputState }

func (TelemetryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Telemetry)(nil)).Elem()
}

func (o TelemetryOutput) ToTelemetryOutput() TelemetryOutput {
	return o
}

func (o TelemetryOutput) ToTelemetryOutputWithContext(ctx context.Context) TelemetryOutput {
	return o
}

func (o TelemetryOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Telemetry) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o TelemetryOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Telemetry) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o TelemetryOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *Telemetry) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Telemetry configuration for workloads. See more details at: https://istio.io/docs/reference/config/telemetry.html
func (o TelemetryOutput) Spec() TelemetrySpecPtrOutput {
	return o.ApplyT(func(v *Telemetry) TelemetrySpecPtrOutput { return v.Spec }).(TelemetrySpecPtrOutput)
}

func (o TelemetryOutput) Status() pulumi.MapOutput {
	return o.ApplyT(func(v *Telemetry) pulumi.MapOutput { return v.Status }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TelemetryInput)(nil)).Elem(), &Telemetry{})
	pulumi.RegisterOutputType(TelemetryOutput{})
}
