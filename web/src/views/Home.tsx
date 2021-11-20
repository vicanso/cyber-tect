import { defineComponent, onBeforeMount } from "vue";
import { NCard, useMessage, NSpin, NButton } from "naive-ui";
import { goTo } from "../routes";
import { names } from "../routes/routes";
import useDetectorState, { getResultSummaries } from "../states/detector";
import { showError } from "../helpers/util";
import { LocationQueryRaw } from "vue-router";

export default defineComponent({
  name: "HomeView",
  setup() {
    const message = useMessage();
    const { detectorResultSummaries } = useDetectorState();

    onBeforeMount(() => {
      const day = 24 * 3600 * 1000;
      getResultSummaries({
        startedAt: new Date(Date.now() - 7 * day).toISOString(),
      }).catch((err) => {
        showError(message, err);
      });
    });
    return {
      detectorResultSummaries,
    };
  },
  render() {
    const { detectorResultSummaries } = this;
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
        <NButton
          bordered={false}
          onClick={() =>
            goTo(item.route, {
              replace: false,
            })
          }
        >
          {item.name}
        </NButton>
      </li>
    ));
    const routes: Record<string, string> = {
      http: names.detectorHTTPResult,
      dns: names.detectorDNSResult,
      tcp: names.detectorTCPResult,
      ping: names.detectorPingResult,
      database: names.detectorDatabaseResult,
    };
    let summaryElement: JSX.Element = <p></p>;
    if (!detectorResultSummaries.processing) {
      if (detectorResultSummaries.items.length !== 0) {
        const summaryList = detectorResultSummaries.items.map((item) => {
          const goToResultView = (result: string) => {
            const query: LocationQueryRaw = {};
            if (result) {
              query.result = result;
            }
            goTo(routes[item.name], {
              replace: false,
              query,
            });
          };
          return (
            <li key={item.name}>
              <NButton bordered={false} onClick={() => goToResultView("")}>
                {item.name.toUpperCase()}({item.success + item.fail})
              </NButton>
              <NButton bordered={false} onClick={() => goToResultView("1")}>
                成功({item.success})
              </NButton>
              <NButton bordered={false} onClick={() => goToResultView("2")}>
                失败({item.fail})
              </NButton>
            </li>
          );
        });
        summaryElement = <ul>{summaryList}</ul>;
      } else {
        summaryElement = (
          <p>
            当前账号未配置接收相关检测结果，请先配置监控检测或添加接收监控告警
          </p>
        );
      }
    }
    return (
      <div>
        <NCard title="简要说明">
          <p>
            CyberTect提供界面式的配置，可以快速配置HTTP接口、TCP端口、DNS域名解析、Ping以及各常用数据库的定时检测告警，
            暂时告警仅通过邮件的形式通知，因此用户均需要在个人信息（顶部用户名处），维护接收邮箱。
          </p>
          <ul>{lists}</ul>
        </NCard>
        <br />
        <NSpin show={detectorResultSummaries.processing}>
          <NCard title="最近7天检测">
            <p>最近7天各类型检测的结果汇总（仅包括已配置为接收告警的检测）</p>
            {summaryElement}
          </NCard>
        </NSpin>
      </div>
    );
  },
});
