#!/bin/bash

apt-get update && \
  DEBIAN_FRONTEND=noninteractive apt-get install -y postfix mailutils && \
  apt-get clean