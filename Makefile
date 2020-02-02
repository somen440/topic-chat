PROTODIR=pb
PROJECT=topic-chat

genproto:
	protoc -I/usr/local/include -I. \
      -I$(GOPATH)/src \
			--go_out=plugins=grpc:src/auth_service \
			--go_out=plugins=grpc:src/frontend_service \
			--go_out=plugins=grpc:src/room_service \
			--go_out=plugins=grpc:src/topic_catalog_service \
			-I$(PROTODIR) \
			$(PROTODIR)/topicchat.proto

####################################################################

build: build_frontend build_topic_catalog
.PHONY: build

build_frontend:
	cd src/frontend && docker build -t gcr.io/$(PROJECT)/frontend .
.PHONY: build_frontend

build_topic_catalog:
	cd src/topic_catalog_service && docker build -t gcr.io/$(PROJECT)/topic_catalog .
.PHONY: build_topic_catalog
