import { ChatServiceClient } from "@/pb/TopicchatServiceClientPb"
import {

} from "@/pb/topicchat_pb";

const SCHEME = process.env.SCHEME ?? "http";
const CHAT_SERVICE_ADDR = process.env.CHAT_SERVICE_ADDR ?? "localhost:9090";

export interface ChatState {
  client: ChatServiceClient,
}

//
// state
//
const state = {
  client: new ChatServiceClient(`${SCHEME}://${CHAT_SERVICE_ADDR}`)
}

//
// getters
//
const getters = {}

//
// mutations
//
const mutations = {}

//
// actions
//
const actions = {
}

export const chat = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
};
