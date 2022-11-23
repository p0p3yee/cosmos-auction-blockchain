# auction
**auction** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Send Transaction

#### Create auction

```
auctiond tx auction create-auction "auction_name" start_price duration --from account_name

// Create an auction from alice with name: exampleAuction
// Starting bid is 100token
// Duration is 100 blocks
auctiond tx auction create-auction "exampleAuction" 100token 100 --from alice
```

#### Finalize Auction

Can only be called by the creator of the auction

```
auctiond tx auction finalize-auction auction_id --from account_name

// Finalize auction id 1 using alice account
auctiond tx auction finalize-auction 1 --from alice
```

### Place Bid

Can only place bid that larger than the previous highest bid

You need sufficient balance in your account to place the bid

token will send to module until someone bid price is higher than you

```
auctiond tx auction place-bid auction_id bidPrice --from account_name

// Place bid for auction id 0, with bid price: 2000token using alice account
auctiond tx auction place-bid 0 2000token --from alice
```

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/auction@latest! | sudo bash
```
`username/auction` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
