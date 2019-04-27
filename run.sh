#!/bin/sh -e
# -*- coding: utf-8 -*-

cd backend
dep ensure
exec go run main.go
