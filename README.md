# name-generator

Uses Docker's [names generator package](https://github.com/moby/moby/blob/master/pkg/namesgenerator/names-generator.go)
to generate names anywhere.

You get a random name running:

```sh
# downloads the binary and executes it locally:
curl -sfL https://git.io/name-generator | sh

# get name from a serverless function:
curl -sf https://name-generator.now.sh

# get name from a serverless function (json):
curl -sf -H 'accept: application/json' https://name-generator.now.sh
```

> The source is the `run.sh` file in the root of this repo.
