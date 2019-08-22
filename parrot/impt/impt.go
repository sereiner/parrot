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
	_ "github.com/sereiner/library/cache/memcache"
	_ "github.com/sereiner/library/cache/redis"
	_ "github.com/sereiner/library/mq/mqtt"
	_ "github.com/sereiner/library/mq/redis"
	_ "github.com/sereiner/library/mq/stomp"
	_ "github.com/sereiner/library/mq/xmq"
	_ "github.com/sereiner/library/queue"
	_ "github.com/sereiner/library/queue/mqtt"
	_ "github.com/sereiner/library/queue/redis"
	_ "github.com/sereiner/library/queue/xmq"
)
