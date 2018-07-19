# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: template.proto

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


DESCRIPTOR = _descriptor.FileDescriptor(
  name='template.proto',
  package='types',
  syntax='proto3',
  serialized_pb=_b('\n\x0etemplate.proto\x12\x05types\x1a\x0c\x63ommon.proto\"4\n\x0f\x41nalyzeTemplate\x12!\n\x06\x66ields\x18\x01 \x03(\x0b\x32\x11.types.FieldTypes\"\xb7\x01\n\x11\x41nonymizeTemplate\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0b\x64isplayName\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x12\n\ncreateTime\x18\x04 \x01(\t\x12\x14\n\x0cmodifiedTime\x18\x05 \x01(\t\x12@\n\x18\x66ieldTypeTransformations\x18\x06 \x03(\x0b\x32\x1e.types.FieldTypeTransformation\"k\n\x17\x46ieldTypeTransformation\x12!\n\x06\x66ields\x18\x01 \x03(\x0b\x32\x11.types.FieldTypes\x12-\n\x0etransformation\x18\x02 \x01(\x0b\x32\x15.types.Transformation\"v\n\x0eTransformation\x12\x10\n\x08newValue\x18\x01 \x01(\t\x12)\n\x0creplaceValue\x18\x02 \x01(\x0b\x32\x13.types.ReplaceValue\x12\'\n\x0bredactValue\x18\x03 \x01(\x0b\x32\x12.types.RedactValue\" \n\x0cReplaceValue\x12\x10\n\x08newValue\x18\x01 \x01(\t\"\r\n\x0bRedactValue\"K\n\nDatabinder\x12\x10\n\x08\x62indType\x18\x01 \x01(\t\x12\x18\n\x10\x63onnectionString\x18\x02 \x01(\t\x12\x11\n\ttableName\x18\x03 \x01(\t\";\n\x12\x44\x61tabinderTemplate\x12%\n\ndatabinder\x18\x01 \x03(\x0b\x32\x11.types.Databinder\"S\n\x11\x42lobStorageConfig\x12\x13\n\x0b\x61\x63\x63ountName\x18\x01 \x01(\t\x12\x12\n\naccountKey\x18\x02 \x01(\t\x12\x15\n\rcontainerName\x18\x03 \x01(\t\"S\n\x08S3Config\x12\x10\n\x08\x61\x63\x63\x65ssId\x18\x01 \x01(\t\x12\x11\n\taccessKey\x18\x02 \x01(\t\x12\x0e\n\x06region\x18\x03 \x01(\t\x12\x12\n\nbucketName\x18\x04 \x01(\t\"e\n\x0bInputConfig\x12\x33\n\x11\x62lobStorageConfig\x18\x01 \x01(\x0b\x32\x18.types.BlobStorageConfig\x12!\n\x08s3Config\x18\x02 \x01(\x0b\x32\x0f.types.S3Config\"\xfd\x01\n\x0fScannerTemplate\x12\x0c\n\x04kind\x18\x01 \x01(\t\x12\'\n\x0binputConfig\x18\x02 \x01(\x0b\x32\x12.types.InputConfig\x12\x16\n\x0eminProbability\x18\x03 \x01(\t\x12/\n\x0f\x61nalyzeTemplate\x18\x04 \x01(\x0b\x32\x16.types.AnalyzeTemplate\x12\x33\n\x11\x61nonymizeTemplate\x18\x05 \x01(\x0b\x32\x18.types.AnonymizeTemplate\x12\x35\n\x12\x64\x61tabinderTemplate\x18\x06 \x01(\x0b\x32\x19.types.DatabinderTemplateb\x06proto3')
  ,
  dependencies=[common__pb2.DESCRIPTOR,])




_ANALYZETEMPLATE = _descriptor.Descriptor(
  name='AnalyzeTemplate',
  full_name='types.AnalyzeTemplate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='fields', full_name='types.AnalyzeTemplate.fields', index=0,
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
  serialized_start=39,
  serialized_end=91,
)


