#!/bin/bash

# Endpoint URL
URL="http://localhost:8002/api/students"

firstNames=("John" "Jane" "Alex" "Chris" "Taylor" "Emily" "Jordan" "Morgan" "Sam" "Dana")
lastNames=("Smith" "Johnson" "Williams" "Jones" "Brown" "Davis" "Miller" "Wilson")
firstNamesLen=${#firstNames[@]}
lastNamesLen=${#lastNames[@]}

# Generate random user details and send 10 POST requests
for i in {1..10}; do
  # Generate random numbers
  random_index_f=$((RANDOM % firstNamesLen))
  random_index_l=$((RANDOM % lastNamesLen))
  # Generate random data for each user
  USERNAME="user$(uuidgen | tr -d '-' | head -c 8)"
  PASSWORD="testPassword123"
  FIRST_NAME=$(echo -e ${firstNames[$random_index_f]})
  LAST_NAME=$(echo -e ${lastNames[$random_index_l]})
  EMAIL="${USERNAME}@example.com"

  # Create JSON payload
  DATA=$(jq -n --arg username "$USERNAME" \
                --arg password "$PASSWORD" \
                --arg first_name "$FIRST_NAME" \
                --arg last_name "$LAST_NAME" \
                --arg email "$EMAIL" \
                '{
                  username: $username,
                  password: $password,
                  first_name: $first_name,
                  last_name: $last_name,
                  email: $email
                }')

  # Send POST request
  RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")

  # Check if the request was successful
  if [ "$RESPONSE" -eq 201 ]; then
    echo "Successfully created user: $USERNAME"
  else
    echo "Failed to create user: $USERNAME - Status Code: $RESPONSE"
  fi
done
