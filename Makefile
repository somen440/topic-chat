PROTODIR=pb
PROJECT=topic-chat

AUTH_SERVICE_OUT_DIR=src/auth_service
FRONTEND_OUT_DIR=src/frontend
CHAT_SERVICE_OUT_DIR=src/chat_service
CHAT_CLIENT_OUT_DIR=src/chat_client/src

PB_OUT_DIR=src/common

genproto:
	protoc -I/usr/local/include -I. \
      -I$(GOPATH)/src \
			--go_out=plugins=grpc:$(PB_OUT_DIR) \
			--js_out=import_style=commonjs:$(CHAT_CLIENT_OUT_DIR) \
			--grpc-web_out=import_style=typescript,mode=grpcwebtext:$(CHAT_CLIENT_OUT_DIR) \
			-I$(PROTODIR) \
			$(PROTODIR)/topicchat.proto

####################################################################

build: build_frontend build_topic_catalog build_auth build_chat build_chat_client
.PHONY: build

build_frontend:
	cd src/frontend && docker build -t gcr.io/$(PROJECT)/frontend .
.PHONY: build_frontend

build_topic_catalog:
	cd src/topic_catalog_service && docker build -t gcr.io/$(PROJECT)/topic_catalog .
.PHONY: build_topic_catalog

build_auth:
	cd src/auth_service && docker build -t gcr.io/$(PROJECT)/auth .
.PHONY: build_auth

build_chat:
	cd src/chat_service && docker build -t gcr.io/$(PROJECT)/chat .
.PHONY: build_chat

build_chat_client:
	cd src/chat_client && \
		export NODE_ENV=debug && \
		yarn build && \
		docker build -t gcr.io/$(PROJECT)/chat_client .
.PHONY: build_chat_client

####################################################################

test:
	cd src/topic_catalog_service/ && make test
	cd src/auth_service/ && make test
.PHONY: test

fmt:
	cd src/topic_catalog_service/ && make fmt
	cd src/auth_service/ && make fmt
.PHONY: fmt

####################################################################

debug_catalog_list:
	grpcurl -plaintext \
		-import-path pb/ \
		-proto topicchat.proto \
		localhost:8081 \
		topicchat.TopicCatalogService/ListTopics
.PHONY: debug_catalog_list

debug_chat_create:
	grpcurl -import-path pb/ \
		-proto topicchat.proto \
		-d "{\"topicId\":$(TOPIC_ID)}" \
		-plaintext -v localhost:8083 topicchat.ChatService/CreateRoom
.PHONY: debug_chat_create

debug_chat_list:
	grpcurl -import-path pb/ \
		-proto topicchat.proto \
		-plaintext -v localhost:8083 topicchat.ChatService/ListRoom
.PHONY: debug_chat_list

debug_chat_join:
	grpcurl -import-path pb/ \
		-proto topicchat.proto \
		-d '{"userId": $(USER_ID), "topicId": $(TOPIC_ID)}' \
		-plaintext -v localhost:8083 topicchat.ChatService/JoinRoom
	grpcurl -import-path pb/ \
		-proto topicchat.proto \
		-d '{"userId": $(USER_ID), "topicId": $(TOPIC_ID)}' \
		-plaintext -v localhost:8083 topicchat.ChatService/RecvMessage
.PHONY: debug_chat_recv

debug_chat_send:
	grpcurl -import-path pb/ \
		-proto topicchat.proto \
		-d '{"message":{"text":"$(TEXT)","user":{"id":$(USER_ID),"name":"hoge","loggedIn":"true"}},"topicId":$(TOPIC_ID)}' \
		-plaintext -v localhost:8083 topicchat.ChatService/SendMessage
.PHONY: debug_chat_send

debug_user_get:
	grpcurl -import-path pb/ \
		-proto topicchat.proto \
		-d '{"userId": $(USER_ID)}' \
		-plaintext -v localhost:8082 topicchat.AuthService/GetUser
.PHONY: debug_user_get

debug_user_get_all:
	grpcurl -import-path pb/ \
		-proto topicchat.proto \
		-plaintext -v localhost:8082 topicchat.AuthService/GetUserAll
.PHONY: debug_user_get_all

debug_topic_get_all:
	grpcurl -import-path pb/ \
		-proto topicchat.proto \
		-plaintext -v localhost:8081 topicchat.TopicCatalogService/ListTopics
.PHONY: debug_topic_get_all

debug_topic_get:
	grpcurl -import-path pb/ \
		-proto topicchat.proto \
		-d '{"topicId":$(TOPIC_ID)}' \
		-plaintext -v localhost:8081 topicchat.TopicCatalogService/GetTopic
.PHONY: debug_topic_get
