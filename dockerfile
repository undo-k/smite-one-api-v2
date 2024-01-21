FROM alpine as ssh-key-stage
ARG SSH_PRIVATE_KEY
RUN mkdir /root/.ssh/ && echo "$SSH_PRIVATE_KEY" > /root/.ssh/id_rsa && chmod 600 /root/.ssh/id_rsa

FROM node:latest as vue-build
WORKDIR /app
RUN git clone https://github.com/undo-k/smite-one.git .
RUN npm install
RUN npm run build

FROM golang:1.20 as go-build
COPY --from=ssh-key-stage /root/ssh/id_rsa /root/.ssh/id_rsa
WORKDIR /server
RUN git clone https://github.com/undo-k/smite-one-api-v2.git .
RUN go build -o goserver

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=go-build /server/goserver .
COPY --from=vue-build /app/dist /root/web/app/dist
EXPOSE 8080
CMD ["./goserver"]
