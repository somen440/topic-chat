import { ActionContext } from "vuex";
import { State as RootState } from "@/store/state";
import { getStoreAccessors } from "vuex-typescript";
import { AuthServiceClient } from "@/pb/TopicchatServiceClientPb";
import { User, JoinRequest } from "@/pb/topicchat_pb";
import { ClientReadableStream } from "grpc-web";

const SCHEME = process.env.SCHEME ?? "http";
const AUTH_SERVICE_ADDR = process.env.AUTH_SERVICE_ADDR ?? "localhost:9093";

export interface AuthState {
  client: AuthServiceClient;
}

type AuthContext = ActionContext<AuthState, RootState>;

const { commit, read, dispatch } = getStoreAccessors<AuthState, RootState>(
  "auth"
);

//
// state
//
const state = {
  client: new AuthServiceClient(`${SCHEME}://${AUTH_SERVICE_ADDR}`)
};

//
// getters
//
const getters = {
  getClient(state: AuthState): AuthServiceClient {
    return state.client;
  }
};

const readClient = read(getters.getClient);

//
// mutations
//
const mutations = {};

//
// actions
//
const actions = {
  join(
    context: AuthContext,
    name: string
  ): Promise<ClientReadableStream<User>> {
    const req = new JoinRequest();
    req.setName(name);
    return new Promise(resolve => {
      const stream = readClient(context).join(req, {}, err => {
        console.log(err);
      });
      resolve(stream);
    });
  }
};

export const dispatchJoin = dispatch(actions.join);

export const auth = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
};
