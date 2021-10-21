import { TableColumn } from "naive-ui/lib/data-table/src/interface";
import { defineComponent, PropType } from "vue";
import { LocationQuery } from "vue-router";

import { FormItemTypes } from "./ExFormInterface";
import ExTable from "./ExTable";
import { getFromQuery } from "../helpers/util";

function getFilters(query: LocationQuery) {
  console.dir(getFromQuery(query, "result"));
  return [
    {
      name: "结果：",
      key: "result",
      placeholder: "请选择要筛选的结果",
      type: FormItemTypes.Select,
      defaultValue: getFromQuery(query, "result"),
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
        filters={getFilters(this.$route.query)}
        columns={columns}
        data={data}
        fetch={this.filterFetch}
      />
    );
  },
});
