FROM golang:1.16-alpine

#"Bee" (Watcher)
RUN go get -u github.com/beego/bee

#Delve (Debugger)
RUN go install github.com/go-delve/delve/cmd/dlv@latest

#Delve (Debugger)
RUN apk add --no-cache libc6-compat

ARG APP_USER=app
ARG APP_HOME=/app
ARG GROUP_ID
ARG USER_ID

RUN addgroup -g $GROUP_ID -S $APP_USER && \
    adduser -u $USER_ID -S $APP_USER -G $APP_USER

RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME

USER $APP_USER
WORKDIR $APP_HOME

ENV CGO_ENABLED 0
ENV GO111MODULE on
ENV GOFLAGS "-mod=vendor"

CMD ["bee", "run"]