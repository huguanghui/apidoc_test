/**
 * @apiDefine test 测试
 */

/**
 * @api {get} /api/v1/getstatus 获取当前状态
 * @apiGroup test
 * @apiName GetStatus
 * @apiParam {String} url RTSP源地址
 * @apiParam {String=TCP,UDP} [transType=TCP] 拉流传输模式
 * @apiSuccess (200) {String} ID  ID号
 * @apiSuccess (200) {String} Status 状态
 * @apiSuccess (200) {String} Time 时间
 * @apiUse simpleSuccess
 */