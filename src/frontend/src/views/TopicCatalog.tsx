import * as tsx from "vue-tsx-support";
import * as topicCatalog from "@/store/modules/topic-catalog";
import { Topic } from "@/pb/topicchat_pb";
import * as user from "@/store/modules/user";

interface TopicCatalogData {
  title: string;
  topics: Topic[];
}

export default tsx.component({
  name: "TopicCatalog",
  data(): TopicCatalogData {
    return {
      title: "Topic Catalog",
      topics: []
    };
  },
  created() {
    topicCatalog.dispatchGetAll(this.$store).then(res => {
      res.on("data", res => {
        this.topics = res.getTopicsList();
      });
    });
  },
  render() {
    return (
      <div>
        <h1>{this.title}</h1>

        <div class="list-group">
          {this.topics.map(e => (
            <a href="#" class="list-group-item list-group-item-action">
              <div class="d-flex w-100 justify-content-between">
                <h5 class="mb-1">
                  {e.getId()}: {e.getName()}
                </h5>
                <small>timestamp</small>
              </div>
              <div class="d-flex w-100 justify-content-between">
                <p class="mb-1">{ e.getDescription() }</p>
                <button
                  class="btn btn-outline-primary"
                  onClick={() => {
                    user.commitSelectTopic(this.$store, e);
                    this.$router.push("/room");
                  }}
                >
                  Join
                </button>
              </div>
            </a>
          ))}
        </div>
      </div>
    );
  }
});
