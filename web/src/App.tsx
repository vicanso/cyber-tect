import { NLayout, NLayoutSider, useLoadingBar, useMessage } from "naive-ui";
import { css } from "@linaria/core";
import { defineComponent, onMounted, watch } from "vue";
import AppHeader from "./AppHeader";
import AppNavigation from "./AppNavigation";
import {
  mainHeaderHeight,
  mainNavigationWidth,
  padding,
} from "./constants/style";
import "./main.css";
import { setLoadingEvent } from "./routes/router";
import useCommonState, { commonUpdateSettingCollapsed } from "./states/common";
import useUserState from "./states/user";

const layoutClass = css`
  top: ${mainHeaderHeight}px !important;
`;

const contentLayoutClass = css`
  padding: ${2 * padding}px;
`;

import { userMeDetail } from "./states/user";
import { toast } from "./helpers/util";

export default defineComponent({
  name: "App",
  setup() {
    const message = useMessage();
    const { settings } = useCommonState();
    const { info } = useUserState();
    watch(
      () => info.account,
      (account) => {
        // 未登录
        if (!account) {
          return;
        }
        userMeDetail().then((userInfo) => {
          if (!userInfo.account) {
            return;
          }
          if (!userInfo.alarmURL && !userInfo.email) {
            toast(
              message,
              "未设置接收告警信息的邮箱，请先在个人信息中设置(右上角个人信息)"
            );
          }
        });
      }
    );

    const loadingBar = useLoadingBar();
    if (loadingBar != undefined) {
      setLoadingEvent(loadingBar.start, loadingBar.finish);
      onMounted(() => {
        loadingBar.finish();
      });
    }
    return {
      settings,
    };
  },
  render() {
    const { settings } = this;
    return (
      <div>
        <AppHeader />
        <NLayout hasSider position="absolute" class={layoutClass}>
          <NLayoutSider
            bordered
            collapseMode="width"
            collapsed={settings.collapsed}
            collapsedWidth={64}
            width={mainNavigationWidth}
            showTrigger
            onCollapse={() => {
              commonUpdateSettingCollapsed(true);
            }}
            onExpand={() => {
              commonUpdateSettingCollapsed(false);
            }}
          >
            <AppNavigation />
          </NLayoutSider>
          <NLayout class={contentLayoutClass}>
            <router-view />
          </NLayout>
        </NLayout>
      </div>
    );
  },
});
