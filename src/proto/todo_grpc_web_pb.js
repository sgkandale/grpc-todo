/**
 * @fileoverview gRPC-Web generated client stub for todo
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.todo = require('./todo_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.todo.TodoServiceClient =
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
proto.todo.TodoServicePromiseClient =
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
 *   !proto.todo.GetItemsRequest,
 *   !proto.todo.GetItemsResponse>}
 */
const methodDescriptor_TodoService_GetItems = new grpc.web.MethodDescriptor(
  '/todo.TodoService/GetItems',
  grpc.web.MethodType.UNARY,
  proto.todo.GetItemsRequest,
  proto.todo.GetItemsResponse,
  /**
   * @param {!proto.todo.GetItemsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.GetItemsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.todo.GetItemsRequest,
 *   !proto.todo.GetItemsResponse>}
 */
const methodInfo_TodoService_GetItems = new grpc.web.AbstractClientBase.MethodInfo(
  proto.todo.GetItemsResponse,
  /**
   * @param {!proto.todo.GetItemsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.GetItemsResponse.deserializeBinary
);


/**
 * @param {!proto.todo.GetItemsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.todo.GetItemsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.todo.GetItemsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.TodoServiceClient.prototype.getItems =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.TodoService/GetItems',
      request,
      metadata || {},
      methodDescriptor_TodoService_GetItems,
      callback);
};


/**
 * @param {!proto.todo.GetItemsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.todo.GetItemsResponse>}
 *     Promise that resolves to the response
 */
proto.todo.TodoServicePromiseClient.prototype.getItems =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.TodoService/GetItems',
      request,
      metadata || {},
      methodDescriptor_TodoService_GetItems);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.todo.Item,
 *   !proto.todo.Item>}
 */
const methodDescriptor_TodoService_CreateItem = new grpc.web.MethodDescriptor(
  '/todo.TodoService/CreateItem',
  grpc.web.MethodType.UNARY,
  proto.todo.Item,
  proto.todo.Item,
  /**
   * @param {!proto.todo.Item} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.Item.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.todo.Item,
 *   !proto.todo.Item>}
 */
const methodInfo_TodoService_CreateItem = new grpc.web.AbstractClientBase.MethodInfo(
  proto.todo.Item,
  /**
   * @param {!proto.todo.Item} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.Item.deserializeBinary
);


/**
 * @param {!proto.todo.Item} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.todo.Item)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.todo.Item>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.TodoServiceClient.prototype.createItem =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.TodoService/CreateItem',
      request,
      metadata || {},
      methodDescriptor_TodoService_CreateItem,
      callback);
};


/**
 * @param {!proto.todo.Item} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.todo.Item>}
 *     Promise that resolves to the response
 */
proto.todo.TodoServicePromiseClient.prototype.createItem =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.TodoService/CreateItem',
      request,
      metadata || {},
      methodDescriptor_TodoService_CreateItem);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.todo.Item,
 *   !proto.todo.GeneralResponse>}
 */
const methodDescriptor_TodoService_CloseItem = new grpc.web.MethodDescriptor(
  '/todo.TodoService/CloseItem',
  grpc.web.MethodType.UNARY,
  proto.todo.Item,
  proto.todo.GeneralResponse,
  /**
   * @param {!proto.todo.Item} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.GeneralResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.todo.Item,
 *   !proto.todo.GeneralResponse>}
 */
const methodInfo_TodoService_CloseItem = new grpc.web.AbstractClientBase.MethodInfo(
  proto.todo.GeneralResponse,
  /**
   * @param {!proto.todo.Item} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.GeneralResponse.deserializeBinary
);


/**
 * @param {!proto.todo.Item} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.todo.GeneralResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.todo.GeneralResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.TodoServiceClient.prototype.closeItem =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.TodoService/CloseItem',
      request,
      metadata || {},
      methodDescriptor_TodoService_CloseItem,
      callback);
};


/**
 * @param {!proto.todo.Item} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.todo.GeneralResponse>}
 *     Promise that resolves to the response
 */
proto.todo.TodoServicePromiseClient.prototype.closeItem =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.TodoService/CloseItem',
      request,
      metadata || {},
      methodDescriptor_TodoService_CloseItem);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.todo.Item,
 *   !proto.todo.GeneralResponse>}
 */
const methodDescriptor_TodoService_DeleteItem = new grpc.web.MethodDescriptor(
  '/todo.TodoService/DeleteItem',
  grpc.web.MethodType.UNARY,
  proto.todo.Item,
  proto.todo.GeneralResponse,
  /**
   * @param {!proto.todo.Item} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.GeneralResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.todo.Item,
 *   !proto.todo.GeneralResponse>}
 */
const methodInfo_TodoService_DeleteItem = new grpc.web.AbstractClientBase.MethodInfo(
  proto.todo.GeneralResponse,
  /**
   * @param {!proto.todo.Item} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.GeneralResponse.deserializeBinary
);


/**
 * @param {!proto.todo.Item} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.todo.GeneralResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.todo.GeneralResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.TodoServiceClient.prototype.deleteItem =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.TodoService/DeleteItem',
      request,
      metadata || {},
      methodDescriptor_TodoService_DeleteItem,
      callback);
};


/**
 * @param {!proto.todo.Item} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.todo.GeneralResponse>}
 *     Promise that resolves to the response
 */
proto.todo.TodoServicePromiseClient.prototype.deleteItem =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.TodoService/DeleteItem',
      request,
      metadata || {},
      methodDescriptor_TodoService_DeleteItem);
};


module.exports = proto.todo;

