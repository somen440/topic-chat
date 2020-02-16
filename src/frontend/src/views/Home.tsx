import * as tsx from "vue-tsx-support";

export default tsx.component({
  name: "Home",
  data() {
    return {
      title: "Home"
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
