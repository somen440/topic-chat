import * as tsx from "vue-tsx-support";
import Header from "@/components/Header";
import Sidebar from "@/components/Sidebar";
import * as user from "@/store/modules/user";

export default tsx.component({
  name: "App",
  render() {
    return (
      <div id="app">
        <Header />

        {user.readIsLoggedIn(this.$store) ? (
          <main role="main">
            <section class="jumbotron text-center">
              <div class="container-fluid">
                <router-view />
              </div>
            </section>
          </main>
        ) : (
          <main role="main">
            <section class="jumbotron text-center">
              <div class="container">
                <router-view />
              </div>
            </section>
          </main>
        )}
      </div>
    );
  }
});
