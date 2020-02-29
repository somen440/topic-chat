import * as tsx from "vue-tsx-support";
import * as topicCatalog from "@/store/modules/topic-catalog";
import { Topic } from "@/pb/topicchat_pb";
import * as user from "@/store/modules/user";
import CardList from "@/components/CardList";

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
    if (!user.readIsLoggedIn(this.$store)) {
      this.$router.push("/");
      alert("login");
      return;
    }
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

        <CardList propTopics={this.topics} />
      </div>
    );
  }
});
