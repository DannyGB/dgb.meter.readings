FROM golang:1.19.1 AS build

RUN mkdir -p /app/elec

COPY . /app/elec

WORKDIR /app/elec/cmd/meter.readings

RUN go build -o meter.readings .

FROM golang:1.19.1 AS final

RUN mkdir -p /app/elec/cmd/meter.readings
WORKDIR /app/elec/cmd/meter.readings

COPY --from=build /app/elec/cmd/meter.readings/meter.readings .

ARG ENVIRONMENT=dev
ENV METER_READINGS_ENVIRONMENT $ENVIRONMENT

ARG HTTP_PORT=80
ENV METER_READINGS_HTTP_PORT $HTTP_PORT

ARG MONGO_COLLECTION=readings
ENV METER_READINGS_MONGO_COLLECTION $MONGO_COLLECTION

ARG MONGO_DB=meter
ENV METER_READINGS_MONGO_DB $MONGO_DB

ARG MONGO_CONNECTION
ENV METER_READINGS_MONGO_CONNECTION $MONGO_CONNECTION

ARG CORS_CLIENTS="*"
ENV METER_READINGS_CORS_CLIENTS $CORS_CLIENTS

CMD [ "/app/elec/cmd/meter.readings/meter.readings" ]
