# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/street.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto import city_pb2 as proto_dot_city__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/street.proto',
  package='street',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x12proto/street.proto\x12\x06street\x1a\x10proto/city.proto\"7\n\x06Street\x12\x13\n\x0bstreet_name\x18\x01 \x01(\t\x12\x18\n\x04\x63ity\x18\x02 \x01(\x0b\x32\n.city.Cityb\x06proto3'
  ,
  dependencies=[proto_dot_city__pb2.DESCRIPTOR,])




_STREET = _descriptor.Descriptor(
  name='Street',
  full_name='street.Street',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='street_name', full_name='street.Street.street_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='city', full_name='street.Street.city', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=48,
  serialized_end=103,
)

_STREET.fields_by_name['city'].message_type = proto_dot_city__pb2._CITY
DESCRIPTOR.message_types_by_name['Street'] = _STREET
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Street = _reflection.GeneratedProtocolMessageType('Street', (_message.Message,), {
  'DESCRIPTOR' : _STREET,
  '__module__' : 'proto.street_pb2'
  # @@protoc_insertion_point(class_scope:street.Street)
  })
_sym_db.RegisterMessage(Street)


# @@protoc_insertion_point(module_scope)
