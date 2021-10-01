import { defineComponent } from "vue";
import { css } from "@linaria/core";
import { EyeRegular } from "@vicons/fa";
import { NCard, NDataTable, NPopover, NIcon } from "naive-ui";
import { RowData, TableColumn } from "naive-ui/lib/data-table/src/interface";

import useDetectorState, { tcpDetectorResultList } from "../../states/detector";
import ExTable, { newListColumn } from "../../components/ExTable";
import { formatDate } from "../../helpers/util";

const popupClass = css`
  max-width: 980px;
  white-space: nowrap;
`;

export default defineComponent({
  name: "TCPResult",
  setup() {
    const fetch = async (params: Record<string, unknown>) => {
      await tcpDetectorResultList(params);
    };
    return {
      tcpDetectorResults: useDetectorState().tcpDetectorResults,
      fetch,
    };
  },
  render() {
    const { tcpDetectorResults, fetch } = this;
    const columns: TableColumn[] = [
      {
        title: "名称",
        key: "taskName",
      },
      {
        title: "结果",
        key: "resultDesc",
      },
      newListColumn({
        title: "检测地址",
        key: "addrs",
      }),
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
              title: "结果",
              key: "resultDesc",
            },
            {
              title: "耗时(ms)",
              key: "duration",
            },
            {
              title: "失败信息",
              key: "message",
            },
            // newListColumn({
            //   key: "timeline",
            //   title: "耗时",
            //   width: 180,
            // }),
            // {
            //   title: "HTTP协议",
            //   key: "protocol",
            // },
            // {
            //   title: "TLS",
            //   key: "tlsVersion",
            // },
            // {
            //   title: "TLS加密",
            //   key: "tlsCipherSuite",
            //   ellipsis: {
            //     tooltip: true,
            //   },
            //   width: 100,
            // },
            // newListColumn({
            //   key: "certificateDNSNames",
            //   title: "证书域名",
            // }),
            // newListColumn({
            //   key: "certificateExpirationDates",
            //   title: "证书有效期",
            // }),
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
      <NCard title={"TCP检测结果"}>
        <ExTable columns={columns} data={tcpDetectorResults} fetch={fetch} />
      </NCard>
    );
  },
});
