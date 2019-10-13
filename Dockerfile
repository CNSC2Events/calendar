FROM golang:1.13.0-alpine AS BUILDER

WORKDIR /builder

ADD . /builder

RUN go build -o calendar *.go

FROM alpine

WORKDIR /sc2_gcal

COPY --from=BUILDER  /builder/calendar /sc2_gcal/calendar

ADD resource/credentials.json /sc2_gcal/resource/credentials.json

ADD events_gcal.db /sc2_gcal/events_gcal.db

RUN apk add tzdata && cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" >  /etc/timezone

COPY cronjobs /etc/crontabs/root

CMD ["crond", "-f", "-d", "8"]
