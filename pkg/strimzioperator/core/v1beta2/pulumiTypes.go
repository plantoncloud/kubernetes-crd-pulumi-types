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

var _ = utilities.GetEnvOrDefault

type StrimziPodSetType struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// The specification of the StrimziPodSet.
	Spec *StrimziPodSetSpec `pulumi:"spec"`
	// The status of the StrimziPodSet.
	Status *StrimziPodSetStatus `pulumi:"status"`
}

type StrimziPodSetMetadata struct {
}

// The specification of the StrimziPodSet.
type StrimziPodSetSpec struct {
	// The Pods managed by this StrimziPodSet.
	Pods []map[string]interface{} `pulumi:"pods"`
	// Selector is a label query which matches all the pods managed by this `StrimziPodSet`. Only `matchLabels` is supported. If `matchExpressions` is set, it will be ignored.
	Selector StrimziPodSetSpecSelector `pulumi:"selector"`
}

// StrimziPodSetSpecInput is an input type that accepts StrimziPodSetSpecArgs and StrimziPodSetSpecOutput values.
// You can construct a concrete instance of `StrimziPodSetSpecInput` via:
//
//	StrimziPodSetSpecArgs{...}
type StrimziPodSetSpecInput interface {
	pulumi.Input

	ToStrimziPodSetSpecOutput() StrimziPodSetSpecOutput
	ToStrimziPodSetSpecOutputWithContext(context.Context) StrimziPodSetSpecOutput
}

// The specification of the StrimziPodSet.
type StrimziPodSetSpecArgs struct {
	// The Pods managed by this StrimziPodSet.
	Pods pulumi.MapArrayInput `pulumi:"pods"`
	// Selector is a label query which matches all the pods managed by this `StrimziPodSet`. Only `matchLabels` is supported. If `matchExpressions` is set, it will be ignored.
	Selector StrimziPodSetSpecSelectorInput `pulumi:"selector"`
}

func (StrimziPodSetSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StrimziPodSetSpec)(nil)).Elem()
}

func (i StrimziPodSetSpecArgs) ToStrimziPodSetSpecOutput() StrimziPodSetSpecOutput {
	return i.ToStrimziPodSetSpecOutputWithContext(context.Background())
}

func (i StrimziPodSetSpecArgs) ToStrimziPodSetSpecOutputWithContext(ctx context.Context) StrimziPodSetSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrimziPodSetSpecOutput)
}

func (i StrimziPodSetSpecArgs) ToStrimziPodSetSpecPtrOutput() StrimziPodSetSpecPtrOutput {
	return i.ToStrimziPodSetSpecPtrOutputWithContext(context.Background())
}

func (i StrimziPodSetSpecArgs) ToStrimziPodSetSpecPtrOutputWithContext(ctx context.Context) StrimziPodSetSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrimziPodSetSpecOutput).ToStrimziPodSetSpecPtrOutputWithContext(ctx)
}

// StrimziPodSetSpecPtrInput is an input type that accepts StrimziPodSetSpecArgs, StrimziPodSetSpecPtr and StrimziPodSetSpecPtrOutput values.
// You can construct a concrete instance of `StrimziPodSetSpecPtrInput` via:
//
//	        StrimziPodSetSpecArgs{...}
//
//	or:
//
//	        nil
type StrimziPodSetSpecPtrInput interface {
	pulumi.Input

	ToStrimziPodSetSpecPtrOutput() StrimziPodSetSpecPtrOutput
	ToStrimziPodSetSpecPtrOutputWithContext(context.Context) StrimziPodSetSpecPtrOutput
}

type strimziPodSetSpecPtrType StrimziPodSetSpecArgs

func StrimziPodSetSpecPtr(v *StrimziPodSetSpecArgs) StrimziPodSetSpecPtrInput {
	return (*strimziPodSetSpecPtrType)(v)
}

func (*strimziPodSetSpecPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StrimziPodSetSpec)(nil)).Elem()
}

func (i *strimziPodSetSpecPtrType) ToStrimziPodSetSpecPtrOutput() StrimziPodSetSpecPtrOutput {
	return i.ToStrimziPodSetSpecPtrOutputWithContext(context.Background())
}

func (i *strimziPodSetSpecPtrType) ToStrimziPodSetSpecPtrOutputWithContext(ctx context.Context) StrimziPodSetSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrimziPodSetSpecPtrOutput)
}

// The specification of the StrimziPodSet.
type StrimziPodSetSpecOutput struct{ *pulumi.OutputState }

func (StrimziPodSetSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StrimziPodSetSpec)(nil)).Elem()
}

func (o StrimziPodSetSpecOutput) ToStrimziPodSetSpecOutput() StrimziPodSetSpecOutput {
	return o
}

func (o StrimziPodSetSpecOutput) ToStrimziPodSetSpecOutputWithContext(ctx context.Context) StrimziPodSetSpecOutput {
	return o
}

func (o StrimziPodSetSpecOutput) ToStrimziPodSetSpecPtrOutput() StrimziPodSetSpecPtrOutput {
	return o.ToStrimziPodSetSpecPtrOutputWithContext(context.Background())
}

