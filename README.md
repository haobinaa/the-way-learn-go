### the-way-to-go 学习

学习 go 语言的一些语法特性， 常见库的用法



#### 语法糖/语言特性

- [slice 内存共享](./code-block/slice_share_memory.go)
- [接口完整性检测(静态)/动态类型断言](./code-block/interface_full_check.go)
- [err处理模式-屏蔽过程中的错误](./code-block/mask_process_err.go)
- [funcOptions-函数式编程解决重载场景](./code-block/func_option.go)
- [map/reduce-控制逻辑和业务逻辑分离](./code-block/func_option.go)
- [接口的动态类型和动态值](./code-block/itab_data.go)
- [encoding/json unmarshal float64精度问题/slice、map等类型 marshal为nil问题](./code-block/json_marshal.go)
- [实现goroutines的守护协程](./code-block/protect_goroutines.go)

#### 限流实现
- [固定窗口](./code-block/limiter/fixed_window_limiter.go)
- [滑动窗口](./code-block/limiter/sliding_window_limiter.go)
- [漏桶](./code-block/limiter/leaky_bucket.go)
- [令牌桶](./code-block/limiter/token_bucket.go)

#### go 实现设计模式:
- [责任链模式](./design-patterns/chain-of-responsibility/chain-of-responsibility.go)
- [命令模式](./design-patterns/command/command.go)
- [迭代器模式](./design-patterns/iterator/iterator.go)
- [中介模式](./design-patterns/mediator/mediator.go)
- [备忘录模式](./design-patterns/memento/memento.go)
- [观察者模式](./design-patterns/observer/observer.go)
- [状态模式](./design-patterns/state/state.go)
- [策略模式](./design-patterns/strategy/strategy.go)
- [模版模式](./design-patterns/template-method/template_method.go)
- [访问者模式](./design-patterns/visitor/visitor.go)