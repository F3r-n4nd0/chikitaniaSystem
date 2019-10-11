#!/usr/bin/env bash


protoc ../protobuf/*.proto --go_out=plugins=grpc:../ --proto_path=../