# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: stream.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import template_pb2 as template__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='stream.proto',
  package='types',
  syntax='proto3',
  serialized_pb=_b('\n\x0cstream.proto\x12\x05types\x1a\x0etemplate.proto\"\xfc\x01\n\rStreamRequest\x12\x0c\n\x04kind\x18\x01 \x01(\t\x12(\n\x0binputConfig\x18\x02 \x01(\x0b\x32\x13.types.StreamConfig\x12\x16\n\x0eminProbability\x18\x03 \x01(\t\x12/\n\x0f\x61nalyzeTemplate\x18\x04 \x01(\x0b\x32\x16.types.AnalyzeTemplate\x12\x33\n\x11\x61nonymizeTemplate\x18\x05 \x01(\x0b\x32\x18.types.AnonymizeTemplate\x12\x35\n\x12\x64\x61tabinderTemplate\x18\x06 \x01(\x0b\x32\x19.types.DatabinderTemplateb\x06proto3')
  ,
  dependencies=[template__pb2.DESCRIPTOR,])




_STREAMREQUEST = _descriptor.Descriptor(
  name='StreamRequest',
  full_name='types.StreamRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='kind', full_name='types.StreamRequest.kind', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='inputConfig', full_name='types.StreamRequest.inputConfig', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='minProbability', full_name='types.StreamRequest.minProbability', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='analyzeTemplate', full_name='types.StreamRequest.analyzeTemplate', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='anonymizeTemplate', full_name='types.StreamRequest.anonymizeTemplate', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='databinderTemplate', full_name='types.StreamRequest.databinderTemplate', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=40,
  serialized_end=292,
)

_STREAMREQUEST.fields_by_name['inputConfig'].message_type = template__pb2._STREAMCONFIG
_STREAMREQUEST.fields_by_name['analyzeTemplate'].message_type = template__pb2._ANALYZETEMPLATE
_STREAMREQUEST.fields_by_name['anonymizeTemplate'].message_type = template__pb2._ANONYMIZETEMPLATE
_STREAMREQUEST.fields_by_name['databinderTemplate'].message_type = template__pb2._DATABINDERTEMPLATE
DESCRIPTOR.message_types_by_name['StreamRequest'] = _STREAMREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

StreamRequest = _reflection.GeneratedProtocolMessageType('StreamRequest', (_message.Message,), dict(
  DESCRIPTOR = _STREAMREQUEST,
  __module__ = 'stream_pb2'
  # @@protoc_insertion_point(class_scope:types.StreamRequest)
  ))
_sym_db.RegisterMessage(StreamRequest)


# @@protoc_insertion_point(module_scope)
