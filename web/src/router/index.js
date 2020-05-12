import Vue from 'vue'
import VueRouter from 'vue-router'

import store from '@/store'

import Home from '@/views/Home.vue'
import Login from '@/views/Login.vue'
import Register from '@/views/Register.vue'
import UserInfo from '@/views/UserInfo.vue'

import AddHTTP from '@/views/AddHTTP.vue'
import ListHTTP from '@/views/ListHTTP.vue'
import UpdateHTTP from '@/views/UpdateHTTP.vue'
import ListHTTPDetectorResult from '@/views/ListHTTPDetectorResult.vue'

import AddDNS from '@/views/AddDNS.vue'
import ListDNS from '@/views/ListDNS.vue'
import UpdateDNS from '@/views/UpdateDNS.vue'
import ListDNSDetectorResult from '@/views/ListDNSDetectorResult.vue'

import AddTCP from '@/views/AddTCP.vue'
import ListTCP from '@/views/ListTCP.vue'
import UpdateTCP from '@/views/UpdateTCP.vue'
import ListTCPDetectorResult from '@/views/ListTCPDetectorResult.vue'

import AddPing from '@/views/AddPing.vue'
import ListPing from '@/views/ListPing.vue'
import UpdatePing from '@/views/UpdatePing.vue'
import ListPingDetectorResult from '@/views/ListPingDetectorResult.vue'

import {
  HOME,
  LOGIN,
  REGISTER,
  USER_INFO,
  ADD_HTTP,
  LIST_HTTP,
  UPDATE_HTTP,
  LIST_HTTP_RESULT,
  ADD_DNS,
  LIST_DNS,
  UPDATE_DNS,
  LIST_DNS_RESULT,
  ADD_TCP,
  LIST_TCP,
  UPDATE_TCP,
  LIST_TCP_RESULT,
  ADD_PING,
  LIST_PING,
  UPDATE_PING,
  LIST_PING_RESULT
} from '@/paths'

Vue.use(VueRouter)

export const ROUTE_HOME = 'home'
export const ROUTE_LOGIN = 'login'
export const ROUTE_REGISTER = 'register'
export const ROUTE_USER_INFO = 'userInfo'

export const ROUTE_ADD_HTTP = 'addHTTP'
export const ROUTE_LIST_HTTP = 'listHTTP'
export const ROUTE_UPDATE_HTTP = 'updateHTTP'
export const ROUTE_LIST_HTTP_DETECTOR_RESULT = 'listHTTPDetectorResult'

export const ROUTE_ADD_DNS = 'addDNS'
export const ROUTE_LIST_DNS = 'listDNS'
export const ROUTE_UPDATE_DNS = 'updateDNS'
export const ROUTE_LIST_DNS_DETECTOR_RESULT = 'listDNSDetectorResult'

export const ROUTE_ADD_PING = 'addPing'
export const ROUTE_LIST_PING = 'listPing'
export const ROUTE_UPDATE_PING = 'updatePing'
export const ROUTE_LIST_PING_DETECTOR_RESULT = 'listPingDetectorResult'

export const ROUTE_ADD_TCP = 'addTCP'
export const ROUTE_LIST_TCP = 'listTCP'
export const ROUTE_UPDATE_TCP = 'updateTCP'
export const ROUTE_LIST_TCP_DETECTOR_RESULT = 'listTCPDetectorResult'

const routes = [
  {
    path: HOME,
    name: ROUTE_HOME,
    component: Home
  },
  {
    path: LOGIN,
    name: ROUTE_LOGIN,
    component: Login
  },
  {
    path: REGISTER,
    name: ROUTE_REGISTER,
    component: Register
  },
  {
    path: USER_INFO,
    name: ROUTE_USER_INFO,
    component: UserInfo,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: ADD_HTTP,
    name: ROUTE_ADD_HTTP,
    component: AddHTTP,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: LIST_HTTP,
    name: ROUTE_LIST_HTTP,
    component: ListHTTP,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: UPDATE_HTTP,
    name: ROUTE_UPDATE_HTTP,
    component: UpdateHTTP,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: LIST_HTTP_RESULT,
    name: ROUTE_LIST_HTTP_DETECTOR_RESULT,
    component: ListHTTPDetectorResult,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: ADD_DNS,
    name: ROUTE_ADD_DNS,
    component: AddDNS,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: LIST_DNS,
    name: ROUTE_LIST_DNS,
    component: ListDNS,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: LIST_DNS_RESULT,
    name: ROUTE_LIST_DNS_DETECTOR_RESULT,
    component: ListDNSDetectorResult,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: UPDATE_DNS,
    name: ROUTE_UPDATE_DNS,
    component: UpdateDNS,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: ADD_TCP,
    name: ROUTE_ADD_TCP,
    component: AddTCP,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: LIST_TCP,
    name: ROUTE_LIST_TCP,
    component: ListTCP,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: UPDATE_TCP,
    name: ROUTE_UPDATE_TCP,
    component: UpdateTCP,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: LIST_TCP_RESULT,
    name: ROUTE_LIST_TCP_DETECTOR_RESULT,
    component: ListTCPDetectorResult,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: ADD_PING,
    name: ROUTE_ADD_PING,
    component: AddPing,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: LIST_PING,
    name: ROUTE_LIST_PING,
    component: ListPing,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: UPDATE_PING,
    name: ROUTE_UPDATE_PING,
    component: UpdatePing,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: LIST_PING_RESULT,
    name: ROUTE_LIST_PING_DETECTOR_RESULT,
    component: ListPingDetectorResult,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '@/views/About.vue')
  }
]

const router = new VueRouter({
  mode: 'hash',
  base: process.env.BASE_URL,
  routes
})

let fetchedUserInfo = false
function waitForFetchingUserInfo () {
  const check = (resolve) => {
    if (fetchedUserInfo) {
      resolve()
      return
    }
    if (!store.state.user.fetching) {
      fetchedUserInfo = true
    }
    setTimeout(() => {
      check(resolve)
    }, 30)
  }

  return new Promise(resolve => {
    check(resolve)
  })
}

router.beforeEach(async (to, from, next) => {
  if (!fetchedUserInfo) {
    await waitForFetchingUserInfo()
  }
  if (!to.meta.requiresAuth) {
    return next()
  }
  if (!store.state.user.account) {
    return next({
      name: ROUTE_LOGIN
    })
  }
  return next()
})

export default router
