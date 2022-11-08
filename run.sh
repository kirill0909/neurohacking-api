#!/bin/bash

./.bin/neurohacking-api&

processID=$(pidof neurohacking-api)
echo "process $processID was successfully launched"
