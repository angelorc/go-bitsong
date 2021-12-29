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
