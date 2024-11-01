#!/bin/bash

# Endpoint URL
URL="http://localhost:8002/api/students"

# Generate random user details and send 10 POST requests
for i in {1..10}; do
  # Generate random data for each user
  USERNAME="user$(uuidgen | tr -d '-' | head -c 8)"
  PASSWORD="testPassword123"
  FIRST_NAME=$(echo -e "John\nJane\nAlex\nChris\nTaylor\nJordan\nMorgan\nSam\nDana\nSkyler" | awk 'BEGIN{srand()} {print rand(), $0}' | sort -n | head -n 1 | cut -d' ' -f2-)
  LAST_NAME=$(echo -e "Smith\nJohnson\nWilliams\nJones\nBrown\nDavis\nMiller\nWilson\nMoore\nTaylor" | awk 'BEGIN{srand()} {print rand(), $0}' | sort -n | head -n 1 | cut -d ' ' -f2-)
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
