# True Tickets Sum Metric Service Problem

### `Running application`
```sh
$   go run main.go
The app will start at localhost:8080.
```

### `Configuring environment variables`
```sh
In .env file you can configure the time each action will occur.

The default value to metrics expiration is one hour (3600 seconds). You can adjust it to lower values to debug the application in METRICS DURATION IN SECONDS.

The env variable TIME_TICKER_IN_SECONDS defines the worker time frequency to check metrics expiration and updated it if applied.
```

### `App operation`
```sh
There is a worker that from time to time checks if the metrics in the queue head had expired.

If so those metrics are updated in the memory repository and dequeued.

When adding a new metric it´s added to the end of the queue and the metric value is updated.

When retrieving the metric it´s read directly in the memory repository where it´s up to date.
```

### `Testing`
```sh
There is a Insomnia collection in the folder ./collections that can be used to tests
```
[Download Insomnia](https://insomnia.rest/download)

| Route                 | Http verb | Description                                                         |
| --------------------- | --------- | ------------------------------------------------------------------- |
|`/metric/{key}/sum`    | `GET`     |  `return a json with the metric value {"value" : value}`            |
|`/metric/{key}`        | `POST`    | `send in the request body a json in the format {"value" : value} `  |