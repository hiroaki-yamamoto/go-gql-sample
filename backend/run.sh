#!/bin/sh -e
# -*- coding: utf-8 -*-

dep ensure
go build
exec ./backend
