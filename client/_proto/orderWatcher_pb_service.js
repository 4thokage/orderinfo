// package: protos
// file: orderWatcher.proto

var orderWatcher_pb = require("./orderWatcher_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var OrderWatcher = (function () {
  function OrderWatcher() {}
  OrderWatcher.serviceName = "protos.OrderWatcher";
  return OrderWatcher;
}());

OrderWatcher.Subscribe = {
  methodName: "Subscribe",
  service: OrderWatcher,
  requestStream: false,
  responseStream: true,
  requestType: orderWatcher_pb.Request,
  responseType: orderWatcher_pb.Response
};

OrderWatcher.Unsubscribe = {
  methodName: "Unsubscribe",
  service: OrderWatcher,
  requestStream: false,
  responseStream: false,
  requestType: orderWatcher_pb.Request,
  responseType: orderWatcher_pb.Response
};

exports.OrderWatcher = OrderWatcher;

function OrderWatcherClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

OrderWatcherClient.prototype.subscribe = function subscribe(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(OrderWatcher.Subscribe, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

OrderWatcherClient.prototype.unsubscribe = function unsubscribe(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(OrderWatcher.Unsubscribe, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.OrderWatcherClient = OrderWatcherClient;

