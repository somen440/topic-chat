/**
 * @fileoverview gRPC-Web generated client stub for topicchat
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {
  ChatMessage,
  Empty,
  GetTopicRequest,
  JoinRequest,
  ListTopicsResponse,
  LoggedInRequest,
  RoomJoinRequest,
  SignoutRequest,
  Topic,
  User} from './topicchat_pb';

export class AuthServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoJoin = new grpcWeb.AbstractClientBase.MethodInfo(
    User,
    (request: JoinRequest) => {
      return request.serializeBinary();
    },
    User.deserializeBinary
  );

  join(
    request: JoinRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: User) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/topicchat.AuthService/Join',
      request,
      metadata || {},
      this.methodInfoJoin,
      callback);
  }

  methodInfoLoggedIn = new grpcWeb.AbstractClientBase.MethodInfo(
    User,
    (request: LoggedInRequest) => {
      return request.serializeBinary();
    },
    User.deserializeBinary
  );

  loggedIn(
    request: LoggedInRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: User) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/topicchat.AuthService/LoggedIn',
      request,
      metadata || {},
      this.methodInfoLoggedIn,
      callback);
  }

  methodInfoSignout = new grpcWeb.AbstractClientBase.MethodInfo(
    Empty,
    (request: SignoutRequest) => {
      return request.serializeBinary();
    },
    Empty.deserializeBinary
  );

  signout(
    request: SignoutRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Empty) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/topicchat.AuthService/Signout',
      request,
      metadata || {},
      this.methodInfoSignout,
      callback);
  }

}

export class RoomServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoJoin = new grpcWeb.AbstractClientBase.MethodInfo(
    Empty,
    (request: RoomJoinRequest) => {
      return request.serializeBinary();
    },
    Empty.deserializeBinary
  );

  join(
    request: RoomJoinRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Empty) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/topicchat.RoomService/Join',
      request,
      metadata || {},
      this.methodInfoJoin,
      callback);
  }

  methodInfoLeft = new grpcWeb.AbstractClientBase.MethodInfo(
    Empty,
    (request: Empty) => {
      return request.serializeBinary();
    },
    Empty.deserializeBinary
  );

  left(
    request: Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Empty) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/topicchat.RoomService/Left',
      request,
      metadata || {},
      this.methodInfoLeft,
      callback);
  }

}

export class ChatServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoRecvMessage = new grpcWeb.AbstractClientBase.MethodInfo(
    ChatMessage,
    (request: Empty) => {
      return request.serializeBinary();
    },
    ChatMessage.deserializeBinary
  );

  recvMessage(
    request: Empty,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/topicchat.ChatService/RecvMessage',
      request,
      metadata || {},
      this.methodInfoRecvMessage);
  }

}

export class TopicCatalogServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoListTopics = new grpcWeb.AbstractClientBase.MethodInfo(
    ListTopicsResponse,
    (request: Empty) => {
      return request.serializeBinary();
    },
    ListTopicsResponse.deserializeBinary
  );

  listTopics(
    request: Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ListTopicsResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/topicchat.TopicCatalogService/ListTopics',
      request,
      metadata || {},
      this.methodInfoListTopics,
      callback);
  }

  methodInfoGetTopic = new grpcWeb.AbstractClientBase.MethodInfo(
    Topic,
    (request: GetTopicRequest) => {
      return request.serializeBinary();
    },
    Topic.deserializeBinary
  );

  getTopic(
    request: GetTopicRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Topic) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/topicchat.TopicCatalogService/GetTopic',
      request,
      metadata || {},
      this.methodInfoGetTopic,
      callback);
  }

}

