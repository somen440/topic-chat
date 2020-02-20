import {
  CounterState,
  TopicCatalogState,
  AuthState,
  ChatState,
  UserState,
  RoomState
} from "./modules";

export interface State {
  auth: AuthState;
  chat: ChatState;
  counter: CounterState;
  room: RoomState;
  topicCatalog: TopicCatalogState;
  user: UserState;
}
