//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2024 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivePeriod) DeepCopyInto(out *ActivePeriod) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivePeriod.
func (in *ActivePeriod) DeepCopy() *ActivePeriod {
	if in == nil {
		return nil
	}
	out := new(ActivePeriod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Between) DeepCopyInto(out *Between) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Between.
func (in *Between) DeepCopy() *Between {
	if in == nil {
		return nil
	}
	out := new(Between)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BufferPolicy) DeepCopyInto(out *BufferPolicy) {
	*out = *in
	out.BufferSize = in.BufferSize
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BufferPolicy.
func (in *BufferPolicy) DeepCopy() *BufferPolicy {
	if in == nil {
		return nil
	}
	out := new(BufferPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChainEntry) DeepCopyInto(out *ChainEntry) {
	*out = *in
	out.Schedule = in.Schedule
	in.Policy.DeepCopyInto(&out.Policy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChainEntry.
func (in *ChainEntry) DeepCopy() *ChainEntry {
	if in == nil {
		return nil
	}
	out := new(ChainEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ChainPolicy) DeepCopyInto(out *ChainPolicy) {
	{
		in := &in
		*out = make(ChainPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChainPolicy.
func (in ChainPolicy) DeepCopy() ChainPolicy {
	if in == nil {
		return nil
	}
	out := new(ChainPolicy)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CounterPolicy) DeepCopyInto(out *CounterPolicy) {
	*out = *in
	out.BufferSize = in.BufferSize
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CounterPolicy.
func (in *CounterPolicy) DeepCopy() *CounterPolicy {
	if in == nil {
		return nil
	}
	out := new(CounterPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FixedIntervalSync) DeepCopyInto(out *FixedIntervalSync) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FixedIntervalSync.
func (in *FixedIntervalSync) DeepCopy() *FixedIntervalSync {
	if in == nil {
		return nil
	}
	out := new(FixedIntervalSync)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscaleRequest) DeepCopyInto(out *FleetAutoscaleRequest) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscaleRequest.
func (in *FleetAutoscaleRequest) DeepCopy() *FleetAutoscaleRequest {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscaleRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscaleResponse) DeepCopyInto(out *FleetAutoscaleResponse) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscaleResponse.
func (in *FleetAutoscaleResponse) DeepCopy() *FleetAutoscaleResponse {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscaleResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscaleReview) DeepCopyInto(out *FleetAutoscaleReview) {
	*out = *in
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(FleetAutoscaleRequest)
		(*in).DeepCopyInto(*out)
	}
	if in.Response != nil {
		in, out := &in.Response, &out.Response
		*out = new(FleetAutoscaleResponse)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscaleReview.
func (in *FleetAutoscaleReview) DeepCopy() *FleetAutoscaleReview {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscaleReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscaler) DeepCopyInto(out *FleetAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscaler.
func (in *FleetAutoscaler) DeepCopy() *FleetAutoscaler {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscalerList) DeepCopyInto(out *FleetAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FleetAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscalerList.
func (in *FleetAutoscalerList) DeepCopy() *FleetAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscalerPolicy) DeepCopyInto(out *FleetAutoscalerPolicy) {
	*out = *in
	if in.Buffer != nil {
		in, out := &in.Buffer, &out.Buffer
		*out = new(BufferPolicy)
		**out = **in
	}
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		*out = new(WebhookPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Counter != nil {
		in, out := &in.Counter, &out.Counter
		*out = new(CounterPolicy)
		**out = **in
	}
	if in.List != nil {
		in, out := &in.List, &out.List
		*out = new(ListPolicy)
		**out = **in
	}
	if in.Chain != nil {
		in, out := &in.Chain, &out.Chain
		*out = make(ChainPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscalerPolicy.
func (in *FleetAutoscalerPolicy) DeepCopy() *FleetAutoscalerPolicy {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscalerPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscalerSpec) DeepCopyInto(out *FleetAutoscalerSpec) {
	*out = *in
	in.Policy.DeepCopyInto(&out.Policy)
	if in.Sync != nil {
		in, out := &in.Sync, &out.Sync
		*out = new(FleetAutoscalerSync)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscalerSpec.
func (in *FleetAutoscalerSpec) DeepCopy() *FleetAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscalerStatus) DeepCopyInto(out *FleetAutoscalerStatus) {
	*out = *in
	if in.LastScaleTime != nil {
		in, out := &in.LastScaleTime, &out.LastScaleTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscalerStatus.
func (in *FleetAutoscalerStatus) DeepCopy() *FleetAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscalerSync) DeepCopyInto(out *FleetAutoscalerSync) {
	*out = *in
	out.FixedInterval = in.FixedInterval
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscalerSync.
func (in *FleetAutoscalerSync) DeepCopy() *FleetAutoscalerSync {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscalerSync)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListPolicy) DeepCopyInto(out *ListPolicy) {
	*out = *in
	out.BufferSize = in.BufferSize
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListPolicy.
func (in *ListPolicy) DeepCopy() *ListPolicy {
	if in == nil {
		return nil
	}
	out := new(ListPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Schedule) DeepCopyInto(out *Schedule) {
	*out = *in
	out.Between = in.Between
	out.ActivePeriod = in.ActivePeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Schedule.
func (in *Schedule) DeepCopy() *Schedule {
	if in == nil {
		return nil
	}
	out := new(Schedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookPolicy) DeepCopyInto(out *WebhookPolicy) {
	*out = *in
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(admissionregistrationv1.ServiceReference)
		(*in).DeepCopyInto(*out)
	}
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookPolicy.
func (in *WebhookPolicy) DeepCopy() *WebhookPolicy {
	if in == nil {
		return nil
	}
	out := new(WebhookPolicy)
	in.DeepCopyInto(out)
	return out
}
