///
//  Generated code. Do not modify.
//  source: auth_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'auth_service.pbenum.dart';

export 'auth_service.pbenum.dart';

class TokenLoginRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'TokenLoginRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmosgov_grpc'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'token')
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chatId', protoName: 'chatId')
    ..e<TokenLoginRequest_Type>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'TYPE', $pb.PbFieldType.OE, protoName: 'TYPE', defaultOrMaker: TokenLoginRequest_Type.TELEGRAM, valueOf: TokenLoginRequest_Type.valueOf, enumValues: TokenLoginRequest_Type.values)
    ..hasRequiredFields = false
  ;

  TokenLoginRequest._() : super();
  factory TokenLoginRequest({
    $core.String? token,
    $fixnum.Int64? chatId,
    TokenLoginRequest_Type? tYPE,
  }) {
    final _result = create();
    if (token != null) {
      _result.token = token;
    }
    if (chatId != null) {
      _result.chatId = chatId;
    }
    if (tYPE != null) {
      _result.tYPE = tYPE;
    }
    return _result;
  }
  factory TokenLoginRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TokenLoginRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TokenLoginRequest clone() => TokenLoginRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TokenLoginRequest copyWith(void Function(TokenLoginRequest) updates) => super.copyWith((message) => updates(message as TokenLoginRequest)) as TokenLoginRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TokenLoginRequest create() => TokenLoginRequest._();
  TokenLoginRequest createEmptyInstance() => create();
  static $pb.PbList<TokenLoginRequest> createRepeated() => $pb.PbList<TokenLoginRequest>();
  @$core.pragma('dart2js:noInline')
  static TokenLoginRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TokenLoginRequest>(create);
  static TokenLoginRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get token => $_getSZ(0);
  @$pb.TagNumber(1)
  set token($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearToken() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get chatId => $_getI64(1);
  @$pb.TagNumber(2)
  set chatId($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasChatId() => $_has(1);
  @$pb.TagNumber(2)
  void clearChatId() => clearField(2);

  @$pb.TagNumber(3)
  TokenLoginRequest_Type get tYPE => $_getN(2);
  @$pb.TagNumber(3)
  set tYPE(TokenLoginRequest_Type v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasTYPE() => $_has(2);
  @$pb.TagNumber(3)
  void clearTYPE() => clearField(3);
}

class TokenLoginResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'TokenLoginResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'cosmosgov_grpc'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'accessToken')
    ..hasRequiredFields = false
  ;

  TokenLoginResponse._() : super();
  factory TokenLoginResponse({
    $core.String? accessToken,
  }) {
    final _result = create();
    if (accessToken != null) {
      _result.accessToken = accessToken;
    }
    return _result;
  }
  factory TokenLoginResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TokenLoginResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TokenLoginResponse clone() => TokenLoginResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TokenLoginResponse copyWith(void Function(TokenLoginResponse) updates) => super.copyWith((message) => updates(message as TokenLoginResponse)) as TokenLoginResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TokenLoginResponse create() => TokenLoginResponse._();
  TokenLoginResponse createEmptyInstance() => create();
  static $pb.PbList<TokenLoginResponse> createRepeated() => $pb.PbList<TokenLoginResponse>();
  @$core.pragma('dart2js:noInline')
  static TokenLoginResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TokenLoginResponse>(create);
  static TokenLoginResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get accessToken => $_getSZ(0);
  @$pb.TagNumber(1)
  set accessToken($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccessToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccessToken() => clearField(1);
}

