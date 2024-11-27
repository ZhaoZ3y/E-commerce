FROM golang:1.23

WORKDIR /usr/src/gomall

# Set the GOPROXY environment variable to use the Go module proxy in China
ENV GOPROXY=https://goproxy.cn,direct

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY app/frontend/go.mod app/frontend/go.sum ./app/frontend/
COPY rpc_gen rpc_gen

RUN cd app/frontend && go mod download && go mod verify

COPY app/frontend app/frontend/

RUN cd app/frontend && go build -v -o /opt/gomall/frontend/server

COPY app/frontend/conf /opt/gomall/frontend/conf
COPY app/frontend/template /opt/gomall/frontend/template
COPY app/frontend/static /opt/gomall/frontend/static

EXPOSE 8080

CMD ["/opt/gomall/frontend/server"]
