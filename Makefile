PROTODIR=pb

genproto:
	protoc -I/usr/local/include -I. \
      -I$(GOPATH)/src \
			--go_out=plugins=grpc:src/auth_service \
			--go_out=plugins=grpc:src/frontend_service \
			--go_out=plugins=grpc:src/room_service \
			--go_out=plugins=grpc:src/topic_catalog_service \
			-I$(PROTODIR) \
			$(PROTODIR)/topicchat.proto
