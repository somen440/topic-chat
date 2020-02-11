/**
 * @fileoverview gRPC-Web generated client stub for topicchat
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {
  AddTopicRequest,
  ChatMessage,
  CreateRoomRequest,
  DeleteTopicRequest,
  Empty,
  GetTopicRequest,
  GetUserAllResponse,
  GetUserRequest,
  JoinRequest,
  JoinRoomRequest,
  JoinRoomResponse,
  ListTopicsResponse,
  LoggedInRequest,
  RecvMessageRequest,
  SendMessageRequest,
  SignoutRequest,
  Topic,
  UpdateTopicRequest,
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

  methodInfoGetUser = new grpcWeb.AbstractClientBase.MethodInfo(
    User,
    (request: GetUserRequest) => {
      return request.serializeBinary();
    },
    User.deserializeBinary
  );

  getUser(
    request: GetUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: User) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/topicchat.AuthService/GetUser',
      request,
      metadata || {},
      this.methodInfoGetUser,
      callback);
  }

  methodInfoGetUserAll = new grpcWeb.AbstractClientBase.MethodInfo(
    GetUserAllResponse,
    (request: Empty) => {
      return request.serializeBinary();
    },
    GetUserAllResponse.deserializeBinary
  );

  getUserAll(
    request: Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: GetUserAllResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/topicchat.AuthService/GetUserAll',
      request,
      metadata || {},
      this.methodInfoGetUserAll,
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
    (request: RecvMessageRequest) => {
      return request.serializeBinary();
    },
    ChatMessage.deserializeBinary
  );

  recvMessage(
    request: RecvMessageRequest,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/topicchat.ChatService/RecvMessage',
      request,
      metadata || {},
      this.methodInfoRecvMessage);
  }

  methodInfoSendMessage = new grpcWeb.AbstractClientBase.MethodInfo(
    Empty,
    (request: SendMessageRequest) => {
      return request.serializeBinary();
    },
    Empty.deserializeBinary
  );

  sendMessage(
    request: SendMessageRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Empty) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/topicchat.ChatService/SendMessage',
      request,
      metadata || {},
      this.methodInfoSendMessage,
      callback);
  }

  methodInfoCreateRoom = new grpcWeb.AbstractClientBase.MethodInfo(
    Empty,
    (request: CreateRoomRequest) => {
      return request.serializeBinary();
    },
    Empty.deserializeBinary
  );

  createRoom(
    request: CreateRoomRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Empty) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/topicchat.ChatService/CreateRoom',
      request,
      metadata || {},
      this.methodInfoCreateRoom,
      callback);
  }

  methodInfoJoinRoom = new grpcWeb.AbstractClientBase.MethodInfo(
    JoinRoomResponse,
    (request: JoinRoomRequest) => {
      return request.serializeBinary();
    },
    JoinRoomResponse.deserializeBinary
  );

  joinRoom(
    request: JoinRoomRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: JoinRoomResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/topicchat.ChatService/JoinRoom',
      request,
      metadata || {},
      this.methodInfoJoinRoom,
      callback);
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

  methodInfoAddTopic = new grpcWeb.AbstractClientBase.MethodInfo(
    Topic,
    (request: AddTopicRequest) => {
      return request.serializeBinary();
    },
    Topic.deserializeBinary
  );

  addTopic(
    request: AddTopicRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Topic) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/topicchat.TopicCatalogService/AddTopic',
      request,
      metadata || {},
      this.methodInfoAddTopic,
      callback);
  }

  methodInfoUpdateTopic = new grpcWeb.AbstractClientBase.MethodInfo(
    Empty,
    (request: UpdateTopicRequest) => {
      return request.serializeBinary();
    },
    Empty.deserializeBinary
  );

  updateTopic(
    request: UpdateTopicRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Empty) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/topicchat.TopicCatalogService/UpdateTopic',
      request,
      metadata || {},
      this.methodInfoUpdateTopic,
      callback);
  }

  methodInfoDeleteTopic = new grpcWeb.AbstractClientBase.MethodInfo(
    Empty,
    (request: DeleteTopicRequest) => {
      return request.serializeBinary();
    },
    Empty.deserializeBinary
  );

  deleteTopic(
    request: DeleteTopicRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Empty) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/topicchat.TopicCatalogService/DeleteTopic',
      request,
      metadata || {},
      this.methodInfoDeleteTopic,
      callback);
  }

}

