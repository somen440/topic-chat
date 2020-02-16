import * as tsx from "vue-tsx-support";
import Header from "@/components/Header";
import * as user from "@/store/modules/user";

export default tsx.component({
  name: "App",
  render() {
    return (
      <div id="app">
        <Header />
        <h2>link</h2>
        <ul>
          <li>
            <router-link tag="li" to="/">
              <a>home</a>
            </router-link>
          </li>
          {user.readIsLoggedIn(this.$store) ? (
            <li>
              <router-link tag="li" to="/topic">
                <a>topic</a>
              </router-link>
            </li>
          ) : (
            <li>
            <router-link tag="li" to="/join">
              <a>join</a>
            </router-link>
          </li>
          )}
        </ul>
        {user.readIsLoggedIn(this.$store) && (
          <div>
            <h2>Profile</h2>
            <p>name: {user.readGetMine(this.$store)?.getName()}</p>
          </div>
        )}
        <router-view />
      </div>
    );
  }
});