func (o StrimziPodSetSpecOutput) ToStrimziPodSetSpecPtrOutputWithContext(ctx context.Context) StrimziPodSetSpecPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StrimziPodSetSpec) *StrimziPodSetSpec {
		return &v
	}).(StrimziPodSetSpecPtrOutput)
}

// The Pods managed by this StrimziPodSet.
func (o StrimziPodSetSpecOutput) Pods() pulumi.MapArrayOutput {
	return o.ApplyT(func(v StrimziPodSetSpec) []map[string]interface{} { return v.Pods }).(pulumi.MapArrayOutput)
}

// Selector is a label query which matches all the pods managed by this `StrimziPodSet`. Only `matchLabels` is supported. If `matchExpressions` is set, it will be ignored.
func (o StrimziPodSetSpecOutput) Selector() StrimziPodSetSpecSelectorOutput {
	return o.ApplyT(func(v StrimziPodSetSpec) StrimziPodSetSpecSelector { return v.Selector }).(StrimziPodSetSpecSelectorOutput)
}

type StrimziPodSetSpecPtrOutput struct{ *pulumi.OutputState }

func (StrimziPodSetSpecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StrimziPodSetSpec)(nil)).Elem()
}

func (o StrimziPodSetSpecPtrOutput) ToStrimziPodSetSpecPtrOutput() StrimziPodSetSpecPtrOutput {
	return o
}

func (o StrimziPodSetSpecPtrOutput) ToStrimziPodSetSpecPtrOutputWithContext(ctx context.Context) StrimziPodSetSpecPtrOutput {
	return o
}

func (o StrimziPodSetSpecPtrOutput) Elem() StrimziPodSetSpecOutput {
	return o.ApplyT(func(v *StrimziPodSetSpec) StrimziPodSetSpec {
		if v != nil {
			return *v
		}
		var ret StrimziPodSetSpec
		return ret
	}).(StrimziPodSetSpecOutput)
}

// The Pods managed by this StrimziPodSet.
func (o StrimziPodSetSpecPtrOutput) Pods() pulumi.MapArrayOutput {
	return o.ApplyT(func(v *StrimziPodSetSpec) []map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.Pods
	}).(pulumi.MapArrayOutput)
}

// Selector is a label query which matches all the pods managed by this `StrimziPodSet`. Only `matchLabels` is supported. If `matchExpressions` is set, it will be ignored.
func (o StrimziPodSetSpecPtrOutput) Selector() StrimziPodSetSpecSelectorPtrOutput {
	return o.ApplyT(func(v *StrimziPodSetSpec) *StrimziPodSetSpecSelector {
		if v == nil {
			return nil
		}
		return &v.Selector
	}).(StrimziPodSetSpecSelectorPtrOutput)
}

// Selector is a label query which matches all the pods managed by this `StrimziPodSet`. Only `matchLabels` is supported. If `matchExpressions` is set, it will be ignored.
type StrimziPodSetSpecSelector struct {
	MatchExpressions []StrimziPodSetSpecSelectorMatchExpressions `pulumi:"matchExpressions"`
	MatchLabels      map[string]string                           `pulumi:"matchLabels"`
}

// StrimziPodSetSpecSelectorInput is an input type that accepts StrimziPodSetSpecSelectorArgs and StrimziPodSetSpecSelectorOutput values.
// You can construct a concrete instance of `StrimziPodSetSpecSelectorInput` via:
//
//	StrimziPodSetSpecSelectorArgs{...}
type StrimziPodSetSpecSelectorInput interface {
	pulumi.Input

	ToStrimziPodSetSpecSelectorOutput() StrimziPodSetSpecSelectorOutput
	ToStrimziPodSetSpecSelectorOutputWithContext(context.Context) StrimziPodSetSpecSelectorOutput
}

// Selector is a label query which matches all the pods managed by this `StrimziPodSet`. Only `matchLabels` is supported. If `matchExpressions` is set, it will be ignored.
type StrimziPodSetSpecSelectorArgs struct {
	MatchExpressions StrimziPodSetSpecSelectorMatchExpressionsArrayInput `pulumi:"matchExpressions"`
	MatchLabels      pulumi.StringMapInput                               `pulumi:"matchLabels"`
}

func (StrimziPodSetSpecSelectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StrimziPodSetSpecSelector)(nil)).Elem()
}

func (i StrimziPodSetSpecSelectorArgs) ToStrimziPodSetSpecSelectorOutput() StrimziPodSetSpecSelectorOutput {
	return i.ToStrimziPodSetSpecSelectorOutputWithContext(context.Background())
}

func (i StrimziPodSetSpecSelectorArgs) ToStrimziPodSetSpecSelectorOutputWithContext(ctx context.Context) StrimziPodSetSpecSelectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrimziPodSetSpecSelectorOutput)
}

