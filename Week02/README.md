## 作业

我们在数据库操作的时候，比如 `dao` 层中当遇到一个 `sql.ErrNoRows` 的时候，是否应该 `Wrap` 这个 `error`，抛给上层。为什么？应该怎么做请写出代码

## 个人理解

场景举例：

- `PUT {host}/user/:user_id` 如果找不到该ID，应该报错
- `GET {host}/user/:user_id` 如果找不到该ID，应该报错
- `DEL {host}/user/:user_id` 如果找不到该ID，应该报错

总结：报错合适，如果有某一些特定场景下认为不应该报错，那么应该在该业务的service层内单独消化这种情况