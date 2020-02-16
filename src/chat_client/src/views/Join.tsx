import * as tsx from "vue-tsx-support";

export default tsx.component({
  name: "Join",
  data() {
    return {
      title: "Join"
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
