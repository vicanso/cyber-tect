import { TableColumn } from "naive-ui/lib/data-table/src/interface";
import { defineComponent, PropType } from "vue";
import { FormItemTypes } from "./ExFormInterface";
import ExTable from "./ExTable";

function getFilters() {
  return [
    {
      name: "结果：",
      key: "result",
      placeholder: "请选择要筛选的结果",
      type: FormItemTypes.Select,
      options: [
        {
          label: "所有",
          value: "",
        },
        {
          label: "成功",
          value: "1",
        },
        {
          label: "失败",
          value: "2",
        },
      ],
      span: 8,
    },
    {
      name: "耗时：",
      key: "duration",
      placeholder: "请输入过滤时长（大于等于）",
      type: FormItemTypes.InputDuration,
      span: 8,
    },
  ];
}

export default defineComponent({
  name: "ExDetectorResultTable",
  props: {
    columns: {
      type: Array as PropType<TableColumn[]>,
      required: true,
    },
    data: {
      type: Object,
      required: true,
    },
    fetch: {
      type: Function,
      required: true,
    },
  },
  setup(props) {
    const filterFetch = (params: Record<string, unknown>) => {
      return props.fetch(params);
    };
    return {
      filterFetch,
    };
  },
  render() {
    const { columns, data } = this.$props;
    return (
      <ExTable
        filters={getFilters()}
        columns={columns}
        data={data}
        fetch={this.filterFetch}
      />
    );
  },
});
