// +build !oci

package impt

import (
	_ "github.com/sereiner/parrot/engines"
	_ "github.com/sereiner/parrot/registry/local"
	_ "github.com/sereiner/parrot/registry/zookeeper"
	_ "github.com/sereiner/parrot/rpc"
	_ "github.com/sereiner/parrot/servers/cron"
	_ "github.com/sereiner/parrot/servers/http"
	_ "github.com/sereiner/parrot/servers/mqc"
	_ "github.com/sereiner/parrot/servers/rpc"
	_ "github.com/sereiner/parrot/servers/ws"
	_ "github.com/sereiner/lib/cache/memcache"
	_ "github.com/sereiner/lib/cache/redis"
	_ "github.com/sereiner/lib/mq/mqtt"
	_ "github.com/sereiner/lib/mq/redis"
	_ "github.com/sereiner/lib/mq/stomp"
	_ "github.com/sereiner/lib/mq/xmq"
	_ "github.com/sereiner/lib/queue"
	_ "github.com/sereiner/lib/queue/redis"
	_ "github.com/sereiner/lib/queue/xmq"
)