_ANONYMIZETEMPLATE = _descriptor.Descriptor(
  name='AnonymizeTemplate',
  full_name='types.AnonymizeTemplate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='types.AnonymizeTemplate.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='displayName', full_name='types.AnonymizeTemplate.displayName', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='types.AnonymizeTemplate.description', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='createTime', full_name='types.AnonymizeTemplate.createTime', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='modifiedTime', full_name='types.AnonymizeTemplate.modifiedTime', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fieldTypeTransformations', full_name='types.AnonymizeTemplate.fieldTypeTransformations', index=5,
      number=6, type=11, cpp_type=10, label=3,
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
  serialized_start=94,
  serialized_end=277,
)


_FIELDTYPETRANSFORMATION = _descriptor.Descriptor(
  name='FieldTypeTransformation',
  full_name='types.FieldTypeTransformation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='fields', full_name='types.FieldTypeTransformation.fields', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='transformation', full_name='types.FieldTypeTransformation.transformation', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_start=279,
  serialized_end=386,
)


_TRANSFORMATION = _descriptor.Descriptor(
  name='Transformation',
  full_name='types.Transformation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='newValue', full_name='types.Transformation.newValue', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='replaceValue', full_name='types.Transformation.replaceValue', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='redactValue', full_name='types.Transformation.redactValue', index=2,
      number=3, type=11, cpp_type=10, label=1,
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
  serialized_start=388,
  serialized_end=506,
)


_REPLACEVALUE = _descriptor.Descriptor(
  name='ReplaceValue',
  full_name='types.ReplaceValue',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='newValue', full_name='types.ReplaceValue.newValue', index=0,
      number=1, type=9, cpp_type=9, label=1,
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
  serialized_start=508,
  serialized_end=540,
)


_REDACTVALUE = _descriptor.Descriptor(
  name='RedactValue',
  full_name='types.RedactValue',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
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
  serialized_start=542,
  serialized_end=555,
)


_DATABINDER = _descriptor.Descriptor(
  name='Databinder',
  full_name='types.Databinder',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='bindType', full_name='types.Databinder.bindType', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='connectionString', full_name='types.Databinder.connectionString', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tableName', full_name='types.Databinder.tableName', index=2,
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
  serialized_start=557,
  serialized_end=632,
)


_DATABINDERTEMPLATE = _descriptor.Descriptor(
  name='DatabinderTemplate',
  full_name='types.DatabinderTemplate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='databinder', full_name='types.DatabinderTemplate.databinder', index=0,
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
  serialized_start=634,
  serialized_end=693,
)


_BLOBSTORAGECONFIG = _descriptor.Descriptor(
  name='BlobStorageConfig',
  full_name='types.BlobStorageConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='accountName', full_name='types.BlobStorageConfig.accountName', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='accountKey', full_name='types.BlobStorageConfig.accountKey', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='containerName', full_name='types.BlobStorageConfig.containerName', index=2,
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
  serialized_start=695,
  serialized_end=778,
)


_S3CONFIG = _descriptor.Descriptor(
  name='S3Config',
  full_name='types.S3Config',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='accessId', full_name='types.S3Config.accessId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='accessKey', full_name='types.S3Config.accessKey', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='region', full_name='types.S3Config.region', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bucketName', full_name='types.S3Config.bucketName', index=3,
      number=4, type=9, cpp_type=9, label=1,
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
  serialized_start=780,
  serialized_end=863,
)


_INPUTCONFIG = _descriptor.Descriptor(
  name='InputConfig',
  full_name='types.InputConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='blobStorageConfig', full_name='types.InputConfig.blobStorageConfig', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='s3Config', full_name='types.InputConfig.s3Config', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_start=865,
  serialized_end=966,
)


