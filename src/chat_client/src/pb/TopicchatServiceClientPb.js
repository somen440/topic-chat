/* eslint-disable */
"use strict";
/**
 * @fileoverview gRPC-Web generated client stub for topicchat
 * @enhanceable
 * @public
 */
exports.__esModule = true;
// GENERATED CODE -- DO NOT EDIT!
var grpcWeb = require("grpc-web");
var topicchat_pb_1 = require("./topicchat_pb");
var AuthServiceClient = /** @class */ (function () {
    function AuthServiceClient(hostname, credentials, options) {
        this.methodInfoJoin = new grpcWeb.AbstractClientBase.MethodInfo(topicchat_pb_1.User, function (request) {
            return request.serializeBinary();
        }, topicchat_pb_1.User.deserializeBinary);
        this.methodInfoLoggedIn = new grpcWeb.AbstractClientBase.MethodInfo(topicchat_pb_1.User, function (request) {
            return request.serializeBinary();
        }, topicchat_pb_1.User.deserializeBinary);
        this.methodInfoSignout = new grpcWeb.AbstractClientBase.MethodInfo(topicchat_pb_1.Empty, function (request) {
            return request.serializeBinary();
        }, topicchat_pb_1.Empty.deserializeBinary);
        if (!options)
            options = {};
        if (!credentials)
            credentials = {};
        options['format'] = 'text';
        this.client_ = new grpcWeb.GrpcWebClientBase(options);
        this.hostname_ = hostname;
        this.credentials_ = credentials;
        this.options_ = options;
    }
    AuthServiceClient.prototype.join = function (request, metadata, callback) {
        return this.client_.rpcCall(this.hostname_ +
            '/topicchat.AuthService/Join', request, metadata || {}, this.methodInfoJoin, callback);
    };
    AuthServiceClient.prototype.loggedIn = function (request, metadata, callback) {
        return this.client_.rpcCall(this.hostname_ +
            '/topicchat.AuthService/LoggedIn', request, metadata || {}, this.methodInfoLoggedIn, callback);
    };
    AuthServiceClient.prototype.signout = function (request, metadata, callback) {
        return this.client_.rpcCall(this.hostname_ +
            '/topicchat.AuthService/Signout', request, metadata || {}, this.methodInfoSignout, callback);
    };
    return AuthServiceClient;
}());
exports.AuthServiceClient = AuthServiceClient;
var RoomServiceClient = /** @class */ (function () {
    function RoomServiceClient(hostname, credentials, options) {
        this.methodInfoJoin = new grpcWeb.AbstractClientBase.MethodInfo(topicchat_pb_1.Empty, function (request) {
            return request.serializeBinary();
        }, topicchat_pb_1.Empty.deserializeBinary);
        this.methodInfoLeft = new grpcWeb.AbstractClientBase.MethodInfo(topicchat_pb_1.Empty, function (request) {
            return request.serializeBinary();
        }, topicchat_pb_1.Empty.deserializeBinary);
        if (!options)
            options = {};
        if (!credentials)
            credentials = {};
        options['format'] = 'text';
        this.client_ = new grpcWeb.GrpcWebClientBase(options);
        this.hostname_ = hostname;
        this.credentials_ = credentials;
        this.options_ = options;
    }
    RoomServiceClient.prototype.join = function (request, metadata, callback) {
        return this.client_.rpcCall(this.hostname_ +
            '/topicchat.RoomService/Join', request, metadata || {}, this.methodInfoJoin, callback);
    };
    RoomServiceClient.prototype.left = function (request, metadata, callback) {
        return this.client_.rpcCall(this.hostname_ +
            '/topicchat.RoomService/Left', request, metadata || {}, this.methodInfoLeft, callback);
    };
    return RoomServiceClient;
}());
exports.RoomServiceClient = RoomServiceClient;
var ChatServiceClient = /** @class */ (function () {
    function ChatServiceClient(hostname, credentials, options) {
        this.methodInfoRecvMessage = new grpcWeb.AbstractClientBase.MethodInfo(topicchat_pb_1.ChatMessage, function (request) {
            return request.serializeBinary();
        }, topicchat_pb_1.ChatMessage.deserializeBinary);
        if (!options)
            options = {};
        if (!credentials)
            credentials = {};
        options['format'] = 'text';
        this.client_ = new grpcWeb.GrpcWebClientBase(options);
        this.hostname_ = hostname;
        this.credentials_ = credentials;
        this.options_ = options;
    }
    ChatServiceClient.prototype.recvMessage = function (request, metadata) {
        return this.client_.serverStreaming(this.hostname_ +
            '/topicchat.ChatService/RecvMessage', request, metadata || {}, this.methodInfoRecvMessage);
    };
    return ChatServiceClient;
}());
exports.ChatServiceClient = ChatServiceClient;
var TopicCatalogServiceClient = /** @class */ (function () {
    function TopicCatalogServiceClient(hostname, credentials, options) {
        this.methodInfoListTopics = new grpcWeb.AbstractClientBase.MethodInfo(topicchat_pb_1.ListTopicsResponse, function (request) {
            return request.serializeBinary();
        }, topicchat_pb_1.ListTopicsResponse.deserializeBinary);
        this.methodInfoGetTopic = new grpcWeb.AbstractClientBase.MethodInfo(topicchat_pb_1.Topic, function (request) {
            return request.serializeBinary();
        }, topicchat_pb_1.Topic.deserializeBinary);
        if (!options)
            options = {};
        if (!credentials)
            credentials = {};
        options['format'] = 'text';
        this.client_ = new grpcWeb.GrpcWebClientBase(options);
        this.hostname_ = hostname;
        this.credentials_ = credentials;
        this.options_ = options;
    }
    TopicCatalogServiceClient.prototype.listTopics = function (request, metadata, callback) {
        return this.client_.rpcCall(this.hostname_ +
            '/topicchat.TopicCatalogService/ListTopics', request, metadata || {}, this.methodInfoListTopics, callback);
    };
    TopicCatalogServiceClient.prototype.getTopic = function (request, metadata, callback) {
        return this.client_.rpcCall(this.hostname_ +
            '/topicchat.TopicCatalogService/GetTopic', request, metadata || {}, this.methodInfoGetTopic, callback);
    };
    return TopicCatalogServiceClient;
}());
exports.TopicCatalogServiceClient = TopicCatalogServiceClient;
