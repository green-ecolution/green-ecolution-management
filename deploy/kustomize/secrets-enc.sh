#! /bin/bash

sops -e ./components/mongodb/secrets.yaml > ./components/mongodb/secrets.enc.yaml
sops -e ./components/mongodb/mongo-express/secrets.yaml > ./components/mongodb/mongo-express/secrets.enc.yaml

sops -e ./dev/secrets.yaml > ./dev/secrets.enc.yaml
