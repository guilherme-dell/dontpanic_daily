FROM ubuntu:20.04 AS build

RUN apt-get update && \
    apt-get -y upgrade && \
    apt-get install -y wget && \
    wget https://go.dev/dl/go1.19.2.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz && \
    rm go1.19.2.linux-amd64.tar.gz
    
ENV PATH "$PATH:/usr/local/go/bin"
WORKDIR /app

ADD . .
RUN ["go", "build"]


FROM ubuntu:20.04

WORKDIR /app
COPY --from=build /app/dontpanic .
COPY --from=build /app/public/ ./public

EXPOSE 8081
ENTRYPOINT [ "/app/dontpanic"]