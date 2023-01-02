/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "grid.grid";

export interface MsgSubmitPower {
  creator: string;
  watt: string;
}

export interface MsgSubmitPowerResponse {
}

export interface MsgSubmitEstimate {
  creator: string;
  watt: string;
}

export interface MsgSubmitEstimateResponse {
}

function createBaseMsgSubmitPower(): MsgSubmitPower {
  return { creator: "", watt: "" };
}

export const MsgSubmitPower = {
  encode(message: MsgSubmitPower, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.watt !== "") {
      writer.uint32(18).string(message.watt);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitPower {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitPower();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.watt = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitPower {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      watt: isSet(object.watt) ? String(object.watt) : "",
    };
  },

  toJSON(message: MsgSubmitPower): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.watt !== undefined && (obj.watt = message.watt);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSubmitPower>, I>>(object: I): MsgSubmitPower {
    const message = createBaseMsgSubmitPower();
    message.creator = object.creator ?? "";
    message.watt = object.watt ?? "";
    return message;
  },
};

function createBaseMsgSubmitPowerResponse(): MsgSubmitPowerResponse {
  return {};
}

export const MsgSubmitPowerResponse = {
  encode(_: MsgSubmitPowerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitPowerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitPowerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSubmitPowerResponse {
    return {};
  },

  toJSON(_: MsgSubmitPowerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSubmitPowerResponse>, I>>(_: I): MsgSubmitPowerResponse {
    const message = createBaseMsgSubmitPowerResponse();
    return message;
  },
};

function createBaseMsgSubmitEstimate(): MsgSubmitEstimate {
  return { creator: "", watt: "" };
}

export const MsgSubmitEstimate = {
  encode(message: MsgSubmitEstimate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.watt !== "") {
      writer.uint32(18).string(message.watt);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitEstimate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitEstimate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.watt = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitEstimate {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      watt: isSet(object.watt) ? String(object.watt) : "",
    };
  },

  toJSON(message: MsgSubmitEstimate): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.watt !== undefined && (obj.watt = message.watt);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSubmitEstimate>, I>>(object: I): MsgSubmitEstimate {
    const message = createBaseMsgSubmitEstimate();
    message.creator = object.creator ?? "";
    message.watt = object.watt ?? "";
    return message;
  },
};

function createBaseMsgSubmitEstimateResponse(): MsgSubmitEstimateResponse {
  return {};
}

export const MsgSubmitEstimateResponse = {
  encode(_: MsgSubmitEstimateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitEstimateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitEstimateResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSubmitEstimateResponse {
    return {};
  },

  toJSON(_: MsgSubmitEstimateResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSubmitEstimateResponse>, I>>(_: I): MsgSubmitEstimateResponse {
    const message = createBaseMsgSubmitEstimateResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  SubmitPower(request: MsgSubmitPower): Promise<MsgSubmitPowerResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SubmitEstimate(request: MsgSubmitEstimate): Promise<MsgSubmitEstimateResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.SubmitPower = this.SubmitPower.bind(this);
    this.SubmitEstimate = this.SubmitEstimate.bind(this);
  }
  SubmitPower(request: MsgSubmitPower): Promise<MsgSubmitPowerResponse> {
    const data = MsgSubmitPower.encode(request).finish();
    const promise = this.rpc.request("grid.grid.Msg", "SubmitPower", data);
    return promise.then((data) => MsgSubmitPowerResponse.decode(new _m0.Reader(data)));
  }

  SubmitEstimate(request: MsgSubmitEstimate): Promise<MsgSubmitEstimateResponse> {
    const data = MsgSubmitEstimate.encode(request).finish();
    const promise = this.rpc.request("grid.grid.Msg", "SubmitEstimate", data);
    return promise.then((data) => MsgSubmitEstimateResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
