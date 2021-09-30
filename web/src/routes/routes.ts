import { Component } from "vue";

export interface Router {
  path: string;
  name: string;
  component: Component | Promise<Component>;
}

export const names = {
  home: "home",
  login: "logon",
  register: "register",
  users: "users",
  logins: "logins",
  userTrackers: "userTrackers",
  httpErrors: "httpErrors",
  requests: "requests",
  mockTime: "mockTime",
  configs: "configs",
  blockIPs: "blockIPs",
  signedKeys: "signedKeys",
  routerMocks: "routerMocks",
  requestConcurrencies: "requestConcurrencies",
  sessionInterceptors: "sessionInterceptors",
  caches: "caches",
  emails: "emails",
  httpServerInterceptors: "httpServerInterceptors",
  detectorHTTP: "detectorHTTP",
  detectorDNS: "detectorDNS",
  detectorTCP: "detectorTCP",
  detectorPing: "detectorPing",
  detectorHTTPResult: "detectorHTTPResult",
};

export const routes: Router[] = [
  {
    path: "/",
    name: names.home,
    component: () => import("../views/Home"),
  },
  {
    path: "/login",
    name: names.login,
    component: () => import("../views/Login"),
  },
  {
    path: "/register",
    name: names.register,
    component: () => import("../views/Register"),
  },
  {
    path: "/users",
    name: names.users,
    component: () => import("../views/Users"),
  },
  {
    path: "/logins",
    name: names.logins,
    component: () => import("../views/stats/Logins"),
  },
  {
    path: "/user-trackers",
    name: names.userTrackers,
    component: () => import("../views/stats/UserTrackers"),
  },
  {
    path: "/http-errors",
    name: names.httpErrors,
    component: () => import("../views/stats/HTTPErrors"),
  },
  {
    path: "/requests",
    name: names.requests,
    component: () => import("../views/stats/Requests"),
  },
  {
    path: "/mock-time",
    name: names.mockTime,
    component: () => import("../views/configurations/MockTime"),
  },
  {
    path: "/configs",
    name: names.configs,
    component: () => import("../views/configurations/Configs"),
  },
  {
    path: "/block-ips",
    name: names.blockIPs,
    component: () => import("../views/configurations/BlockIPs"),
  },
  {
    path: "/signed-keys",
    name: names.signedKeys,
    component: () => import("../views/configurations/SignedKeys"),
  },
  {
    path: "/router-mocks",
    name: names.routerMocks,
    component: () => import("../views/configurations/RouterMocks"),
  },
  {
    path: "/request-concurrencies",
    name: names.requestConcurrencies,
    component: () => import("../views/configurations/RequestConcurrencies"),
  },
  {
    path: "/session-interceptors",
    name: names.sessionInterceptors,
    component: () => import("../views/configurations/SessionInterceptors"),
  },
  {
    path: "/caches",
    name: names.caches,
    component: () => import("../views/Caches"),
  },
  {
    path: "/emails",
    name: names.emails,
    component: () => import("../views/configurations/Emails"),
  },
  {
    path: "/http-server-interceptors",
    name: names.httpServerInterceptors,
    component: () => import("../views/configurations/HTTPServerInterceptors"),
  },
  {
    path: "/detectors/http",
    name: names.detectorHTTP,
    component: () => import("../views/detectors/HTTP"),
  },
  {
    path: "/detectors/dns",
    name: names.detectorDNS,
    component: () => import("../views/detectors/DNS"),
  },
  {
    path: "/detectors/tcp",
    name: names.detectorTCP,
    component: () => import("../views/detectors/TCP"),
  },
  {
    path: "/detectors/ping",
    name: names.detectorPing,
    component: () => import("../views/detectors/Ping"),
  },
  {
    path: "/detectors/http/results",
    name: names.detectorHTTPResult,
    component: () => import("../views/detectors/HTTPResult"),
  },
];
