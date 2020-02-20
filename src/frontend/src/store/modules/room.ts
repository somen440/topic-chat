import { ActionContext } from "vuex";
import { State as RootState } from "@/store/state";
import { getStoreAccessors } from "vuex-typescript";
import { User } from '@/pb/topicchat_pb';

export interface RoomState {
  member: User[];
}

type RoomContext = ActionContext<RoomState, RootState>;

const { commit, read, dispatch } = getStoreAccessors<
  RoomState,
  RootState
>("room");

//
// state
//
const state = {
  member: []
};

//
// getters
//
const getters = {
  getMembers(state: RoomState): User[] {
    return state.member;
  },
  isNotEmptyMember(state: RoomState): boolean {
    return state.member.length !== 0;
  }
};

export const readMember = read(getters.getMembers);
export const readIsNotEmptyMember = read(getters.isNotEmptyMember);

//
// mutations
//
const mutations = {
  initializeMember(state: RoomState, users: User[]): void {
    state.member = users;
  },
  addMember(state: RoomState, user: User): void {
    state.member.push(user);
  },
};

export const commitInitializeMember = commit(mutations.initializeMember);
export const commitAddMember = commit(mutations.addMember);

//
// actions
//
const actions = {};

export const room = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
};
