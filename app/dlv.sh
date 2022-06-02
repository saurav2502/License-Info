#!/bin/sh
#
# Copyright (c)  by Saurav from 2022
#
dlv debug --headless --log -l 0.0.0.0:2345 --api-version=2 --accept-multiclient --execute app