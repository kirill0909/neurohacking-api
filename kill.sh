#!/bin/bash

processID=$(pidof neurohacking-api)
kill $processID
echo "process $processID was killed"
