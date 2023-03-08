# DummyGameBackend
Dummy backend for a game


docker run -d -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgrespw -e POSTGRES_DB=mainDb -p 5433:5432 --name postgres postgres