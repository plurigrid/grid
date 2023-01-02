# Plurigrid
2022-01-01 Testnet #0: Edith Clarke
2022-02-01 Testnet #1: Marie Curie

<img width="1587" alt="clarke" src="https://user-images.githubusercontent.com/1236584/210226945-e9bd74cf-09bf-4a1d-9009-d2515f4ffe6e.png">

## To contribute, make a Pull Request and join [#testnet on our Zulip](https://plurigrid.zulipchat.com/join/5vffrr3jwddnywzopwyagkud/)
Record batched data about power output from solar panels and other renewable energy sources. ☀️⚡️🔋

Record batched data about what private capacity model estimates for any given time period (using https://github.com/Plurigrid/moose).
## Get started

##  Set up Celestia Light Node locally
https://docs.celestia.org/nodes/light-node

## Set up Plurigrid testet node locally
Clone this repo and make sure you install Ignite:
```
curl https://get.ignite.com/cli | bash
sudo mv ignite /usr/local/bin/
```

Then:
```
ignite chain build
```

`build` command installs dependencies, builds.

## Run Plurigrid test node locally

```
bash init.sh
```

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```
