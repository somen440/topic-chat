import * as tsx from "vue-tsx-support";
import * as chat from "@/store/modules/chat";
import * as user from "@/store/modules/user";
import { User, ChatMessage, Topic } from "@/pb/topicchat_pb";
import * as moment from "moment-timezone";
import * as room from "@/store/modules/room";

interface RoomData {
  title: string;
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
    getTopicInfo(): string {
      if (this.topic === undefined) {
        return "unknown";
      }
      return `${this.topic.getId()}: ${this.topic.getName()}`
    },
    getMemberUser(userId: number): User {
      const user = room.readMember(this.$store)
        .find(e => e.getId() === userId)
      if (user === undefined) {
        const unUser = new User();
        unUser.setName("system");
        return unUser;
      }
      return user;
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
        console.log(res.getMemberList());
        room.commitInitializeMember(this.$store, res.getMemberList());
        chat.dispatchRecvMessage(this.$store, payload).then(res => {
          res.on("data", res => {
            this.messages.push(res);
          });
        });
        chat.dispatchRecvMember(this.$store, payload).then(res => {
          res.on("data", res => {
            const systemJoinMessage = new ChatMessage();
            systemJoinMessage.setText(`${res.getName()} さんが入室しました。`);

            const now = moment.tz("Asia/Tokyo").format("YYYY-MM-DD HH:mm:ss");
            systemJoinMessage.setActionedAt(now);

            this.messages.push(systemJoinMessage);
            room.commitAddMember(this.$store, res);
          });
        });
      });
    });
  },
  render() {
    return (
      <div>
        {this.messages.map(e => (
          <div class="media my-sm-2">
            <img
              src="//storage.googleapis.com/topic-chat/images/guest_user.png"
              width="64"
              height="64"
              class="bd-placeholder-img mr-3"
            />
            <div class="media-body text-monospace">
              <h5 class="mt-0">
                { this.getMemberUser(e.getUserId())?.getName() ?? "system" }{" "}
                <small class="text-muted">{ e.getActionedAt() }</small>
              </h5>
              { e.getText() }
            </div>
          </div>
        ))}

        <div
          class="input-group mb-3 mx-sm-2 position-absolute"
          style="bottom: 0; right: 0;"
        >
          <input
            type="text"
            class="form-control"
            placeholder="Message"
            onInput={e => {
              this.sendText = e.target.value?.toString() ?? "";
            }}
          />
          <div class="input-group-append">
            <button
              class="btn btn-outline-secondary"
              type="button"
              onClick={() => {
                const payload = {
                  topicId: user.readGetSelectedTopic(this.$store)?.getId() ?? -1,
                  message: this.message()
                };
                chat.dispatchSendMessage(this.$store, payload);
              }}
            >
              Send
            </button>
          </div>
        </div>
      </div>
    );
  }
});
