FROM golang as builder
WORKDIR /src/rms-post
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build  -ldflags "-X main.Version=`git tag --sort=-version:refname | head -n 1`" -o rms-media-discovery -a -installsuffix cgo rms-post.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
RUN mkdir /app
WORKDIR /app
COPY --from=builder /src/rms-post/rms-post .
COPY --from=builder /src/rms-post/configs/rms-post.json /etc/rms/
EXPOSE 8080/tcp
EXPOSE 2112/tcp
CMD ["./rms-post"]