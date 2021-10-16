import { defineComponent } from "vue";
import { TableColumn } from "naive-ui/lib/data-table/src/interface";

import { FormItem, FormItemTypes } from "../../components/ExFormInterface";

import useDetectorState, {
  databaseDetectorCreate,
  databaseDetectorFindByID,
  databaseDetectorList,
  databaseDetectorUpdateByID,
} from "../../states/detector";

import { newListColumn } from "../../components/ExTable";
import Detector from "./Detector";

export default defineComponent({
  name: "Database",
  setup() {
    const databaseDetectors = useDetectorState().databaseDetectors;
    const findByID = async (id: number) => {
      const result = await databaseDetectorFindByID(id);
      return result as Record<string, unknown>;
    };
    return {
      findByID,
      databaseDetectors,
    };
  },
  render() {
    const { databaseDetectors, findByID } = this;
    const columns: TableColumn[] = [
      newListColumn({
        key: "uris",
        title: "连接串列表",
      }),
      {
        title: "cret pem",
        key: "certPem",
        width: 120,
        ellipsis: {
          tooltip: true,
        },
      },
      {
        title: "key pem",
        key: "keyPem",
        width: 120,
        ellipsis: {
          tooltip: true,
        },
      },
    ];
    const formItems: FormItem[] = [
      {
        type: FormItemTypes.DynamicInput,
        name: "连接串列表：",
        key: "uris",
        span: 24,
        placeholder:
          "请输入对应的数据库连接串，如：reids://user:pwd@host1:port1,host2:port2/",
      },
      {
        type: FormItemTypes.TextArea,
        name: "Cert PEM block：",
        key: "certPem",
        span: 24,
        placeholder:
          "请填写tls证书对应的cert pem数据，若不需要使用tls则无需要配置，若已原有配置需要删除则配置为空格",
      },
      {
        type: FormItemTypes.TextArea,
        name: "Key PEM block:",
        key: "keyPem",
        span: 24,
        placeholder:
          "请填写tls证书对应的key pem数据，若不需要使用tls则无需要配置，若已原有配置需要删除则配置为空格",
      },
    ];
    return (
      <Detector
        columns={columns}
        fetch={databaseDetectorList}
        findByID={findByID}
        updateByID={databaseDetectorUpdateByID}
        create={databaseDetectorCreate}
        title={"数据库检测配置"}
        description={"指定数据库的连接串，定时检测是否正常"}
        formItems={formItems}
        data={databaseDetectors}
      />
    );
  },
});
