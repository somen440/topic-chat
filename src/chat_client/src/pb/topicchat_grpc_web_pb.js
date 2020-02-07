/* eslint-disable */
/**
 * @fileoverview gRPC-Web generated client stub for topicchat
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.topicchat = require('./topicchat_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.topicchat.AuthServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.topicchat.AuthServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.topicchat.JoinRequest,
 *   !proto.topicchat.User>}
 */
const methodDescriptor_AuthService_Join = new grpc.web.MethodDescriptor(
  '/topicchat.AuthService/Join',
  grpc.web.MethodType.UNARY,
  proto.topicchat.JoinRequest,
  proto.topicchat.User,
  /**
   * @param {!proto.topicchat.JoinRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.topicchat.User.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.topicchat.JoinRequest,
 *   !proto.topicchat.User>}
 */
const methodInfo_AuthService_Join = new grpc.web.AbstractClientBase.MethodInfo(
  proto.topicchat.User,
  /**
   * @param {!proto.topicchat.JoinRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.topicchat.User.deserializeBinary
);


/**
 * @param {!proto.topicchat.JoinRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.topicchat.User)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.topicchat.User>|undefined}
 *     The XHR Node Readable Stream
 */
proto.topicchat.AuthServiceClient.prototype.join =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/topicchat.AuthService/Join',
      request,
      metadata || {},
      methodDescriptor_AuthService_Join,
      callback);
};


/**
 * @param {!proto.topicchat.JoinRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.topicchat.User>}
 *     A native promise that resolves to the response
 */
proto.topicchat.AuthServicePromiseClient.prototype.join =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/topicchat.AuthService/Join',
      request,
      metadata || {},
      methodDescriptor_AuthService_Join);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.topicchat.LoggedInRequest,
 *   !proto.topicchat.User>}
 */
const methodDescriptor_AuthService_LoggedIn = new grpc.web.MethodDescriptor(
  '/topicchat.AuthService/LoggedIn',
  grpc.web.MethodType.UNARY,
  proto.topicchat.LoggedInRequest,
  proto.topicchat.User,
  /**
   * @param {!proto.topicchat.LoggedInRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.topicchat.User.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.topicchat.LoggedInRequest,
 *   !proto.topicchat.User>}
 */
const methodInfo_AuthService_LoggedIn = new grpc.web.AbstractClientBase.MethodInfo(
  proto.topicchat.User,
  /**
   * @param {!proto.topicchat.LoggedInRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.topicchat.User.deserializeBinary
);


/**
 * @param {!proto.topicchat.LoggedInRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.topicchat.User)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.topicchat.User>|undefined}
 *     The XHR Node Readable Stream
 */
proto.topicchat.AuthServiceClient.prototype.loggedIn =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/topicchat.AuthService/LoggedIn',
      request,
      metadata || {},
      methodDescriptor_AuthService_LoggedIn,
      callback);
};


/**
 * @param {!proto.topicchat.LoggedInRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.topicchat.User>}
 *     A native promise that resolves to the response
 */
proto.topicchat.AuthServicePromiseClient.prototype.loggedIn =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/topicchat.AuthService/LoggedIn',
      request,
      metadata || {},
      methodDescriptor_AuthService_LoggedIn);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.topicchat.SignoutRequest,
 *   !proto.topicchat.Empty>}
 */
