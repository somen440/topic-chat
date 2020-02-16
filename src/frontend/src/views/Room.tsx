import * as tsx from "vue-tsx-support";
import * as chat from "@/store/modules/chat";
import * as user from "@/store/modules/user";
import { User, ChatMessage } from '@/pb/topicchat_pb';

interface RoomData {
  title: string
  member: User[]
  messages: ChatMessage[],
  sendText: string
}

export default tsx.component({
  name: "Room",
  data(): RoomData {
    return {
      title: "Room",
      member: [],
      messages: [],
      sendText: ""
    };
  },
  methods: {
    validSend(): boolean {
      return this.sendText !== ""
    },
    message(): ChatMessage {
      const message = new ChatMessage();
      message.setText(this.sendText);
      message.setUser(user.readGetMine(this.$store));
      return message;
    }
  },
  created() {
    const userId = user.readGetMine(this.$store)?.getId() ?? -1;
    const topicId = user.readGetSelectedTopic(this.$store)?.getId() ?? -1;
    const payload = { userId, topicId };

    chat.dispatchJoinRoom(this.$store, payload).then(res => {
      res.on("data", (res) => {
        this.member = res.getMemberList();
        chat.dispatchRecvMessage(this.$store, payload).then(res => {
          res.on("data", (res) => {
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
        <h2>member</h2>
        <ul>
          {this.member.map(e => (
            <li>
            {user.isEqualsUser(user.readGetMine(this.$store), e) ? (
              <b>[my] { e.getId() }: { e.getName() }</b>
            ) : (
              <div>{ e.getId() }: { e.getName() }</div>
            )}
            </li>
          ))}
        </ul>
        <h2>messages</h2>
        <ul>
          {this.messages.map(e => (
            <li>{ e.getUser()?.getName() ?? "unknown" }: { e.getText() }</li>
          ))}
        </ul>
        <input
          type="text"
          onInput={e => {
            this.sendText = e.target.value?.toString() ?? ""
          }}
        />
        {this.validSend() && (
          <button
          onClick={() => {
            const payload = {
              topicId: user.readGetSelectedTopic(this.$store)?.getId() ?? -1,
              message: this.message(),
            }
            chat.dispatchSendMessage(this.$store, payload);
          }}
          >send</button>
        )}
      </div>
    );
  }
});
