import { defineComponent } from "vue";
import { NCard } from "naive-ui";
import { css } from "@linaria/core";

import { goTo } from "../routes";
import { names } from "../routes/routes";

const anchorClass = css`
  cursor: pointer;
  margin: 10px 0;
  display: block;
`;

export default defineComponent({
  name: "Home",
  render() {
    const configs = [
      {
        route: names.detectorHTTP,
        name: "立即配置 HTTP 检测",
      },
      {
        route: names.detectorTCP,
        name: "立即配置 TCP 检测",
      },
      {
        route: names.detectorPing,
        name: "立即配置 Ping 检测",
      },
      {
        route: names.detectorDatabase,
        name: "立即配置 Database 检测",
      },
      {
        route: names.detectorDNS,
        name: "立即配置 DNS 检测",
      },
    ];
    const lists = configs.map((item) => (
      <li key={item.name}>
        <a
          class={anchorClass}
          onClick={() =>
            goTo(item.route, {
              replace: false,
            })
          }
        >
          {item.name}
        </a>
      </li>
    ));
    return (
      <div>
        <NCard title="简要说明">
          <p>
            CyberTect提供界面式的配置，可以快速配置HTTP接口、TCP端口、DNS域名解析、Ping以及各常用数据库的定时检测告警，
            暂时告警仅通过邮件的形式通知，因此用户均需要在个人信息（顶部用户名处），维护接收邮箱。
          </p>
          <ul>{lists}</ul>
        </NCard>
      </div>
    );
  },
});
