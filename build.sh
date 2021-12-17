#!/bin/bash
docker build -f Dockerfile -t motor-mini-backend-web .
docker run -idt -p 8000:8000 motor-mini-backend-web