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

build_frontend:
	cd src/frontend && go build -o frontend

build_topic_catalog:
	cd src/topic_catalog_service && go build -o topic_catalog

####################################################################

build_image: build_frontend_image build_topic_catalog_image
.PHONY: build_image

build_frontend_image:
	cd src/frontend && docker build -t gcr.io/$(PROJECT)/frontend .
.PHONY: build_frontend_image

build_topic_catalog_image:
	cd src/topic_catalog_service && docker build -t gcr.io/$(PROJECT)/topic_catalog .
.PHONY: build_topic_catalog_image
