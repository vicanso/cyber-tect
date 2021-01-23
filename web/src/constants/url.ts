// 用户相关url
// 用户信息
export const USERS_ME = "/users/v1/me";
// 用户登录
export const USERS_LOGIN = "/users/v1/me/login";
// 用户行为
export const USERS_ACTIONS = "/users/v1/actions";
// 用户登录记录
export const USERS_LOGINS = "/users/v1/login-records";
// 用户角色列表
export const USERS_ROLES = "/users/v1/roles";
// 用户列表
export const USERS = "/users/v1";
// 根据ID查询用户信息
export const USERS_ID = "/users/v1/:id";

// flux相关查询
// 用户行为日志列表
export const FLUXES_TRACKERS = "/fluxes/v1/trackers";
// http出错列表
export const FLUXES_HTTP_ERRORS = "/fluxes/v1/http-errors";
// tag value列表
export const FLUXES_TAG_VALUES = "/fluxes/v1/tag-values/:measurement/:tag";

// 通用接口相关url
// 图形验证码
export const COMMONS_CAPTCHA = "/commons/captcha";
// schema状态列表
export const COMMONS_STATUSES = "/commons/schema-statuses";
// 路由列表
export const COMMONS_ROUTERS = "/commons/routers";
// 随机字符串
export const COMMONS_RANDOM_KEYS = "/commons/random-keys";

// 系统配置相关url
// 配置列表
export const CONFIGS = "/configurations/v1";
// 根据ID查询或更新配置
export const CONFIGS_ID = "/configurations/v1/:id";
// 当前有效配置
export const CONFIGS_CURRENT_VALID = "/configurations/v1/current-valid";

// 检测配置url
// http检测配置
export const DETECTORS_HTTPS = "/detectors/v1/https";
// 更新HTTP检测配置
export const DETECTORS_HTTPS_UPDATE = "/detectors/v1/https/:id";
// 获取接收者列表
export const DETECTORS_RECEIVERS = "/detectors/v1/receivers";
