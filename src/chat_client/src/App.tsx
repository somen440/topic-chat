import * as tsx from "vue-tsx-support";
import { ChatServiceClient } from "@/pb/TopicchatServiceClientPb";
import { ChatMessage, RecvMessageRequest, User, SendMessageRequest } from "@/pb/topicchat_pb";

export default tsx.component({
  name: "App",
  data() {
    return {
      message: "hoge",
      messages: ["a", "b", "c"],
      client: new ChatServiceClient("http://localhost:9090"),
      topicId: 0
    };
  },
  created() {
    this.topicId = Number.parseInt(this.$route.query.topic_id.toString());
    console.log(this.topicId);

    const req = new RecvMessageRequest();
    req.setTopicId(this.topicId);

    const user = new User();
    user.setId(Math.random().toString(32).substring(2));
    req.setUser(user);

    const stream = this.client.recvMessage(req, {});
    stream.on("data", response => {
      console.log("data");
      console.log(response.getText());
      this.messages.push(response.getText());
    });
    stream.on("status", status => {
      console.log("status");
      console.log(status);
    });
    stream.on("error", err => {
      console.log("error");
      console.log(err);
    });
  },
  render() {
    return (
      <div id="app">
        <h1>app</h1>
        <ul>
          {this.messages.map(value => (
            <li>{value}</li>
          ))}
        </ul>
        <input
          type="text"
          value={this.message}
          onInput={e => {
            this.message = e.target.value?.toString() ?? "undefined";
          }}
        />
        <button
          onClick={() => {
            console.log("click");
            console.log(this.message);
            const req = new SendMessageRequest();
            req.setTopicId(this.topicId);
            const msg = new ChatMessage();
            msg.setText(this.message);
            req.setMessage(msg);
            this.client.sendMessage(req, null, (err, _) => {
              console.log(err);
            });
          }}
        >
          send
        </button>
      </div>
    );
  }
});