func (i StrimziPodSetSpecSelectorArgs) ToStrimziPodSetSpecSelectorPtrOutput() StrimziPodSetSpecSelectorPtrOutput {
	return i.ToStrimziPodSetSpecSelectorPtrOutputWithContext(context.Background())
}

func (i StrimziPodSetSpecSelectorArgs) ToStrimziPodSetSpecSelectorPtrOutputWithContext(ctx context.Context) StrimziPodSetSpecSelectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrimziPodSetSpecSelectorOutput).ToStrimziPodSetSpecSelectorPtrOutputWithContext(ctx)
}

// StrimziPodSetSpecSelectorPtrInput is an input type that accepts StrimziPodSetSpecSelectorArgs, StrimziPodSetSpecSelectorPtr and StrimziPodSetSpecSelectorPtrOutput values.
// You can construct a concrete instance of `StrimziPodSetSpecSelectorPtrInput` via:
//
//	        StrimziPodSetSpecSelectorArgs{...}
//
//	or:
//
//	        nil
type StrimziPodSetSpecSelectorPtrInput interface {
	pulumi.Input

	ToStrimziPodSetSpecSelectorPtrOutput() StrimziPodSetSpecSelectorPtrOutput
	ToStrimziPodSetSpecSelectorPtrOutputWithContext(context.Context) StrimziPodSetSpecSelectorPtrOutput
}

type strimziPodSetSpecSelectorPtrType StrimziPodSetSpecSelectorArgs

func StrimziPodSetSpecSelectorPtr(v *StrimziPodSetSpecSelectorArgs) StrimziPodSetSpecSelectorPtrInput {
	return (*strimziPodSetSpecSelectorPtrType)(v)
}

func (*strimziPodSetSpecSelectorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StrimziPodSetSpecSelector)(nil)).Elem()
}

func (i *strimziPodSetSpecSelectorPtrType) ToStrimziPodSetSpecSelectorPtrOutput() StrimziPodSetSpecSelectorPtrOutput {
	return i.ToStrimziPodSetSpecSelectorPtrOutputWithContext(context.Background())
}

func (i *strimziPodSetSpecSelectorPtrType) ToStrimziPodSetSpecSelectorPtrOutputWithContext(ctx context.Context) StrimziPodSetSpecSelectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrimziPodSetSpecSelectorPtrOutput)
}

// Selector is a label query which matches all the pods managed by this `StrimziPodSet`. Only `matchLabels` is supported. If `matchExpressions` is set, it will be ignored.
type StrimziPodSetSpecSelectorOutput struct{ *pulumi.OutputState }

func (StrimziPodSetSpecSelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StrimziPodSetSpecSelector)(nil)).Elem()
}

func (o StrimziPodSetSpecSelectorOutput) ToStrimziPodSetSpecSelectorOutput() StrimziPodSetSpecSelectorOutput {
	return o
}

func (o StrimziPodSetSpecSelectorOutput) ToStrimziPodSetSpecSelectorOutputWithContext(ctx context.Context) StrimziPodSetSpecSelectorOutput {
	return o
}

func (o StrimziPodSetSpecSelectorOutput) ToStrimziPodSetSpecSelectorPtrOutput() StrimziPodSetSpecSelectorPtrOutput {
	return o.ToStrimziPodSetSpecSelectorPtrOutputWithContext(context.Background())
}

func (o StrimziPodSetSpecSelectorOutput) ToStrimziPodSetSpecSelectorPtrOutputWithContext(ctx context.Context) StrimziPodSetSpecSelectorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StrimziPodSetSpecSelector) *StrimziPodSetSpecSelector {
		return &v
	}).(StrimziPodSetSpecSelectorPtrOutput)
}

func (o StrimziPodSetSpecSelectorOutput) MatchExpressions() StrimziPodSetSpecSelectorMatchExpressionsArrayOutput {
	return o.ApplyT(func(v StrimziPodSetSpecSelector) []StrimziPodSetSpecSelectorMatchExpressions {
		return v.MatchExpressions
	}).(StrimziPodSetSpecSelectorMatchExpressionsArrayOutput)
}

func (o StrimziPodSetSpecSelectorOutput) MatchLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v StrimziPodSetSpecSelector) map[string]string { return v.MatchLabels }).(pulumi.StringMapOutput)
}

type StrimziPodSetSpecSelectorPtrOutput struct{ *pulumi.OutputState }

func (StrimziPodSetSpecSelectorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StrimziPodSetSpecSelector)(nil)).Elem()
}

func (o StrimziPodSetSpecSelectorPtrOutput) ToStrimziPodSetSpecSelectorPtrOutput() StrimziPodSetSpecSelectorPtrOutput {
	return o
}

func (o StrimziPodSetSpecSelectorPtrOutput) ToStrimziPodSetSpecSelectorPtrOutputWithContext(ctx context.Context) StrimziPodSetSpecSelectorPtrOutput {
	return o
}

func (o StrimziPodSetSpecSelectorPtrOutput) Elem() StrimziPodSetSpecSelectorOutput {
	return o.ApplyT(func(v *StrimziPodSetSpecSelector) StrimziPodSetSpecSelector {
		if v != nil {
			return *v
		}
		var ret StrimziPodSetSpecSelector
		return ret
	}).(StrimziPodSetSpecSelectorOutput)
}

