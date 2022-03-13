docker run -d -p 4222:4222 --rm --name yuta-nats -v nats-data:/var/lib/nats  nats:latest --jetstream --store_dir /var/lib/nats --cluster_name yuta

nats account info
nats stream list

