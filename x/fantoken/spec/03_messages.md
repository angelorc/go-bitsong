<!-- 
order: 3
-->

# Messages

Messages (`msg`s) are objects that trigger state transitions. Messages are wrapped in transactions (`tx`s) that clients submit to the network. The BitSong SDK wraps and unwraps `fantoken` module messages from transactions.

## MsgIssueFanToken
The `MsgIssueFanToken` message is used to issue a new fan token. It takes in input `Symbol`, `Name`, `MaxSupply`, `Description`, `Owner`, and `IssueFee` (described in [fan token definition](01_concepts.md#Fan-token)). Thanks to these values, the module can verify if the `Owner` can issue a new token (it is not a blocked address). At this point, it proceeds with token issuing and emitting of corresponding events. More specifically, the module manages the `issuing fees`, calculates the `denom`, generates the `metadata`, and finally creates the fan token. At this point, an `EventTypeIssueFanToken` event is emitted.

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
The `MsgEditFanToken` message is used to modify an existing fan token. It takes in input `Denom`, `Mintable`, and `Owner` (described in [fan token definition](01_concepts.md#Fan-token)). Thanks to these values, the module can verify whether the modifications are lawful (i.e., requested by the `Owner` and in accord with the state transition definition). The message permits to change the "*mintability*" of the fan token. In particular, at the issuing, the fan token can be minted (in fact the `Mintable` bool is set to true). Later on, during the lifecycle of the fan token, the owner is able to disable the possibility to mint it (check the [relative docs](01_concepts.md#Lifecycle-of-a-fan-token) for more details). After disabling the mintability, the `max supply` of the token is updated. At this point, an `EventTypeEditFanToken` event is emitted.

```go
type MsgEditFanToken struct {
	Denom		string
	Mintable	bool
	Owner		string
}
```

## MsgMintFanToken

The `MsgEditFanToken` message is used to modify an existing fan token. It takes in input `Recipient`, `Denom`, `Amount`, and `Owner` (all described in [fan token definition](01_concepts.md#Fan-token) except the `Amount`, which is the quantity of token to mint). In such a message, the `Recipient` is not required and, its default value is the same of `Owner`. 
Thanks to these values, the module can verify whether the minting operation is lawful (i.e., requested: by the owner, on a mintable fan token, and for a quantity that allow to do not overcome the maximum supply), recalling that only the owner of the fan token can mint the token to a specified account. 
At this point, the token is minted, the result is sent to the recipient, and an `EventTypeMintFanToken` event is emitted.

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