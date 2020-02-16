import { CounterState, TopicCatalogState, AuthState, ChatState, UserState } from "./modules";

export interface State {
  auth: AuthState,
  chat: ChatState,
  counter: CounterState,
  topicCatalog: TopicCatalogState,
  user: UserState,
}
