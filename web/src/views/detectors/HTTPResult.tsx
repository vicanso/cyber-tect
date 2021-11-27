import { defineComponent } from "vue";
import { css } from "@linaria/core";
import { EyeRegular } from "@vicons/fa";
import { NCard, NDataTable, NPopover, NIcon } from "naive-ui";
import { RowData, TableColumn } from "naive-ui/lib/data-table/src/interface";

import useDetectorState, {
  httpDetectorResultList,
} from "../../states/detector";
import { newListColumn, newLevelValueColumn } from "../../components/ExTable";
import ExDetectorResultTable from "../../components/ExDetectorResultTable";
import { formatDate } from "../../helpers/util";

const popupClass = css`
  max-width: 980px;
  white-space: nowrap;
`;

export default defineComponent({
  name: "HTTPResult",
  setup() {
    const fetch = async (params: Record<string, unknown>) => {
      await httpDetectorResultList(params);
    };
    return {
      httpDetectorResults: useDetectorState().httpDetectorResults,
      fetch,
    };
  },
  render() {
    const { httpDetectorResults, fetch } = this;
    const columns: TableColumn[] = [
      {
        title: "名称",
        key: "taskName",
        width: 120,
      },
      newLevelValueColumn({
        title: "结果",
        key: "result.desc",
        width: 80,
      }),
      {
        title: "检测URL",
        key: "url",
        width: 300,
      },
      {
        title: "最大耗时(ms)",
        key: "maxDuration",
        width: 120,
      },
      newListColumn({
        title: "失败信息",
        key: "messages",
      }),
      {
        title: "更新于",
        key: "updatedAt",
        width: 180,
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
              title: "地址",
              key: "addr",
              fixed: "left",
              width: 200,
            },
            {
              title: "代理地址",
              key: "proxy",
              width: 200,
            },
            newLevelValueColumn({
              title: "结果",
              key: "result.desc",
            }),
            {
              title: "失败信息",
              key: "message",
              ellipsis: {
                tooltip: true,
              },
              width: 200,
            },
            newListColumn({
              key: "timeline",
              title: "耗时(ms)",
              width: 180,
            }),
            {
              title: "HTTP协议",
              key: "protocol",
            },
            {
              title: "TLS",
              key: "tlsVersion",
            },
            {
              title: "TLS加密",
              key: "tlsCipherSuite",
              ellipsis: {
                tooltip: true,
              },
              width: 100,
            },
            newListColumn({
              key: "certificateDNSNames",
              title: "证书域名",
            }),
            newListColumn({
              key: "certificateExpirationDates",
              title: "证书有效期",
            }),
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
                  scrollX={1300}
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
      <NCard title={"HTTP检测结果"}>
        <ExDetectorResultTable
          columns={columns}
          data={httpDetectorResults}
          fetch={fetch}
          category="http"
        />
      </NCard>
    );
  },
});
