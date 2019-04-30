#!/bin/sh -e
# -*- coding: utf-8 -*-

/app/prerun_hook.sh
exec /app/bin/prisma-local