_SCANNERTEMPLATE = _descriptor.Descriptor(
  name='ScannerTemplate',
  full_name='types.ScannerTemplate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='kind', full_name='types.ScannerTemplate.kind', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='inputConfig', full_name='types.ScannerTemplate.inputConfig', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='minProbability', full_name='types.ScannerTemplate.minProbability', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='analyzeTemplate', full_name='types.ScannerTemplate.analyzeTemplate', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='anonymizeTemplate', full_name='types.ScannerTemplate.anonymizeTemplate', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='databinderTemplate', full_name='types.ScannerTemplate.databinderTemplate', index=5,
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
  serialized_start=969,
  serialized_end=1222,
)

_ANALYZETEMPLATE.fields_by_name['fields'].message_type = common__pb2._FIELDTYPES
_ANONYMIZETEMPLATE.fields_by_name['fieldTypeTransformations'].message_type = _FIELDTYPETRANSFORMATION
_FIELDTYPETRANSFORMATION.fields_by_name['fields'].message_type = common__pb2._FIELDTYPES
_FIELDTYPETRANSFORMATION.fields_by_name['transformation'].message_type = _TRANSFORMATION
_TRANSFORMATION.fields_by_name['replaceValue'].message_type = _REPLACEVALUE
_TRANSFORMATION.fields_by_name['redactValue'].message_type = _REDACTVALUE
_DATABINDERTEMPLATE.fields_by_name['databinder'].message_type = _DATABINDER
_INPUTCONFIG.fields_by_name['blobStorageConfig'].message_type = _BLOBSTORAGECONFIG
_INPUTCONFIG.fields_by_name['s3Config'].message_type = _S3CONFIG
_SCANNERTEMPLATE.fields_by_name['inputConfig'].message_type = _INPUTCONFIG
_SCANNERTEMPLATE.fields_by_name['analyzeTemplate'].message_type = _ANALYZETEMPLATE
_SCANNERTEMPLATE.fields_by_name['anonymizeTemplate'].message_type = _ANONYMIZETEMPLATE
_SCANNERTEMPLATE.fields_by_name['databinderTemplate'].message_type = _DATABINDERTEMPLATE
DESCRIPTOR.message_types_by_name['AnalyzeTemplate'] = _ANALYZETEMPLATE
DESCRIPTOR.message_types_by_name['AnonymizeTemplate'] = _ANONYMIZETEMPLATE
DESCRIPTOR.message_types_by_name['FieldTypeTransformation'] = _FIELDTYPETRANSFORMATION
DESCRIPTOR.message_types_by_name['Transformation'] = _TRANSFORMATION
DESCRIPTOR.message_types_by_name['ReplaceValue'] = _REPLACEVALUE
DESCRIPTOR.message_types_by_name['RedactValue'] = _REDACTVALUE
DESCRIPTOR.message_types_by_name['Databinder'] = _DATABINDER
DESCRIPTOR.message_types_by_name['DatabinderTemplate'] = _DATABINDERTEMPLATE
DESCRIPTOR.message_types_by_name['BlobStorageConfig'] = _BLOBSTORAGECONFIG
DESCRIPTOR.message_types_by_name['S3Config'] = _S3CONFIG
DESCRIPTOR.message_types_by_name['InputConfig'] = _INPUTCONFIG
DESCRIPTOR.message_types_by_name['ScannerTemplate'] = _SCANNERTEMPLATE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AnalyzeTemplate = _reflection.GeneratedProtocolMessageType('AnalyzeTemplate', (_message.Message,), dict(
  DESCRIPTOR = _ANALYZETEMPLATE,
  __module__ = 'template_pb2'
  # @@protoc_insertion_point(class_scope:types.AnalyzeTemplate)
  ))
_sym_db.RegisterMessage(AnalyzeTemplate)

