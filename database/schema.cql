CREATE KEYSPACE "grpctodo"
WITH replication = {'class': GeoReplication, 'replication_factor' : 3}
AND durable_writes = true;

CREATE TABLE items(
    id TEXT,
    title TEXT,
    description TEXT,
    closed BOOLEAN,
    PRIMARY KEY (id)
  );