func (o StrimziPodSetSpecSelectorPtrOutput) MatchExpressions() StrimziPodSetSpecSelectorMatchExpressionsArrayOutput {
	return o.ApplyT(func(v *StrimziPodSetSpecSelector) []StrimziPodSetSpecSelectorMatchExpressions {
		if v == nil {
			return nil
		}
		return v.MatchExpressions
	}).(StrimziPodSetSpecSelectorMatchExpressionsArrayOutput)
}

func (o StrimziPodSetSpecSelectorPtrOutput) MatchLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StrimziPodSetSpecSelector) map[string]string {
		if v == nil {
			return nil
		}
		return v.MatchLabels
	}).(pulumi.StringMapOutput)
}

type StrimziPodSetSpecSelectorMatchExpressions struct {
	Key      *string  `pulumi:"key"`
	Operator *string  `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}

// StrimziPodSetSpecSelectorMatchExpressionsInput is an input type that accepts StrimziPodSetSpecSelectorMatchExpressionsArgs and StrimziPodSetSpecSelectorMatchExpressionsOutput values.
// You can construct a concrete instance of `StrimziPodSetSpecSelectorMatchExpressionsInput` via:
//
//	StrimziPodSetSpecSelectorMatchExpressionsArgs{...}
type StrimziPodSetSpecSelectorMatchExpressionsInput interface {
	pulumi.Input

	ToStrimziPodSetSpecSelectorMatchExpressionsOutput() StrimziPodSetSpecSelectorMatchExpressionsOutput
	ToStrimziPodSetSpecSelectorMatchExpressionsOutputWithContext(context.Context) StrimziPodSetSpecSelectorMatchExpressionsOutput
}

type StrimziPodSetSpecSelectorMatchExpressionsArgs struct {
	Key      pulumi.StringPtrInput   `pulumi:"key"`
	Operator pulumi.StringPtrInput   `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (StrimziPodSetSpecSelectorMatchExpressionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StrimziPodSetSpecSelectorMatchExpressions)(nil)).Elem()
}

func (i StrimziPodSetSpecSelectorMatchExpressionsArgs) ToStrimziPodSetSpecSelectorMatchExpressionsOutput() StrimziPodSetSpecSelectorMatchExpressionsOutput {
	return i.ToStrimziPodSetSpecSelectorMatchExpressionsOutputWithContext(context.Background())
}

func (i StrimziPodSetSpecSelectorMatchExpressionsArgs) ToStrimziPodSetSpecSelectorMatchExpressionsOutputWithContext(ctx context.Context) StrimziPodSetSpecSelectorMatchExpressionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrimziPodSetSpecSelectorMatchExpressionsOutput)
}

// StrimziPodSetSpecSelectorMatchExpressionsArrayInput is an input type that accepts StrimziPodSetSpecSelectorMatchExpressionsArray and StrimziPodSetSpecSelectorMatchExpressionsArrayOutput values.
// You can construct a concrete instance of `StrimziPodSetSpecSelectorMatchExpressionsArrayInput` via:
//
//	StrimziPodSetSpecSelectorMatchExpressionsArray{ StrimziPodSetSpecSelectorMatchExpressionsArgs{...} }
type StrimziPodSetSpecSelectorMatchExpressionsArrayInput interface {
	pulumi.Input

	ToStrimziPodSetSpecSelectorMatchExpressionsArrayOutput() StrimziPodSetSpecSelectorMatchExpressionsArrayOutput
	ToStrimziPodSetSpecSelectorMatchExpressionsArrayOutputWithContext(context.Context) StrimziPodSetSpecSelectorMatchExpressionsArrayOutput
}

type StrimziPodSetSpecSelectorMatchExpressionsArray []StrimziPodSetSpecSelectorMatchExpressionsInput

func (StrimziPodSetSpecSelectorMatchExpressionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StrimziPodSetSpecSelectorMatchExpressions)(nil)).Elem()
}

func (i StrimziPodSetSpecSelectorMatchExpressionsArray) ToStrimziPodSetSpecSelectorMatchExpressionsArrayOutput() StrimziPodSetSpecSelectorMatchExpressionsArrayOutput {
	return i.ToStrimziPodSetSpecSelectorMatchExpressionsArrayOutputWithContext(context.Background())
}

func (i StrimziPodSetSpecSelectorMatchExpressionsArray) ToStrimziPodSetSpecSelectorMatchExpressionsArrayOutputWithContext(ctx context.Context) StrimziPodSetSpecSelectorMatchExpressionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrimziPodSetSpecSelectorMatchExpressionsArrayOutput)
}

type StrimziPodSetSpecSelectorMatchExpressionsOutput struct{ *pulumi.OutputState }

func (StrimziPodSetSpecSelectorMatchExpressionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StrimziPodSetSpecSelectorMatchExpressions)(nil)).Elem()
}

