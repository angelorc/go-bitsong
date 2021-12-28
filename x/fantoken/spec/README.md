# `fantoken`

## Abstract

This document specifies the _fantoken_ module of the BitSong chain.

The _fantoken_ module enables the BitSong chain to support fan tokens, allowing actors in the music industry to create their economy. In this sense, they can generate new ways to monetize their music and brand and provide a unique and innovative channel to engage with fans. Thanks to this module, artists, labels, producers, and all the actors in the music industry can start minting their _fan tokens_ (which are **fungible tokens**) and list them within a few minutes for low fees.

We can identify each _fan token_ through two keys: the **denom** and the **symbol**.
Even if they both represent the token, we generate them in different ways, and, for this reason, they also share separate information.

More specifically:

- **denom** is calculated by the tendermint crypto hash through the creator, symbol, and name. For this reason, it is _unique_;
- **symbol**, on the other hand, is defined by the user and can be any string.

Finally, thanks to the _fantoken_ module, users on BitSong can:

- manage _fan tokens_, issuing, minting, burning, and transferring them;
- build applications that use the _fan tokens_ API to create completely new and custom artists' economies.

Features that may be added in the future are described in Future Improvements.

## Table of Contents

1. **[Concepts](01_concepts.md)**
   - [Fan Token](01_concepts.md#FanToken)
2. **[State](02_state.md)**
   - [Fan Token](02_state.md#FanToken)
   - [Params](02_state.md#Params)
     <!--
     State Transitions
     -->
     <!--
     Keeper
     -->
3. **[Messages](03_messages.md)**
   - [MsgIssueFanToken](03_messages.md#MsgIssueFanToken)
   - [MsgEditFanToken](03_messages.md#MsgEditFanToken)
   - [MsgMintFanToken](03_messages.md#MsgMintFanToken)
   - [MsgBurnFanToken](03_messages.md#MsgBurnFanToken)
   - [MsgTransferFanTokenOwner](03_messages.md#MsgTransferFanTokenOwner)
     <!--
     Begin-Block
     -->
     <!--
     End-Block
     -->
4. **[Events](04_events.md)**
   - [MsgIssueFanToken](04_events.md#MsgIssueFanToken)
   - [MsgEditFanToken](04_events.md#MsgEditFanToken)
   - [MsgMintFanToken](04_events.md#MsgMintFanToken)
   - [MsgBurnFanToken](04_events.md#MsgBurnFanToken)
   - [MsgTransferFanTokenOwner](04_events.md#MsgTransferFanTokenOwner)
5. **[Parameters](05_parameters.md)**
   <!--
   Test Cases
   -->
   <!--
   Benchmarks
   -->
6. **[Future Improvements](06_future_improvements.md)**
