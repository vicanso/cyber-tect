import { defineComponent } from "vue";
import { css } from "@linaria/core";
import { NCard, NDataTable, NPopover } from "naive-ui";
import { RowData, TableColumn } from "naive-ui/lib/data-table/src/interface";

import useDetectorState, {
  pingDetectorResultList,
} from "../../states/detector";
import { newListColumn, newLevelValueColumn } from "../../components/ExTable";
import { formatSimpleDate } from "../../helpers/util";
import ExDetectorResultTable, {
  newShowMoreIcon,
} from "../../components/ExDetectorResultTable";

const popupClass = css`
  max-width: 800px;
  white-space: nowrap;
`;

export default defineComponent({
  name: "PingResult",
  setup() {
    const fetch = async (params: Record<string, unknown>) => {
      await pingDetectorResultList(params);
    };
    return {
      pingDetectorResults: useDetectorState().pingDetectorResults,
      fetch,
    };
  },
  render() {
    const { pingDetectorResults, fetch } = this;
    const columns: TableColumn[] = [
      {
        title: "名称",
        key: "taskName",
      },
      newLevelValueColumn({
        title: "结果",
        key: "result.desc",
      }),
      newListColumn({
        title: "检测IP",
        key: "ips",
      }),
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
          return formatSimpleDate(row.updatedAt as string);
        },
      },
      {
        title: "更多",
        key: "",
        render(data: Record<string, unknown>) {
          const columns: TableColumn[] = [
            {
              title: "IP",
              key: "ip",
              fixed: "left",
              width: 200,
            },
            newLevelValueColumn({
              title: "结果",
              key: "result.desc",
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
            trigger: newShowMoreIcon,
          };
          return (
            <NPopover v-slots={slots} placement="left-end" trigger="click">
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
      <NCard title={"Ping检测结果"}>
        <ExDetectorResultTable
          columns={columns}
          data={pingDetectorResults}
          fetch={fetch}
          category="ping"
        />
      </NCard>
    );
  },
});
