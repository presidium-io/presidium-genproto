# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: analyze.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import common_pb2 as common__pb2
import template_pb2 as template__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='analyze.proto',
  package='types',
  syntax='proto3',
  serialized_pb=_b('\n\ranalyze.proto\x12\x05types\x1a\x0c\x63ommon.proto\x1a\x0etemplate.proto\"<\n\x11\x41nalyzeApiRequest\x12\x0c\n\x04text\x18\x01 \x01(\t\x12\x19\n\x11\x61nalyzeTemplateId\x18\x02 \x01(\t\"g\n\x0e\x41nalyzeRequest\x12\x0c\n\x04text\x18\x01 \x01(\t\x12/\n\x0f\x61nalyzeTemplate\x18\x02 \x01(\x0b\x32\x16.types.AnalyzeTemplate\x12\x16\n\x0eminProbability\x18\x03 \x01(\t\"?\n\x0f\x41nalyzeResponse\x12,\n\x0e\x61nalyzeResults\x18\x01 \x03(\x0b\x32\x14.types.AnalyzeResult2J\n\x0e\x41nalyzeService\x12\x38\n\x05\x41pply\x12\x15.types.AnalyzeRequest\x1a\x16.types.AnalyzeResponse\"\x00\x62\x06proto3')
  ,
  dependencies=[common__pb2.DESCRIPTOR,template__pb2.DESCRIPTOR,])




_ANALYZEAPIREQUEST = _descriptor.Descriptor(
  name='AnalyzeApiRequest',
  full_name='types.AnalyzeApiRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='text', full_name='types.AnalyzeApiRequest.text', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='analyzeTemplateId', full_name='types.AnalyzeApiRequest.analyzeTemplateId', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=54,
  serialized_end=114,
)


_ANALYZEREQUEST = _descriptor.Descriptor(
  name='AnalyzeRequest',
  full_name='types.AnalyzeRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='text', full_name='types.AnalyzeRequest.text', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='analyzeTemplate', full_name='types.AnalyzeRequest.analyzeTemplate', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='minProbability', full_name='types.AnalyzeRequest.minProbability', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=116,
  serialized_end=219,
)


_ANALYZERESPONSE = _descriptor.Descriptor(
  name='AnalyzeResponse',
  full_name='types.AnalyzeResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='analyzeResults', full_name='types.AnalyzeResponse.analyzeResults', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=221,
  serialized_end=284,
)

_ANALYZEREQUEST.fields_by_name['analyzeTemplate'].message_type = template__pb2._ANALYZETEMPLATE
_ANALYZERESPONSE.fields_by_name['analyzeResults'].message_type = common__pb2._ANALYZERESULT
DESCRIPTOR.message_types_by_name['AnalyzeApiRequest'] = _ANALYZEAPIREQUEST
DESCRIPTOR.message_types_by_name['AnalyzeRequest'] = _ANALYZEREQUEST
DESCRIPTOR.message_types_by_name['AnalyzeResponse'] = _ANALYZERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AnalyzeApiRequest = _reflection.GeneratedProtocolMessageType('AnalyzeApiRequest', (_message.Message,), dict(
  DESCRIPTOR = _ANALYZEAPIREQUEST,
  __module__ = 'analyze_pb2'
  # @@protoc_insertion_point(class_scope:types.AnalyzeApiRequest)
  ))
_sym_db.RegisterMessage(AnalyzeApiRequest)

AnalyzeRequest = _reflection.GeneratedProtocolMessageType('AnalyzeRequest', (_message.Message,), dict(
  DESCRIPTOR = _ANALYZEREQUEST,
  __module__ = 'analyze_pb2'
  # @@protoc_insertion_point(class_scope:types.AnalyzeRequest)
  ))
_sym_db.RegisterMessage(AnalyzeRequest)

AnalyzeResponse = _reflection.GeneratedProtocolMessageType('AnalyzeResponse', (_message.Message,), dict(
  DESCRIPTOR = _ANALYZERESPONSE,
  __module__ = 'analyze_pb2'
  # @@protoc_insertion_point(class_scope:types.AnalyzeResponse)
  ))
_sym_db.RegisterMessage(AnalyzeResponse)



_ANALYZESERVICE = _descriptor.ServiceDescriptor(
  name='AnalyzeService',
  full_name='types.AnalyzeService',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=286,
  serialized_end=360,
  methods=[
  _descriptor.MethodDescriptor(
    name='Apply',
    full_name='types.AnalyzeService.Apply',
    index=0,
    containing_service=None,
    input_type=_ANALYZEREQUEST,
    output_type=_ANALYZERESPONSE,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_ANALYZESERVICE)

DESCRIPTOR.services_by_name['AnalyzeService'] = _ANALYZESERVICE

# @@protoc_insertion_point(module_scope)
