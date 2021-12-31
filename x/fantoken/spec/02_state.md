<!--
order: 2
-->

# State

The `fantoken` module keeps track of [**parameters**](#Params), [**fan tokens**](#Token), and **burned coins**.

```
Params:      sdk.Coin
Tokens:      []types.FanToken
BurnedCoins: []sdk.Coin
```

## Params

In the state definition, we can find the **params**. This section corresponds to a module-wide configuration structure that stores system parameters. In particular, it defines the overall fantoken module functioning and contains the **issuePrice** for the fan token. Such an implementation allows governance to decide the issue price for the tokens arbitrarily since proposals can modify it.

```go
type Params struct {
	IssuePrice	sdk.Coin
}
```

## Token

Definition of data structure of Fungible Token

```go
type FanToken struct {
	Name		string
	MaxSupply	sdk.Int
	Mintable	bool
	Owner		string
	MetaData	bank.Metadata
}
```
