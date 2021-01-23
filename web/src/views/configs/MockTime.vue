<template lang="pug">
config-editor(
  name="添加/更新MockTime配置"
  summary="针对应用时间Mock，用于测试环境中调整应用时间"
  :category="category"
  :defaultValue="defaultValue"
  :backDisabled="true"
  v-if="!processing"
  :id="currentID"
  :back="noop"
): template(
  #data="configProps"
): mock-time-data(
  :data="configProps.form.data"
  @change="configProps.form.data = $event"
)
</template>
<script lang="ts">
import { defineComponent } from "vue";

import ConfigEditor from "../../components/configs/Editor.vue";
import MockTimeData from "../../components/configs/MockTimeData.vue";
import { MOCK_TIME } from "../../constants/common";
import { useConfigStore } from "../../store";

export default defineComponent({
  name: "MockTime",
  components: {
    MockTimeData,
    ConfigEditor,
  },
  setup() {
    const configStore = useConfigStore();
    return {
      configs: configStore.state.configs,
      list: (params) => configStore.dispatch("list", params),
    };
  },
  data() {
    return {
      defaultValue: {
        name: MOCK_TIME,
        category: MOCK_TIME,
      },
      processing: true,
      currentID: 0,
      category: MOCK_TIME,
    };
  },
  async mounted() {
    const { $route, $router } = this;
    this.processing = true;
    try {
      const { configurations } = await this.list({
        name: MOCK_TIME,
      });
      if (configurations && configurations.length !== 0) {
        let currentID = null;
        if ($route.query.id) {
          currentID = Number($route.query.id);
        }
        if (currentID !== configurations[0].id) {
          $router.replace({
            query: {
              id: configurations[0].id,
            },
          });
        }
      }
    } catch (err) {
      this.$error(err);
    } finally {
      this.processing = false;
    }
  },
  methods: {
    // eslint-ignore-next-line
    noop() {
      return;
    },
  },
});
</script>
