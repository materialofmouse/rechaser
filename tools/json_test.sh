#!/bin/bash
curl -X POST http://localhost:8080/post \
	-H "Content-Type: application/json" \
	-d '{ "data": { "session": "10191019", "command": { "action": "WALK", "direction": "UP"}}}'
