import { defineComponent } from "vue";
import { TableColumn } from "naive-ui/lib/data-table/src/interface";

import { FormItem, FormItemTypes } from "../../components/ExFormInterface";
import useDetectorState, {
  httpDetectorCreate,
  httpDetectorFindByID,
  httpDetectorList,
  httpDetectorUpdateByID,
} from "../../states/detector";
import { newListColumn } from "../../components/ExTable";
import Detector from "./Detector";

export default defineComponent({
  name: "HTTP",
  setup() {
    const httpDetectors = useDetectorState().httpDetectors;
    const findByID = async (id: number) => {
      const result = await httpDetectorFindByID(id);
      return result as Record<string, unknown>;
    };
    return {
      fetch: httpDetectorList,
      findByID,
      httpDetectors,
      updateByID: httpDetectorUpdateByID,
      create: httpDetectorCreate,
    };
  },
  render() {
    const { httpDetectors, fetch, findByID, updateByID, create } = this;
    const columns: TableColumn[] = [
      {
        title: "URL",
        key: "url",
      },
      newListColumn({
        key: "ips",
        title: "IP列表",
      }),
    ];

    const formItems: FormItem[] = [
      {
        type: FormItemTypes.DynamicInput,
        name: "IP列表：",
        key: "ips",
        span: 12,
        placeholder: "请输入对应的IP地址",
        min: 1,
      },
      {
        name: "检测地址：",
        key: "url",
        span: 12,
        placeholder: "请输入要检测的URL",
      },
    ];
    return (
      <Detector
        columns={columns}
        fetch={fetch}
        findByID={findByID}
        updateByID={updateByID}
        create={create}
        title={"HTTP检测配置"}
        description={"指定HTTP检测URL以及IP列表，定时检测该URL是否可正常访问"}
        formItems={formItems}
        data={httpDetectors}
      />
    );
  },
});