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
import {
  getDefaultFormRules,
  newRequireRule,
  newListRequireRules,
} from "../../components/ExDetectorEditor";

export default defineComponent({
  name: "HTTP",
  setup() {
    const httpDetectors = useDetectorState().httpDetectors;
    const findByID = async (id: number) => {
      const result = await httpDetectorFindByID(id);
      return result as Record<string, unknown>;
    };
    return {
      findByID,
      httpDetectors,
    };
  },
  render() {
    const { httpDetectors, findByID } = this;
    const columns: TableColumn[] = [
      {
        title: "URL",
        key: "url",
        width: 100,
        ellipsis: {
          tooltip: true,
        },
      },
      newListColumn({
        key: "ips",
        title: "IP列表",
      }),
      {
        title: "检测脚本",
        key: "script",
      },
    ];

    const formItems: FormItem[] = [
      {
        type: FormItemTypes.DynamicInput,
        name: "IP列表：",
        key: "ips",
        span: 12,
        placeholder: "请输入对应的IP地址，若不指定则输入0.0.0.0",
        min: 1,
      },
      {
        name: "检测地址：",
        key: "url",
        span: 12,
        placeholder: "请输入要检测的URL",
      },
      {
        name: "检测脚本：",
        key: "script",
        span: 24,
        placeholder:
          "请输入对应的检测脚本，响应数据为resp，可针其数据检测(如果是更新，如果是需要清空，需要设置为空格)",
        type: FormItemTypes.TextArea,
      },
    ];
    const rules = getDefaultFormRules({
      url: newRequireRule("检测地址不能为空"),
      ips: newListRequireRules("IP地址列表不能为空"),
    });
    return (
      <Detector
        columns={columns}
        fetch={httpDetectorList}
        findByID={findByID}
        updateByID={httpDetectorUpdateByID}
        create={httpDetectorCreate}
        title={"HTTP检测配置"}
        description={"指定HTTP检测URL以及IP列表，定时检测该URL是否可正常访问"}
        formItems={formItems}
        data={httpDetectors}
        rules={rules}
      />
    );
  },
});
