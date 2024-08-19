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
    cat head.txt &&
    jq . response.json &&
    apikey=$(jq .ApiKey response.json) &&
    echo &&
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
    echo &&
    echo 'Testing bad api key...' &&
    echo &&
    curl \
        --silent \
        --dump-header head.txt \
        --output response.json \
        --header "Authorization: ApiKey 1111" \
        --data @./payloadfeedtest.json \
        http://localhost:8080/v1/feeds &&
    cat head.txt &&
    jq . response.json &&
