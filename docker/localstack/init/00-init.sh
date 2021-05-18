#!/bin/bash
awslocal sqs create-queue --queue-name batchem --cli-input-json file:///localstack_files/sqs/batchem.json
