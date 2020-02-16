import { ActionContext } from "vuex";
import { State as RootState } from "@/store/state";
import { getStoreAccessors } from "vuex-typescript";
import { User, Topic } from "@/pb/topicchat_pb";

type UnUser = User | undefined;
type UnTopic = Topic | undefined;

export interface UserState {
  mine: UnUser;
  selectedTopicId: UnTopic;
}

type UserContext = ActionContext<UserState, RootState>;

const { commit, read, dispatch } = getStoreAccessors<UserState, RootState>(
  "user"
);

//
// state
//
const state = {
  mine: undefined,
  selectedTopicId: undefined
};

//
// getters
//
const getters = {
  getMine(state: UserState): UnUser {
    return state.mine;
  },
  isLoggedIn(state: UserState): boolean {
    return state.mine !== undefined;
  },

  getSelectedTopic(state: UserState): UnTopic {
    return state.selectedTopicId;
  },
  isSelectedTopic(state: UserState): boolean {
    return state.selectedTopicId !== undefined;
  }
};

export const readGetMine = read(getters.getMine);
export const readIsLoggedIn = read(getters.isLoggedIn);
export const readGetSelectedTopic = read(getters.getSelectedTopic);
export const readIsSelectedTopic = read(getters.isSelectedTopic);

//
// mutations
//
const mutations = {
  loggedIn(state: UserState, mine: User): void {
    state.mine = mine;
  },
  selectTopic(state: UserState, topic: Topic): void {
    state.selectedTopicId = topic;
  }
};

export const commitLoggedIn = commit(mutations.loggedIn);
export const commitSelectTopic = commit(mutations.selectTopic);

//
// actions
//
const actions = {};

export const user = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
};

export function isEqualsUser(a: UnUser, b: UnUser): boolean {
  if (a === undefined || b === undefined) {
    return false;
  }
  return a.getId() === b.getId();
}
