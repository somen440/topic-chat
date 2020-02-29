import * as tsx from "vue-tsx-support";
import * as auth from "@/store/modules/auth";
import * as user from "@/store/modules/user";

export default tsx.component({
  name: "JoinForm",
  data() {
    return {
      title: "JoinForm",
      name: "",
      avatorUrl:
        "https://storage.googleapis.com/topic-chat/images/animals/ris.png",
      defaultAvators: [
        // todo: あとでサムネイル選択できるようにしてから消す
        {
          url:
            "https://storage.googleapis.com/topic-chat/images/animals/ris.png",
          value: "リスさん",
          selected: true
        },
        {
          url:
            "https://storage.googleapis.com/topic-chat/images/animals/pig.png",
          value: "ブタさん",
          selected: false
        },
        {
          url:
            "https://storage.googleapis.com/topic-chat/images/animals/penguin.png",
          value: "ペンギンさん",
          selected: false
        },
        {
          url:
            "https://storage.googleapis.com/topic-chat/images/animals/panda.png",
          value: "パンダさん",
          selected: false
        },
        {
          url:
            "https://storage.googleapis.com/topic-chat/images/animals/kuma.png",
          value: "クマさん",
          selected: false
        },
        {
          url:
            "https://storage.googleapis.com/topic-chat/images/animals/elephant.png",
          value: "ゾウさん",
          selected: false
        },
        {
          url:
            "https://storage.googleapis.com/topic-chat/images/animals/chiken.png",
          value: "ニワトリさん",
          selected: false
        },
        {
          url:
            "https://storage.googleapis.com/topic-chat/images/animals/bake.png",
          value: "バケモノさん",
          selected: false
        }
      ]
    };
  },
  methods: {
    validName(): boolean {
      return this.name !== "";
    }
  },
  render() {
    return (
      <div>
        <div class="form-group">
          <label for="inputName">
            こんにちは。あなたのお名前を教えてください。
          </label>
          <input
            id="inputName"
            class="form-control"
            type="text"
            onInput={e => {
              this.name = e.target.value?.toString() ?? "";
            }}
          />
        </div>

        <div class="input-group mb-3">
          <div class="input-group-prepend">
            <label class="input-group-text" for="avatorSelect">
              Avator
            </label>
          </div>
          <select
            class="custom-select"
            id="avatorSelect"
            onChange={e => {
              if (e.target.value === undefined) {
                return;
              }
              this.avatorUrl = e.target.value?.toString();
            }}
          >
            {this.defaultAvators.map(({ url, selected, value }) => (
              <option value={url} selected={selected}>
                {value}
              </option>
            ))}
          </select>
        </div>

        {this.validName && (
          <button
            class="btn btn-outline-primary btn-block"
            onClick={() => {
              const payload = { name: this.name, avator: this.avatorUrl };
              auth.dispatchJoin(this.$store, payload).then(res => {
                res.on("data", res => {
                  user.commitLoggedIn(this.$store, res);
                  this.$router.push("/topic");
                });
              });
            }}
          >
            register
          </button>
        )}
      </div>
    );
  }
});
