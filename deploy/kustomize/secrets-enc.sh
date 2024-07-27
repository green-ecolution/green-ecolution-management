#! /bin/bash

sops -e ./components/mongodb/secrets.yaml > ./components/mongodb/secrets.enc.yaml
sops -e ./components/mongo-express/secrets.yaml > ./components/mongo-express/secrets.enc.yaml

sops -e ./dev/secrets.yaml > ./dev/secrets.enc.yaml
