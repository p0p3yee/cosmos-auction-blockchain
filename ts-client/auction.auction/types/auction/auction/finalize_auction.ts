/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "auction.auction";

export interface FinalizeAuction {
  creator: string;
  id: number;
  auctionId: number;
  bidId: number;
  finalPrice: number;
}

function createBaseFinalizeAuction(): FinalizeAuction {
  return { creator: "", id: 0, auctionId: 0, bidId: 0, finalPrice: 0 };
}

export const FinalizeAuction = {
  encode(message: FinalizeAuction, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.auctionId !== 0) {
      writer.uint32(24).uint64(message.auctionId);
    }
    if (message.bidId !== 0) {
      writer.uint32(32).uint64(message.bidId);
    }
    if (message.finalPrice !== 0) {
      writer.uint32(40).uint64(message.finalPrice);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FinalizeAuction {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFinalizeAuction();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.auctionId = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.bidId = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.finalPrice = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FinalizeAuction {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
      auctionId: isSet(object.auctionId) ? Number(object.auctionId) : 0,
      bidId: isSet(object.bidId) ? Number(object.bidId) : 0,
      finalPrice: isSet(object.finalPrice) ? Number(object.finalPrice) : 0,
    };
  },

  toJSON(message: FinalizeAuction): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.auctionId !== undefined && (obj.auctionId = Math.round(message.auctionId));
    message.bidId !== undefined && (obj.bidId = Math.round(message.bidId));
    message.finalPrice !== undefined && (obj.finalPrice = Math.round(message.finalPrice));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FinalizeAuction>, I>>(object: I): FinalizeAuction {
    const message = createBaseFinalizeAuction();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    message.auctionId = object.auctionId ?? 0;
    message.bidId = object.bidId ?? 0;
    message.finalPrice = object.finalPrice ?? 0;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
