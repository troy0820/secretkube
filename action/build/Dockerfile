FROM golang:1.16.2-alpine

LABEL "com.github.action.name"="Go-workflow"
LABEL "com.github.actions.description"="Go build, test,output"
LABEL "com.github.actions.icon"="mic"
LABEL "com.github.actions.color"="purple"

ENV GO111MODULE=on
RUN apk add --no-cache git gcc clang
ADD entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
