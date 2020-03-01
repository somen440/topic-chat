import * as tsx from "vue-tsx-support";
import { Topic } from "@/pb/topicchat_pb";
import * as user from "@/store/modules/user";

type Topics = Topic[];
type TopicsList = Topics[];

interface CardListData {
  topicsList: TopicsList;
}

const WIDTH = 3;

export default tsx.component({
  name: "CardList",
  props: {
    propTopics: {
      type: Array as () => Topics,
      required: true
    }
  },
  data(): CardListData {
    return {
      topicsList: []
    };
  },
  watch: {
    propTopics() {
      if (this.propTopics.length === 0) {
        return;
      }
      let topicsListIndex = 0;
      let topicsIndex = 0;
      let propIndex = 0;
      let topics: Topics = [];
      for (;;) {
        topics[topicsIndex++] = this.propTopics[propIndex++];
        if (propIndex === this.propTopics.length) {
          this.topicsList[topicsListIndex] = topics;
          break;
        }
        if (topicsIndex === WIDTH) {
          this.topicsList[topicsListIndex++] = topics;
          topicsIndex = 0;
          topics = [];
        }
      }
      this.$forceUpdate();
    }
  },
  methods: {
    toCardCol(topics: Topics): string {
      const col = Math.floor(12 / topics.length);
      return `card shadow-sm md-${col}`;
    }
  },
  render() {
    return (
      <div class="row">
        {this.topicsList.map(topics => {
          return topics.map(topic => (
            <div class="col-md-4">
              <div class="card shadow-sm md-4">
                <img
                  class="card-img-top"
                  width="225"
                  height="225"
                  src={topic.getImage()}
                  data-holder-rendered="true"
                />

                <div class="card-body">
                  <h5 class="card-title">{topic.getName()}</h5>
                  <p class="card-text">{topic.getDescription()}</p>
                  <div class="d-flex justify-content-between align-items-center">
                    <div class="btn-group">
                      <button
                        type="button"
                        class="btn btn-sm  btn-outline-secondary"
                        onClick={() => {
                          user.commitSelectTopic(this.$store, topic);
                          this.$router.push("/room");
                        }}
                      >
                        Join
                      </button>
                    </div>
                    <small class="text-muted">n mins</small>
                  </div>
                </div>
              </div>
            </div>
          ));
        })}
      </div>
    );
  }
});
