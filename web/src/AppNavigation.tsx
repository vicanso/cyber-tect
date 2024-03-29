import { Cogs, User, Tasks } from "@vicons/fa";
import { css } from "@linaria/core";
import { NButton, NIcon, NMenu } from "naive-ui";
import { Component, defineComponent, h } from "vue";
import { containsAny } from "./helpers/util";
import { goTo, goToLogin } from "./routes";
import { names } from "./routes/routes";
import useCommonState from "./states/common";
import useUserState from "./states/user";

function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

const loginButtonClass = css`
  margin: 50px auto;
  text-align: center;
`;

const navigationOptions = [
  {
    label: "检测配置",
    key: "detector",
    icon: renderIcon(Cogs),
    disabled: false,
    children: [
      {
        label: "HTTP",
        key: names.detectorHTTP,
      },
      {
        label: "TCP",
        key: names.detectorTCP,
      },
      {
        label: "Ping",
        key: names.detectorPing,
      },
      {
        label: "Database",
        key: names.detectorDatabase,
      },
      {
        label: "DNS",
        key: names.detectorDNS,
      },
    ],
  },
  {
    label: "检测结果",
    key: "detectorResult",
    icon: renderIcon(Tasks),
    disabled: false,
    children: [
      {
        label: "HTTP",
        key: names.detectorHTTPResult,
      },
      {
        label: "TCP",
        key: names.detectorTCPResult,
      },
      {
        label: "Ping",
        key: names.detectorPingResult,
      },
      {
        label: "Database",
        key: names.detectorDatabaseResult,
      },
      {
        label: "DNS",
        key: names.detectorDNSResult,
      },
    ],
  },
  {
    label: "用户",
    key: "user",
    icon: renderIcon(User),
    disabled: true,
    children: [
      {
        label: "用户列表",
        key: names.users,
      },
      {
        label: "登录记录",
        key: names.logins,
      },
    ],
  },

  // {
  //   label: "统计",
  //   key: "stats",
  //   icon: renderIcon(ChartBar),
  //   children: [
  //     {
  //       label: "用户行为",
  //       key: names.userTrackers,
  //     },
  //     {
  //       label: "响应出错记录",
  //       key: names.httpErrors,
  //     },
  //     {
  //       label: "后端HTTP调用",
  //       key: names.requests,
  //     },
  //   ],
  // },
  // {
  //   label: "配置",
  //   key: "settings",
  //   disabled: true,
  //   icon: renderIcon(Cogs),
  //   children: [
  //     {
  //       label: "所有配置",
  //       key: names.configs,
  //     },
  //     {
  //       label: "MockTime配置",
  //       key: names.mockTime,
  //     },
  //     {
  //       label: "黑名单IP",
  //       key: names.blockIPs,
  //     },
  //     {
  //       label: "SignedKey配置",
  //       key: names.signedKeys,
  //     },
  //     {
  //       label: "路由Mock配置",
  //       key: names.routerMocks,
  //     },
  //     {
  //       label: "HTTP实例并发配置",
  //       key: names.requestConcurrencies,
  //     },
  //     {
  //       label: "Session拦截配置",
  //       key: names.sessionInterceptors,
  //     },
  //     {
  //       label: "HTTP服务拦截配置",
  //       key: names.httpServerInterceptors,
  //     },
  //     {
  //       label: "接收邮箱配置",
  //       key: names.emails,
  //     },
  //   ],
  // },
  // {
  //   label: "其它",
  //   key: "others",
  //   disabled: true,
  //   icon: renderIcon(Deezer),
  //   children: [
  //     {
  //       label: "缓存",
  //       key: names.caches,
  //     },
  //   ],
  // },
];

export default defineComponent({
  name: "AppNavigation",
  setup() {
    const { info } = useUserState();
    const { settings } = useCommonState();
    return {
      settings,
      userInfo: info,
      handleNavigation(key: string): void {
        goTo(key, {
          replace: false,
        });
      },
    };
  },
  render() {
    const { userInfo, $router, settings } = this;
    if (userInfo.processing) {
      return <p class="tac">...</p>;
    }
    if (!userInfo.account) {
      if (settings.collapsed) {
        return <div />;
      }
      return (
        <div class={loginButtonClass}>
          <NButton type="info" onClick={() => goToLogin()}>
            立即登录
          </NButton>
        </div>
      );
    }
    const options = navigationOptions.slice(0);
    if (containsAny(userInfo.roles, ["su", "admin"])) {
      options.forEach((item) => {
        if (item.disabled) {
          item.disabled = false;
        }
      });
    }
    const currentRoute = $router.currentRoute.value.name?.toString();
    return (
      <NMenu
        value={currentRoute}
        defaultExpandAll={true}
        onUpdate:value={this.handleNavigation}
        options={options}
        collapsedWidth={64}
        collapsed={settings.collapsed}
      />
    );
  },
});
