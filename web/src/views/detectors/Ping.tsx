import { defineComponent } from "vue";
import { TableColumn } from "naive-ui/lib/data-table/src/interface";

import { FormItem, FormItemTypes } from "../../components/ExFormInterface";
import useDetectorState, {
  pingDetectorCreate,
  pingDetectorFindByID,
  pingDetectorList,
  pingDetectorUpdateByID,
} from "../../states/detector";
import { newListColumn } from "../../components/ExTable";
import Detector from "./Detector";
import {
  getDefaultFormRules,
  newListRequireRules,
} from "../../components/ExDetectorEditor";

export default defineComponent({
  name: "PingDetector",
  setup() {
    const pingDetectors = useDetectorState().pingDetectors;
    const findByID = async (id: number) => {
      const result = await pingDetectorFindByID(id);
      return result as Record<string, unknown>;
    };
    return {
      findByID,
      pingDetectors,
    };
  },
  render() {
    const { pingDetectors, findByID } = this;
    const columns: TableColumn[] = [
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
    ];
    const rules = getDefaultFormRules({
      ips: newListRequireRules("IP地址列表不能为空"),
    });
    return (
      <Detector
        columns={columns}
        fetch={pingDetectorList}
        findByID={findByID}
        updateByID={pingDetectorUpdateByID}
        create={pingDetectorCreate}
        title={"Ping检测配置"}
        description={"指定Ping检测IP列表，定时Ping该IP是否正常"}
        formItems={formItems}
        data={pingDetectors}
        rules={rules}
      />
    );
  },
});
