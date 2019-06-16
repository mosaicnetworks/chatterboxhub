#Chatterbox Hub

**NB This repository is not currently complete, and is being actively amended. It is definitely a work in progress**


The [Monet Project](https://monet.network/)  sent a team to the Cosmos Hackathon in Berlin over the weekend of June 15th and 16th. This repository is part of their efforts. 

##The building blocks

MONET is an open network architecture for mobile blockchains on demand. A blockchain on demand means that participants of a certain activity can form a temporary network for the duration of their interaction without a need of any centralized third party. Your can read more about MONET [here](https://monet.network/about.html)

We already have a mobile [implementation](http://docs.babble.io/en/latest/) of the hashgraph consensus algorithm. We also have a [ChatterBox app](https://monet.network/chatterbox.html) 

We have a roadmap to support IBC communication between the mobile adhoc blockchains and a central hub. The [Cosmos SDK](https://cosmos.network/docs/intro/) implementation of IBC seemed an ideal fit.

##Our Aims for the Hackathon

We want to learn more about the Cosmos SDK whilst prototyping a key piece of our development roadmap, i.e. the ability to preserve the state of a transient mobile blockchain on a more permanent central hub. 

We aim to enable IBC communication between a mobile ad hoc blockchain and a central Cosmos hub using IBC. The mobile blockchain is within our Chatterbox app, which forms an adhoc blockchain using WiFi direct. We will host out central hub on AWS. 

We will add a karma system to the chatterbox app. Each message will have thumbs up and thumbs down button added to it. Each button press will generate a transaction to increase / decrease the karma of the user making the comment. These transactions pass through consensus and are applied to the block chain. These transactions are summarised and communicated via IBC to the central hub where they are stored in the state of the central hub. 

## How's it going

We are attempting to implement a hub using the Cosmos SDK from a standing start (in terms of Cosmos knowledge at least). This is no small challenge. It has already been immensely useful to aid our internal understanding of Cosmos. Nothing helps getting to know a product better than a failing build. 

Thus far we have melted 2 laptops, one with a faulty temperature sensor (unrelated to the quite mild weather) one going rogue and filling its hard disk necessitating recovery disks and a complete reinstallation of the Go development environment (exquisite timing we thought). We have recovered from both, but lost a significant amount of time in the process. 

We have the karma module mostly complete. The App and UI changes are mostly complete. The EC2 Cosmos instance is in place. The IBC dragon spotting is in progress. 



### To Do
The ibc module code has been copied into our repository. Tidy that up and fork it properly. 
