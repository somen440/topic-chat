import * as tsx from "vue-tsx-support";
import * as auth from "@/store/modules/auth";
import * as user from "@/store/modules/user";

export default tsx.component({
  name: "Join",
  data() {
    return {
      title: "Join",
      name: ""
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
        <h1>{this.title}</h1>
        <input
          type="text"
          onInput={(e) => {
            this.name = e.target.value?.toString() ?? "";
          }}
        />
        {this.validName && (
          <button
            onClick={() => {
              console.log("click");
              auth.dispatchJoin(this.$store, this.name)
                .then(res => {
                  res.on("data", res => {
                    user.commitLoggedIn(this.$store, res);
                    this.$router.push("/");
                  });
                });
            }}
          >register</button>
        )}
      </div>
    );
  }
});
