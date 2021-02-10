FROM golang:1.15 as build

WORKDIR /src
COPY src /src/
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -o /bin/workoutReminderBot
RUN chmod +x /bin/workoutReminderBot

FROM scratch
# golang image already have certificates https://stackoverflow.com/a/52979541
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /bin/workoutReminderBot /bin/workoutReminderBot
ENTRYPOINT ["/bin/workoutReminderBot"]