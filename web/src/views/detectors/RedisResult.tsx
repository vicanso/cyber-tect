import { defineComponent } from "vue";
import { css } from "@linaria/core";
import { EyeRegular } from "@vicons/fa";
import { NCard, NDataTable, NPopover, NIcon } from "naive-ui";
import { RowData, TableColumn } from "naive-ui/lib/data-table/src/interface";

import useDetectorState, {
  redisDetectorResultList,
} from "../../states/detector";

import { newListColumn, newLevelValueColumn } from "../../components/ExTable";
import ExDetectorResultTable from "../../components/ExDetectorResultTable";
import { formatDate } from "../../helpers/util";

const popupClass = css`
  max-width: 800px;
  white-space: nowrap;
`;

export default defineComponent({
  name: "RedisResult",
  setup() {
    const fetch = async (params: Record<string, unknown>) => {
      await redisDetectorResultList(params);
    };
    return {
      redisDetectorResults: useDetectorState().redisDetectorResults,
      fetch,
    };
  },
  render() {
    const { redisDetectorResults, fetch } = this;
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
          return formatDate(row.updatedAt as string);
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
      <NCard title={"Redis检测结果"}>
        <ExDetectorResultTable
          columns={columns}
          data={redisDetectorResults}
          fetch={fetch}
        />
      </NCard>
    );
  },
});
