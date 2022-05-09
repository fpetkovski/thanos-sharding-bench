# Thanos query sharding benchmark

Use `make block` to generate a new TSDB block with 100K series. 
Deploy the setup with `make start`. The sharded query frontend implementation is avilable on localhost:30902 