func (o StrimziPodSetSpecSelectorMatchExpressionsOutput) ToStrimziPodSetSpecSelectorMatchExpressionsOutput() StrimziPodSetSpecSelectorMatchExpressionsOutput {
	return o
}

func (o StrimziPodSetSpecSelectorMatchExpressionsOutput) ToStrimziPodSetSpecSelectorMatchExpressionsOutputWithContext(ctx context.Context) StrimziPodSetSpecSelectorMatchExpressionsOutput {
	return o
}

func (o StrimziPodSetSpecSelectorMatchExpressionsOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StrimziPodSetSpecSelectorMatchExpressions) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o StrimziPodSetSpecSelectorMatchExpressionsOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StrimziPodSetSpecSelectorMatchExpressions) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o StrimziPodSetSpecSelectorMatchExpressionsOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StrimziPodSetSpecSelectorMatchExpressions) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type StrimziPodSetSpecSelectorMatchExpressionsArrayOutput struct{ *pulumi.OutputState }

func (StrimziPodSetSpecSelectorMatchExpressionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StrimziPodSetSpecSelectorMatchExpressions)(nil)).Elem()
}

func (o StrimziPodSetSpecSelectorMatchExpressionsArrayOutput) ToStrimziPodSetSpecSelectorMatchExpressionsArrayOutput() StrimziPodSetSpecSelectorMatchExpressionsArrayOutput {
	return o
}

func (o StrimziPodSetSpecSelectorMatchExpressionsArrayOutput) ToStrimziPodSetSpecSelectorMatchExpressionsArrayOutputWithContext(ctx context.Context) StrimziPodSetSpecSelectorMatchExpressionsArrayOutput {
	return o
}

func (o StrimziPodSetSpecSelectorMatchExpressionsArrayOutput) Index(i pulumi.IntInput) StrimziPodSetSpecSelectorMatchExpressionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StrimziPodSetSpecSelectorMatchExpressions {
		return vs[0].([]StrimziPodSetSpecSelectorMatchExpressions)[vs[1].(int)]
	}).(StrimziPodSetSpecSelectorMatchExpressionsOutput)
}

type StrimziPodSetSpecSelectorMatchLabels struct {
}

// The status of the StrimziPodSet.
type StrimziPodSetStatus struct {
	// List of status conditions.
	Conditions []StrimziPodSetStatusConditions `pulumi:"conditions"`
	// Number of pods managed by this `StrimziPodSet` resource that have the current revision.
	CurrentPods *int `pulumi:"currentPods"`
	// The generation of the CRD that was last reconciled by the operator.
	ObservedGeneration *int `pulumi:"observedGeneration"`
	// Number of pods managed by this `StrimziPodSet` resource.
	Pods *int `pulumi:"pods"`
	// Number of pods managed by this `StrimziPodSet` resource that are ready.
	ReadyPods *int `pulumi:"readyPods"`
}

// StrimziPodSetStatusInput is an input type that accepts StrimziPodSetStatusArgs and StrimziPodSetStatusOutput values.
// You can construct a concrete instance of `StrimziPodSetStatusInput` via:
//
//	StrimziPodSetStatusArgs{...}
type StrimziPodSetStatusInput interface {
	pulumi.Input

	ToStrimziPodSetStatusOutput() StrimziPodSetStatusOutput
	ToStrimziPodSetStatusOutputWithContext(context.Context) StrimziPodSetStatusOutput
}

// The status of the StrimziPodSet.
type StrimziPodSetStatusArgs struct {
	// List of status conditions.
	Conditions StrimziPodSetStatusConditionsArrayInput `pulumi:"conditions"`
	// Number of pods managed by this `StrimziPodSet` resource that have the current revision.
	CurrentPods pulumi.IntPtrInput `pulumi:"currentPods"`
	// The generation of the CRD that was last reconciled by the operator.
	ObservedGeneration pulumi.IntPtrInput `pulumi:"observedGeneration"`
	// Number of pods managed by this `StrimziPodSet` resource.
	Pods pulumi.IntPtrInput `pulumi:"pods"`
	// Number of pods managed by this `StrimziPodSet` resource that are ready.
	ReadyPods pulumi.IntPtrInput `pulumi:"readyPods"`
}

func (StrimziPodSetStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StrimziPodSetStatus)(nil)).Elem()
}

func (i StrimziPodSetStatusArgs) ToStrimziPodSetStatusOutput() StrimziPodSetStatusOutput {
	return i.ToStrimziPodSetStatusOutputWithContext(context.Background())
}

func (i StrimziPodSetStatusArgs) ToStrimziPodSetStatusOutputWithContext(ctx context.Context) StrimziPodSetStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrimziPodSetStatusOutput)
}

func (i StrimziPodSetStatusArgs) ToStrimziPodSetStatusPtrOutput() StrimziPodSetStatusPtrOutput {
	return i.ToStrimziPodSetStatusPtrOutputWithContext(context.Background())
}

func (i StrimziPodSetStatusArgs) ToStrimziPodSetStatusPtrOutputWithContext(ctx context.Context) StrimziPodSetStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrimziPodSetStatusOutput).ToStrimziPodSetStatusPtrOutputWithContext(ctx)
}

