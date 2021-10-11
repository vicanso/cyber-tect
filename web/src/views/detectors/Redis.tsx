import { defineComponent } from "vue";
import { TableColumn } from "naive-ui/lib/data-table/src/interface";

import { FormItem, FormItemTypes } from "../../components/ExFormInterface";

import useDetectorState, {
  redisDetectorCreate,
  redisDetectorFindByID,
  redisDetectorList,
  redisDetectorUpdateByID,
} from "../../states/detector";

import { newListColumn } from "../../components/ExTable";
import Detector from "./Detector";

export default defineComponent({
  name: "Redis",
  setup() {
    const redisDetectors = useDetectorState().redisDetectors;
    const findByID = async (id: number) => {
      const result = await redisDetectorFindByID(id);
      return result as Record<string, unknown>;
    };
    return {
      findByID,
      redisDetectors,
    };
  },
  render() {
    const { redisDetectors, findByID } = this;
    const columns: TableColumn[] = [
      newListColumn({
        key: "uris",
        title: "连接串列表",
      }),
    ];
    const formItems: FormItem[] = [
      {
        type: FormItemTypes.DynamicInput,
        name: "连接串列表：",
        key: "uris",
        span: 24,
        placeholder:
          "请输入对应的redis连接串，如：reids://user:pwd@host1:port1,host2:port2/",
      },
    ];
    return (
      <Detector
        columns={columns}
        fetch={redisDetectorList}
        findByID={findByID}
        updateByID={redisDetectorUpdateByID}
        create={redisDetectorCreate}
        title={"Redis检测配置"}
        description={"指定Redis的连接串，定时调用Ping命令检测是否正常"}
        formItems={formItems}
        data={redisDetectors}
      />
    );
  },
});
