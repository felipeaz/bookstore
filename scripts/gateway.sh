printf "Creating Services...\n\n"

curl --location --request POST 'http://localhost:8001/services' \
--form 'name="account-service"' \
--form 'protocol="http"' \
--form 'host="account-service"' \
--form 'port="8081"' \
--form 'path="/"'

curl --location --request POST 'http://localhost:8001/services' \
--form 'name="inventory-service"' \
--form 'protocol="http"' \
--form 'host="inventory-service"' \
--form 'port="8082"' \
--form 'path="/"'

curl --location --request POST 'http://localhost:8001/services' \
--form 'name="orders-service"' \
--form 'protocol="http"' \
--form 'host="orders-service"' \
--form 'port="8083"' \
--form 'path="/"'

printf "Creating Routes...\n\n"

curl --location --request POST 'http://localhost:8001/services/account-service/routes' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=account' \
--data-urlencode 'paths=/auth' \
--data-urlencode 'methods=GET' \
--data-urlencode 'methods=POST' \
--data-urlencode 'methods=PUT' \
--data-urlencode 'methods=DELETE'

curl --location --request POST 'http://localhost:8001/services/inventory-service/routes' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=manager' \
--data-urlencode 'paths=/mgr' \
--data-urlencode 'methods=GET' \
--data-urlencode 'methods=POST' \
--data-urlencode 'methods=PUT' \
--data-urlencode 'methods=DELETE'

curl --location --request POST 'http://localhost:8001/services/orders-service/routes' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=orders' \
--data-urlencode 'paths=/' \
--data-urlencode 'methods=GET' \
--data-urlencode 'methods=POST' \
--data-urlencode 'methods=PUT' \
--data-urlencode 'methods=DELETE'

printf "Setting up authentication...\n\n"

curl --location --request POST 'http://localhost:8001/services/orders-service/plugins' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=jwt' \
--data-urlencode 'config.secret_is_base64=false' \
--data-urlencode 'config.run_on_preflight=true'

curl --location --request POST 'http://localhost:8001/services/inventory-service/plugins' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=jwt' \
--data-urlencode 'config.secret_is_base64=false' \
--data-urlencode 'config.run_on_preflight=true'