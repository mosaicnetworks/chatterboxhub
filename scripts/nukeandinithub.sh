#!/bin/bash

# Hard coded path. #HACKATHON_SHORTCUT
CBDDIR=~/go/src/github.com/mosaicnetworks/chatterboxhub
# Softcode in production
CHAINID=monetchain
HUBNAME=monetcosmoshub

echo "Changing cwd to $CBDDIR"
cd $CBDDIR

echo "Clearing out old configuration"
rm -rf ~/.karmad
rm -rf ~/.karmacli

echo "Generate Genesis File"
karmad init $HUBNAME --chain-id $CHAINID


for name in validator alice beth charlie dave ethel franz gilbert harriet iris julio kenny larry maeve norbert olive petal quentin rachel sacha thierry umberto violet will 
do
figlet $name
	karmacli keys add $name <<!
password
password
!
	karmad add-genesis-account $(karmacli keys show $name -a) 1000nametoken,100000000stake
done


karmacli config chain-id $CHAINID
karmacli config output json
karmacli config indent true
karmacli config trust-node true

karmad gentx --name validator <<!
password
!

karmad collect-gentxs

karmad validate-genesis
 

