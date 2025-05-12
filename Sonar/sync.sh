#!/bin/bash

cd ..

git subtree push --prefix=Frontend/server sonar-server main
git subtree push --prefix=Frontend/client sonar-client main