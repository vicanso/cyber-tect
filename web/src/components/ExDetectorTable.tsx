import { TableColumn } from "naive-ui/lib/data-table/src/interface";
import { newListColumn } from "../components/ExTable";
import { formatDate } from "../helpers/util";

export function getDefaultColumns(): TableColumn[] {
  return [
    {
      title: "名称",
      key: "name",
      width: 100,
    },
    {
      title: "状态",
      key: "status",
      render(row: Record<string, unknown>) {
        return row.statusDesc as string;
      },
      width: 60,
    },
    {
      title: "超时设置",
      key: "timeout",
      width: 80,
    },
    newListColumn({
      key: "owners",
      title: "所有人",
      width: 120,
    }),
    newListColumn({
      key: "receivers",
      title: "告警接收人",
      width: 120,
    }),
    {
      title: "更新于",
      key: "updatedAt",
      render(row: Record<string, unknown>) {
        return formatDate(row.updatedAt as string);
      },
    },
    {
      title: "配置描述",
      key: "description",
      width: 100,
      ellipsis: {
        tooltip: true,
      },
    },
  ];
}
