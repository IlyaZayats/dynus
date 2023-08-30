
cd /path/to/source
docker-compose up

bearer="sgadgnkaslgtjkbnsMC,ssoeugfjk"
endpopint="http://IP:port"

curl -XGET -H"Content-Type: application/json" $endpopint/api/v1/{user_id}

curl -XPUT -H"Content-Type: application/json" $endpopint/api/v1/slug \
    -d '{"name": "AVITO_VOICE_MESSAGES"}'

curl -XDELETE -H"Content-Type: application/json" $endpopint/api/v1/slug \
    -d '{"name": "AVITO_VOICE_MESSAGES"}'

curl -XPOST -H"Content-Type: application/json" $endpopint/api/v1/{user_id}/slug \
    -d '{"user_id": 1000,
    "slug": {
        "add": ["AVITO_VOICE_MESSAGES"],
        "remove": ["AVITO_VOICE_MESSAGES"]
        }
    }'

curl -XPOST -H"Content-Type: application/json" localhost:8080/slugs/2 \-d '{"insert_slugs": ["AVITO_VOICE_MESSAGES", "AVITO_DISCOUNT_30"], "delete_slugs": []}'



postgres://dynus:dynus@localhost:5432/dynus

#

curl -XGET -H"Content-Type: application/json" localhost:8090/slugs/2

curl -XPUT -H"Content-Type: application/json" localhost:8090/slugs \-d '{"name": "AVITO_VOICE_MESSAGES", "chance": "0.1"}'

curl -XPOST -H"Content-Type: application/json" localhost:8080/slugs/6 \-d '{"insert_slugs": ["AVITO_VOICE_MESSAGES2"], "delete_slugs": ["AVITO_VOICE_MESSAGES9"], "ttl": {"AVITO_VOICE_MESSAGES2" : "1 hour"}}'

curl -XDELETE -H"Content-Type: application/json" localhost:8090/slugs \-d '{"name": "AVITO_VOICE_MESSAGES"}'

curl -XGET -H"Content-Type: application/json" localhost:8090/slugs/history/2023-08'

C:/Users/Mylky/sdk/go1.19/bin/go.exe get