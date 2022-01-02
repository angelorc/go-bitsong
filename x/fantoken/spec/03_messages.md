<!-- 
order: 3
-->

# Messages

Messages (`msg`s) are objects that trigger state transitions. Messages are wrapped in transactions (`tx`s) that clients submit to the network. The Cosmos SDK wraps and unwraps `fantoken` module messages from transactions.

## MsgIssueFanToken
The `MsgIssueFanToken` message is used to issue a new fan token. It takes in input `Symbol`, `Name`, `MaxSupply`, `Description`, `Owner`, and `IssueFee` (described in [fan token definition](01_concepts.md#Fan-token)). Thanks to these values, the module can verify if the owner can issue a new token (it is not a blocked address). At this point, it proceeds with token issuing and emitting of corresponding events. More specifically, the module manages the `issuing fees`, calculates the `denom`, generates the `metadata`, and finally creates the fan token. At this point, a `EventTypeIssueFanToken` event is emitted.

```go
type MsgIssueFanToken struct {
	Symbol		string
	Name		string
	MaxSupply	sdk.Int
	Description string
	Owner		string
	IssueFee	sdk.Coin
}
```

## MsgEditFanToken
The `MsgEditFanToken` message is used to modify an existing fan token. It takes in input `Denom`, `Mintable`, and `Owner` (described in [fan token definition](01_concepts.md#Fan-token)). Thanks to these values, the module can verify whether the modifications are lawful (e.g., requested by the owner or in accord with the state transition definition). The message permits to change the "*mintability*" of the fan token. In particular, at the issuing, the fan token can be minted (in fact the `mintable` bool is set to true). Later on, during the lifecycle of the fan token, the owner is able to disable the possibility to mint it (check the [relative docs](01_concepts.md#Lifecycle-of-a-fan-token) for more details).

```go
type MsgEditFanToken struct {
	Denom		string
	Mintable	bool
	Owner		string
}
```

## MsgMintFanToken
Only the owner of the fantoken can mint new fantoken to a specified account. It fails if the total supply > max supply

```go
type MsgMintFanToken struct {
	Recipient	string
	Denom		string
	Amount		sdk.Int
	Owner		string
}
```

## MsgBurnFanToken
The action will be completed if the sender balance > amount to burn

```go
type MsgBurnFanToken struct {
	Denom		string
	Amount		sdk.Int
	Sender		string
}
```

## MsgTransferFanTokenOwner

Transfer the ownership of the fantoken to another account owner

```go
type MsgTransferFanTokenOwner struct {
	Denom		string
	SrcOwner	string
	DstOwner	string
}
```