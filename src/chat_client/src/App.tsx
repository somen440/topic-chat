import * as tsx from "vue-tsx-support";
import { ChatServiceClient } from "../pb/TopicchatServiceClientPb";
import { Empty } from "../pb/topicchat_pb";

export default tsx.component({
  name: "App",
  data() {
    return {
      message: "hoge",
      messages: ["a", "b", "c"]
    }
  },
  render() {
    const client = new ChatServiceClient('http://localhost:9090');
    const stream = client.recvMessage(new Empty(), {})
    stream
      .on("data", (response) => {
        console.log("data");
        console.log(response.getText());
        this.messages.push(response.getText());
      });
    stream
      .on("status", (status) => {
        console.log("status");
        console.log(status);
      });
    stream
      .on("error", (err) => {
        console.log("error");
        console.log(err);
      });

    return (
      <div id="app">
        <h1>app</h1>
        <ul>
          { this.messages.map((value) => <li>{value}</li>) }
        </ul>
        <input
          type="text"
          value={this.message}
          onInput={(e) => {
            this.message = e.target.value?.toString() ?? "undefined";
          }}
        />
        <button
          onClick={() => {
            console.log("click");
            console.log(this.message);
          }}
        >send</button>
      </div>
    );
  }
});
