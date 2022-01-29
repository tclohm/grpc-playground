#!/bin/bash

protoc -I=greet/greetpb --go_out=. greet/greetpb/greet.proto