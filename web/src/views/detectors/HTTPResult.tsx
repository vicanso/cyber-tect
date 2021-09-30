import { defineComponent, ref, PropType } from "vue";
import { NButton, NCard, useMessage, NDataTable } from "naive-ui";
import { RowData, TableColumn } from "naive-ui/lib/data-table/src/interface";

import useDetectorState, {
  httpDetectorResultList,
} from "../../states/detector";
import { showError } from "../../helpers/util";
import ExTable, { newListColumn } from "../../components/ExTable";
import { formatDate } from "../../helpers/util";

export default defineComponent({
  name: "HTTPResult",
  setup() {
    const message = useMessage();
    const fetch = async (params: Record<string, unknown>) => {
      try {
        await httpDetectorResultList(params);
      } catch (err) {
        showError(message, err);
      }
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
      },
      {
        type: "expand",
        expandable: (data: Record<string, unknown>) => true,
        renderExpand: (data: Record<string, unknown>) => {
          const columns: TableColumn[] = [
            {
              title: "地址",
              key: "addr",
            },
            {
              title: "结果",
              key: "resultDesc",
            },
            {
              title: "失败信息",
              key: "message",
            },
            newListColumn({
              key: "timeline",
              title: "耗时",
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
              title: "TLS加密套件",
              key: "tlsCipherSuite",
              ellipsis: true,
              width: 50,
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
          return (
            <NDataTable columns={columns} data={data.results as RowData[]} />
          );
        },
      },
      {
        title: "结果",
        key: "resultDesc",
      },
      {
        title: "检测URL",
        key: "url",
      },
      {
        title: "最大耗时(ms)",
        key: "maxDuration",
      },
      {
        title: "失败信息",
        key: "message",
      },
      {
        title: "更新于",
        key: "updatedAt",
        render(row: Record<string, unknown>) {
          return formatDate(row.updatedAt as string);
        },
      },
    ];
    return (
      <NCard title={"HTTP检测结果"}>
        <ExTable columns={columns} data={httpDetectorResults} fetch={fetch} />
      </NCard>
    );
  },
});
