echo &&
    curl \
        --silent \
        --dump-header head.txt \
        --output response.json \
        --header "Authorization: ApiKey $apikey" \
        http://localhost:8080/v1/posts &&
    cat head.txt &&
    jq . response.json &&
    echo &&
    rm head.txt &&
    rm response.json