// StrimziPodSetStatusPtrInput is an input type that accepts StrimziPodSetStatusArgs, StrimziPodSetStatusPtr and StrimziPodSetStatusPtrOutput values.
// You can construct a concrete instance of `StrimziPodSetStatusPtrInput` via:
//
//	        StrimziPodSetStatusArgs{...}
//
//	or:
//
//	        nil
type StrimziPodSetStatusPtrInput interface {
	pulumi.Input

	ToStrimziPodSetStatusPtrOutput() StrimziPodSetStatusPtrOutput
	ToStrimziPodSetStatusPtrOutputWithContext(context.Context) StrimziPodSetStatusPtrOutput
}

type strimziPodSetStatusPtrType StrimziPodSetStatusArgs

func StrimziPodSetStatusPtr(v *StrimziPodSetStatusArgs) StrimziPodSetStatusPtrInput {
	return (*strimziPodSetStatusPtrType)(v)
}

func (*strimziPodSetStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StrimziPodSetStatus)(nil)).Elem()
}

func (i *strimziPodSetStatusPtrType) ToStrimziPodSetStatusPtrOutput() StrimziPodSetStatusPtrOutput {
	return i.ToStrimziPodSetStatusPtrOutputWithContext(context.Background())
}

func (i *strimziPodSetStatusPtrType) ToStrimziPodSetStatusPtrOutputWithContext(ctx context.Context) StrimziPodSetStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrimziPodSetStatusPtrOutput)
}

// The status of the StrimziPodSet.
type StrimziPodSetStatusOutput struct{ *pulumi.OutputState }

func (StrimziPodSetStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StrimziPodSetStatus)(nil)).Elem()
}

func (o StrimziPodSetStatusOutput) ToStrimziPodSetStatusOutput() StrimziPodSetStatusOutput {
	return o
}

func (o StrimziPodSetStatusOutput) ToStrimziPodSetStatusOutputWithContext(ctx context.Context) StrimziPodSetStatusOutput {
	return o
}

func (o StrimziPodSetStatusOutput) ToStrimziPodSetStatusPtrOutput() StrimziPodSetStatusPtrOutput {
	return o.ToStrimziPodSetStatusPtrOutputWithContext(context.Background())
}

func (o StrimziPodSetStatusOutput) ToStrimziPodSetStatusPtrOutputWithContext(ctx context.Context) StrimziPodSetStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StrimziPodSetStatus) *StrimziPodSetStatus {
		return &v
	}).(StrimziPodSetStatusPtrOutput)
}

// List of status conditions.
func (o StrimziPodSetStatusOutput) Conditions() StrimziPodSetStatusConditionsArrayOutput {
	return o.ApplyT(func(v StrimziPodSetStatus) []StrimziPodSetStatusConditions { return v.Conditions }).(StrimziPodSetStatusConditionsArrayOutput)
}

// Number of pods managed by this `StrimziPodSet` resource that have the current revision.
func (o StrimziPodSetStatusOutput) CurrentPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StrimziPodSetStatus) *int { return v.CurrentPods }).(pulumi.IntPtrOutput)
}

// The generation of the CRD that was last reconciled by the operator.
func (o StrimziPodSetStatusOutput) ObservedGeneration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StrimziPodSetStatus) *int { return v.ObservedGeneration }).(pulumi.IntPtrOutput)
}

// Number of pods managed by this `StrimziPodSet` resource.
func (o StrimziPodSetStatusOutput) Pods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StrimziPodSetStatus) *int { return v.Pods }).(pulumi.IntPtrOutput)
}

// Number of pods managed by this `StrimziPodSet` resource that are ready.
func (o StrimziPodSetStatusOutput) ReadyPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StrimziPodSetStatus) *int { return v.ReadyPods }).(pulumi.IntPtrOutput)
}

type StrimziPodSetStatusPtrOutput struct{ *pulumi.OutputState }

func (StrimziPodSetStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StrimziPodSetStatus)(nil)).Elem()
}

func (o StrimziPodSetStatusPtrOutput) ToStrimziPodSetStatusPtrOutput() StrimziPodSetStatusPtrOutput {
	return o
}

func (o StrimziPodSetStatusPtrOutput) ToStrimziPodSetStatusPtrOutputWithContext(ctx context.Context) StrimziPodSetStatusPtrOutput {
	return o
}

func (o StrimziPodSetStatusPtrOutput) Elem() StrimziPodSetStatusOutput {
	return o.ApplyT(func(v *StrimziPodSetStatus) StrimziPodSetStatus {
		if v != nil {
			return *v
		}
		var ret StrimziPodSetStatus
		return ret
	}).(StrimziPodSetStatusOutput)
}

// List of status conditions.
func (o StrimziPodSetStatusPtrOutput) Conditions() StrimziPodSetStatusConditionsArrayOutput {
	return o.ApplyT(func(v *StrimziPodSetStatus) []StrimziPodSetStatusConditions {
		if v == nil {
			return nil
		}
		return v.Conditions
	}).(StrimziPodSetStatusConditionsArrayOutput)
}

