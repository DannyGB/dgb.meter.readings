# 🦊 dgb.meter.readings

A playground for the Go language

# Installation

Requires:

- [Go](https://go.dev/)
- A [MongoDB](https://www.mongodb.com) instance

# Config

An environment variable called `METER_READINGS_ENVIRONMENT` must be set.
This will form the prefix to a required file called _config.json.

For instance in development you would have a file called `dev_config.json` and you would set the `METER_READINGS_ENVIRONMENT` environment variable to `dev`

The `dev_config.json` file should be in the directory you are running the app from. I usually place it in the `cmd/meter.readings` directory

The `_config.json` file should look like this:

```json
{
    "MONGO_CONNECTION": "<your mongodb connection string>",
    "MONGO_COLLECTION": "<the name of the meter readings collection in mongo>",
    "MONGO_DB": "<the name of the mongo db containing the readings collection",
    "HTTP_PORT": "<the port the Api should listen on"
}
```
