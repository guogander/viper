package viper

import "github.com/spf13/pflag"

// FlagValueSet is an interface that users can implement
// to bind a set of flags to viper.
// FlagValueSet是一个用户可以实现的将flags绑定在viper上的接口
type FlagValueSet interface {
	VisitAll(fn func(FlagValue))
}

// FlagValue is an interface that users can implement
// to bind different flags to viper.
// FlagValueSet是一个用户可以实现的将不同的flags绑定在viper上的接口
type FlagValue interface {
	HasChanged() bool
	Name() string
	ValueString() string
	ValueType() string
}

// pflagValueSet is a wrapper around *pflag.ValueSet
// that implements FlagValueSet.
// pflagValueSet对*pflag.ValueSet进行包装，实现FlagValueSet这个方法
type pflagValueSet struct {
	flags *pflag.FlagSet
}

// VisitAll iterates over all *pflag.Flag inside the *pflag.FlagSet.
// VisitAll遍历所有包含在*pflag.FlagSet中的*pflag.Flag
func (p pflagValueSet) VisitAll(fn func(flag FlagValue)) {
	p.flags.VisitAll(func(flag *pflag.Flag) {
		fn(pflagValue{flag})
	})
}

// pflagValue is a wrapper aroung *pflag.flag
// that implements FlagValue
// pflagValue对*pflag.flag进行包装，实现FlagValue这个方法
type pflagValue struct {
	flag *pflag.Flag
}

// HasChanged returns whether the flag has changes or not.
// HasChanged返回flag标志是否有变化
func (p pflagValue) HasChanged() bool {
	return p.flag.Changed
}

// Name returns the name of the flag.
// Name返回flag标志的名称。
func (p pflagValue) Name() string {
	return p.flag.Name
}

// ValueString returns the value of the flag as a string.
// ValueString以字符串形式返回flag标志值。
func (p pflagValue) ValueString() string {
	return p.flag.Value.String()
}

// ValueType returns the type of the flag as a string.
// ValueType以字符串形式返回flag标志类型。
func (p pflagValue) ValueType() string {
	return p.flag.Value.Type()
}
