import { ChatServiceClient } from "./pb/TopicchatServiceClientPb";
import { Empty } from "./pb/topicchat_pb";

const client = new ChatServiceClient('http://localhost:9090');
const stream = client.recvMessage(new Empty(), {})
stream
  .on("data", (response) => {
    console.log("data");
    console.log(response.getText());
    const li = document.createElement("li");
    li.innerHTML = response.getText();
    document.getElementById("messages").appendChild(li)
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
