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
// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: protoc-gen-openapiv2/options/annotations.proto

#include "protoc-gen-openapiv2/options/annotations.pb.h"

#include <algorithm>
#include "google/protobuf/io/coded_stream.h"
#include "google/protobuf/extension_set.h"
#include "google/protobuf/wire_format_lite.h"
#include "google/protobuf/descriptor.h"
#include "google/protobuf/generated_message_reflection.h"
#include "google/protobuf/reflection_ops.h"
#include "google/protobuf/wire_format.h"
// @@protoc_insertion_point(includes)

// Must be included last.
#include "google/protobuf/port_def.inc"
PROTOBUF_PRAGMA_INIT_SEG
namespace _pb = ::PROTOBUF_NAMESPACE_ID;
namespace _pbi = ::PROTOBUF_NAMESPACE_ID::internal;
namespace grpc {
namespace gateway {
namespace protoc_gen_openapiv2 {
namespace options {
}  // namespace options
}  // namespace protoc_gen_openapiv2
}  // namespace gateway
}  // namespace grpc
static constexpr const ::_pb::EnumDescriptor**
    file_level_enum_descriptors_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto = nullptr;
static constexpr const ::_pb::ServiceDescriptor**
    file_level_service_descriptors_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto = nullptr;
const ::uint32_t TableStruct_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto::offsets[1] = {};
static constexpr ::_pbi::MigrationSchema* schemas = nullptr;
static constexpr ::_pb::Message* const* file_default_instances = nullptr;
const char descriptor_table_protodef_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
    "\n.protoc-gen-openapiv2/options/annotatio"
    "ns.proto\022)grpc.gateway.protoc_gen_openap"
    "iv2.options\032 google/protobuf/descriptor."
    "proto\032,protoc-gen-openapiv2/options/open"
    "apiv2.proto:l\n\021openapiv2_swagger\022\034.googl"
    "e.protobuf.FileOptions\030\222\010 \001(\01322.grpc.gat"
    "eway.protoc_gen_openapiv2.options.Swagge"
    "r:r\n\023openapiv2_operation\022\036.google.protob"
    "uf.MethodOptions\030\222\010 \001(\01324.grpc.gateway.p"
    "rotoc_gen_openapiv2.options.Operation:m\n"
    "\020openapiv2_schema\022\037.google.protobuf.Mess"
    "ageOptions\030\222\010 \001(\01321.grpc.gateway.protoc_"
    "gen_openapiv2.options.Schema:g\n\ropenapiv"
    "2_tag\022\037.google.protobuf.ServiceOptions\030\222"
    "\010 \001(\0132..grpc.gateway.protoc_gen_openapiv"
    "2.options.Tag:n\n\017openapiv2_field\022\035.googl"
    "e.protobuf.FieldOptions\030\222\010 \001(\01325.grpc.ga"
    "teway.protoc_gen_openapiv2.options.JSONS"
    "chemaBHZFgithub.com/grpc-ecosystem/grpc-"
    "gateway/v2/protoc-gen-openapiv2/optionsb"
    "\006proto3"
};
static const ::_pbi::DescriptorTable* const descriptor_table_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto_deps[2] =
    {
        &::descriptor_table_google_2fprotobuf_2fdescriptor_2eproto,
        &::descriptor_table_protoc_2dgen_2dopenapiv2_2foptions_2fopenapiv2_2eproto,
};
static ::absl::once_flag descriptor_table_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto = {
    false,
    false,
    807,
    descriptor_table_protodef_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto,
    "protoc-gen-openapiv2/options/annotations.proto",
    &descriptor_table_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto_once,
    descriptor_table_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto_deps,
    2,
    0,
    schemas,
    file_default_instances,
    TableStruct_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto::offsets,
    nullptr,
    file_level_enum_descriptors_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto,
    file_level_service_descriptors_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto,
};

// This function exists to be marked as weak.
// It can significantly speed up compilation by breaking up LLVM's SCC
// in the .pb.cc translation units. Large translation units see a
// reduction of more than 35% of walltime for optimized builds. Without
// the weak attribute all the messages in the file, including all the
// vtables and everything they use become part of the same SCC through
// a cycle like:
// GetMetadata -> descriptor table -> default instances ->
//   vtables -> GetMetadata
// By adding a weak function here we break the connection from the
// individual vtables back into the descriptor table.
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto_getter() {
  return &descriptor_table_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto;
}
// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2
static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto(&descriptor_table_protoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto);
namespace grpc {
namespace gateway {
namespace protoc_gen_openapiv2 {
namespace options {
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 ::PROTOBUF_NAMESPACE_ID::internal::ExtensionIdentifier< ::PROTOBUF_NAMESPACE_ID::FileOptions,
    ::PROTOBUF_NAMESPACE_ID::internal::MessageTypeTraits< ::grpc::gateway::protoc_gen_openapiv2::options::Swagger >, 11, false>
  openapiv2_swagger(kOpenapiv2SwaggerFieldNumber, ::grpc::gateway::protoc_gen_openapiv2::options::Swagger::default_instance(), nullptr);
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 ::PROTOBUF_NAMESPACE_ID::internal::ExtensionIdentifier< ::PROTOBUF_NAMESPACE_ID::MethodOptions,
    ::PROTOBUF_NAMESPACE_ID::internal::MessageTypeTraits< ::grpc::gateway::protoc_gen_openapiv2::options::Operation >, 11, false>
  openapiv2_operation(kOpenapiv2OperationFieldNumber, ::grpc::gateway::protoc_gen_openapiv2::options::Operation::default_instance(), nullptr);
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 ::PROTOBUF_NAMESPACE_ID::internal::ExtensionIdentifier< ::PROTOBUF_NAMESPACE_ID::MessageOptions,
    ::PROTOBUF_NAMESPACE_ID::internal::MessageTypeTraits< ::grpc::gateway::protoc_gen_openapiv2::options::Schema >, 11, false>
  openapiv2_schema(kOpenapiv2SchemaFieldNumber, ::grpc::gateway::protoc_gen_openapiv2::options::Schema::default_instance(), nullptr);
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 ::PROTOBUF_NAMESPACE_ID::internal::ExtensionIdentifier< ::PROTOBUF_NAMESPACE_ID::ServiceOptions,
    ::PROTOBUF_NAMESPACE_ID::internal::MessageTypeTraits< ::grpc::gateway::protoc_gen_openapiv2::options::Tag >, 11, false>
  openapiv2_tag(kOpenapiv2TagFieldNumber, ::grpc::gateway::protoc_gen_openapiv2::options::Tag::default_instance(), nullptr);
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 ::PROTOBUF_NAMESPACE_ID::internal::ExtensionIdentifier< ::PROTOBUF_NAMESPACE_ID::FieldOptions,
    ::PROTOBUF_NAMESPACE_ID::internal::MessageTypeTraits< ::grpc::gateway::protoc_gen_openapiv2::options::JSONSchema >, 11, false>
  openapiv2_field(kOpenapiv2FieldFieldNumber, ::grpc::gateway::protoc_gen_openapiv2::options::JSONSchema::default_instance(), nullptr);
// @@protoc_insertion_point(namespace_scope)
}  // namespace options
}  // namespace protoc_gen_openapiv2
}  // namespace gateway
}  // namespace grpc
PROTOBUF_NAMESPACE_OPEN
PROTOBUF_NAMESPACE_CLOSE
// @@protoc_insertion_point(global_scope)
#include "google/protobuf/port_undef.inc"
