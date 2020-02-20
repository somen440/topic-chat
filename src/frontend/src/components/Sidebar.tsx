import * as tsx from "vue-tsx-support";
import * as user from "@/store/modules/user";
import * as room from "@/store/modules/room";

export default tsx.component({
  name: "Sidebar",
  props: {
    current: {
      type: String
    },
  },
  methods: {
    routerLinkClass(to: string): string {
      if (to === this.current) {
        return "list-group-item list-group-item-primary";
      }
      return "list-group-item";
    }
  },
  render() {
    return (
      <div>
        {user.readIsLoggedIn(this.$store) && (
          <div>
            <span class="badge badge-primary">{ user.readGetMine(this.$store)?.getName() }</span>
            <hr />
          </div>
        )}

        {user.readIsSelectedTopic(this.$store) && (
          <div>
            <div class="card">
              <img src="https://storage.googleapis.com/topic-chat/images/no_img.png" width="128" height="128" />
              <div class="card-body">
                <h5 class="card-title">{ user.readGetSelectedTopic(this.$store)?.getName() }</h5>
                <p class="card-text">topic detail</p>
              </div>
            </div>
            <hr />
          </div>
        )}

        {room.readIsNotEmptyMember(this.$store) &&  (
          <div>
            <h3>members</h3>
            <ul class="list-group">
              {room.readMember(this.$store).map(e => (
                <li class="list-group-item">
                  { e.getName() }{ user.readGetMine(this.$store)?.getId() === e.getId() && (<span> (you)</span>)}
                </li>
              ))}
            </ul>
            <hr />
          </div>
        )}

        <h3>Link</h3>
        <ul class="list-group">
          <router-link
            tag="li"
            to="/"
            class={ this.routerLinkClass("Home") }
            style="cursor: pointer;"
          >
            home
          </router-link>
          {user.readIsLoggedIn(this.$store) ? (
            <router-link
              tag="li"
              to="/topic"
              class={ this.routerLinkClass("TopicCatalog") }
              style="cursor: pointer;"
            >
              topic
            </router-link>
          ) : (
            <router-link
              tag="li"
              to="/join"
              class={ this.routerLinkClass("Join") }
              style="cursor: pointer;"
            >
              join
            </router-link>
          )}
        </ul>
      </div>
    );
  }
});
