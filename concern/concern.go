// 别忘记改package name
package my_concern

import (
	"github.com/Sora233/DDBOT/lsp/concern"
	"github.com/Sora233/DDBOT/lsp/concern_type"
	"github.com/Sora233/DDBOT/lsp/mmsg"
	"github.com/Sora233/MiraiGo-Template/utils"
)

const (
	// 这个名字是日志中的名字，如果不知道取什么名字，可以和Site一样
	ConcernName = "my-great-concern-name"

	// 插件支持的网站名
	Site = "mysite"
	// 这个插件支持的订阅类型可以像这样自定义，然后在 Types 中返回
	Type1 concern_type.Type = "type1"
	// 当像这样定义的时候，支持 /watch -s mysite -t type1 id
	// 当实现的时候，请修改上面的定义
)

var logger = utils.GetModuleLogger(ConcernName)

type myStateManager struct {
	*concern.StateManager
}

// GetGroupConcernConfig 重写 concern.StateManager 的GetGroupConcernConfig方法，让我们自己定义的 GroupConcernConfig 生效
func (c *myStateManager) GetGroupConcernConfig(groupCode int64, id interface{}) concern.IConfig {
	return NewGroupConcernConfig(c.StateManager.GetGroupConcernConfig(groupCode, id))
}

type myConcern struct {
	*myStateManager
	*extraKey
}

func (c *myConcern) Site() string {
	return Site
}

func (c *myConcern) Types() []concern_type.Type {
	return []concern_type.Type{Type1}
}

func (c *myConcern) ParseId(s string) (interface{}, error) {
	// 在这里解析id
	// 此处返回的id类型，即是其他地方id interface{}的类型
	// 其他所有地方的id都由此函数生成
	// 推荐在string 或者 int64类型中选择其一
	// 如果订阅源有uid等数字唯一标识，请选择int64，如 bilibili
	// 如果订阅源有数字并且有字符，请选择string， 如 douyu
	panic("implement me")
}

func (c *myConcern) Add(ctx mmsg.IMsgCtx, groupCode int64, id interface{}, ctype concern_type.Type) (concern.IdentityInfo, error) {
	// 这里是添加订阅的函数
	// 可以使 c.StateManager.AddGroupConcern(groupCode, id, ctype) 来添加这个订阅
	// 通常在添加订阅前还需要通过id访问网站上的个人信息页面，来确定id是否存在，是否可以正常订阅
	panic("implement me")
}

func (c *myConcern) Remove(ctx mmsg.IMsgCtx, groupCode int64, id interface{}, ctype concern_type.Type) (concern.IdentityInfo, error) {
	// 大部分时候简单的删除即可
	// 如果还有更复杂的逻辑可以自由实现
	_, err := c.GetStateManager().RemoveGroupConcern(groupCode, id.(string), ctype)
	if err != nil {
		return nil, err
	}
	return c.Get(id)
}

func (c *myConcern) Get(id interface{}) (concern.IdentityInfo, error) {
	// 查看一个订阅的信息
	// 通常是查看数据库中是否有id的信息，如果没有可以去网页上获取
	panic("implement me")
}

func (c *myConcern) Start() error {
	// 如果需要启用轮询器，可以使用下面的方法
	//c.UseEmitQueue()
	// 下面两个函数是订阅的关键，需要实现，请阅读文档
	//c.StateManager.UseFreshFunc()
	//c.StateManager.UseNotifyGeneratorFunc()
	return c.StateManager.Start()
}

func (c *myConcern) GetStateManager() concern.IStateManager {
	return c.StateManager
}

func newConcern(notifyChan chan<- concern.Notify) *myConcern {
	c := &myConcern{extraKey: new(extraKey)}
	// 默认是string格式的id
	c.StateManager = concern.NewStateManagerWithStringID(Site, notifyChan)
	// 如果要使用int64格式的id，可以用下面的
	//c.StateManager = concern.NewStateManagerWithInt64ID(Site, notifyChan)
	return c
}
