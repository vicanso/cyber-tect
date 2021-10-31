import { defineComponent } from "vue";
import { TableColumn } from "naive-ui/lib/data-table/src/interface";

import { FormItem, FormItemTypes } from "../../components/ExFormInterface";

import useDetectorState, {
  tcpDetectorCreate,
  tcpDetectorFindByID,
  tcpDetectorList,
  tcpDetectorUpdateByID,
} from "../../states/detector";
import { newListColumn } from "../../components/ExTable";
import Detector from "./Detector";
import {
  getDefaultFormRules,
  newListRequireRules,
} from "../../components/ExDetectorEditor";

export default defineComponent({
  name: "TCP",
  setup() {
    const tcpDetectors = useDetectorState().tcpDetectors;
    const findByID = async (id: number) => {
      const result = await tcpDetectorFindByID(id);
      return result as Record<string, unknown>;
    };
    return {
      findByID,
      tcpDetectors,
    };
  },
  render() {
    const { tcpDetectors, findByID } = this;
    const columns: TableColumn[] = [
      newListColumn({
        key: "addrs",
        title: "检测地址",
      }),
    ];
    const formItems: FormItem[] = [
      {
        type: FormItemTypes.DynamicInput,
        name: "地址列表：",
        key: "addrs",
        span: 12,
        placeholder: "请输入需要检测的地址，如(IP:Port)",
        min: 1,
      },
    ];
    const rules = getDefaultFormRules({
      addrs: newListRequireRules("检测地址列表不能为空"),
    });

    return (
      <Detector
        columns={columns}
        fetch={tcpDetectorList}
        findByID={findByID}
        updateByID={tcpDetectorUpdateByID}
        create={tcpDetectorCreate}
        title={"TCP检测配置"}
        description={"指定IP与端口，定时检测是否可用"}
        formItems={formItems}
        data={tcpDetectors}
        rules={rules}
      />
    );
  },
});
