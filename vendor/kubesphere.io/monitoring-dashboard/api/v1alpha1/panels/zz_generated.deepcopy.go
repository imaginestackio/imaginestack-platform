// +build !ignore_autogenerated

/*
Copyright 2023 The ImagineKube authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package panels

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Graph) DeepCopyInto(out *Graph) {
	*out = *in
	if in.Colors != nil {
		in, out := &in.Colors, &out.Colors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Yaxes != nil {
		in, out := &in.Yaxes, &out.Yaxes
		*out = make([]Yaxis, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Graph.
func (in *Graph) DeepCopy() *Graph {
	if in == nil {
		return nil
	}
	out := new(Graph)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Row) DeepCopyInto(out *Row) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Row.
func (in *Row) DeepCopy() *Row {
	if in == nil {
		return nil
	}
	out := new(Row)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SingleStat) DeepCopyInto(out *SingleStat) {
	*out = *in
	if in.Decimals != nil {
		in, out := &in.Decimals, &out.Decimals
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleStat.
func (in *SingleStat) DeepCopy() *SingleStat {
	if in == nil {
		return nil
	}
	out := new(SingleStat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Target) DeepCopyInto(out *Target) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Target.
func (in *Target) DeepCopy() *Target {
	if in == nil {
		return nil
	}
	out := new(Target)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Yaxis) DeepCopyInto(out *Yaxis) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Yaxis.
func (in *Yaxis) DeepCopy() *Yaxis {
	if in == nil {
		return nil
	}
	out := new(Yaxis)
	in.DeepCopyInto(out)
	return out
}
