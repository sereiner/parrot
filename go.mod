module github.com/sereiner/parrot

go 1.12

replace (
	github.com/zkfy/archiver => github.com/mholt/archiver v2.1.0+incompatible
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20180910181607-0e37d006457b
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190301231843-5614ed5bae6f
	golang.org/x/net => github.com/golang/net v0.0.0-20180911220305-26e67e76b6c3
	golang.org/x/sync => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20180909124046-d0be0721c37e
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20180820150726-fbb02b2291d28
	google.golang.org/appengine => github.com/golang/appengine v1.1.0
	google.golang.org/genproto => github.com/ilisin/genproto v0.0.0-20181026194446-8b5d7a19e2d9
	google.golang.org/grpc => github.com/grpc/grpc-go v1.16.0
)

require (
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gin-gonic/gin v1.4.0
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/golang/protobuf v1.3.1
	github.com/golang/snappy v0.0.1
	github.com/gorilla/websocket v1.4.0
	github.com/json-iterator/go v1.1.6
	github.com/nwaples/rardecode v1.0.0 // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/pkg/profile v1.3.0
	github.com/sereiner/lib v0.1.0
	github.com/sereiner/log v0.0.3
	github.com/stretchr/testify v1.3.0
	github.com/ugorji/go v1.1.4
	github.com/urfave/cli v1.20.0
	github.com/wule61/log v0.0.0-20190426025328-54b7fa0d64ad
	github.com/zkfy/archiver v0.0.0-00010101000000-000000000000
	github.com/zkfy/cron v0.0.0-20170309132418-df38d32658d8
	github.com/zkfy/go-oci8 v0.0.0-20180327092318-ad9f59dedff0
	github.com/zkfy/log v0.0.0-20180312054228-b2704c3ef896
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/sys v0.0.0-20190429190828-d89cdac9e872
	google.golang.org/grpc v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.2.2
)
