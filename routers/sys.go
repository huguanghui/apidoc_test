/**
 * @apiDefine sys 系统
 */


/**
 * @api {get} /api/v1/getserverinfo 获取平台运行信息
 * @apiGroup sys
 * @apiName GetServerInfo
 * @apiSuccess (200) {String} Hardware 硬件信息
 * @apiSuccess (200) {String} RunningTime 运行时间
 * @apiSuccess (200) {String} StartUpTime 启动时间
 * @apiSuccess (200) {String} Server 软件信息
 */

/**
 * @api {get} /api/v1/restart 重启服务
 * @apiGroup sys
 * @apiName Restart
 * @apiUse simpleSuccess
 */

/**
 * @apiDefine userInfo
 * @apiSuccess (200) {String} id
 * @apiSuccess (200) {String} name 用户名
 * @apiSuccess (200) {String[]} [roles] 角色列表
 */

/**
 * @api {get} /api/v1/login 登录
 * @apiGroup sys
 * @apiName Login
 * @apiParam {String} username 用户名
 * @apiParam {String} password 密码(经过md5加密,32位长度,不带中划线,不区分大小写)
 * @apiSuccessExample 成功
 * HTTP/1.1 200 OK
 * Set-Cookie: token=s%3ArkyMbQE0M.5AKAOXbW8c7iP%2BOo0venPkCYiEiPK9FY31mB6AlFQak;//用着后续接口调用的 token
 */

/**
 * @api {get} /api/v1/userInfo 获取当前登录用户信息
 * @apiGroup sys
 * @apiName UserInfo
 * @apiUse userInfo
 */

/**
 * @api {get} /api/v1/logout 登出
 * @apiGroup sys
 * @apiName Logout
 * @apiUse simpleSuccess
 */
