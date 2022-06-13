# ShipLens Example App

This is an example app used to test out [shiplens.io](https://shiplens.io).

This project is deployed automatically to [beta.example.shiplens.io](https://beta.example.shiplens.io).

It can be manually deployed to [example.shiplens.io](https://example.shiplens.io) by running `make deploy-production`.

## Environments

| Environment | URL                                 |
|-------------|-------------------------------------|
| Production  | https://example.shiplens.io         |
| Staging     | https://staging-example.shiplens.io |

## ShipLens Config

This configuration file can be used to test ShipLens against this project:

```yaml
# TODO
```

## API

This app exposes version information from multiple sources in various formats.

### `HEAD /`

```
Server: example/3188c9337f28805a281dc869fcf12c0dd9f9f578
```

### `GET /json`

```json
{
  "git": {
    "sha": "3188c9337f28805a281dc869fcf12c0dd9f9f578"
  },
  "time": {
    "iso8601": "2022-06-06T20:10:25Z"
  }
}
```
