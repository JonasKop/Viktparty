FROM migrate/migrate:v4.15.2

COPY ./migrations /migrations

ENTRYPOINT [ "" ]
CMD ["sh", "-c", "migrate -path=/migrations/ -database=postgres://$PGUSER:$PGPASSWORD@$PGHOST:5432/$PGDATABASE?sslmode=disable up"]
