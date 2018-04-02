#!/usr/bin/env bash
mysql food -uroot -h172.17.0.3 -p$MYSQL_ROOT_PASSWORD < ./exampledb.sql
