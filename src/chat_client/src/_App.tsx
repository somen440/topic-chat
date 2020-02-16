import * as tsx from "vue-tsx-support";
import {
  ChatServiceClient,
  AuthServiceClient
} from "@/pb/TopicchatServiceClientPb";
import {
  ChatMessage,
  RecvMessageRequest,
  User,
  SendMessageRequest,
  JoinRoomRequest,
  Topic
} from "@/pb/topicchat_pb";

interface AppData {
  message: string;
  messages: ChatMessage[];
  client: ChatServiceClient;
  topicId: number;
  userId: number;
  topic: Topic | undefined;
  user: User | undefined;
  users: User[];
}

const CHAT_SERVICE_ADDR = process.env.CHAT_SERVICE_ADDR ?? "localhost:9090";

export default tsx.component({
  name: "App",
  data(): AppData {
    const client = new ChatServiceClient(`http://${CHAT_SERVICE_ADDR}`);
    const topicId = Number.parseInt(this.$route.query.topic_id.toString());
    const userId = Number.parseInt(this.$route.query.user_id.toString()); // todo: セキュアな方法に

    return {
      message: "hoge",
      messages: [],
      client: client,
      topicId: topicId,
      userId: userId,
      topic: undefined,
      user: undefined,
      users: []
    };
  },
  methods: {
    joinRoom() {
      const req = new JoinRoomRequest();
      req.setUserId(this.userId);
      req.setTopicId(this.topicId);

      return this.client.joinRoom(req, {}, err => {
        console.log(err);
      });
    },
    recvMessage() {
      const req = new RecvMessageRequest();
      req.setTopicId(this.topicId);
      req.setUserId(this.userId);

      const stream = this.client.recvMessage(req, {});
      stream.on("data", response => {
        console.log("data");
        console.log(response.getText());
        this.messages.push(response);
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
    validSend() {
      return this.user !== undefined;
    }
  },
  created() {
    this.joinRoom().on("data", res => {
      this.topic = res.getTopic();
      this.users = res.getMemberList();
      this.user = res.getPerson();
      this.recvMessage();
    });
  },
  render() {
    return (
      <div id="app">
        <h1>app</h1>
        <h2>topic</h2>
        <ul>
          <li>ID: {this.topic?.getId()}</li>
          <li>Name: {this.topic?.getName()}</li>
        </ul>

        <h2>users</h2>
        <ul>
          {this.users.map(value => (
            <li>
              <ul>
                <li>{value.getId()}</li>
                <li>{value.getName()}</li>
              </ul>
            </li>
          ))}
        </ul>

        <h2>messages</h2>
        <ul>
          {this.messages.map(value => (
            <li>
              [{value.getUser()?.getName() ?? "undefined"}] {value.getText()}
            </li>
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
            if (!this.validSend()) {
              return;
            }

            const req = new SendMessageRequest();
            req.setTopicId(this.topicId);

            const msg = new ChatMessage();
            msg.setText(this.message);
            msg.setUser(this.user);

            req.setMessage(msg);
            this.client
              .sendMessage(req, null, err => {
                console.log(err);
              })
              .on("error", err => {
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
