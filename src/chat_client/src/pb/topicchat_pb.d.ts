import * as jspb from "google-protobuf"

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

export class User extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getName(): string;
  setName(value: string): void;

  getLoggedin(): boolean;
  setLoggedin(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: number,
    name: string,
    loggedin: boolean,
  }
}

export class Topic extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Topic.AsObject;
  static toObject(includeInstance: boolean, msg: Topic): Topic.AsObject;
  static serializeBinaryToWriter(message: Topic, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Topic;
  static deserializeBinaryFromReader(message: Topic, reader: jspb.BinaryReader): Topic;
}

export namespace Topic {
  export type AsObject = {
    id: number,
    name: string,
  }
}

export class ChatMessage extends jspb.Message {
  getText(): string;
  setText(value: string): void;

  getUser(): User | undefined;
  setUser(value?: User): void;
  hasUser(): boolean;
  clearUser(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChatMessage.AsObject;
  static toObject(includeInstance: boolean, msg: ChatMessage): ChatMessage.AsObject;
  static serializeBinaryToWriter(message: ChatMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChatMessage;
  static deserializeBinaryFromReader(message: ChatMessage, reader: jspb.BinaryReader): ChatMessage;
}

export namespace ChatMessage {
  export type AsObject = {
    text: string,
    user?: User.AsObject,
  }
}

export class Room extends jspb.Message {
  getTopicId(): number;
  setTopicId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Room.AsObject;
  static toObject(includeInstance: boolean, msg: Room): Room.AsObject;
  static serializeBinaryToWriter(message: Room, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Room;
  static deserializeBinaryFromReader(message: Room, reader: jspb.BinaryReader): Room;
}

export namespace Room {
  export type AsObject = {
    topicId: number,
  }
}

export class JoinRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JoinRequest.AsObject;
  static toObject(includeInstance: boolean, msg: JoinRequest): JoinRequest.AsObject;
  static serializeBinaryToWriter(message: JoinRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JoinRequest;
  static deserializeBinaryFromReader(message: JoinRequest, reader: jspb.BinaryReader): JoinRequest;
}

export namespace JoinRequest {
  export type AsObject = {
    name: string,
  }
}

export class LoggedInRequest extends jspb.Message {
  getUserId(): number;
  setUserId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoggedInRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoggedInRequest): LoggedInRequest.AsObject;
  static serializeBinaryToWriter(message: LoggedInRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoggedInRequest;
  static deserializeBinaryFromReader(message: LoggedInRequest, reader: jspb.BinaryReader): LoggedInRequest;
}

export namespace LoggedInRequest {
  export type AsObject = {
    userId: number,
  }
}

export class SignoutRequest extends jspb.Message {
  getUserId(): number;
  setUserId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SignoutRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SignoutRequest): SignoutRequest.AsObject;
  static serializeBinaryToWriter(message: SignoutRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SignoutRequest;
  static deserializeBinaryFromReader(message: SignoutRequest, reader: jspb.BinaryReader): SignoutRequest;
}

export namespace SignoutRequest {
  export type AsObject = {
    userId: number,
  }
}

export class GetUserRequest extends jspb.Message {
  getUserId(): number;
  setUserId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserRequest): GetUserRequest.AsObject;
  static serializeBinaryToWriter(message: GetUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserRequest;
  static deserializeBinaryFromReader(message: GetUserRequest, reader: jspb.BinaryReader): GetUserRequest;
}

export namespace GetUserRequest {
  export type AsObject = {
    userId: number,
  }
}

export class GetUserAllResponse extends jspb.Message {
  getUsersList(): Array<User>;
  setUsersList(value: Array<User>): void;
  clearUsersList(): void;
  addUsers(value?: User, index?: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserAllResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserAllResponse): GetUserAllResponse.AsObject;
  static serializeBinaryToWriter(message: GetUserAllResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserAllResponse;
  static deserializeBinaryFromReader(message: GetUserAllResponse, reader: jspb.BinaryReader): GetUserAllResponse;
}

export namespace GetUserAllResponse {
  export type AsObject = {
    usersList: Array<User.AsObject>,
  }
}

export class CreateRoomRequest extends jspb.Message {
  getTopicId(): number;
  setTopicId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRoomRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRoomRequest): CreateRoomRequest.AsObject;
  static serializeBinaryToWriter(message: CreateRoomRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRoomRequest;
  static deserializeBinaryFromReader(message: CreateRoomRequest, reader: jspb.BinaryReader): CreateRoomRequest;
}

export namespace CreateRoomRequest {
  export type AsObject = {
    topicId: number,
  }
}

export class RecvMessageRequest extends jspb.Message {
  getTopicId(): number;
  setTopicId(value: number): void;

  getUserId(): number;
  setUserId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RecvMessageRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RecvMessageRequest): RecvMessageRequest.AsObject;
  static serializeBinaryToWriter(message: RecvMessageRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RecvMessageRequest;
  static deserializeBinaryFromReader(message: RecvMessageRequest, reader: jspb.BinaryReader): RecvMessageRequest;
}

export namespace RecvMessageRequest {
  export type AsObject = {
    topicId: number,
    userId: number,
  }
}

export class SendMessageRequest extends jspb.Message {
  getTopicId(): number;
  setTopicId(value: number): void;

  getMessage(): ChatMessage | undefined;
  setMessage(value?: ChatMessage): void;
  hasMessage(): boolean;
  clearMessage(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SendMessageRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SendMessageRequest): SendMessageRequest.AsObject;
  static serializeBinaryToWriter(message: SendMessageRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SendMessageRequest;
  static deserializeBinaryFromReader(message: SendMessageRequest, reader: jspb.BinaryReader): SendMessageRequest;
}

export namespace SendMessageRequest {
  export type AsObject = {
    topicId: number,
    message?: ChatMessage.AsObject,
  }
}

export class JoinRoomRequest extends jspb.Message {
  getUserId(): number;
  setUserId(value: number): void;

  getTopicId(): number;
  setTopicId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JoinRoomRequest.AsObject;
  static toObject(includeInstance: boolean, msg: JoinRoomRequest): JoinRoomRequest.AsObject;
  static serializeBinaryToWriter(message: JoinRoomRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JoinRoomRequest;
  static deserializeBinaryFromReader(message: JoinRoomRequest, reader: jspb.BinaryReader): JoinRoomRequest;
}

export namespace JoinRoomRequest {
  export type AsObject = {
    userId: number,
    topicId: number,
  }
}

export class JoinRoomResponse extends jspb.Message {
  getTopic(): Topic | undefined;
  setTopic(value?: Topic): void;
  hasTopic(): boolean;
  clearTopic(): void;

  getPerson(): User | undefined;
  setPerson(value?: User): void;
  hasPerson(): boolean;
  clearPerson(): void;

  getMemberList(): Array<User>;
  setMemberList(value: Array<User>): void;
  clearMemberList(): void;
  addMember(value?: User, index?: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JoinRoomResponse.AsObject;
  static toObject(includeInstance: boolean, msg: JoinRoomResponse): JoinRoomResponse.AsObject;
  static serializeBinaryToWriter(message: JoinRoomResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JoinRoomResponse;
  static deserializeBinaryFromReader(message: JoinRoomResponse, reader: jspb.BinaryReader): JoinRoomResponse;
}

export namespace JoinRoomResponse {
  export type AsObject = {
    topic?: Topic.AsObject,
    person?: User.AsObject,
    memberList: Array<User.AsObject>,
  }
}

export class ListTopicsResponse extends jspb.Message {
  getTopicsList(): Array<Topic>;
  setTopicsList(value: Array<Topic>): void;
  clearTopicsList(): void;
  addTopics(value?: Topic, index?: number): Topic;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTopicsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListTopicsResponse): ListTopicsResponse.AsObject;
  static serializeBinaryToWriter(message: ListTopicsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTopicsResponse;
  static deserializeBinaryFromReader(message: ListTopicsResponse, reader: jspb.BinaryReader): ListTopicsResponse;
}

export namespace ListTopicsResponse {
  export type AsObject = {
    topicsList: Array<Topic.AsObject>,
  }
}

export class GetTopicRequest extends jspb.Message {
  getTopicId(): number;
  setTopicId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTopicRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetTopicRequest): GetTopicRequest.AsObject;
  static serializeBinaryToWriter(message: GetTopicRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTopicRequest;
  static deserializeBinaryFromReader(message: GetTopicRequest, reader: jspb.BinaryReader): GetTopicRequest;
}

export namespace GetTopicRequest {
  export type AsObject = {
    topicId: number,
  }
}

export class AddTopicRequest extends jspb.Message {
  getTopic(): Topic | undefined;
  setTopic(value?: Topic): void;
  hasTopic(): boolean;
  clearTopic(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddTopicRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddTopicRequest): AddTopicRequest.AsObject;
  static serializeBinaryToWriter(message: AddTopicRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddTopicRequest;
  static deserializeBinaryFromReader(message: AddTopicRequest, reader: jspb.BinaryReader): AddTopicRequest;
}

export namespace AddTopicRequest {
  export type AsObject = {
    topic?: Topic.AsObject,
  }
}

export class UpdateTopicRequest extends jspb.Message {
  getTopic(): Topic | undefined;
  setTopic(value?: Topic): void;
  hasTopic(): boolean;
  clearTopic(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateTopicRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateTopicRequest): UpdateTopicRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateTopicRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateTopicRequest;
  static deserializeBinaryFromReader(message: UpdateTopicRequest, reader: jspb.BinaryReader): UpdateTopicRequest;
}

export namespace UpdateTopicRequest {
  export type AsObject = {
    topic?: Topic.AsObject,
  }
}

export class DeleteTopicRequest extends jspb.Message {
  getTopicId(): number;
  setTopicId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteTopicRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteTopicRequest): DeleteTopicRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteTopicRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteTopicRequest;
  static deserializeBinaryFromReader(message: DeleteTopicRequest, reader: jspb.BinaryReader): DeleteTopicRequest;
}

export namespace DeleteTopicRequest {
  export type AsObject = {
    topicId: number,
  }
}

