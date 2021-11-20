import { defineComponent } from "vue";
import { css } from "@linaria/core";
import { NButton, NPopover } from "naive-ui";
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
import {
  getDefaultFormRules,
  newListRequireRules,
} from "../../components/ExDetectorEditor";

const connectionURIListClass = css`
  margin: 0;
  padding: 0;
  list-style: inside;
  line-height: 2.5em;
`;

export default defineComponent({
  name: "DatabaseDetectors",
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
          "请输入对应的数据库连接串，如：reids://user:pwd@host1:port1,host2:port2/, postgres://user:pwd@host1:port1,host2:port2/mydb",
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
    const slots = {
      trigger: () => (
        <NButton class="mbottom15">查看数据库连接串配置示例</NButton>
      ),
    };

    const descriptionDetail = (
      <NPopover v-slots={slots} placement="bottom-end" trigger="click">
        <ul class={connectionURIListClass}>
          <li>redis://[:pass@]host1,port1,host2:port2,host3:port3/</li>
          <li>
            postgres://[jack:secret@]foo.example.com:5432[,...bar.example.com:5432]/mydb
          </li>
          <li>
            mysql://[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
          </li>
          <li>
            mongodb://[username:password@]host1[:port1][,...hostN[:portN]][/[defaultauthdb][?options]]
          </li>
        </ul>
      </NPopover>
    );
    const rules = getDefaultFormRules({
      uris: newListRequireRules("连接串列表不能为空"),
    });
    return (
      <Detector
        columns={columns}
        fetch={databaseDetectorList}
        findByID={findByID}
        updateByID={databaseDetectorUpdateByID}
        create={databaseDetectorCreate}
        title={"数据库检测配置"}
        description={
          "指定数据库的连接串，定时检测是否正常，支持redis、postgres、mysql以及mongodb"
        }
        descriptionDetail={descriptionDetail}
        formItems={formItems}
        data={databaseDetectors}
        rules={rules}
      />
    );
  },
});
