import Vue from "vue";
import VueRouter from "vue-router";
import App from "./App";
import router from "./router";
import store from "./store";

Vue.use(VueRouter);

Vue.config.productionTip = false;

new Vue({
  store,
  router,
  render: h => h(App)
}).$mount("#app");
