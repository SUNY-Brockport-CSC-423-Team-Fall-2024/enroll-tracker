#!/bin/bash

# Endpoint URL
URL="http://localhost:8002/api/teachers"

firstNames=("John" "Jane" "Alex" "Chris" "Taylor" "Emily" "Jordan" "Morgan" "Sam" "Dana")
lastNames=("Smith" "Johnson" "Williams" "Jones" "Brown" "Davis" "Miller" "Wilson")
offices=("Brown Bldg #123" "Brown Bldg #124" "Westchester Hall #124" "Wosinski Bldg #14")
firstNamesLen=${#firstNames[@]}
lastNamesLen=${#lastNames[@]}
officesLen=${#offices[@]}

# Generate random user details and send 10 POST requests
for i in {1..10}; do
  # Generate random numbers
  random_index_f=$((RANDOM % firstNamesLen))
  random_index_l=$((RANDOM % lastNamesLen))
  random_index_o=$((RANDOM % officesLen))
  # Generate random data for each user
  USERNAME="user$(uuidgen | tr -d '-' | head -c 8)"
  PASSWORD="testPassword123"
  FIRST_NAME=$(echo -e ${firstNames[$random_index_f]})
  LAST_NAME=$(echo -e ${lastNames[$random_index_l]})
  PHONE_NUMBER="585-555-5555"
  EMAIL="${USERNAME}@example.com"
  OFFICE=$(echo -e ${offices[$random_index_o]})

  # Create JSON payload
  DATA=$(jq -n --arg username "$USERNAME" \
                --arg password "$PASSWORD" \
                --arg first_name "$FIRST_NAME" \
                --arg last_name "$LAST_NAME" \
                --arg phone_number "$PHONE_NUMBER" \
                --arg email "$EMAIL" \
                --arg office "$OFFICE" \
                '{
                  username: $username,
                  password: $password,
                  first_name: $first_name,
                  last_name: $last_name,
                  phone_number: $phone_number,
                  email: $email,
                  office: $office
                }')

  # Send POST request
  RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")

  # Check if the request was successful
  if [ "$RESPONSE" -eq 201 ]; then
    echo "Successfully created teacher: $USERNAME"
  else
    echo "Failed to create teacher: $USERNAME - Status Code: $RESPONSE"
  fi
done
