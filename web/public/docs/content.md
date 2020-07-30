## Usage

Once sign up, a `Project ID` is generated for your project, to use within your app to make requests to tezoslink.io.

Then, add the tezoslink.io RPC endpoint to your prefered Tezos JS library.

> i.e with [Sotez](https://github.com/AndrewKishino/sotez) :
>```js
> const sotez = new Sotez('https://<NETWORK>.tezoslink.io/v1/<YOUR_PROJECT_ID>');
> ```

## Networks

Use one of these endpoints as your Tezos client provider.

|NETWORK|DESCRIPTION|URL                                              |
|-------|-----------|-------------------------------------------------|
|Mainnet| JSON/RPC  |https://mainnet.tezoslink.io/v1/<YOUR_PROJECT_ID>|
|Carthagenet| JSON/RPC  |https://carthagenet.tezoslink.io/v1/<YOUR_PROJECT_ID>|

## Make requests

```bash
# Be sure to replace YOUR-PROJECT-ID with a Project ID from your Tezos Link dashboard
$ curl https://mainnet.tezoslink.io/v1/<YOUR_PROJECT_ID>/chains/main/blocks/head
```

You should receive the last received block.

## Security

The `Project ID` authorize requests.

## RPC Endpoints

### Whitelisted

All requests of type `/chains/main/blocks(.*?)` are accepted.

>Example of valid paths:
>- `/chains/main/blocks/head/context/contracts/<ADDRESS>/balance`
>- `/chains/main/blocks/head/context/contracts/<ADDRESS>/delegate`
>- `/chains/main/blocks/head/context/contracts/<ADDRESS>/manager_key`
>- `/chains/main/blocks/head/context/contracts/<ADDRESS>/counter`
>- `/chains/main/blocks/head/context/delegates/<ADDRESS>`
>- `/chains/main/blocks/head/header`
>- `/chains/main/blocks/head/votes/proposals`
>- `/chains/main/blocks/head/votes/current_quorum` 

[More about the Tezos `JSON/RPC` endpoints](https://tezos.gitlab.io/api/rpc.html) 

## Nodes

Tezos has three types of nodes:
- Full mode (default mode)
- **Rolling mode**
- **Archive mode**

We use two types of mode:
- **Archive** to store the whole blockchain. Archive is the heaviest mode as it keeps the whole chain data to be able to query any information stored on the chain since the genesis. It is particularly suitable for indexers or block explorer, that is why we use archive nodes.
- **Rolling** to store last blocks (and scale them faster)

> [More about history modes](https://blog.nomadic-labs.com/introducing-snapshots-and-history-modes-for-the-tezos-node.html)
