# kubemb
![Go 1.16](https://img.shields.io/badge/Go-v1.16-blue)&nbsp;[![](https://img.shields.io/github/license/Yu-HaoYu/Ferryman?color=red)](https://github.com/Yu-HaoYu/Ferryman/blob/master/LICENSE)

> ​		这是一个kubectl插件，提供查看deployment对象当前使用的镜像版本与所属Pod资源信息等。
>

<br/>

## **功能**

- 查看镜像版本

- 查看Pod资源信息

<br/>

## 下载插件

```bash
curl -LO https://github.com/Yu-HaoYu/kubemb/releases/download/v0.0.2/kubemb
chmod +x ./kubemb
mv ./kubemb /usr/local/bin
ln -s /usr/local/bin/kubemb /usr/local/bin/kubectl-mb

kubemb -v
kubectl mb -v

```

**备注:** `kubemb` 等价 `kubectl mb`

<br/>

## **使用插件**

### 查看镜像版本

```bash
# 查看Deployment的image信息
kubemb image -n default

NAMESPACE	DEPLOY_NAME		CONTAINER_NAME		IMAGE
default		php-apache		php-apache			k8s.gcr.io/hpa-example
default		podinfo			podinfod			stefanprodan/podinfo:0.0.1
default		sample-app		metrics-provider	luxas/autoscale-demo:v0.1.2
```

<br/>

### 查看Pod资源信息

```bash
# 查看Deployment的资源信息
kubemb cap -n default

NAMESPACE	CONTAINER NAME			CPU REQUESTS	MEMORY REQUESTS	CPU LIMITS	MEMORY LIMITS
default		php-apache				200m			0				500m		0
default		podinfod				1m				32Mi			100m		256Mi
default		metrics-provider		0				0				0			0
```



