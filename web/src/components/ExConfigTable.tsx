import { defineComponent, onUnmounted, PropType } from "vue";
import { NButton, NIcon } from "naive-ui";
import { EditRegular } from "@vicons/fa";
import { TableColumn } from "naive-ui/lib/data-table/src/interface";
import ExTable from "../components/ExTable";
import { formatDate } from "../helpers/util";
import useConfigState, {
  configList,
  configListClear,
  ConfigStatus,
} from "../states/configs";

function isJSON(str: string): boolean {
  if (str.length < 2) {
    return false;
  }
  const firstLetter = str[0];
  const lastLetter = str[str.length - 1];
  // 示判断{]的场景
  if (firstLetter !== "{" && firstLetter !== "[") {
    return false;
  }
  if (lastLetter !== "}" && lastLetter !== "]") {
    return false;
  }
  return true;
}

function getColumns(): TableColumn[] {
  return [
    {
      title: "名称",
      key: "name",
    },
    {
      type: "expand",
      expandable: () => true,
      renderExpand: (data: Record<string, unknown>) => {
        const str = data.data as string;
        if (isJSON(str.trim())) {
          return <pre>{JSON.stringify(JSON.parse(str), null, 2)}</pre>;
        }
        return str;
      },
    },
    {
      title: "分类",
      key: "category",
    },
    {
      title: "状态",
      key: "status",
      render(row: Record<string, unknown>) {
        if (row.status === ConfigStatus.Enabled) {
          return "启用";
        }
        return "禁用";
      },
    },
    {
      title: "创建者",
      key: "owner",
    },
    {
      title: "配置生效时间",
      key: "startedAt",
      render(row: Record<string, unknown>) {
        return formatDate(row.startedAt as string);
      },
    },
    {
      title: "配置失效时间",
      key: "endedAt",
      render(row: Record<string, unknown>) {
        return formatDate(row.endedAt as string);
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

function noop(id: number): void {
  // 无操作
}

export default defineComponent({
  name: "Configs",
  props: {
    title: {
      type: String,
      default: "",
    },
    category: {
      type: String,
      default: () => "",
    },
    onUpdate: {
      type: Function as PropType<
      (id: number) => void
      >,
      default: noop,
    },
  },
  setup(props) {
    const { configs } = useConfigState();

    const fetchConfigs = () =>
      configList({
        category: props.category,
      });

    onUnmounted(() => {
      configListClear();
    });
    return {
      fetchConfigs,
      configs,
    };
  },
  render() {
    const { title, onUpdate } = this.$props;
    const { configs, fetchConfigs, $slots } = this;
    const columns = getColumns();
    if (onUpdate !== noop) {
      columns.push({
        title: "操作",
        key: "op",
        render(row: Record<string, unknown>) {
          return (
            <NButton
              bordered={false}
              onClick={() => {
                onUpdate(row.id as number);
              }}
            >
              <NIcon>
                <EditRegular />
              </NIcon>
              更新
            </NButton>
          );
        },
      });
    }
    return (
      <ExTable
        title={title}
        columns={columns}
        data={configs}
        fetch={fetchConfigs}
      >
        {$slots}
      </ExTable>
    );
  },
});
