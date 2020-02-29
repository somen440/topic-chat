import * as tsx from "vue-tsx-support";
import JoinForm from "@/components/JoinForm";
import * as user from "@/store/modules/user";

export default tsx.component({
  name: "Home",
  data() {
    return {
      title: "Home"
    };
  },
  created() {
    if (user.readIsLoggedIn(this.$store)) {
      this.$router.push("/topic");
      return;
    }
  },
  render() {
    return (
      <div>
        <JoinForm />
      </div>
    );
  }
});
