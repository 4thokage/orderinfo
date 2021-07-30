// package: protos
// file: orderWatcher.proto

import * as orderWatcher_pb from "./orderWatcher_pb";
import {grpc} from "@improbable-eng/grpc-web";

type OrderWatcherSubscribe = {
  readonly methodName: string;
  readonly service: typeof OrderWatcher;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof orderWatcher_pb.Request;
  readonly responseType: typeof orderWatcher_pb.Response;
};

type OrderWatcherUnsubscribe = {
  readonly methodName: string;
  readonly service: typeof OrderWatcher;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof orderWatcher_pb.Request;
  readonly responseType: typeof orderWatcher_pb.Response;
};

export class OrderWatcher {
  static readonly serviceName: string;
  static readonly Subscribe: OrderWatcherSubscribe;
  static readonly Unsubscribe: OrderWatcherUnsubscribe;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class OrderWatcherClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  subscribe(requestMessage: orderWatcher_pb.Request, metadata?: grpc.Metadata): ResponseStream<orderWatcher_pb.Response>;
  unsubscribe(
    requestMessage: orderWatcher_pb.Request,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: orderWatcher_pb.Response|null) => void
  ): UnaryResponse;
  unsubscribe(
    requestMessage: orderWatcher_pb.Request,
    callback: (error: ServiceError|null, responseMessage: orderWatcher_pb.Response|null) => void
  ): UnaryResponse;
}

