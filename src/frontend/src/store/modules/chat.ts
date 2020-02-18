import { ActionContext } from "vuex";
import { State as RootState } from "@/store/state";
import { getStoreAccessors } from "vuex-typescript";
import { ChatServiceClient } from "@/pb/TopicchatServiceClientPb";
import {
  SendMessageRequest,
  JoinRoomRequest,
  JoinRoomResponse,
  RecvMessageRequest,
  ChatMessage,
  Empty,
  RecvMemberRequest,
  User
} from "@/pb/topicchat_pb";
import { ClientReadableStream } from "grpc-web";
import * as moment from "moment-timezone";

const SCHEME = process.env.SCHEME ?? "http";
const CHAT_SERVICE_ADDR = process.env.CHAT_SERVICE_ADDR ?? "localhost:9090";

export interface ChatState {
  client: ChatServiceClient;
}

type ChatContext = ActionContext<ChatState, RootState>;

const { commit, read, dispatch } = getStoreAccessors<ChatState, RootState>(
  "chat"
);

//
// state
//
const state = {
  client: new ChatServiceClient(`${SCHEME}://${CHAT_SERVICE_ADDR}`)
};

//
// getters
//
const getters = {
  getClient(state: ChatState) {
    return state.client;
  }
};

export const readGetClient = read(getters.getClient);

//
// mutations
//
const mutations = {};

//
// actions
//
const actions = {
  joinRoom(
    context: ChatContext,
    item: { userId: number; topicId: number }
  ): Promise<ClientReadableStream<JoinRoomResponse>> {
    const req = new JoinRoomRequest();
    req.setUserId(item.userId);
    req.setTopicId(item.topicId);

    return new Promise(resolve => {
      const stream = readGetClient(context).joinRoom(req, {}, err => {
        console.log(err);
      });
      resolve(stream);
    });
  },
  recvMessage(
    context: ChatContext,
    item: { userId: number; topicId: number }
  ): Promise<ClientReadableStream<ChatMessage>> {
    const req = new RecvMessageRequest();
    req.setUserId(item.userId);
    req.setTopicId(item.topicId);

    return new Promise(resolve => {
      const stream = readGetClient(context).recvMessage(req, {});
      resolve(stream);
    });
  },
  sendMessage(
    context: ChatContext,
    item: { topicId: number; message: ChatMessage }
  ): Promise<ClientReadableStream<Empty>> {
    const req = new SendMessageRequest();
    req.setTopicId(item.topicId);

    const now = moment.tz("Asia/Tokyo").format("YYYY-MM-DD HH:mm:ss");
    item.message.setActionedAt(now);
    req.setMessage(item.message);

    return new Promise(resolve => {
      const stream = readGetClient(context).sendMessage(req, {}, err => {
        console.log(err);
      });
      resolve(stream);
    });
  },
  recvMember(
    context: ChatContext,
    item: { topicId: number; userId: number; }
  ): Promise<ClientReadableStream<User>> {
    const req = new RecvMemberRequest();
    req.setUserId(item.userId);
    req.setTopicId(item.topicId);

    return new Promise(resolve => {
      const stream = readGetClient(context).recvMember(req, {});
      resolve(stream);
    });
  }
};

export const dispatchJoinRoom = dispatch(actions.joinRoom);
export const dispatchRecvMessage = dispatch(actions.recvMessage);
export const dispatchSendMessage = dispatch(actions.sendMessage);
export const dispatchRecvMember = dispatch(actions.recvMember);

export const chat = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
};
