FROM golang:1.15 as build

WORKDIR /go/src/github.com/nextuponstream/workoutReminderBot
COPY pkg pkg
COPY cmd cmd
COPY go.mod go.mod
RUN go get -d -v ./...
RUN go test -v ./...

WORKDIR /go/src/github.com/nextuponstream/workoutReminderBot/cmd/bot
RUN CGO_ENABLED=0 go build -v -o /bin/workoutReminderBot

FROM scratch
# golang image already has certificates https://stackoverflow.com/a/52979541
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /bin/workoutReminderBot /bin/workoutReminderBot
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Europe/Berlin
# Note: tested time difference with America/Los_Angeles
ENTRYPOINT ["/bin/workoutReminderBot"]