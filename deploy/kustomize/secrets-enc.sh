#! /bin/bash

sops -e ./dev/secrets.yaml > ./dev/secrets.enc.yaml
