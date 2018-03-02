# Simple tool to check if the server is running #

It used to be hard to talk to grpc endpoint from a command line, now it's easy.

## How to setup a server ##

Add a health check:
```go
import (
  	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

    // ...

	healthsrv := health.NewServer()
	healthpb.RegisterHealthServer(srv, healthsrv)
	healthsrv.SetServingStatus(serviceName, healthpb.HealthCheckResponse_SERVING)
```

## How to use a client ##

If you do it locally, first, forward a port:

```bash
kubectl port-forward POD_NAME 8080:8080
```

Then run:
```bash
grpcping -address "localhost:8080" -service SERVICE_NAME
```

where `SERVICE_NAME` is `serviceName` at the server side.

### Using grpcping in a pod ###


```
make
kubectl cp grpcping POD_NAME:/bin/grpcping
kubectl exec -it POD_NAME /bin/sh
chown root:root /bin/grpcping && chmod a+x /bin/grpcping && grpcping -help
grpcping -address ADDRESS -service SERVICE_NAME
```
