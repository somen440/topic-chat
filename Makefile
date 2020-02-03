PROTODIR=pb
PROJECT=topic-chat

genproto:
	protoc -I/usr/local/include -I. \
      -I$(GOPATH)/src \
			--go_out=plugins=grpc:src/auth_service \
			--go_out=plugins=grpc:src/frontend \
			--go_out=plugins=grpc:src/room_service \
			--go_out=plugins=grpc:src/topic_catalog_service \
			--go_out=plugins=grpc:src/chat_service \
			-I$(PROTODIR) \
			$(PROTODIR)/topicchat.proto

####################################################################

build: build_frontend build_topic_catalog build_auth build_chat
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

####################################################################

debug_catalog_list:
	grpcurl -plaintext \
		-import-path pb/ \
		-proto topicchat.proto \
		localhost:8081 \
		topicchat.TopicCatalogService/ListTopics
.PHONY: debug_catalog_list

debug_chat_stream:
	grpcurl -import-path pb/ \
		-proto topicchat.proto \
		-plaintext -v localhost:8083 topicchat.ChatService/GetStream
.PHONY: debug_chat_stream

debug_chat_send:
	grpcurl -import-path pb/ \
		-proto topicchat.proto \
		-d "{\"text\":\"$(TEXT)\"}" \
		-plaintext -v localhost:8083 topicchat.ChatService/Send
.PHONY: debug_chat_stream
