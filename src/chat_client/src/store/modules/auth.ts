import { AuthServiceClient } from "@/pb/TopicchatServiceClientPb"
import {

} from "@/pb/topicchat_pb";

const SCHEME = process.env.SCHEME ?? "http";
const AUTH_SERVICE_ADDR = process.env.AUTH_SERVICE_ADDR ?? "localhost:9093";

export interface AuthState {
  client: AuthServiceClient,
}

//
// state
//
const state = {
  client: new AuthServiceClient(`${SCHEME}://${AUTH_SERVICE_ADDR}`)
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

export const auth = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
};
