FROM golang:1.19.1

RUN mkdir /app

COPY . /app

WORKDIR /app/cmd/meter.readings

ARG ENVIRONMENT=dev
ENV METER_READINGS_ENVIRONMENT $ENVIRONMENT

RUN go build -o meter.readings .

CMD [ "/app/cmd/meter.readings/meter.readings" ]
