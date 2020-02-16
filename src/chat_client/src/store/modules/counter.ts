import { ActionContext } from "vuex";
import { getStoreAccessors } from "vuex-typescript";
import { State as RootState } from "../state";

export interface CounterState {
  count: number;
}

type CounterContext = ActionContext<CounterState, RootState>;

const { commit, read, dispatch } = getStoreAccessors<CounterState, RootState>(
  "counter"
);

//
// state
//
const state = {
  count: 10
};

//
// getters
//
const getters = {
  getCount(state: CounterState): number {
    return state.count;
  },
  getDoubledCount(state: CounterState): number {
    return state.count * 2;
  }
};

export const readCount = read(getters.getCount);
export const readDoubledCount = read(getters.getDoubledCount);

//
// mutations
//
const mutations = {
  incrementCount(state: CounterState): void {
    state.count++;
  },
  decrementCount(state: CounterState): void {
    state.count--;
  }
};

export const commitIncrementCount = commit(mutations.incrementCount);
export const commitDecrementCount = commit(mutations.decrementCount);

//
// actions
//
const actions = {
  async incrementCountAsync(context: CounterContext): Promise<void> {
    setTimeout(() => {
      commitIncrementCount(context);
    }, 1000);
  },
  async decrementCountAsync(context: CounterContext): Promise<void> {
    setTimeout(() => {
      commitDecrementCount(context);
    }, 1000);
  }
};

export const dispatchIncrementCount = dispatch(actions.incrementCountAsync);
export const dispatchDecrementCount = dispatch(actions.incrementCountAsync);

export const counter = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
};
