import { defineComponent } from "vue";
import { css } from "@linaria/core";
import { EyeRegular } from "@vicons/fa";
import { NCard, NDataTable, NPopover, NIcon } from "naive-ui";
import { RowData, TableColumn } from "naive-ui/lib/data-table/src/interface";

import useDetectorState, { dnsDetectorResultList } from "../../states/detector";
import { newListColumn, newLevelValueColumn } from "../../components/ExTable";
import { formatDate } from "../../helpers/util";
import ExDetectorResultTable from "../../components/ExDetectorResultTable";

const popupClass = css`
  max-width: 800px;
  white-space: nowrap;
`;

export default defineComponent({
  name: "DNSResult",
  setup() {
    const fetch = async (params: Record<string, unknown>) => {
      await dnsDetectorResultList(params);
    };
    return {
      dnsDetectorResults: useDetectorState().dnsDetectorResults,
      fetch,
    };
  },
  render() {
    const { dnsDetectorResults, fetch } = this;
    const columns: TableColumn[] = [
      {
        title: "名称",
        key: "taskName",
      },
      newLevelValueColumn({
        title: "结果",
        key: "result.desc",
      }),
      {
        title: "域名",
        key: "host",
      },
      //   newListColumn({
      //     title: "检测IP",
      //     key: "ips",
      //   }),
      {
        title: "最大耗时(ms)",
        key: "maxDuration",
      },
      newListColumn({
        title: "失败信息",
        key: "messages",
      }),
      {
        title: "更新于",
        key: "updatedAt",
        render(row: Record<string, unknown>) {
          return formatDate(row.updatedAt as string);
        },
      },
      {
        title: "更多",
        key: "",
        render(data: Record<string, unknown>) {
          const columns: TableColumn[] = [
            {
              title: "DNS服务器",
              key: "server",
              fixed: "left",
              width: 200,
            },
            newLevelValueColumn({
              title: "结果",
              key: "result.desc",
            }),
            newListColumn({
              title: "IP列表",
              key: "ips",
            }),
            {
              title: "耗时(ms)",
              key: "duration",
            },
            {
              title: "失败信息",
              key: "message",
            },
          ];
          const slots = {
            trigger: () => (
              <NIcon>
                <EyeRegular />
              </NIcon>
            ),
          };
          return (
            <NPopover v-slots={slots} placement="left-end">
              <div class={popupClass}>
                <NDataTable
                  columns={columns}
                  data={data.results as RowData[]}
                />
              </div>
            </NPopover>
          );
        },
      },
    ];
    return (
      <NCard title={"DNS检测结果"}>
        <ExDetectorResultTable
          columns={columns}
          data={dnsDetectorResults}
          fetch={fetch}
          category="dns"
        />
      </NCard>
    );
  },
});
