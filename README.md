# 🦊 dgb.meter.readings

A playground for the Go language

# Pre-requisites

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

# Usage

The included Dockerfile can be used to run the application

Under linux or Wsl create a symlink to the dockerfile in the root directory using

    ln -s ./build/Dockerfile Dockerfile

Build the image from the root directory

    docker build -t <your tag> .

The docker image contains an environment variable argument called `ENVIRONMENT` that can be supplied to change the config that the application loads on startup. The default value us `dev`

    docker build -t <your tag> --build-arg ENVIRONMENT=prod .

Run the image

    docker run -it -d --rm -p 8000:8000 <your tag>
