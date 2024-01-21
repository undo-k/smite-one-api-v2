FROM alpine as ssh-key-stage
RUN mkdir -p /root/.ssh/ &&  \
    chmod 0700 /root/.ssh/
RUN apk update && apk add --no-cache openssh-client git
RUN --mount=type=secret,id=ssh_key \
    cp /run/secrets/ssh_key /root/.ssh/id_rsa && \
    chmod 600 /root/.ssh/id_rsa && \
    ssh-keyscan github.com > /root/.ssh/known_hosts


FROM node:latest as vue-build
WORKDIR /app
RUN git clone https://github.com/undo-k/smite-one.git .
RUN npm install
RUN npm run build

FROM golang:1.20 as go-build
COPY --from=ssh-key-stage /root/.ssh/id_rsa /root/.ssh/id_rsa
COPY --from=ssh-key-stage /root/.ssh/known_hosts /root/.ssh/known_hosts
WORKDIR /server
RUN git clone git@github.com:undo-k/smite-one-api-v2.git .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -a -o goserver -C ./cmd/smite-one-api-v2/

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=go-build /server/cmd/smite-one-api-v2/goserver .
RUN chmod +x /root/goserver
COPY --from=vue-build /app/dist /root/web/app/dist
EXPOSE 8080
CMD ["/root/goserver"]
