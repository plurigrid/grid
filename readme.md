# Plurigrid
- **2023-01-01** Testnet #0: Edith Clarke
- **2023-02-01** Testnet #1: Marie Curie

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

### gridd
```
gridd
Plurigrid testnet #0: Edith Clarke

Usage:
  gridd [command]

Available Commands:
  add-genesis-account       Add a genesis account to genesis.json
  collect-gentxs            Collect genesis txs and output a genesis.json file
  config                    Create or query an application CLI configuration file
  debug                     Tool for helping with debugging your application
  export                    Export state to JSON
  gentx                     Generate a genesis tx carrying a self delegation
  help                      Help about any command
  init                      Initialize private validator, p2p, genesis, and application configuration files
  keys                      Manage your application's keys
  migrate                   Migrate genesis to a specified target version
  query                     Querying subcommands
  rollback                  rollback cosmos-sdk and tendermint state by one height
  start                     Run the full node
  start-with-http-tunneling Run the full node with http tunneling
  status                    Query remote node for status
  tendermint                Tendermint subcommands
  tx                        Transactions subcommands
  validate-genesis          validates the genesis file at the default location or at the location passed as an arg
  version                   Print the application binary version information

Flags:
  -h, --help                help for gridd
      --home string         directory for config and data (default "/Users/barton/.grid")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```
