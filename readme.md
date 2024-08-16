# cnho
**cnho** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started by Download
### Download the Binary cnho_stables
[Centos7+](https://159.138.232.248/resource/ChainCNHO/cnho_stables1.0.0.bin)
### Download the file genesis.json 
[genesis.json](https://github.com/alashooinc/ChainCNHO/blob/master/mainnet/genesis.json)
#### Build The Node
```
Put them to the home path of host and than do such like bellow
$ mv cnho_stablesx.x.x.bin cnho_stablesx
$ chmod +x cnho_stables
$ ./cnho_stables init cnho_stables-2
$ cp -rf genesis.json ~/.cnho/config/genesis.json
$ vi ~/.cnho/config/config.toml

Modify Field of [seeds] as seeds="94e013d19485dc70a4a20477b7879cda4af47b55@172.105.116.194:26656"
```
### Run The Chain
```
$ ./cnho_stables start
```

## Get started with code from github
### System Requirements
```
Release:Ubuntu 20.04
```
### Build Binary
```
cd ${codePath}
git clone https://github.com/alashooinc/ChainCNHO
export GOROOT=~/gdk/go1.19.13
cd ChainCNHO
go build -o cnho_stables -gcflags all="-N -l" cnho/cmd/cnhod
ls -a cnho_stables
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Build The Node 
```
cd ${codePath}/ChainCNHO
chmod +x cnho_stables
./cnho_stables init cnho_stables-2
cp -rf mainnet/genesis.json ~/.cnho/config/genesis.json
vi ~/.cnho/config/config.toml
Modify Field of [seeds] as seeds="94e013d19485dc70a4a20477b7879cda4af47b55@172.105.116.194:26656"
vi ~/.cnho/config/client.toml
Modify Field of [chain-id] as chain-id = "cnho_stables-1"
```
### Run The Chain
```
cd ${codePath}/ChainCNHO

./cnho_stables start

```
### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).


## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
