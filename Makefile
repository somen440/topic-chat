PROTODIR=pb
PROJECT=topic-chat

AUTH_SERVICE_OUT_DIR=src/auth_service
FRONTEND_OUT_DIR=src/frontend
TOPIC_CATALOG_SERVICE_OUT_DIR=src/topic_catalog_service
CHAT_SERVICE_OUT_DIR=src/chat_service
CHAT_CLIENT_OUT_DIR=src/chat_client/src

genproto:
	protoc -I/usr/local/include -I. \
      -I$(GOPATH)/src \
			--go_out=plugins=grpc:$(AUTH_SERVICE_OUT_DIR) \
			--go_out=plugins=grpc:$(FRONTEND_OUT_DIR) \
			--go_out=plugins=grpc:$(TOPIC_CATALOG_SERVICE_OUT_DIR) \
			--go_out=plugins=grpc:$(CHAT_SERVICE_OUT_DIR) \
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
		yarn build && \
		docker build -t gcr.io/$(PROJECT)/chat_client .
.PHONY: build_chat_client

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

debug_chat_recv:
	grpcurl -import-path pb/ \
		-proto topicchat.proto \
		-d "{\"topicId\":$(TOPIC_ID)}" \
		-plaintext -v localhost:8083 topicchat.ChatService/RecvMessage
.PHONY: debug_chat_recv

debug_chat_send:
	grpcurl -import-path pb/ \
		-proto topicchat.proto \
		-d '{"message":{"text":"$(TEXT)"},"topicId":$(TOPIC_ID)}' \
		-plaintext -v localhost:8083 topicchat.ChatService/SendMessage
.PHONY: debug_chat_send

debug_user_get:
	grpcurl -import-path pb/ \
		-proto topicchat.proto \
		-d '{"userId": $(USER_ID)}' \
		-plaintext -v localhost:8082 topicchat.AuthService/GetUser
.PHONY: debug_chat_send

debug_user_get_all:
	grpcurl -import-path pb/ \
		-proto topicchat.proto \
		-plaintext -v localhost:8082 topicchat.AuthService/GetUserAll
.PHONY: debug_chat_send
