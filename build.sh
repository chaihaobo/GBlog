#!/bin/bash
pwd=${echo pwd}
cd $pwd/api-core/
go build main.log
cd $pwd/web
npm run build
cd $pwd
docker-compose up
