import { defineComponent } from "vue";
import { css } from "@linaria/core";
import { NCard, NDataTable, NPopover } from "naive-ui";
import { RowData, TableColumn } from "naive-ui/lib/data-table/src/interface";

import useDetectorState, {
  databaseDetectorResultList,
} from "../../states/detector";

import { newListColumn, newLevelValueColumn } from "../../components/ExTable";
import ExDetectorResultTable, {
  newShowMoreIcon,
} from "../../components/ExDetectorResultTable";
import { formatSimpleDate } from "../../helpers/util";

const popupClass = css`
  max-width: 800px;
  white-space: nowrap;
`;

export default defineComponent({
  name: "DatabaseResult",
  setup() {
    const fetch = async (params: Record<string, unknown>) => {
      await databaseDetectorResultList(params);
    };
    return {
      databaseDetectorResults: useDetectorState().databaseDetectorResults,
      fetch,
    };
  },
  render() {
    const { databaseDetectorResults, fetch } = this;
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
        title: "连接串列表",
        key: "uris",
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
              title: "连接串",
              key: "uri",
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
      <NCard title={"Database检测结果"}>
        <ExDetectorResultTable
          columns={columns}
          data={databaseDetectorResults}
          fetch={fetch}
          category="database"
        />
      </NCard>
    );
  },
});
