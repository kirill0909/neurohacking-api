#!/bin/bash

./.bin/neurohacking-api&

# get process id
processID=$(pidof neurohacking-api)
echo "process $processID was successfully launched"
