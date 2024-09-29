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
