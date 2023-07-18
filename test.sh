#!/bin/bash

./example/build.sh
(cd example && go test -v ./)
