/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.27.0
// source: proto/greeting_service.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as proto_greeting_service_pb from '../proto/greeting_service_pb'; // proto import: "proto/greeting_service.proto"


export class ImageServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorGenerateGreeting = new grpcWeb.MethodDescriptor(
    '/ImageService/GenerateGreeting',
    grpcWeb.MethodType.UNARY,
    proto_greeting_service_pb.GreetingGenerationRequest,
    proto_greeting_service_pb.GreetingGenerationResponse,
    (request: proto_greeting_service_pb.GreetingGenerationRequest) => {
      return request.serializeBinary();
    },
    proto_greeting_service_pb.GreetingGenerationResponse.deserializeBinary
  );

  generateGreeting(
    request: proto_greeting_service_pb.GreetingGenerationRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_greeting_service_pb.GreetingGenerationResponse>;

  generateGreeting(
    request: proto_greeting_service_pb.GreetingGenerationRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_greeting_service_pb.GreetingGenerationResponse) => void): grpcWeb.ClientReadableStream<proto_greeting_service_pb.GreetingGenerationResponse>;

  generateGreeting(
    request: proto_greeting_service_pb.GreetingGenerationRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_greeting_service_pb.GreetingGenerationResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/ImageService/GenerateGreeting',
        request,
        metadata || {},
        this.methodDescriptorGenerateGreeting,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/ImageService/GenerateGreeting',
    request,
    metadata || {},
    this.methodDescriptorGenerateGreeting);
  }

}
