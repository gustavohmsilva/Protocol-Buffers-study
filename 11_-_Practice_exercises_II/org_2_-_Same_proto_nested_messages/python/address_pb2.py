# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: address.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='address.proto',
  package='eleven.city',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\raddress.proto\x12\x0b\x65leven.city\"\xe5\x01\n\x08\x42uilding\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0e\n\x06number\x18\x02 \x01(\t\x12,\n\x06street\x18\x03 \x01(\x0b\x32\x1c.eleven.city.Building.Street\x1a\x8c\x01\n\x06Street\x12\x13\n\x0bstreet_name\x18\x01 \x01(\t\x12/\n\x04\x63ity\x18\x02 \x01(\x0b\x32!.eleven.city.Building.Street.City\x1a<\n\x04\x43ity\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x10\n\x08zip_code\x18\x02 \x01(\t\x12\x14\n\x0c\x63ountry_code\x18\x03 \x01(\tb\x06proto3'
)




_BUILDING_STREET_CITY = _descriptor.Descriptor(
  name='City',
  full_name='eleven.city.Building.Street.City',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='eleven.city.Building.Street.City.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='zip_code', full_name='eleven.city.Building.Street.City.zip_code', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='country_code', full_name='eleven.city.Building.Street.City.country_code', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=200,
  serialized_end=260,
)

_BUILDING_STREET = _descriptor.Descriptor(
  name='Street',
  full_name='eleven.city.Building.Street',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='street_name', full_name='eleven.city.Building.Street.street_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='city', full_name='eleven.city.Building.Street.city', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_BUILDING_STREET_CITY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=120,
  serialized_end=260,
)

_BUILDING = _descriptor.Descriptor(
  name='Building',
  full_name='eleven.city.Building',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='eleven.city.Building.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='number', full_name='eleven.city.Building.number', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='street', full_name='eleven.city.Building.street', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_BUILDING_STREET, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=31,
  serialized_end=260,
)

_BUILDING_STREET_CITY.containing_type = _BUILDING_STREET
_BUILDING_STREET.fields_by_name['city'].message_type = _BUILDING_STREET_CITY
_BUILDING_STREET.containing_type = _BUILDING
_BUILDING.fields_by_name['street'].message_type = _BUILDING_STREET
DESCRIPTOR.message_types_by_name['Building'] = _BUILDING
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Building = _reflection.GeneratedProtocolMessageType('Building', (_message.Message,), {

  'Street' : _reflection.GeneratedProtocolMessageType('Street', (_message.Message,), {

    'City' : _reflection.GeneratedProtocolMessageType('City', (_message.Message,), {
      'DESCRIPTOR' : _BUILDING_STREET_CITY,
      '__module__' : 'address_pb2'
      # @@protoc_insertion_point(class_scope:eleven.city.Building.Street.City)
      })
    ,
    'DESCRIPTOR' : _BUILDING_STREET,
    '__module__' : 'address_pb2'
    # @@protoc_insertion_point(class_scope:eleven.city.Building.Street)
    })
  ,
  'DESCRIPTOR' : _BUILDING,
  '__module__' : 'address_pb2'
  # @@protoc_insertion_point(class_scope:eleven.city.Building)
  })
_sym_db.RegisterMessage(Building)
_sym_db.RegisterMessage(Building.Street)
_sym_db.RegisterMessage(Building.Street.City)


# @@protoc_insertion_point(module_scope)