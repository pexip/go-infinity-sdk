#!/bin/bash

COOKIES_FILE="cookies.txt" # Define the cookies file name

read -p "Enter base URL (e.g., https://manager.example.com): " BASE_URL
BASE_URL="${BASE_URL%/}"
read -p "Enter basic username: " USERNAME
read -p "Enter basic password: " PASSWORD
BASIC_AUTH=$(echo -n "${USERNAME}:${PASSWORD}" | base64)
read -p "Enter module name (default: configuration): " MODULE

SUB_MODULE="" # Either of participant, conference or platform is valid
if [ "$MODULE" == "command" ]; then
    read -p "Enter sub-module for command (e.g., conference): " SUB_MODULE
fi

# Read the content of cookies.txt into the COOKIES variable.
# This assumes cookies.txt contains the cookie string (e.g., "key=value; key2=value2").
if [ -f "$COOKIES_FILE" ]; then
   COOKIES=$(cat "$COOKIES_FILE")
   echo "Loaded cookies from $COOKIES_FILE"
else
   COOKIES=""
   echo "No $COOKIES_FILE found, proceeding without cookies."
fi

#### Functions starts here ####

download_file() {
 local url="$1"
 local out_file="$2"
 local basic_auth_token="$3"
 local cookies_data="$4"

 echo "Downloading $url -> $out_file"
 curl \
   -H "authorization: Basic ${basic_auth_token}" \
   -H 'cache-control: no-cache' \
   -b "${cookies_data}" \
   -sSf "$url" \
   -o "$out_file"

 if [ $? -ne 0 ]; then
   echo "Failed to download $url"
   return 1 # Indicate failure
 fi
 return 0 # Indicate success
}

#### Main script starts here ####

config_url="${BASE_URL}/api/admin/${MODULE}/v1/"
if [ -n "$SUB_MODULE" ]; then
  config_url="${config_url}${SUB_MODULE}/"
fi
config_output_file="${MODULE}.json"

download_file "$config_url" "$config_output_file" "$BASIC_AUTH" "$COOKIES"
if [ $? -eq 0 ]; then
  echo "configuration.json downloaded successfully to $config_output_file."
else
  echo "Error: Failed to download configuration.json from $config_url."
  exit 1
fi

mkdir -p "$MODULE"

# Extract all schema URLs and resource names from the config
jq -r 'to_entries[] | "\(.key) \(.value.schema)"' "$config_output_file" | while read -r name schema; do
   url="${BASE_URL}${schema}"
   out_file="${MODULE}/${name}.json"
   download_file "$url" "$out_file" "$BASIC_AUTH" "$COOKIES"
   if [ $? -ne 0 ]; then
        echo "Failed to download $url"
        break
   fi
done