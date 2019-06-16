#Chatterbox Hub

The [Monet Project](https://monet.network/)  sent a team to the Cosmos Hackathon in Berlin over the weekend of June 15th and 16th. This repository is part of their efforts. 

##The building blocks

MONET is an open network architecture for mobile blockchains on demand. A blockchain on demand means that participants of a certain activity can form a temporary network for the duration of their interaction without a need of any centralized third party. Your can read more about MONET [here](https://monet.network/about.html)

We already have a mobile [implementation](http://docs.babble.io/en/latest/) of the hashgraph consensus algorithm. We also have a [ChatterBox app](https://monet.network/chatterbox.html) 

We have a roadmap to support IBC communication between the mobile adhoc blockchains and a central hub. The [Cosmos SDK](https://cosmos.network/docs/intro/) implementation of IBC seemed an ideal fit.

##Our Aims for the Hackathon

We aim to enable IBC communication between a mobile ad hoc blockchain and a central Cosmos hub using IBC. The mobile blockchain is within our Chatterbox app, which forms an adhoc blockchain using WiFi direct. We will host out central hub on AWS. 

