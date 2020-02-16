import Vue from "vue";
import Vuex from "vuex";
import { State } from "./state";
import { counter } from "./modules";

Vue.use(Vuex);

export const createStore = () =>
  new Vuex.Store<State>({
    modules: {
      counter
    }
  });

export const store = createStore();
