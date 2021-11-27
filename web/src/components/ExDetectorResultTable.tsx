import { TableColumn } from "naive-ui/lib/data-table/src/interface";
import { defineComponent, onMounted, PropType } from "vue";
import { useMessage } from "naive-ui";
import { LocationQuery } from "vue-router";

import { FormItemTypes } from "./ExFormInterface";
import ExTable from "./ExTable";
import { getFromQuery, showError } from "../helpers/util";
import useDetectorState, { listTask } from "../states/detector";

function getFilters(query: LocationQuery) {
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
      name: "检测任务：",
      key: "filterTasks",
      placeholder: "请选择要筛选的任务",
      type: FormItemTypes.MultiSelect,
      options: [] as {
        label: string;
        value: string;
      }[],
      span: 8,
    },
    {
      name: "耗时：",
      key: "duration",
      placeholder: "请输入过滤时长（大于等于）",
      type: FormItemTypes.InputDuration,
      span: 8,
    },
    {
      key: "startedAt:endedAt",
      name: "开始结束时间：",
      placeholder: "请选择开始时间:请选择结束时间",
      type: FormItemTypes.DateRange,
      span: 18,
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
    category: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    const filterFetch = (params: Record<string, unknown>) => {
      return props.fetch(params);
    };
    const detectorTasks = useDetectorState().detectorTasks;
    onMounted(async () => {
      try {
        await listTask(props.category);
      } catch (err) {
        const message = useMessage();
        showError(message, err);
      }
    });
    return {
      detectorTasks,
      filterFetch,
    };
  },
  render() {
    const { detectorTasks } = this;
    const { columns, data } = this.$props;
    const filters = getFilters(this.$route.query);
    if (!detectorTasks.processing) {
      detectorTasks.items.forEach((item) => {
        const options = filters[1].options;
        if (!options) {
          return;
        }
        options.push({
          label: `${item.name}(${item.id})`,
          value: `${item.id}`,
        });
      });
    }
    return (
      <ExTable
        filters={filters}
        columns={columns}
        data={data}
        fetch={this.filterFetch}
      />
    );
  },
});
