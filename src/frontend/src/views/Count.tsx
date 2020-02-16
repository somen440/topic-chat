import * as tsx from "vue-tsx-support";
import * as counter from "@/store/modules/counter";

export default tsx.component({
  name: "Count",
  data() {
    return {
      title: "Count"
    };
  },
  render() {
    return (
      <div>
        <h1>{this.title}</h1>
        <p>{counter.readCount(this.$store)}</p>
        <p>{counter.readDoubledCount(this.$store)}</p>
        <button onClick={() => counter.commitIncrementCount(this.$store)}>
          +
        </button>
        <button onClick={() => counter.commitDecrementCount(this.$store)}>
          -
        </button>
        <button onClick={() => counter.dispatchIncrementCount(this.$store)}>
          async +
        </button>
        <button onClick={() => counter.dispatchDecrementCount(this.$store)}>
          async -
        </button>
      </div>
    );
  }
});
