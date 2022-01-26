<!--
order: 2
-->

# State

The `fantoken` module keeps track of [**parameters**](#Params), [**fan tokens**](#Token), and [**burned coins**](#BurnedCoins).

```
Params:      sdk.Coin
Tokens:      []types.FanToken
BurnedCoins: []sdk.Coin
```

## Params

In the state definition, we can find the **Params**. This section corresponds to a module-wide configuration structure that stores system parameters. In particular, it defines the overall fantoken module functioning and contains the **issuePrice** for the fan token. Such an implementation allows governance to decide the issue price for the tokens arbitrarily since proposals can modify it.

```go
type Params struct {
	IssuePrice	sdk.Coin
}
```

## Token

The state contains a list of **Tokens**. They are [fan tokens](01_concepts.md#Fan-token) (fungible tokens deriving by the ERC-20 Standard), and their state information are:

- **Name**, which corresponds to the name of the fan token. It is a `string` and _cannot change_ for the whole life of the token;
- **MaxSupply**, that represents the maximum number of possible mintable tokens. It is an `integer number`, expressed in micro unit (![formula](https://render.githubusercontent.com/render/math?math=\color{gray}\mu=10^{-6})) as explained in [concepts](01_concepts.md#Fan-token), and _cannot change_ for the whole life of the token;
- **Mintable**, indicating the ability of the token to be minted. It is a `boolean` value and \*can change **only once\*** during the token lifecycle. At the issuing it is set to true, and the token can be minted. When the owner change this value in the state, the token can be minted no more.
- **Owner**, which is the current owner of the token. It is an address and _can change_ during the token lifecycle thanks to the **ownership transfer**;
- **MetaData**, which contains metadata for the fan token and is made up of the description, the denom, the symbol and a set of denomUnits. It is a `bank.Metadata` object and _cannot change_ for the whole life of the token.

```go
type FanToken struct {
	Name		string
	MaxSupply	sdk.Int
	Mintable	bool
	Owner		string
	MetaData	bank.Metadata
}
```

## BurnedCoins

Another section in this module state is represented by **BurnedCoins**. It contains the total amount of all the burned tokens.

```go
BurnedCoins: []sdk.Coin
```

```go
type Coin struct {
	Denom  string
	Amount Int
}
```
