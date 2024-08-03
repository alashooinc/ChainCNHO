# cnho
**cnho** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started
### System


### Build Binary
```
cd ${codePath}
git clone https://github.com/alashooinc/ChainCNHO
export GOROOT=~/gdk/go1.19.13
git clone https://github.com/alashooinc/ChainCNHO
cd ChainCNHO
go build -o cnho_stables -gcflags all="-N -l" cnho/cmd/cnhod
ls -a cnho_stables
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Build Node 
```
cd ${codePath}/ChainCNHO
chmod +x cnho_stables
./cnho_stables init cnho_stables-2
cp -rf mainnet/genesis.json ~/.cnho/config/genesis.json
vi ~/.cnho/config/config.toml
Modify Field of [seeds] as seeds="c03cc5c48b1b9aba7ec5258510d386de6e88806@172.105.116.194:26656"
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
