FROM golang:1.16 as build
RUN go install github.com/beego/bee/v2@latest

ENV GO111MODULE=on

RUN mkdir -p "foss/app"
WORKDIR /foss
COPY app /foss/app
COPY go.mod ./
RUN go mod download

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN cd /foss/app &&  go build ./main.go

FROM nginx:latest as packager
WORKDIR /
COPY ./resources ./resources/
RUN tar -czvf "FossApp.tar.gz" ./resources/

FROM scratch as runtime
COPY --from=build /foss/app /
COPY --from=packager "FossApp.tar.gz" ./resources/

EXPOSE 8081/tcp
CMD ["./main"]
#ENTRYPOINT ["./main"]