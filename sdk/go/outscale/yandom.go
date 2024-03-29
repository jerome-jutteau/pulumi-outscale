// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package outscale

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Yandom struct {
	pulumi.CustomResourceState

	Length pulumi.IntOutput    `pulumi:"length"`
	Result pulumi.StringOutput `pulumi:"result"`
}

// NewYandom registers a new resource with the given unique name, arguments, and options.
func NewYandom(ctx *pulumi.Context,
	name string, args *YandomArgs, opts ...pulumi.ResourceOption) (*Yandom, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Length == nil {
		return nil, errors.New("invalid value for required argument 'Length'")
	}
	var resource Yandom
	err := ctx.RegisterResource("outscale:index:Yandom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetYandom gets an existing Yandom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetYandom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *YandomState, opts ...pulumi.ResourceOption) (*Yandom, error) {
	var resource Yandom
	err := ctx.ReadResource("outscale:index:Yandom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Yandom resources.
type yandomState struct {
}

type YandomState struct {
}

func (YandomState) ElementType() reflect.Type {
	return reflect.TypeOf((*yandomState)(nil)).Elem()
}

type yandomArgs struct {
	Length int `pulumi:"length"`
}

// The set of arguments for constructing a Yandom resource.
type YandomArgs struct {
	Length pulumi.IntInput
}

func (YandomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*yandomArgs)(nil)).Elem()
}

type YandomInput interface {
	pulumi.Input

	ToYandomOutput() YandomOutput
	ToYandomOutputWithContext(ctx context.Context) YandomOutput
}

func (*Yandom) ElementType() reflect.Type {
	return reflect.TypeOf((**Yandom)(nil)).Elem()
}

func (i *Yandom) ToYandomOutput() YandomOutput {
	return i.ToYandomOutputWithContext(context.Background())
}

func (i *Yandom) ToYandomOutputWithContext(ctx context.Context) YandomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YandomOutput)
}

type YandomOutput struct{ *pulumi.OutputState }

func (YandomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Yandom)(nil)).Elem()
}

func (o YandomOutput) ToYandomOutput() YandomOutput {
	return o
}

func (o YandomOutput) ToYandomOutputWithContext(ctx context.Context) YandomOutput {
	return o
}

func (o YandomOutput) Length() pulumi.IntOutput {
	return o.ApplyT(func(v *Yandom) pulumi.IntOutput { return v.Length }).(pulumi.IntOutput)
}

func (o YandomOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v *Yandom) pulumi.StringOutput { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*YandomInput)(nil)).Elem(), &Yandom{})
	pulumi.RegisterOutputType(YandomOutput{})
}
