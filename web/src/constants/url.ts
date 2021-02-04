// 用户相关url
// 用户信息
export const USERS_ME = "/users/v1/me";
// 用户详细信息
export const USERS_ME_DETAIL = "/users/v1/detail";
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
// 客户端上传的action日志列表
export const FLUXES_ACTIONS = "/fluxes/v1/actions";
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
// 获取接收者列表
export const DETECTORS_RECEIVERS = "/detectors/v1/receivers";
// HTTP检测配置
export const DETECTORS_HTTPS = "/detectors/v1/https";
// 更新HTTP检测配置
export const DETECTORS_HTTPS_UPDATE = "/detectors/v1/https/:id";
// 获取HTTP检测结果
export const DETECTORS_HTTPS_RESULTS = "/detectors/v1/https/results";
// 获取HTTP检测结果详情
export const DETECTORS_HTTPS_RESULTS_DETAIL = "/detectors/v1/https/results/:id";

// DNS检测配置
export const DETECTORS_DNSES = "/detectors/v1/dnses";
// 更新DNS检测配置
export const DETECTORS_DNSES_UPDATE = "/detectors/v1/dnses/:id";
// 获取DNS检测结果
export const DETECTORS_DNSES_RESULTS = "/detectors/v1/dnses/results";
// 获取DNS检测结果详情
export const DETECTORS_DNSES_RESULTS_DEATIL = "/detectors/v1/dnses/results/:id";

// TCP检测配置
export const DETECTORS_TCPS = "/detectors/v1/tcps";
// 更新TCP检测配置
export const DETECTORS_TCPS_UPDATE = "/detectors/v1/tcps/:id";
// 获取TCP检测结果
export const DETECTORS_TCPS_RESULTS = "/detectors/v1/tcps/results";
// 获取TCP检测结果详情
export const DETECTORS_TCPS_RESULTS_DETAIL = "/detectors/v1/tcps/results/:id";

// Ping检测配置
export const DETECTORS_PINGS = "/detectors/v1/pings";
// 更新Ping检测配置
export const DETECTORS_PINGS_UPDATE = "/detectors/v1/pings/:id";
// 获取Ping检测结果
export const DETECTORS_PINGS_RESULTS = "/detectors/v1/pings/results";
// 获取Ping检测结果详情
export const DETECTORS_PINGS_RESULTS_DETAIL = "/detectors/v1/pings/results/:id";
