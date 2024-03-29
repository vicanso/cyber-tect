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
  name: "HTTPDetector",
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
        width: 250,
        ellipsis: {
          tooltip: true,
        },
      },
      newListColumn({
        key: "ips",
        title: "IP列表",
      }),
      newListColumn({
        key: "proxies",
        title: "代理地址",
      }),
      {
        title: "随机QueryString",
        key: "randomQueryString",
        render: (row: Record<string, unknown>) => {
          if (row.randomQueryString === 1) {
            return '是'
          }
          return '否'
        }
      },
      {
        title: "检测脚本",
        key: "script",
        width: 300,
        render: (row: Record<string, unknown>) => {
          return (
            <pre
              style={{
                whiteSpace: "break-spaces",
              }}
            >
              {row["script"]}
            </pre>
          );
        },
      },
    ];

    const formItems: FormItem[] = [
      {
        type: FormItemTypes.DynamicInput,
        name: "IP列表：",
        key: "ips",
        span: 6,
        placeholder: "请输入对应的IP地址，若不指定则输入0.0.0.0",
        min: 1,
      },
      {
        name: "检测地址：",
        key: "url",
        span: 6,
        placeholder: "请输入要检测的URL",
      },
      {
        type: FormItemTypes.Select,
        name: "随机QueryString：",
        key: "randomQueryString",
        span: 6,
        placeholder: "请选择是否使用",
        options: [
          {
            label: "是",
            value: 1,
          },
          {
            label: "否",
            value: 2,
          },
        ],
      },
      {
        type: FormItemTypes.DynamicInput,
        name: "代理地址：",
        key: "proxies",
        span: 6,
        placeholder:
          "请输入使用的代理地址，如：http://127.0.0.1:52206，http://0.0.0.0 表示不使用代理",
      },
      {
        name: "检测脚本：",
        key: "script",
        span: 24,
        placeholder:
          "请输入对应的检测脚本，响应数据为resp，可针其数据检测(如果是更新，如果是需要清空，需要设置为空格)，throw new Error('出错信息')即为检测失败",
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
        tableScrollX={2200}
      />
    );
  },
});
