import Vue from "vue";
import Vuex from "vuex";
import { State } from "@/store/state";
import { auth, chat, counter, topicCatalog } from "@/store/modules";

Vue.use(Vuex);

export const createStore = () =>
  new Vuex.Store<State>({
    modules: {
      auth,
      chat,
      counter,
      topicCatalog
    }
  });

export const store = createStore();
