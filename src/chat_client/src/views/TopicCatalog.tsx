import * as tsx from "vue-tsx-support";
import * as topicCatalog from "@/store/modules/topic-catalog";
import { Topic } from '@/pb/topicchat_pb';

interface TopicCatalogData {
  title: string
  topics: Topic[]
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
    topicCatalog.dispatchGetAll(this.$store)
      .then(res => {
        res.on("data", res => {
          this.topics = res.getTopicsList();
        });
      });
  },
  render() {
    return (
      <div>
        <h1>{this.title}</h1>
        <ul>
        {this.topics.map(e => (
          <li>
            { e.getId() }: { e.getName() }
            <button>
              <router-link tag="button" to="room">
              join
              </router-link>
            </button>
          </li>
        ))}
        </ul>
      </div>
    );
  }
});
