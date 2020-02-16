import * as tsx from "vue-tsx-support";
import Header from "@/components/Header";

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
          <li>
            <router-link tag="li" to="/topic">
              <a>topic</a>
            </router-link>
          </li>
          <li>
            <router-link tag="li" to="/join">
              <a>join</a>
            </router-link>
          </li>
          <li>
            <router-link tag="li" to="/count">
              <a>count</a>
            </router-link>
          </li>
        </ul>
        <router-view />
      </div>
    );
  }
});
