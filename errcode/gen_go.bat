@echo off

set OUTPUT=%CD%

:COMPILE_PROTOBUF
protoc --go_out=plugins=grpc:%OUTPUT% codes.proto

copy /Y %OUTPUT%\huastart.com\pub\errcode\*.go .\
rmdir /S /Q %OUTPUT%\huastart.com

:GEN_CODE
echo Generate pb/src/error_codes.go
set PB_PATH=%OUTPUT%/../../../pb
cd tool
go run .
cd ..

:end

