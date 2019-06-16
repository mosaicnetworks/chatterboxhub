#Chatterbox Hub

The [Monet Project](https://monet.network/)  sent a team to the Cosmos Hackathon in Berlin over the weekend of June 15th and 16th. This repository is part of their efforts. 

##The building blocks

MONET is an open network architecture for mobile blockchains on demand. A blockchain on demand means that participants of a certain activity can form a temporary network for the duration of their interaction without a need of any centralized third party. Your can read more about MONET [here](https://monet.network/about.html)

We already have a mobile [implementation](http://docs.babble.io/en/latest/) of the hashgraph consensus algorithm. We also have a [ChatterBox app](https://monet.network/chatterbox.html) 

We have a roadmap to support IBC communication between the mobile adhoc blockchains and a central hub. The [Cosmos SDK](https://cosmos.network/docs/intro/) implementation of IBC seemed an ideal fit.

##Our Aims for the Hackathon

We aim to enable IBC communication between a mobile ad hoc blockchain and a central Cosmos hub using IBC. The mobile blockchain is within our Chatterbox app, which forms an adhoc blockchain using WiFi direct. We will host out central hub on AWS. 

We will add a karma system to the chatterbox app. Each message will have thumbs up and thumbs down button added to it. Each button press will generate a transaction to increase / decrease the karma of the user making the comment. These transactions pass through consensus and are applied to the block chain. These transactions are summarised and communicated via IBC to the central hub where they are stored in the state of the central hub. 

## How's it going

We are attempting to implement a hub using the Cosmos SDK from a standing start (in terms of Cosmos knowledge at least). This is no small challenge. 

Thus far we have melted 2 laptops, one with a faulty temperature sensor (unrelated to the quite mild weather) one going rogue and filling its hard disk necessitating recovery disks and a complete reinstallation of the Go development environment (exquisite timing we thought). 