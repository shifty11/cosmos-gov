///
//  Generated code. Do not modify.
//  source: auth_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'auth_service.pb.dart' as $0;
export 'auth_service.pb.dart';

class AuthServiceClient extends $grpc.Client {
  static final _$tokenLogin =
      $grpc.ClientMethod<$0.TokenLoginRequest, $0.TokenLoginResponse>(
          '/cosmosgov_grpc.AuthService/TokenLogin',
          ($0.TokenLoginRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.TokenLoginResponse.fromBuffer(value));

  AuthServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.TokenLoginResponse> tokenLogin(
      $0.TokenLoginRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$tokenLogin, request, options: options);
  }
}

abstract class AuthServiceBase extends $grpc.Service {
  $core.String get $name => 'cosmosgov_grpc.AuthService';

  AuthServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.TokenLoginRequest, $0.TokenLoginResponse>(
        'TokenLogin',
        tokenLogin_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.TokenLoginRequest.fromBuffer(value),
        ($0.TokenLoginResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.TokenLoginResponse> tokenLogin_Pre($grpc.ServiceCall call,
      $async.Future<$0.TokenLoginRequest> request) async {
    return tokenLogin(call, await request);
  }

  $async.Future<$0.TokenLoginResponse> tokenLogin(
      $grpc.ServiceCall call, $0.TokenLoginRequest request);
}
