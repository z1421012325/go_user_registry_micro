
## micro server 服务介绍

```shell script
模型
用户表
    用户id
    昵称
    账号  手机号码或者email
    密码
    注册手段 比如 手机号码,email
    注册状态
    注册时间
    注销时间(软删除)

    ...
```

```shell script
功能
    注册账号(未完成确认注册)
      -注册账号查重
      -发送短信验证码,email验证码(放入redis保存验证吗)
      -密码加密,解密

    确认注册
      -根据已保存在mysql中的自动生成id(未完成注册确定)去redis中寻找保存的验证码与发送过来的验证吗确定,request信息与mysql确定

    登录账号,返回token
      -生成redis唯一存证,redis中以用户id为key,存证为value(redis能否抗住并发?)
      -token含有过期时间,用户信息(除密码),redis唯一存证(多个token时作唯一)

    退出登录
      -根据token中的用户id去redis中删除存证
```