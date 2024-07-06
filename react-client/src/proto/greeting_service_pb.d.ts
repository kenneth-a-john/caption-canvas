import * as jspb from 'google-protobuf'



export class GreetingGenerationRequest extends jspb.Message {
  getMessagePrompt(): string;
  setMessagePrompt(value: string): GreetingGenerationRequest;

  getImagePrompt(): string;
  setImagePrompt(value: string): GreetingGenerationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GreetingGenerationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GreetingGenerationRequest): GreetingGenerationRequest.AsObject;
  static serializeBinaryToWriter(message: GreetingGenerationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GreetingGenerationRequest;
  static deserializeBinaryFromReader(message: GreetingGenerationRequest, reader: jspb.BinaryReader): GreetingGenerationRequest;
}

export namespace GreetingGenerationRequest {
  export type AsObject = {
    messagePrompt: string,
    imagePrompt: string,
  }
}

export class GreetingGenerationResponse extends jspb.Message {
  getImagePath(): string;
  setImagePath(value: string): GreetingGenerationResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GreetingGenerationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GreetingGenerationResponse): GreetingGenerationResponse.AsObject;
  static serializeBinaryToWriter(message: GreetingGenerationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GreetingGenerationResponse;
  static deserializeBinaryFromReader(message: GreetingGenerationResponse, reader: jspb.BinaryReader): GreetingGenerationResponse;
}

export namespace GreetingGenerationResponse {
  export type AsObject = {
    imagePath: string,
  }
}

