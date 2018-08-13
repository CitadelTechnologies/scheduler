#!/bin/bash

api-console build doc/jobs/index.raml
api-console serve --hostname 0.0.0.0 -p 8000 build

exec "$@"
