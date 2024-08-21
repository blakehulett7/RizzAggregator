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
    apikey=$(jq .ApiKey response.json) &&
    echo 'Posting a feed to the database...' &&
    echo &&
    curl \
        --silent \
        --dump-header head.txt \
        --output response.json \
        --header "Authorization: ApiKey $apikey" \
        --data @./payloadfeedtest.json \
        http://localhost:8080/v1/feeds &&
    cat head.txt &&
    jq . response.json &&
    feedID=$(jq .ID response.json) &&
    echo &&
    echo 'Getting feeds that I follow...' &&
    echo &&
    curl \
        --silent \
        --dump-header head.txt \
        --output response.json \
        --header "Authorization: ApiKey $apikey" \
        http://localhost:8080/v1/feed_follows &&
    cat head.txt &&
    jq . response.json &&
    echo &&
    rm head.txt response.json
