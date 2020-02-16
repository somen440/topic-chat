import { ActionContext } from "vuex";
import { State as RootState } from "@/store/state"
import { getStoreAccessors } from "vuex-typescript";
import { User } from '@/pb/topicchat_pb';

type UnUser = User | undefined

export interface UserState {
  mine: UnUser
}

type UserContext = ActionContext<UserState, RootState>;

const { commit, read, dispatch } = getStoreAccessors<UserState, RootState>(
  "user"
);

//
// state
//
const state = {
  mine: undefined
}

//
// getters
//
const getters = {
  getMine(state: UserState): UnUser {
    return state.mine;
  },
  isLoggedIn(): boolean {
    return state.mine !== undefined;
  }
}

export const readGetMine = read(getters.getMine);
export const readIsLoggedIn = read(getters.isLoggedIn);

//
// mutations
//
const mutations = {
  loggedIn(state: UserState, mine: User): void {
    state.mine = mine;
  }
};

export const commitLoggedIn = commit(mutations.loggedIn);

//
// actions
//
const actions = {}

export const user = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
};
