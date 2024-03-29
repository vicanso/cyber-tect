// 用户相关url
// 用户信息
export const USERS_ME = "/users/v1/me";
// 用户详细信息
export const USERS_ME_DETAIL = "/users/v1/detail";
// 用户登录
export const USERS_LOGIN = "/users/v1/me/login";
export const USERS_INNER_LOGIN = "/users/inner/v1/me/login";
// 用户列表
export const USERS = "/users/v1";
// 用户登录记录
export const USERS_LOGINS = "/users/v1/login-records";
export const USERS_ID = "/users/v1/:id";

// 通用接口相关url
// 图形验证码
export const COMMONS_CAPTCHA = "/commons/captcha";
// 路由列表
export const COMMONS_ROUTERS = "/commons/routers";
// HTTP性能指标统计
export const COMMONS_HTTP_STATS = "/commons/http-stats";

// flux相关查询
// 用户行为日志列表
export const FLUXES_TRACKERS = "/fluxes/v1/trackers";
// http出错列表
export const FLUXES_HTTP_ERRORS = "/fluxes/v1/http-errors";
// 客户端上传的action日志列表
export const FLUXES_ACTIONS = "/fluxes/v1/actions";
// 后端HTTP调用列表
export const FLUXES_REQUESTS = "/fluxes/v1/requests";
// tag value列表
export const FLUXES_TAG_VALUES = "/fluxes/v1/tag-values/:measurement/:tag";
// flux单条记录查询
export const FLUXES_FIND_ONE = "/fluxes/v1/one/:measurement";

// 系统配置相关url
// 配置列表
export const CONFIGS = "/configurations/v1";
// 根据ID查询或更新配置
export const CONFIGS_ID = "/configurations/v1/:id";
// 当前有效配置
export const CONFIGS_CURRENT_VALID = "/configurations/v1/current-valid";

// 管理员相关接口
export const ADMINS_CACHE_ID = "/@admin/v1/caches/:key";

export const DETECTOR_LIST_USER = "/detectors/v1/users";
export const DETECTOR_RESULT_SUMMARIES = "/detectors/v1/result-summaries";
export const DETECTOR_LIST_TASK = "/detectors/v1/tasks/:category";
// http检测
export const HTTP_DETECTORS = "/http-detectors/v1";
export const HTTP_DETECTORS_ID = "/http-detectors/v1/:id";
export const HTTP_DETECTOR_RESULTS = "/http-detectors/v1/results";
// dns检测
export const DNS_DETECTORS = "/dns-detectors/v1";
export const DNS_DETECTORS_ID = "/dns-detectors/v1/:id";
export const DNS_DETECTOR_RESULTS = "/dns-detectors/v1/results";
// tcp检测
export const TCP_DETECTORS = "/tcp-detectors/v1";
export const TCP_DETECTORS_ID = "/tcp-detectors/v1/:id";
export const TCP_DETECTOR_RESULTS = "/tcp-detectors/v1/results";
// ping检测
export const PING_DETECTORS = "/ping-detectors/v1";
export const PING_DETECTORS_ID = "/ping-detectors/v1/:id";
export const PING_DETECTOR_RESULTS = "/ping-detectors/v1/results";
// database检测
export const DATABASE_DETECTORS = "/database-detectors/v1";
export const DATABASE_DETECTORS_ID = "/database-detectors/v1/:id";
export const DATABASE_DETECTOR_RESULTS = "/database-detectors/v1/results";
