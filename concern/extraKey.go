package my_concern

import (
	localdb "github.com/Sora233/DDBOT/lsp/buntdb"
)

// 由于ddbot内置的是一个键值型数据库，通常需要使用一个key获取一个value，所以当需要存储数据的时候，需要使用额外的自定义key
// 可以在这个文件内实现

type extraKey struct{}

func (e *extraKey) exampleKey(keys ...interface{}) string {
	return localdb.NamedKey("example", keys)
}
