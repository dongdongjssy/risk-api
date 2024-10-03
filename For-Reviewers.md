This project uses the following frameworks and libraries to assist development:

- gin: [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- uuid: [http://pkg.go.dev/github.com/google/uuid](http://pkg.go.dev/github.com/google/uuid)
- swagger: [https://github.com/swaggo/gin-swagger](https://github.com/swaggo/gin-swagger)

## Notes

There are some improvements can be done in the future:

1. Get risks api does not support pagination.
2. Risk state is stored as string, if risks amount is huge, need to change the type to small int to save storing spaces.
3. It uses Gin's build-in logging which should be replaced with other customized tool, also thinking of redirecting logs to a local file while service is deployed in a cluster as a pod.
4. No authentication for apis, will need to add a middleware service to do it.
5. Create a DockerFile to build the service into image so it can be easily deployed to cluster as microservices.
6. In my current daily work, I'm using Nodejs+Express to develop apis, here is an example: [A Fake Api](https://gist.github.com/dongdongjssy/5008e4db3fdd2e5d12d763c461d36fa2)
