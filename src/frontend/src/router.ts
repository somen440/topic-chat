import Vue from "vue";
import Router from "vue-router";
import Home from "./views/Home";
import Room from "./views/Room";
import TopicCatalog from "./views/TopicCatalog";

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "/",
      name: "Home",
      component: Home
    },
    {
      path: "/room",
      name: "Room",
      component: Room
    },
    {
      path: "/topic",
      name: "TopicCatalog",
      component: TopicCatalog
    }
  ]
});
