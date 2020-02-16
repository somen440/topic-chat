import { CounterState, TopicCatalogState, AuthState, ChatState } from "./modules";

export interface State {
  auth: AuthState,
  chat: ChatState,
  counter: CounterState,
  topicCatalog: TopicCatalogState,
}
