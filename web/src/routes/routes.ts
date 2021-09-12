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
    component: () => import("../views/Logins"),
  },
  {
    path: "/user-trackers",
    name: names.userTrackers,
    component: () => import("../views/UserTrackers"),
  },
  {
    path: "/http-errors",
    name: names.httpErrors,
    component: () => import("../views/HTTPErrors"),
  },
  {
    path: "/requests",
    name: names.requests,
    component: () => import("../views/Requests"),
  },
  {
    path: "/mock-time",
    name: names.mockTime,
    component: () => import("../views/MockTime"),
  },
  {
    path: "/configs",
    name: names.configs,
    component: () => import("../views/Configs"),
  },
  {
    path: "/block-ips",
    name: names.blockIPs,
    component: () => import("../views/BlockIPs"),
  },
  {
    path: "/signed-keys",
    name: names.signedKeys,
    component: () => import("../views/SignedKeys"),
  },
  {
    path: "/router-mocks",
    name: names.routerMocks,
    component: () => import("../views/RouterMocks"),
  },
  {
    path: "/request-concurrencies",
    name: names.requestConcurrencies,
    component: () => import("../views/RequestConcurrencies"),
  },
  {
    path: "/session-interceptors",
    name: names.sessionInterceptors,
    component: () => import("../views/SessionInterceptors"),
  },
  {
    path: "/caches",
    name: names.caches,
    component: () => import("../views/Caches"),
  },
  {
    path: "/emails",
    name: names.emails,
    component: () => import("../views/Emails"),
  },
  {
    path: "/http-server-interceptors",
    name: names.httpServerInterceptors,
    component: () => import("../views/HTTPServerInterceptors"),
  },
];
