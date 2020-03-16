package errorcode

// All common error code
var (
	OK = add(0, "success") // 正确

	AppKeyInvalid           = add(1, "app key invalid")                                  // 应用程序不存在或已被封禁
	AccessKeyErr            = add(2, "access key is failed")                             // Access Key错误
	SignCheckErr            = add(3, "signature verification failed")                    // API校验密匙错误
	MethodNoPermission      = add(4, "caller does not have permission on the method")    // 调用方对该Method没有权限
	NoLogin                 = add(101, "account not logged in")                          // 账号未登录
	UserDisabled            = add(102, "account is blocked")                             // 账号被封停
	LackOfScores            = add(103, "integral deficiency")                            // 积分不足
	LackOfCoins             = add(104, "credit is running low")                          // 余额不足
	CaptchaErr              = add(105, "verification code error")                        // 验证码错误
	UserInactive            = add(106, "account not activated")                          // 账号未激活
	UserNoMember            = add(107, "account number informal member or on probation") // 账号非正式会员或在适应期
	AppDenied               = add(108, "application does not exists or is banned")       // 应用不存在或者被封禁
	MobileNoVerfiy          = add(110, "mobile phone not bound")                         // 未绑定手机
	CsrfNotMatchErr         = add(111, "csrf token verification failed")                 // csrf 校验失败
	ServiceUpdate           = add(112, "system upgrading")                               // 系统升级中
	UserIDCheckInvalid      = add(113, "account has not been verified by real name")     // 账号尚未实名认证
	UserIDCheckInvalidPhone = add(114, "please bind your phone first")                   // 请先绑定手机
	UserIDCheckInvalidCard  = add(115, "please complete real name authentication first") // 请先完成实名认证

	NotModified           = add(304, "no change")                                            // 木有改动
	TemporaryRedirect     = add(307, "crash jump")                                           // 撞车跳转
	RequestErr            = add(400, "request error")                                        // 请求错误
	Unauthorized          = add(401, "no authentication")                                    // 未认证
	AccessDenied          = add(403, "insufficient access")                                  // 访问权限不足
	NothingFound          = add(404, "not found")                                            // 啥都木有
	MethodNotAllowed      = add(405, "the method is not supported")                          // 不支持该方法
	Conflict              = add(409, "conflict")                                             // 冲突

	ServerErr             = add(500, "server error")                                         // 服务器错误
	ServiceUnavailable    = add(503, "overload protection, service temporarily unavailable") // 过载保护,服务暂不可用
	Deadline              = add(504, "server request time out")                              // 服务调用超时
	LimitExceed           = add(509, "exceeding the limit")


	FileNotExists         = add(616, "upload file not exists")                               // 上传文件不存在
	FileTooLarge          = add(617, "upload file too long")                                 // 上传文件太大
	FailedTooManyTimes    = add(625, "too many login failures")                              // 登录失败次数太多
	UserNotExist          = add(626, "account not exists")                                   // 用户不存在
	PasswordTooLeak       = add(628, "password is too weak")                                 // 密码太弱
	UsernameOrPasswordErr = add(629, "username or password failed")                          // 用户名或密码错误
	TargetNumberLimit     = add(632, "limit on the number of operands")                      // 操作对象数量限制
	TargetBlocked         = add(643, "is locked")                                            // 被锁定
	UserLevelLow          = add(650, "user level too low")                                   // 用户等级太低
	UserDuplicate         = add(652, "duplicate users")                                      // 重复的用户
	AccessTokenExpires    = add(658, "token expired")                                        // Token 过期
	PasswordHashExpires   = add(662, "password time stamp expired")                          // 密码时间戳过期
	AreaLimit             = add(688, "geographical limits")                                  // 地理区域限制
	CopyrightLimit        = add(689, "copyright restrictions")                               // 版权限制

	ParamsInvalid = add(700, "invalid parameter")   // 参数非法

	DataUnmarshalFailed = add(800, "failed to deserialize data") // 反序列化数据失败
	DataFailed          = add(801, "failed data")                // 反序列化数据失败

	Degrade     = add(1200, "degraded filtered requests")              // 被降级过滤的请求
	RPCNoClient = add(1201, "no client of rpc service is available")   // rpc服务的client都不可用
	RPCNoAuth   = add(1202, "client of rpc service is not authorized") // rpc服务的client没有授权
)