const methodDescriptor_AuthService_Signout = new grpc.web.MethodDescriptor(
  '/topicchat.AuthService/Signout',
  grpc.web.MethodType.UNARY,
  proto.topicchat.SignoutRequest,
  proto.topicchat.Empty,
  /**
   * @param {!proto.topicchat.SignoutRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.topicchat.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.topicchat.SignoutRequest,
 *   !proto.topicchat.Empty>}
 */
const methodInfo_AuthService_Signout = new grpc.web.AbstractClientBase.MethodInfo(
  proto.topicchat.Empty,
  /**
   * @param {!proto.topicchat.SignoutRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.topicchat.Empty.deserializeBinary
);


/**
 * @param {!proto.topicchat.SignoutRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.topicchat.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.topicchat.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.topicchat.AuthServiceClient.prototype.signout =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/topicchat.AuthService/Signout',
      request,
      metadata || {},
      methodDescriptor_AuthService_Signout,
      callback);
};


/**
 * @param {!proto.topicchat.SignoutRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.topicchat.Empty>}
 *     A native promise that resolves to the response
 */
proto.topicchat.AuthServicePromiseClient.prototype.signout =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/topicchat.AuthService/Signout',
      request,
      metadata || {},
      methodDescriptor_AuthService_Signout);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.topicchat.RoomServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.topicchat.RoomServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.topicchat.RoomJoinRequest,
 *   !proto.topicchat.Empty>}
 */
const methodDescriptor_RoomService_Join = new grpc.web.MethodDescriptor(
  '/topicchat.RoomService/Join',
  grpc.web.MethodType.UNARY,
  proto.topicchat.RoomJoinRequest,
  proto.topicchat.Empty,
  /**
   * @param {!proto.topicchat.RoomJoinRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.topicchat.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.topicchat.RoomJoinRequest,
 *   !proto.topicchat.Empty>}
 */
const methodInfo_RoomService_Join = new grpc.web.AbstractClientBase.MethodInfo(
  proto.topicchat.Empty,
  /**
   * @param {!proto.topicchat.RoomJoinRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.topicchat.Empty.deserializeBinary
);


/**
 * @param {!proto.topicchat.RoomJoinRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.topicchat.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.topicchat.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.topicchat.RoomServiceClient.prototype.join =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/topicchat.RoomService/Join',
      request,
      metadata || {},
      methodDescriptor_RoomService_Join,
      callback);
};


/**
 * @param {!proto.topicchat.RoomJoinRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.topicchat.Empty>}
 *     A native promise that resolves to the response
 */
proto.topicchat.RoomServicePromiseClient.prototype.join =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/topicchat.RoomService/Join',
      request,
      metadata || {},
      methodDescriptor_RoomService_Join);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.topicchat.Empty,
 *   !proto.topicchat.Empty>}
 */
const methodDescriptor_RoomService_Left = new grpc.web.MethodDescriptor(
  '/topicchat.RoomService/Left',
  grpc.web.MethodType.UNARY,
  proto.topicchat.Empty,
  proto.topicchat.Empty,
  /**
   * @param {!proto.topicchat.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.topicchat.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.topicchat.Empty,
 *   !proto.topicchat.Empty>}
 */
const methodInfo_RoomService_Left = new grpc.web.AbstractClientBase.MethodInfo(
  proto.topicchat.Empty,
  /**
   * @param {!proto.topicchat.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.topicchat.Empty.deserializeBinary
);


/**
 * @param {!proto.topicchat.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.topicchat.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.topicchat.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.topicchat.RoomServiceClient.prototype.left =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/topicchat.RoomService/Left',
      request,
      metadata || {},
      methodDescriptor_RoomService_Left,
      callback);
};


/**
 * @param {!proto.topicchat.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.topicchat.Empty>}
 *     A native promise that resolves to the response
 */
proto.topicchat.RoomServicePromiseClient.prototype.left =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/topicchat.RoomService/Left',
      request,
      metadata || {},
      methodDescriptor_RoomService_Left);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.topicchat.ChatServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.topicchat.ChatServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.topicchat.ChatMessage,
 *   !proto.topicchat.Empty>}
 */