AnonymizeTemplate = _reflection.GeneratedProtocolMessageType('AnonymizeTemplate', (_message.Message,), dict(
  DESCRIPTOR = _ANONYMIZETEMPLATE,
  __module__ = 'template_pb2'
  # @@protoc_insertion_point(class_scope:types.AnonymizeTemplate)
  ))
_sym_db.RegisterMessage(AnonymizeTemplate)

FieldTypeTransformation = _reflection.GeneratedProtocolMessageType('FieldTypeTransformation', (_message.Message,), dict(
  DESCRIPTOR = _FIELDTYPETRANSFORMATION,
  __module__ = 'template_pb2'
  # @@protoc_insertion_point(class_scope:types.FieldTypeTransformation)
  ))
_sym_db.RegisterMessage(FieldTypeTransformation)

Transformation = _reflection.GeneratedProtocolMessageType('Transformation', (_message.Message,), dict(
  DESCRIPTOR = _TRANSFORMATION,
  __module__ = 'template_pb2'
  # @@protoc_insertion_point(class_scope:types.Transformation)
  ))
_sym_db.RegisterMessage(Transformation)

ReplaceValue = _reflection.GeneratedProtocolMessageType('ReplaceValue', (_message.Message,), dict(
  DESCRIPTOR = _REPLACEVALUE,
  __module__ = 'template_pb2'
  # @@protoc_insertion_point(class_scope:types.ReplaceValue)
  ))
_sym_db.RegisterMessage(ReplaceValue)

RedactValue = _reflection.GeneratedProtocolMessageType('RedactValue', (_message.Message,), dict(
  DESCRIPTOR = _REDACTVALUE,
  __module__ = 'template_pb2'
  # @@protoc_insertion_point(class_scope:types.RedactValue)
  ))
_sym_db.RegisterMessage(RedactValue)

Databinder = _reflection.GeneratedProtocolMessageType('Databinder', (_message.Message,), dict(
  DESCRIPTOR = _DATABINDER,
  __module__ = 'template_pb2'
  # @@protoc_insertion_point(class_scope:types.Databinder)
  ))
_sym_db.RegisterMessage(Databinder)

DatabinderTemplate = _reflection.GeneratedProtocolMessageType('DatabinderTemplate', (_message.Message,), dict(
  DESCRIPTOR = _DATABINDERTEMPLATE,
  __module__ = 'template_pb2'
  # @@protoc_insertion_point(class_scope:types.DatabinderTemplate)
  ))
_sym_db.RegisterMessage(DatabinderTemplate)

BlobStorageConfig = _reflection.GeneratedProtocolMessageType('BlobStorageConfig', (_message.Message,), dict(
  DESCRIPTOR = _BLOBSTORAGECONFIG,
  __module__ = 'template_pb2'
  # @@protoc_insertion_point(class_scope:types.BlobStorageConfig)
  ))
_sym_db.RegisterMessage(BlobStorageConfig)

S3Config = _reflection.GeneratedProtocolMessageType('S3Config', (_message.Message,), dict(
  DESCRIPTOR = _S3CONFIG,
  __module__ = 'template_pb2'
  # @@protoc_insertion_point(class_scope:types.S3Config)
  ))
_sym_db.RegisterMessage(S3Config)

InputConfig = _reflection.GeneratedProtocolMessageType('InputConfig', (_message.Message,), dict(
  DESCRIPTOR = _INPUTCONFIG,
  __module__ = 'template_pb2'
  # @@protoc_insertion_point(class_scope:types.InputConfig)
  ))
_sym_db.RegisterMessage(InputConfig)

ScannerTemplate = _reflection.GeneratedProtocolMessageType('ScannerTemplate', (_message.Message,), dict(
  DESCRIPTOR = _SCANNERTEMPLATE,
  __module__ = 'template_pb2'
  # @@protoc_insertion_point(class_scope:types.ScannerTemplate)
  ))
_sym_db.RegisterMessage(ScannerTemplate)


# @@protoc_insertion_point(module_scope)
