import { defineComponent } from "vue";
import { TableColumn } from "naive-ui/lib/data-table/src/interface";

import { FormItem, FormItemTypes } from "../../components/ExFormInterface";

import useDetectorState, {
  dnsDetectorCreate,
  dnsDetectorFindByID,
  dnsDetectorList,
  dnsDetectorUpdateByID,
} from "../../states/detector";
import { newListColumn } from "../../components/ExTable";
import Detector from "./Detector";
import {
  getDefaultFormRules,
  newRequireRule,
  newListRequireRules,
} from "../../components/ExDetectorEditor";

export default defineComponent({
  name: "DNSDetector",
  setup() {
    const dnsDetectors = useDetectorState().dnsDetectors;
    const findByID = async (id: number) => {
      const result = await dnsDetectorFindByID(id);
      return result as Record<string, unknown>;
    };
    return {
      findByID,
      dnsDetectors,
    };
  },
  render() {
    const { dnsDetectors, findByID } = this;
    const columns: TableColumn[] = [
      {
        title: "域名",
        key: "host",
      },
      newListColumn({
        key: "ips",
        title: "IP列表",
      }),
      newListColumn({
        key: "servers",
        title: "DNS服务器",
      }),
    ];
    const formItems: FormItem[] = [
      {
        name: "域名",
        key: "host",
        placeholder: "请输入要检测的域名",
        span: 8,
      },
      {
        type: FormItemTypes.DynamicInput,
        name: "IP列表：",
        key: "ips",
        span: 8,
        placeholder: "请输入对应的IP解析",
        min: 1,
      },
      {
        type: FormItemTypes.DynamicInput,
        name: "DNS服务器：",
        key: "servers",
        span: 8,
        placeholder: "请输入DNS服务器",
        min: 1,
      },
    ];
    const rules = getDefaultFormRules({
      host: newRequireRule("域名不能为空"),
      servers: newListRequireRules("DNS服务器列表不能为空"),
      ips: newListRequireRules("域名解析IP地址列表不能为空"),
    });
    return (
      <Detector
        columns={columns}
        fetch={dnsDetectorList}
        findByID={findByID}
        updateByID={dnsDetectorUpdateByID}
        create={dnsDetectorCreate}
        title={"DNS检测"}
        description={"指定DNS服务器检测解析IP是否正确"}
        formItems={formItems}
        data={dnsDetectors}
        rules={rules}
      />
    );
  },
});