const methodDescriptor_ChatService_Send = new grpc.web.MethodDescriptor(
  '/topicchat.ChatService/Send',
  grpc.web.MethodType.UNARY,
  proto.topicchat.ChatMessage,
  proto.topicchat.Empty,
  /**
   * @param {!proto.topicchat.ChatMessage} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.topicchat.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.topicchat.ChatMessage,
 *   !proto.topicchat.Empty>}
 */
const methodInfo_ChatService_Send = new grpc.web.AbstractClientBase.MethodInfo(
  proto.topicchat.Empty,
  /**
   * @param {!proto.topicchat.ChatMessage} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.topicchat.Empty.deserializeBinary
);


/**
 * @param {!proto.topicchat.ChatMessage} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.topicchat.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.topicchat.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.topicchat.ChatServiceClient.prototype.send =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/topicchat.ChatService/Send',
      request,
      metadata || {},
      methodDescriptor_ChatService_Send,
      callback);
};


/**
 * @param {!proto.topicchat.ChatMessage} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.topicchat.Empty>}
 *     A native promise that resolves to the response
 */
proto.topicchat.ChatServicePromiseClient.prototype.send =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/topicchat.ChatService/Send',
      request,
      metadata || {},
      methodDescriptor_ChatService_Send);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.topicchat.TopicCatalogServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.topicchat.TopicCatalogServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.topicchat.Empty,
 *   !proto.topicchat.ListTopicsResponse>}
 */
const methodDescriptor_TopicCatalogService_ListTopics = new grpc.web.MethodDescriptor(
  '/topicchat.TopicCatalogService/ListTopics',
  grpc.web.MethodType.UNARY,
  proto.topicchat.Empty,
  proto.topicchat.ListTopicsResponse,
  /**
   * @param {!proto.topicchat.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.topicchat.ListTopicsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.topicchat.Empty,
 *   !proto.topicchat.ListTopicsResponse>}
 */
const methodInfo_TopicCatalogService_ListTopics = new grpc.web.AbstractClientBase.MethodInfo(
  proto.topicchat.ListTopicsResponse,
  /**
   * @param {!proto.topicchat.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.topicchat.ListTopicsResponse.deserializeBinary
);


/**
 * @param {!proto.topicchat.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.topicchat.ListTopicsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.topicchat.ListTopicsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.topicchat.TopicCatalogServiceClient.prototype.listTopics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/topicchat.TopicCatalogService/ListTopics',
      request,
      metadata || {},
      methodDescriptor_TopicCatalogService_ListTopics,
      callback);
};


/**
 * @param {!proto.topicchat.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.topicchat.ListTopicsResponse>}
 *     A native promise that resolves to the response
 */
proto.topicchat.TopicCatalogServicePromiseClient.prototype.listTopics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/topicchat.TopicCatalogService/ListTopics',
      request,
      metadata || {},
      methodDescriptor_TopicCatalogService_ListTopics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.topicchat.GetTopicRequest,
 *   !proto.topicchat.Topic>}
 */
const methodDescriptor_TopicCatalogService_GetTopic = new grpc.web.MethodDescriptor(
  '/topicchat.TopicCatalogService/GetTopic',
  grpc.web.MethodType.UNARY,
  proto.topicchat.GetTopicRequest,
  proto.topicchat.Topic,
  /**
   * @param {!proto.topicchat.GetTopicRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.topicchat.Topic.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.topicchat.GetTopicRequest,
 *   !proto.topicchat.Topic>}
 */
const methodInfo_TopicCatalogService_GetTopic = new grpc.web.AbstractClientBase.MethodInfo(
  proto.topicchat.Topic,
  /**
   * @param {!proto.topicchat.GetTopicRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.topicchat.Topic.deserializeBinary
);


/**
 * @param {!proto.topicchat.GetTopicRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.topicchat.Topic)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.topicchat.Topic>|undefined}
 *     The XHR Node Readable Stream
 */
proto.topicchat.TopicCatalogServiceClient.prototype.getTopic =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/topicchat.TopicCatalogService/GetTopic',
      request,
      metadata || {},
      methodDescriptor_TopicCatalogService_GetTopic,
      callback);
};


/**
 * @param {!proto.topicchat.GetTopicRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.topicchat.Topic>}
 *     A native promise that resolves to the response
 */
proto.topicchat.TopicCatalogServicePromiseClient.prototype.getTopic =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/topicchat.TopicCatalogService/GetTopic',
      request,
      metadata || {},
      methodDescriptor_TopicCatalogService_GetTopic);
};


module.exports = proto.topicchat;
