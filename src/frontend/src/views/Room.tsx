import * as tsx from "vue-tsx-support";
import * as chat from "@/store/modules/chat";
import * as user from "@/store/modules/user";
import { User, ChatMessage, Topic } from "@/pb/topicchat_pb";

interface RoomData {
  title: string;
  member: User[];
  messages: ChatMessage[];
  sendText: string;
  mine: User | undefined;
  topic: Topic | undefined;
}

export default tsx.component({
  name: "Room",
  data(): RoomData {
    return {
      title: "Room",
      member: [],
      messages: [],
      sendText: "",
      mine: undefined,
      topic: undefined
    };
  },
  methods: {
    validSend(): boolean {
      return this.sendText !== "";
    },
    message(): ChatMessage {
      const message = new ChatMessage();
      message.setText(this.sendText);
      message.setUserId(this.getMineId());
      return message;
    },
    getMineId(): number {
      return this.mine?.getId() ?? -1;
    },
    getTopicId(): number {
      return this.topic?.getId() ?? -1;
    },
    getMemberUser(userId: number): User | undefined {
      return this.member.find(v => v.getId() === userId);
    },
    getTopicInfo(): string {
      if (this.topic === undefined) {
        return "unknown";
      }
      return `${this.topic.getId()}: ${this.topic.getName()}`
    }
  },
  created() {
    this.mine = user.readGetMine(this.$store);
    this.topic = user.readGetSelectedTopic(this.$store);
    if (this.mine === undefined || this.topic === undefined) {
      this.$router.push("/");
      alert("login");
    }

    const payload = {
      userId: this.getMineId(),
      topicId: this.getTopicId()
    };

    chat.dispatchJoinRoom(this.$store, payload).then(res => {
      res.on("data", res => {
        this.member = res.getMemberList();
        chat.dispatchRecvMessage(this.$store, payload).then(res => {
          res.on("data", res => {
            this.messages.push(res);
          });
        });
      });
    });
  },
  render() {
    return (
      <div>
        <h1>{this.title}</h1>
        <h2>room info</h2>
        <h3>topic</h3>
        <p>{this.getTopicInfo()}</p>
        <h3>member</h3>
        <ul>
          {this.member.map(e => (
            <li>
              {user.isEqualsUser(user.readGetMine(this.$store), e) ? (
                <b>
                  [my] {e.getId()}: {e.getName()}
                </b>
              ) : (
                <div>
                  {e.getId()}: {e.getName()}
                </div>
              )}
            </li>
          ))}
        </ul>
        <h2>messages</h2>
        <ul>
          {this.messages.map(e => (
            <li>
              {e.getActionedAt()}:{" "}
              {this.getMemberUser(e.getUserId())?.getName() ?? "unknown"}:{" "}
              {e.getText()}
            </li>
          ))}
        </ul>
        <input
          type="text"
          onInput={e => {
            this.sendText = e.target.value?.toString() ?? "";
          }}
        />
        {this.validSend() && (
          <button
            onClick={() => {
              const payload = {
                topicId: user.readGetSelectedTopic(this.$store)?.getId() ?? -1,
                message: this.message()
              };
              chat.dispatchSendMessage(this.$store, payload);
            }}
          >
            send
          </button>
        )}
      </div>
    );
  }
});
