FROM golang:latest

LABEL "com.github.action.name"="Go-workflow"
LABEL "com.github.actions.description"="Go build, test,output"
LABEL "com.github.actions.icon"="mic"
LABEL "com.github.actions.color"="purple"

ENV GO111MODULE=on

ADD entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
