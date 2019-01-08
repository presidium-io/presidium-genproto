# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: anonymizeJson.proto

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
  name='anonymizeJson.proto',
  package='types',
  syntax='proto3',
  serialized_pb=_b('\n\x13\x61nonymizeJson.proto\x12\x05types\x1a\x0etemplate.proto\"\x92\x02\n\x17\x41nonymizeJsonApiRequest\x12\x0c\n\x04json\x18\x01 \x01(\t\x12\x14\n\x0cjsonSchemaId\x18\x02 \x01(\t\x12\x19\n\x11\x61nalyzeTemplateId\x18\x03 \x01(\t\x12\x1b\n\x13\x61nonymizeTemplateId\x18\x04 \x01(\t\x12\x35\n\x12jsonSchemaTemplate\x18\x05 \x01(\x0b\x32\x19.types.JsonSchemaTemplate\x12/\n\x0f\x61nalyzeTemplate\x18\x06 \x01(\x0b\x32\x16.types.AnalyzeTemplate\x12\x33\n\x11\x61nonymizeTemplate\x18\x07 \x01(\x0b\x32\x18.types.AnonymizeTemplate\"\xb9\x01\n\x14\x41nonymizeJsonRequest\x12\x0c\n\x04json\x18\x01 \x01(\t\x12-\n\njsonSchema\x18\x02 \x01(\x0b\x32\x19.types.JsonSchemaTemplate\x12/\n\x0f\x61nalyzeTemplate\x18\x03 \x01(\x0b\x32\x16.types.AnalyzeTemplate\x12\x33\n\x11\x61nonymizeTemplate\x18\x04 \x01(\x0b\x32\x18.types.AnonymizeTemplateb\x06proto3')
  ,
  dependencies=[template__pb2.DESCRIPTOR,])




_ANONYMIZEJSONAPIREQUEST = _descriptor.Descriptor(
  name='AnonymizeJsonApiRequest',
  full_name='types.AnonymizeJsonApiRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='json', full_name='types.AnonymizeJsonApiRequest.json', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='jsonSchemaId', full_name='types.AnonymizeJsonApiRequest.jsonSchemaId', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='analyzeTemplateId', full_name='types.AnonymizeJsonApiRequest.analyzeTemplateId', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='anonymizeTemplateId', full_name='types.AnonymizeJsonApiRequest.anonymizeTemplateId', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='jsonSchemaTemplate', full_name='types.AnonymizeJsonApiRequest.jsonSchemaTemplate', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='analyzeTemplate', full_name='types.AnonymizeJsonApiRequest.analyzeTemplate', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='anonymizeTemplate', full_name='types.AnonymizeJsonApiRequest.anonymizeTemplate', index=6,
      number=7, type=11, cpp_type=10, label=1,
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
  serialized_start=47,
  serialized_end=321,
)


_ANONYMIZEJSONREQUEST = _descriptor.Descriptor(
  name='AnonymizeJsonRequest',
  full_name='types.AnonymizeJsonRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='json', full_name='types.AnonymizeJsonRequest.json', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='jsonSchema', full_name='types.AnonymizeJsonRequest.jsonSchema', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='analyzeTemplate', full_name='types.AnonymizeJsonRequest.analyzeTemplate', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='anonymizeTemplate', full_name='types.AnonymizeJsonRequest.anonymizeTemplate', index=3,
      number=4, type=11, cpp_type=10, label=1,
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
  serialized_start=324,
  serialized_end=509,
)

_ANONYMIZEJSONAPIREQUEST.fields_by_name['jsonSchemaTemplate'].message_type = template__pb2._JSONSCHEMATEMPLATE
_ANONYMIZEJSONAPIREQUEST.fields_by_name['analyzeTemplate'].message_type = template__pb2._ANALYZETEMPLATE
_ANONYMIZEJSONAPIREQUEST.fields_by_name['anonymizeTemplate'].message_type = template__pb2._ANONYMIZETEMPLATE
_ANONYMIZEJSONREQUEST.fields_by_name['jsonSchema'].message_type = template__pb2._JSONSCHEMATEMPLATE
_ANONYMIZEJSONREQUEST.fields_by_name['analyzeTemplate'].message_type = template__pb2._ANALYZETEMPLATE
_ANONYMIZEJSONREQUEST.fields_by_name['anonymizeTemplate'].message_type = template__pb2._ANONYMIZETEMPLATE
DESCRIPTOR.message_types_by_name['AnonymizeJsonApiRequest'] = _ANONYMIZEJSONAPIREQUEST
DESCRIPTOR.message_types_by_name['AnonymizeJsonRequest'] = _ANONYMIZEJSONREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AnonymizeJsonApiRequest = _reflection.GeneratedProtocolMessageType('AnonymizeJsonApiRequest', (_message.Message,), dict(
  DESCRIPTOR = _ANONYMIZEJSONAPIREQUEST,
  __module__ = 'anonymizeJson_pb2'
  # @@protoc_insertion_point(class_scope:types.AnonymizeJsonApiRequest)
  ))
_sym_db.RegisterMessage(AnonymizeJsonApiRequest)

AnonymizeJsonRequest = _reflection.GeneratedProtocolMessageType('AnonymizeJsonRequest', (_message.Message,), dict(
  DESCRIPTOR = _ANONYMIZEJSONREQUEST,
  __module__ = 'anonymizeJson_pb2'
  # @@protoc_insertion_point(class_scope:types.AnonymizeJsonRequest)
  ))
_sym_db.RegisterMessage(AnonymizeJsonRequest)


# @@protoc_insertion_point(module_scope)
