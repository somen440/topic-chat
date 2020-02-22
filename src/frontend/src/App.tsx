import * as tsx from "vue-tsx-support";
import Header from "@/components/Header";
import Sidebar from "./components/Sidebar";

export default tsx.component({
  name: "App",
  render() {
    return (
      <div id="app">
        <Header />
        <div class="container-fluid">
          <div class="row">
            <div class="col-sm-3">
              <Sidebar current={this.$router.currentRoute.name} />
            </div>
            <div class="col-sm-9">
              <router-view />
            </div>
          </div>
        </div>
      </div>
    );
  }
});
