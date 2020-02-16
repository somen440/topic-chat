import { ActionContext } from "vuex";
import { State as RootState } from "@/store/state"
import { getStoreAccessors } from "vuex-typescript";
import { TopicCatalogServiceClient } from "@/pb/TopicchatServiceClientPb"
import { Topic, AddTopicRequest, DeleteTopicRequest, GetTopicRequest, Empty, ListTopicsResponse, UpdateTopicRequest } from "@/pb/topicchat_pb";
import { ClientReadableStream } from 'grpc-web';

const SCHEME = process.env.SCHEME ?? "http";
const TOPIC_CATALOG_SERVICE_ADDR = process.env.TOPIC_CATALOG_SERVICE_ADDR ?? "localhost:9092";

export interface TopicCatalogState {
  client: TopicCatalogServiceClient,
}

type TopicCatalogContext = ActionContext<TopicCatalogState, RootState>;

const { commit, read, dispatch } = getStoreAccessors<TopicCatalogState, RootState>(
  "topicCatalog"
);

//
// state
//
const state = {
  client: new TopicCatalogServiceClient(`${SCHEME}://${TOPIC_CATALOG_SERVICE_ADDR}`)
}

//
// getters
//
const getters = {
  getClient(state: TopicCatalogState): TopicCatalogServiceClient {
    return state.client;
  }
}

const readClient = read(getters.getClient);

//
// mutations
//
const mutations = {};

//
// actions
//
const actions = {
  add(context: TopicCatalogContext, topic: Topic): ClientReadableStream<Topic> {
    const req = new AddTopicRequest();
    req.setTopic(topic);
    return readClient(context).addTopic(req, {}, (err) => {
      console.log({ err });
    });
  },
  delete(context: TopicCatalogContext, topicId: number): ClientReadableStream<Empty> {
    const req = new DeleteTopicRequest();
    req.setTopicId(topicId);
    return readClient(context).deleteTopic(req, {}, (err) => {
      console.log({ err });
    });
  },
  get(context: TopicCatalogContext, topicId: number): ClientReadableStream<Topic> {
    const req = new GetTopicRequest();
    req.setTopicId(topicId);
    return readClient(context).getTopic(req, {}, (err) => {
      console.log({ err });
    });
  },
  getAll(context: TopicCatalogContext): Promise<ClientReadableStream<ListTopicsResponse>> {
    const req = new Empty();
    console.log("req");

    return new Promise((resolve) => {
      const stream = readClient(context).listTopics(req, {}, (err) => {
        console.log({ err });
      });
      resolve(stream);
    });
  },
  update(context: TopicCatalogContext, topic: Topic) {
    const req = new UpdateTopicRequest();
    req.setTopic(topic);
    return readClient(context).updateTopic(req, {}, (err) => {
      console.log({ err });
    });
  }
}

export const dispatchGetAll = dispatch(actions.getAll)

export const topicCatalog = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
};
