echo &&
    echo 'Posting a user to the database...' &&
    echo &&
    curl \
        --silent \
        --dump-header head.txt \
        --output response.json \
        --header "Content-Type: application/json" \
        --data @./payloadtest.json \
        http://localhost:8080/v1/users &&
    apikey1=$(jq .ApiKey response.json) &&
    echo 'Posting a feed to the database...' &&
    echo &&
    curl \
        --silent \
        --dump-header head.txt \
        --output response.json \
        --header "Authorization: ApiKey $apikey1" \
        --data @./payloadfeedtest.json \
        http://localhost:8080/v1/feeds &&
    cat head.txt &&
    jq . response.json &&
    feedID=$(jq .feed.ID response.json) &&
    echo &&
    echo 'Posting a second user to the database...' &&
    echo &&
    curl \
        --silent \
        --dump-header head.txt \
        --output response.json \
        --header "Content-Type: application/json" \
        --data @./payloadtest.json \
        http://localhost:8080/v1/users &&
    apikey2=$(jq .ApiKey response.json) &&
    echo 'Following a feed...' &&
    echo &&
    echo "{\"feed_id\": $feedID}" >>payloadfollow.json &&
    curl \
        --dump-header head.txt \
        --output response.json \
        --header "Authorization: ApiKey $apikey2" \
        --data @./payloadfollow.json \
        http://localhost:8080/v1/feed_follows &&
    cat head.txt &&
    jq . response.json &&
    rm head.txt response.json payloadfollow.json &&
    echo