// Number of pods managed by this `StrimziPodSet` resource that have the current revision.
func (o StrimziPodSetStatusPtrOutput) CurrentPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StrimziPodSetStatus) *int {
		if v == nil {
			return nil
		}
		return v.CurrentPods
	}).(pulumi.IntPtrOutput)
}

// The generation of the CRD that was last reconciled by the operator.
func (o StrimziPodSetStatusPtrOutput) ObservedGeneration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StrimziPodSetStatus) *int {
		if v == nil {
			return nil
		}
		return v.ObservedGeneration
	}).(pulumi.IntPtrOutput)
}

// Number of pods managed by this `StrimziPodSet` resource.
func (o StrimziPodSetStatusPtrOutput) Pods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StrimziPodSetStatus) *int {
		if v == nil {
			return nil
		}
		return v.Pods
	}).(pulumi.IntPtrOutput)
}

// Number of pods managed by this `StrimziPodSet` resource that are ready.
func (o StrimziPodSetStatusPtrOutput) ReadyPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StrimziPodSetStatus) *int {
		if v == nil {
			return nil
		}
		return v.ReadyPods
	}).(pulumi.IntPtrOutput)
}

type StrimziPodSetStatusConditions struct {
	// Last time the condition of a type changed from one status to another. The required format is 'yyyy-MM-ddTHH:mm:ssZ', in the UTC time zone.
	LastTransitionTime *string `pulumi:"lastTransitionTime"`
	// Human-readable message indicating details about the condition's last transition.
	Message *string `pulumi:"message"`
	// The reason for the condition's last transition (a single word in CamelCase).
	Reason *string `pulumi:"reason"`
	// The status of the condition, either True, False or Unknown.
	Status *string `pulumi:"status"`
	// The unique identifier of a condition, used to distinguish between other conditions in the resource.
	Type *string `pulumi:"type"`
}

// StrimziPodSetStatusConditionsInput is an input type that accepts StrimziPodSetStatusConditionsArgs and StrimziPodSetStatusConditionsOutput values.
// You can construct a concrete instance of `StrimziPodSetStatusConditionsInput` via:
//
//	StrimziPodSetStatusConditionsArgs{...}
type StrimziPodSetStatusConditionsInput interface {
	pulumi.Input

	ToStrimziPodSetStatusConditionsOutput() StrimziPodSetStatusConditionsOutput
	ToStrimziPodSetStatusConditionsOutputWithContext(context.Context) StrimziPodSetStatusConditionsOutput
}

type StrimziPodSetStatusConditionsArgs struct {
	// Last time the condition of a type changed from one status to another. The required format is 'yyyy-MM-ddTHH:mm:ssZ', in the UTC time zone.
	LastTransitionTime pulumi.StringPtrInput `pulumi:"lastTransitionTime"`
	// Human-readable message indicating details about the condition's last transition.
	Message pulumi.StringPtrInput `pulumi:"message"`
	// The reason for the condition's last transition (a single word in CamelCase).
	Reason pulumi.StringPtrInput `pulumi:"reason"`
	// The status of the condition, either True, False or Unknown.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The unique identifier of a condition, used to distinguish between other conditions in the resource.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (StrimziPodSetStatusConditionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StrimziPodSetStatusConditions)(nil)).Elem()
}

func (i StrimziPodSetStatusConditionsArgs) ToStrimziPodSetStatusConditionsOutput() StrimziPodSetStatusConditionsOutput {
	return i.ToStrimziPodSetStatusConditionsOutputWithContext(context.Background())
}

func (i StrimziPodSetStatusConditionsArgs) ToStrimziPodSetStatusConditionsOutputWithContext(ctx context.Context) StrimziPodSetStatusConditionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrimziPodSetStatusConditionsOutput)
}

// StrimziPodSetStatusConditionsArrayInput is an input type that accepts StrimziPodSetStatusConditionsArray and StrimziPodSetStatusConditionsArrayOutput values.
// You can construct a concrete instance of `StrimziPodSetStatusConditionsArrayInput` via:
//
//	StrimziPodSetStatusConditionsArray{ StrimziPodSetStatusConditionsArgs{...} }
type StrimziPodSetStatusConditionsArrayInput interface {
	pulumi.Input

	ToStrimziPodSetStatusConditionsArrayOutput() StrimziPodSetStatusConditionsArrayOutput
	ToStrimziPodSetStatusConditionsArrayOutputWithContext(context.Context) StrimziPodSetStatusConditionsArrayOutput
}

type StrimziPodSetStatusConditionsArray []StrimziPodSetStatusConditionsInput

func (StrimziPodSetStatusConditionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StrimziPodSetStatusConditions)(nil)).Elem()
}

func (i StrimziPodSetStatusConditionsArray) ToStrimziPodSetStatusConditionsArrayOutput() StrimziPodSetStatusConditionsArrayOutput {
	return i.ToStrimziPodSetStatusConditionsArrayOutputWithContext(context.Background())
}

