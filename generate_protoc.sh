#!/bin/bash

protoc calculatorpb/calculator.proto --go_out=plugins=grpc:.
