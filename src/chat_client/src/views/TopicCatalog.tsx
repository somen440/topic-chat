import * as tsx from "vue-tsx-support";

export default tsx.component({
  name: "TopicCatalog",
  data() {
    return {
      title: "Topic Catalog"
    };
  },
  render() {
    return (
      <div>
        <h1>{this.title}</h1>
      </div>
    );
  }
});
