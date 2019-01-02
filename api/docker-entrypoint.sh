#!/bin/sh

../api

#Keep the container awake just in case it dies and we need to do some testing
while true; do sleep 1000; done
