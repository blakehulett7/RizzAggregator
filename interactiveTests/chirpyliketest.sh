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
    echo 'Getting all feeds...' &&
    echo &&
    curl \
        --silent \
        --dump-header head.txt \
        --output response.json \
        --header "Authorization: ApiKey $apikey" \
        http://localhost:8080/v1/feeds &&
    cat head.txt &&
    jq . response.json &&
    echo &&
    rm head.txt &&
    rm response.json
