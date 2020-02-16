import * as tsx from "vue-tsx-support";

export default tsx.component({
  name: "Room",
  data() {
    return {
      title: "Room"
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
