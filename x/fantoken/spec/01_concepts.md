<!--
order: 1
-->

# Concepts

## Conventions

By looking at numbers, we separate the decimals by point and the thousands by comma. For instance, the number _one thousand two hundred thirty-four and fifty-six hundredths_, is written as:

![formula](https://render.githubusercontent.com/render/math?math=\color{gray}1,234.56)

## Fan token

Fan tokens, conceptually based on the [ERC-20 Standard](https://ethereum.org/it/developers/docs/standards/tokens/erc-20), are **fungible tokens** issued for fan communities. They borns to create new connections between fans and any content creator, like star performers, actors, designers, musicians, photographers, writers, models, influencers, etc.
They enable the growth of a private and (most importantly) custom economy creating new channels for fans' engagement.
Fan tokens have enormous potential. By using them, you can build myriad applications allowing fans a deeper interaction in the artistic life of their top performers.

To provide you with some examples, you can think that it is possible to use them for creating loyalty programs to provide privileged access to exclusive content. To allow your fan to crowdfund a tour or studio album and share part of the revenue with your fans. To enable your fans with the opportunity to vote on the cities for an upcoming tour. Or even to accept fan tokens as payment for NFTs.

A **fan token** is characterized by:
| Attribute | Type | Description |
| --------------------- | ---------------------------- | ---------------------------------------------- |
| symbol | `string` | It is chosen by the user. It should follow the ISO standard for the [alphabetic code](https://www.iso.org/iso-4217-currency-codes.html) (e.g., USD, EUR, BTSG, etc.).|
| name | `string` | It is chosen by the user. It should correspond to the long name the user want to associate to the symbol (e.g., Dollar, Euro, BitSong). |
| maxSupply | `sdk.Int` | It is chosen by the user. It is the maximum supply value of mintable tokens from its definition. It is expressed in micro unit (![formula](https://render.githubusercontent.com/render/math?math=\color{gray}\mu=10^{-6})). For this reason, to indicate a maximum supply of ![formula](https://render.githubusercontent.com/render/math?math=\color{gray}456) tokens, this value must be equal to ![formula](https://render.githubusercontent.com/render/math?math=\color{gray}456\cdot10^{6}=456,000,000).|
| description | `string` | It is chosen by the user. It is a small description about the fan token the user is gonna issuing. |
| mintable | `boolean` | It is true at issuing. It can be later changed by the owner. If it is `true`, the fan token owner can mint the token to a particular address. |
| owner | `sdk.AccAddress` | It is the owner of the fan token. It can be changed to trasfer the ownership of the token during the time. It is mainly used to verify the ability to perform operations.|
| denom | `string` | It is an hash calculated on Owner, Symbol, and Name of the fan token. It is the hash identifying the fan token and is used to [prevent the creation of identical tokens](#Uniqueness-of-the-denom). |
| metadata | `banktypes.Metadata` | It is made up of the description, the denom, the symbol and a set of denomUnits. It is generated and it is not directly editable.|
| issueFee | `sdk.Coin` | It is chosen by the user. It describes the issuing fee both for for the `amount` and for the `coin`. An example value could be `1000000ubtsg`.|

## Lifecycle of a fan token

It is possible to entirely represent the lifecycle of a fan token through Finite State Machine (FSM) diagrams. We will present two representations:

- the first refers to the fan token **object**. We can compare such a definition with that of currency (e.g., Euro, Dollar, BitSong);
- the second, instead, is referred to the lifecycle of the fan token **instance**. Such definition is comparable with that of coin/money (e.g., the specific 1 Euro coin you could have in your pocket at a particular moment in time).

We can describe the lifecycle of a fan token **object** through two states.

![Fantoken object lifecycle](img/fantoken_object_lifecycle.png "Fantoken object lifecycle")

Referring to the figure above, once the fan token is detailed in the documentation, to "create" the fan token, we need to **issue it**. This operation leads to the birth of the object and thus to its first state, state _1_. Here, the token is related to an owner, which can mint it. From this state, the owner can perform two actions on the object:

- to **transfer the ownership**, which produces the changing of the owner, without modifying the landing state;
- to **disable the minting ability**, which produces a state change to the state _2_. Here, the owner cannot mint the fan token anymore.

Once the fan token lands in state _2_, the only possible action is to transfer its ownership. Here, the owner **can enable the minting ability** no more.

Also referring to the lifecycle of a fan token **instance**, it is possible to identify two states.

![Fantoken instance lifecycle](img/fantoken_instance_lifecycle.png "Fantoken instance lifecycle")

Concerning to the figure above, when the fan token object is issued, we can **mint** it. Minting leads to the birth of a new instance, moving the fan token instance to state _1_. In this state, the token can be:

- **traded**, which produces the changing of the owner, without modifying the landing state;
- **burned**, which produces a state change to the state _2_, where the owner cannot operate on the fan token instance anymore.

## Uniqueness of the denom

The _denom_ is calculated on the Owner, Symbol, and Name of the fan token.

```go
func GetFantokenDenom(creator sdk.AccAddress, symbol, name string) string {
	return "ft" + tmcrypto.AddressHash([]byte(creator.String()+symbol+name)).String()
}
```

The _denom_ of every fan token starts with the prefix `ft`. Follows a **hash** of Owner, Symbol, and Name of the fan token. This _denom_ is used as base denom for the fan token, and, for this reason, it should be **unique**. In this sense, since the hash also depends on the creator, multiple fan tokens with the same name and symbol can co-exist. In this case, they can not be created by the same account.
