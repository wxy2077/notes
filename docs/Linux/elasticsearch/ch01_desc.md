# Elasticsearch 介绍


## Docker 安装运行 ES

docker 官方 es 镜像地址: https://hub.docker.com/_/elasticsearch


拉取镜像
```
docker pull elasticsearch:7.10.1
```

启动
```shell
docker run --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -d elasticsearch:7.10.1
```
本地访问以下地址: `http://127.0.0.1:9200/`， 可以访问出现以下内容则说明安装正常。
```json
{
    "name": "cc4951b1e3a3",
    "cluster_name": "docker-cluster",
    "cluster_uuid": "Yn2FIYPURbmWLaMcQQDTJQ",
    "version": {
        "number": "7.10.1",
        "build_flavor": "default",
        "build_type": "docker",
        "build_hash": "1c34507e66d7db1211f66f3513706fdf548736aa",
        "build_date": "2020-12-05T01:00:33.671820Z",
        "build_snapshot": false,
        "lucene_version": "8.7.0",
        "minimum_wire_compatibility_version": "6.8.0",
        "minimum_index_compatibility_version": "6.0.0-beta1"
    },
    "tagline": "You Know, for Search"
}
```

## 安装jk

https://github.com/medcl/elasticsearch-analysis-ik/releases
