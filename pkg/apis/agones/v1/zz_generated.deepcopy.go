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
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AggregatedCounterStatus) DeepCopyInto(out *AggregatedCounterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AggregatedCounterStatus.
func (in *AggregatedCounterStatus) DeepCopy() *AggregatedCounterStatus {
	if in == nil {
		return nil
	}
	out := new(AggregatedCounterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AggregatedListStatus) DeepCopyInto(out *AggregatedListStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AggregatedListStatus.
func (in *AggregatedListStatus) DeepCopy() *AggregatedListStatus {
	if in == nil {
		return nil
	}
	out := new(AggregatedListStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AggregatedPlayerStatus) DeepCopyInto(out *AggregatedPlayerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AggregatedPlayerStatus.
func (in *AggregatedPlayerStatus) DeepCopy() *AggregatedPlayerStatus {
	if in == nil {
		return nil
	}
	out := new(AggregatedPlayerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationOverflow) DeepCopyInto(out *AllocationOverflow) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationOverflow.
func (in *AllocationOverflow) DeepCopy() *AllocationOverflow {
	if in == nil {
		return nil
	}
	out := new(AllocationOverflow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CounterStatus) DeepCopyInto(out *CounterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CounterStatus.
func (in *CounterStatus) DeepCopy() *CounterStatus {
	if in == nil {
		return nil
	}
	out := new(CounterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Eviction) DeepCopyInto(out *Eviction) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Eviction.
func (in *Eviction) DeepCopy() *Eviction {
	if in == nil {
		return nil
	}
	out := new(Eviction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fleet) DeepCopyInto(out *Fleet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fleet.
func (in *Fleet) DeepCopy() *Fleet {
	if in == nil {
		return nil
	}
	out := new(Fleet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Fleet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetList) DeepCopyInto(out *FleetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Fleet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetList.
func (in *FleetList) DeepCopy() *FleetList {
	if in == nil {
		return nil
	}
	out := new(FleetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetSpec) DeepCopyInto(out *FleetSpec) {
	*out = *in
	if in.AllocationOverflow != nil {
		in, out := &in.AllocationOverflow, &out.AllocationOverflow
		*out = new(AllocationOverflow)
		(*in).DeepCopyInto(*out)
	}
	in.Strategy.DeepCopyInto(&out.Strategy)
	if in.Priorities != nil {
		in, out := &in.Priorities, &out.Priorities
		*out = make([]Priority, len(*in))
		copy(*out, *in)
	}
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetSpec.
func (in *FleetSpec) DeepCopy() *FleetSpec {
	if in == nil {
		return nil
	}
	out := new(FleetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetStatus) DeepCopyInto(out *FleetStatus) {
	*out = *in
	if in.Players != nil {
		in, out := &in.Players, &out.Players
		*out = new(AggregatedPlayerStatus)
		**out = **in
	}
	if in.Counters != nil {
		in, out := &in.Counters, &out.Counters
		*out = make(map[string]AggregatedCounterStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Lists != nil {
		in, out := &in.Lists, &out.Lists
		*out = make(map[string]AggregatedListStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetStatus.
func (in *FleetStatus) DeepCopy() *FleetStatus {
	if in == nil {
		return nil
	}
	out := new(FleetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServer) DeepCopyInto(out *GameServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServer.
func (in *GameServer) DeepCopy() *GameServer {
	if in == nil {
		return nil
	}
	out := new(GameServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GameServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerList) DeepCopyInto(out *GameServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GameServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerList.
func (in *GameServerList) DeepCopy() *GameServerList {
	if in == nil {
		return nil
	}
	out := new(GameServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GameServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerPort) DeepCopyInto(out *GameServerPort) {
	*out = *in
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerPort.
func (in *GameServerPort) DeepCopy() *GameServerPort {
	if in == nil {
		return nil
	}
	out := new(GameServerPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerSet) DeepCopyInto(out *GameServerSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerSet.
func (in *GameServerSet) DeepCopy() *GameServerSet {
	if in == nil {
		return nil
	}
	out := new(GameServerSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GameServerSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerSetList) DeepCopyInto(out *GameServerSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GameServerSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerSetList.
func (in *GameServerSetList) DeepCopy() *GameServerSetList {
	if in == nil {
		return nil
	}
	out := new(GameServerSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GameServerSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerSetSpec) DeepCopyInto(out *GameServerSetSpec) {
	*out = *in
	if in.AllocationOverflow != nil {
		in, out := &in.AllocationOverflow, &out.AllocationOverflow
		*out = new(AllocationOverflow)
		(*in).DeepCopyInto(*out)
	}
	if in.Priorities != nil {
		in, out := &in.Priorities, &out.Priorities
		*out = make([]Priority, len(*in))
		copy(*out, *in)
	}
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerSetSpec.
func (in *GameServerSetSpec) DeepCopy() *GameServerSetSpec {
	if in == nil {
		return nil
	}
	out := new(GameServerSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerSetStatus) DeepCopyInto(out *GameServerSetStatus) {
	*out = *in
	if in.Players != nil {
		in, out := &in.Players, &out.Players
		*out = new(AggregatedPlayerStatus)
		**out = **in
	}
	if in.Counters != nil {
		in, out := &in.Counters, &out.Counters
		*out = make(map[string]AggregatedCounterStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Lists != nil {
		in, out := &in.Lists, &out.Lists
		*out = make(map[string]AggregatedListStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerSetStatus.
func (in *GameServerSetStatus) DeepCopy() *GameServerSetStatus {
	if in == nil {
		return nil
	}
	out := new(GameServerSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerSpec) DeepCopyInto(out *GameServerSpec) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]GameServerPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Health = in.Health
	out.SdkServer = in.SdkServer
	in.Template.DeepCopyInto(&out.Template)
	if in.Players != nil {
		in, out := &in.Players, &out.Players
		*out = new(PlayersSpec)
		**out = **in
	}
	if in.Counters != nil {
		in, out := &in.Counters, &out.Counters
		*out = make(map[string]CounterStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Lists != nil {
		in, out := &in.Lists, &out.Lists
		*out = make(map[string]ListStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Eviction != nil {
		in, out := &in.Eviction, &out.Eviction
		*out = new(Eviction)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerSpec.
func (in *GameServerSpec) DeepCopy() *GameServerSpec {
	if in == nil {
		return nil
	}
	out := new(GameServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerStatus) DeepCopyInto(out *GameServerStatus) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]GameServerStatusPort, len(*in))
		copy(*out, *in)
	}
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]corev1.NodeAddress, len(*in))
		copy(*out, *in)
	}
	if in.ReservedUntil != nil {
		in, out := &in.ReservedUntil, &out.ReservedUntil
		*out = (*in).DeepCopy()
	}
	if in.Players != nil {
		in, out := &in.Players, &out.Players
		*out = new(PlayerStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Counters != nil {
		in, out := &in.Counters, &out.Counters
		*out = make(map[string]CounterStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Lists != nil {
		in, out := &in.Lists, &out.Lists
		*out = make(map[string]ListStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Eviction != nil {
		in, out := &in.Eviction, &out.Eviction
		*out = new(Eviction)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerStatus.
func (in *GameServerStatus) DeepCopy() *GameServerStatus {
	if in == nil {
		return nil
	}
	out := new(GameServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerStatusPort) DeepCopyInto(out *GameServerStatusPort) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerStatusPort.
func (in *GameServerStatusPort) DeepCopy() *GameServerStatusPort {
	if in == nil {
		return nil
	}
	out := new(GameServerStatusPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerTemplateSpec) DeepCopyInto(out *GameServerTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerTemplateSpec.
func (in *GameServerTemplateSpec) DeepCopy() *GameServerTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(GameServerTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Health) DeepCopyInto(out *Health) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Health.
func (in *Health) DeepCopy() *Health {
	if in == nil {
		return nil
	}
	out := new(Health)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListStatus) DeepCopyInto(out *ListStatus) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListStatus.
func (in *ListStatus) DeepCopy() *ListStatus {
	if in == nil {
		return nil
	}
	out := new(ListStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlayerStatus) DeepCopyInto(out *PlayerStatus) {
	*out = *in
	if in.IDs != nil {
		in, out := &in.IDs, &out.IDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlayerStatus.
func (in *PlayerStatus) DeepCopy() *PlayerStatus {
	if in == nil {
		return nil
	}
	out := new(PlayerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlayersSpec) DeepCopyInto(out *PlayersSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlayersSpec.
func (in *PlayersSpec) DeepCopy() *PlayersSpec {
	if in == nil {
		return nil
	}
	out := new(PlayersSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Priority) DeepCopyInto(out *Priority) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Priority.
func (in *Priority) DeepCopy() *Priority {
	if in == nil {
		return nil
	}
	out := new(Priority)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdkServer) DeepCopyInto(out *SdkServer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdkServer.
func (in *SdkServer) DeepCopy() *SdkServer {
	if in == nil {
		return nil
	}
	out := new(SdkServer)
	in.DeepCopyInto(out)
	return out
}
