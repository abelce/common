#!/bin/bash
docker build -t abelce/article_db:1.0.0 .
docker push abelce/article_db:1.0.0