func (i StrimziPodSetStatusConditionsArray) ToStrimziPodSetStatusConditionsArrayOutputWithContext(ctx context.Context) StrimziPodSetStatusConditionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StrimziPodSetStatusConditionsArrayOutput)
}

type StrimziPodSetStatusConditionsOutput struct{ *pulumi.OutputState }

func (StrimziPodSetStatusConditionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StrimziPodSetStatusConditions)(nil)).Elem()
}

func (o StrimziPodSetStatusConditionsOutput) ToStrimziPodSetStatusConditionsOutput() StrimziPodSetStatusConditionsOutput {
	return o
}

func (o StrimziPodSetStatusConditionsOutput) ToStrimziPodSetStatusConditionsOutputWithContext(ctx context.Context) StrimziPodSetStatusConditionsOutput {
	return o
}

// Last time the condition of a type changed from one status to another. The required format is 'yyyy-MM-ddTHH:mm:ssZ', in the UTC time zone.
func (o StrimziPodSetStatusConditionsOutput) LastTransitionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StrimziPodSetStatusConditions) *string { return v.LastTransitionTime }).(pulumi.StringPtrOutput)
}

// Human-readable message indicating details about the condition's last transition.
func (o StrimziPodSetStatusConditionsOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StrimziPodSetStatusConditions) *string { return v.Message }).(pulumi.StringPtrOutput)
}

// The reason for the condition's last transition (a single word in CamelCase).
func (o StrimziPodSetStatusConditionsOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StrimziPodSetStatusConditions) *string { return v.Reason }).(pulumi.StringPtrOutput)
}

// The status of the condition, either True, False or Unknown.
func (o StrimziPodSetStatusConditionsOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StrimziPodSetStatusConditions) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The unique identifier of a condition, used to distinguish between other conditions in the resource.
func (o StrimziPodSetStatusConditionsOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StrimziPodSetStatusConditions) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type StrimziPodSetStatusConditionsArrayOutput struct{ *pulumi.OutputState }

func (StrimziPodSetStatusConditionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StrimziPodSetStatusConditions)(nil)).Elem()
}

func (o StrimziPodSetStatusConditionsArrayOutput) ToStrimziPodSetStatusConditionsArrayOutput() StrimziPodSetStatusConditionsArrayOutput {
	return o
}

func (o StrimziPodSetStatusConditionsArrayOutput) ToStrimziPodSetStatusConditionsArrayOutputWithContext(ctx context.Context) StrimziPodSetStatusConditionsArrayOutput {
	return o
}

func (o StrimziPodSetStatusConditionsArrayOutput) Index(i pulumi.IntInput) StrimziPodSetStatusConditionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StrimziPodSetStatusConditions {
		return vs[0].([]StrimziPodSetStatusConditions)[vs[1].(int)]
	}).(StrimziPodSetStatusConditionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StrimziPodSetSpecInput)(nil)).Elem(), StrimziPodSetSpecArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StrimziPodSetSpecPtrInput)(nil)).Elem(), StrimziPodSetSpecArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StrimziPodSetSpecSelectorInput)(nil)).Elem(), StrimziPodSetSpecSelectorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StrimziPodSetSpecSelectorPtrInput)(nil)).Elem(), StrimziPodSetSpecSelectorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StrimziPodSetSpecSelectorMatchExpressionsInput)(nil)).Elem(), StrimziPodSetSpecSelectorMatchExpressionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StrimziPodSetSpecSelectorMatchExpressionsArrayInput)(nil)).Elem(), StrimziPodSetSpecSelectorMatchExpressionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StrimziPodSetStatusInput)(nil)).Elem(), StrimziPodSetStatusArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StrimziPodSetStatusPtrInput)(nil)).Elem(), StrimziPodSetStatusArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StrimziPodSetStatusConditionsInput)(nil)).Elem(), StrimziPodSetStatusConditionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StrimziPodSetStatusConditionsArrayInput)(nil)).Elem(), StrimziPodSetStatusConditionsArray{})
	pulumi.RegisterOutputType(StrimziPodSetSpecOutput{})
	pulumi.RegisterOutputType(StrimziPodSetSpecPtrOutput{})
	pulumi.RegisterOutputType(StrimziPodSetSpecSelectorOutput{})
	pulumi.RegisterOutputType(StrimziPodSetSpecSelectorPtrOutput{})
	pulumi.RegisterOutputType(StrimziPodSetSpecSelectorMatchExpressionsOutput{})
	pulumi.RegisterOutputType(StrimziPodSetSpecSelectorMatchExpressionsArrayOutput{})
	pulumi.RegisterOutputType(StrimziPodSetStatusOutput{})
	pulumi.RegisterOutputType(StrimziPodSetStatusPtrOutput{})
	pulumi.RegisterOutputType(StrimziPodSetStatusConditionsOutput{})
	pulumi.RegisterOutputType(StrimziPodSetStatusConditionsArrayOutput{})
}
