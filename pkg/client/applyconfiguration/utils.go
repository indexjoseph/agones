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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1 "agones.dev/agones/pkg/apis/agones/v1"
	autoscalingv1 "agones.dev/agones/pkg/apis/autoscaling/v1"
	multiclusterv1 "agones.dev/agones/pkg/apis/multicluster/v1"
	agonesv1 "agones.dev/agones/pkg/client/applyconfiguration/agones/v1"
	applyconfigurationautoscalingv1 "agones.dev/agones/pkg/client/applyconfiguration/autoscaling/v1"
	applyconfigurationmulticlusterv1 "agones.dev/agones/pkg/client/applyconfiguration/multicluster/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=agones.dev, Version=v1
	case v1.SchemeGroupVersion.WithKind("AggregatedCounterStatus"):
		return &agonesv1.AggregatedCounterStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AggregatedListStatus"):
		return &agonesv1.AggregatedListStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AggregatedPlayerStatus"):
		return &agonesv1.AggregatedPlayerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AllocationOverflow"):
		return &agonesv1.AllocationOverflowApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CounterStatus"):
		return &agonesv1.CounterStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Eviction"):
		return &agonesv1.EvictionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Fleet"):
		return &agonesv1.FleetApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FleetSpec"):
		return &agonesv1.FleetSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FleetStatus"):
		return &agonesv1.FleetStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GameServer"):
		return &agonesv1.GameServerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GameServerPort"):
		return &agonesv1.GameServerPortApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GameServerSet"):
		return &agonesv1.GameServerSetApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GameServerSetSpec"):
		return &agonesv1.GameServerSetSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GameServerSetStatus"):
		return &agonesv1.GameServerSetStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GameServerSpec"):
		return &agonesv1.GameServerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GameServerStatus"):
		return &agonesv1.GameServerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GameServerStatusPort"):
		return &agonesv1.GameServerStatusPortApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GameServerTemplateSpec"):
		return &agonesv1.GameServerTemplateSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Health"):
		return &agonesv1.HealthApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ListStatus"):
		return &agonesv1.ListStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PlayersSpec"):
		return &agonesv1.PlayersSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PlayerStatus"):
		return &agonesv1.PlayerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Priority"):
		return &agonesv1.PriorityApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SdkServer"):
		return &agonesv1.SdkServerApplyConfiguration{}

		// Group=autoscaling.agones.dev, Version=v1
	case autoscalingv1.SchemeGroupVersion.WithKind("BufferPolicy"):
		return &applyconfigurationautoscalingv1.BufferPolicyApplyConfiguration{}
	case autoscalingv1.SchemeGroupVersion.WithKind("CounterPolicy"):
		return &applyconfigurationautoscalingv1.CounterPolicyApplyConfiguration{}
	case autoscalingv1.SchemeGroupVersion.WithKind("FixedIntervalSync"):
		return &applyconfigurationautoscalingv1.FixedIntervalSyncApplyConfiguration{}
	case autoscalingv1.SchemeGroupVersion.WithKind("FleetAutoscaler"):
		return &applyconfigurationautoscalingv1.FleetAutoscalerApplyConfiguration{}
	case autoscalingv1.SchemeGroupVersion.WithKind("FleetAutoscalerPolicy"):
		return &applyconfigurationautoscalingv1.FleetAutoscalerPolicyApplyConfiguration{}
	case autoscalingv1.SchemeGroupVersion.WithKind("FleetAutoscalerSpec"):
		return &applyconfigurationautoscalingv1.FleetAutoscalerSpecApplyConfiguration{}
	case autoscalingv1.SchemeGroupVersion.WithKind("FleetAutoscalerStatus"):
		return &applyconfigurationautoscalingv1.FleetAutoscalerStatusApplyConfiguration{}
	case autoscalingv1.SchemeGroupVersion.WithKind("FleetAutoscalerSync"):
		return &applyconfigurationautoscalingv1.FleetAutoscalerSyncApplyConfiguration{}
	case autoscalingv1.SchemeGroupVersion.WithKind("ListPolicy"):
		return &applyconfigurationautoscalingv1.ListPolicyApplyConfiguration{}
	case autoscalingv1.SchemeGroupVersion.WithKind("WebhookPolicy"):
		return &applyconfigurationautoscalingv1.WebhookPolicyApplyConfiguration{}

		// Group=multicluster.agones.dev, Version=v1
	case multiclusterv1.SchemeGroupVersion.WithKind("ClusterConnectionInfo"):
		return &applyconfigurationmulticlusterv1.ClusterConnectionInfoApplyConfiguration{}
	case multiclusterv1.SchemeGroupVersion.WithKind("GameServerAllocationPolicy"):
		return &applyconfigurationmulticlusterv1.GameServerAllocationPolicyApplyConfiguration{}
	case multiclusterv1.SchemeGroupVersion.WithKind("GameServerAllocationPolicySpec"):
		return &applyconfigurationmulticlusterv1.GameServerAllocationPolicySpecApplyConfiguration{}

	}
	return nil
}
