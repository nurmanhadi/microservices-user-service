docker build -t microservices/user-service:latest .
docker run -d \
--name user-service \
-p 4000:4000 \
--network user-service_user-service \
-e DB_HOST=user-service-db \
-e DB_PORT=5432 \
-e DB_DATABASE=user_service \
-e DB_USERNAME=user_service \
-e DB_PASSWORD=user_service \
-e DB_MAX_POOL_CONNS=10 \
-e DB_MAX_IDLE_CONNS=5 \
-e DB_CONN_MAX_LIFETIME=300 \
-e JWT_ACCESS_SECRET=testing \
-e JWT_ACCESS_EXP=15 \
-e JWT_REFRESH_SECRET=testing \
-e JWT_REFRESH_EXP=7 \
microservices/user-service