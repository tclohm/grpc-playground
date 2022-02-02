#!/bin/bash

protoc greet/greetpb/greet.proto --go_out=. --go-grpc_out=.

protoc calc/calcpb/calc.proto --go_out=. --go-grpc_out=.

protoc blog/blogpb/blog.proto --go_out=. --go-grpc_out=.