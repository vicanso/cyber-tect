<template lang="pug">
el-card.baseEditor(
  v-loading="processing"  
)
  template(
    #header
  )
    i(
      v-if="$props.icon"
      :class="$props.icon"
    )
    span {{ $props.title }}
  el-form(
    v-if="inited"
    :label-width="$props.labelWidth"
    ref="baseEditorForm"
    :rules="rules"
    :model="current"
  ): el-row(
    :gutter="15"
  )
    el-col(
      v-for="field in $props.fields"
      :span="field.span || 8"
      :key="field.key"
    ): el-form-item(
      :label="field.label"
      :label-width="field.labelWidth"
      :class="field.itemClass"
      :prop="field.key"
    )
      //- 选择列表
      el-select.select(
        v-if="field.type === 'select'"
        :placeholder="field.placeholder"
        v-model="current[field.key]"
        :multiple="field.multiple || false"
      ): el-option(
        v-for="item in field.options"
        :key="item.key || item.value"
        :label="item.label || item.name"
        :value="item.value"
      )
      //- 带单位选择的输入框
      el-input(
        v-else-if="field.type === 'specsUnit'"
        :placeholder="field.placeholder"
        v-model="current[field.key]"
        :clearable="field.clearable"
      ): el-select.inputSelect(
        #append
        v-model="current[field.selectKey]"
        :placeholder="field.selectPlaceholder"
      ): el-option(
        v-for="item in field.options"
        :key="item.name"
        :label="item.name"
        :value="item.value"
      )
      //- 输入区域textarea
      el-input(
        type="textarea"
        v-else-if="field.type === 'textarea'"
        v-model="current[field.key]"
        :placeholder="field.placeholder"
        :autosize="field.autosize"
      )
      //- 日期选择
      el-date-picker(
        v-else-if="field.type === 'datePicker'"
        v-model="current[field.key]"
        :type="field.pickerType || 'date'"
        :placeholder="field.placeholder"
      )
      //- 输入框
      el-input(
        v-else
        v-model="current[field.key]"
        :clearable="field.clearable"
        :disabled="field.disabled || false"
        :placeholder="field.placeholder"
      )
    //- 提交
    el-col(
      :span="12"
    ): el-form-item: ex-button(
      :onClick="submit"
    ) {{ submitText }}
    //- 返回
    el-col(
      :span="12"
    ): el-form-item: el-button.btn(
      @click="goBack"
    ) 返回

</template>
<script lang="ts">
import { defineComponent } from "vue";

import ExButton from "../ExButton.vue";
import { diff, validateForm, omitNil, getFieldRules } from "../../helpers/util";

export default defineComponent({
  name: "BaseEditor",
  components: {
    ExButton,
  },
  props: {
    icon: {
      type: String,
      default: "",
    },
    title: {
      type: String,
      required: true,
    },
    labelWidth: {
      type: String,
      default: "80px",
    },
    fields: {
      type: Array,
      required: true,
    },
    id: {
      type: Number,
      default: 0,
    },
    findByID: {
      type: Function,
      default: null,
    },
    updateByID: {
      type: Function,
      default: null,
    },
    add: {
      type: Function,
      default: null,
    },
  },
  data() {
    const { id, fields } = this.$props;
    const submitText = id ? "更新" : "添加";
    const current = {};
    fields.forEach((item) => {
      current[item.key] = null;
    });

    return {
      inited: false,
      originData: null,
      processing: false,
      submitText,
      current,
      rules: getFieldRules(fields),
    };
  },
  async beforeMount() {
    const { id, findByID } = this.$props;
    if (!id) {
      this.inited = true;
      return;
    }
    try {
      const data = await findByID(id);
      this.originData = data;
      Object.assign(this.current, data);
    } catch (err) {
      this.$error(err);
    } finally {
      this.inited = true;
    }
  },
  methods: {
    handleUpload(files) {
      this.current.files = files;
    },
    // 添加
    async handleAdd(data): Promise<boolean> {
      const { add } = this.$props;
      const { rules } = this;
      this.processing = true;
      let isSuccess = false;
      try {
        if (rules) {
          await validateForm(this.$refs.baseEditorForm);
        }
        await add(data);
        this.$message.info("已成功添加");
        this.goBack();
        isSuccess = true;
      } catch (err) {
        this.$error(err);
      } finally {
        this.processing = false;
      }
      return isSuccess;
    },
    // 更新
    async handleUpdate(data): Promise<boolean> {
      let isSuccess = false;
      const { id, updateByID } = this.$props;
      const { originData, rules } = this;
      const updateInfo = diff(omitNil(data), originData);
      if (updateInfo.modifiedCount === 0) {
        this.$message.warning("请先修改要更新的信息");
        return isSuccess;
      }

      this.processing = true;
      try {
        if (rules) {
          await validateForm(this.$refs.baseEditorForm);
        }
        await updateByID({
          id,
          data: updateInfo.data,
        });
        this.$message.info("已成功更新");
        this.goBack();
        isSuccess = true;
      } catch (err) {
        this.$error(err);
      } finally {
        this.processing = false;
      }
      return isSuccess;
    },
    // 提交
    async submit(): Promise<boolean> {
      const { current } = this;
      const { id, fields } = this.$props;
      const data = Object.assign({}, current);
      fields.forEach((item) => {
        if (item.dataType === "number") {
          data[item.key] = Number(data[item.key]);
        }
      });
      if (!id) {
        return await this.handleAdd(data);
      }
      return await this.handleUpdate(data);
    },
    goBack() {
      this.$router.back();
    },
  },
});
</script>
<style lang="stylus" scoped>
.baseEditor
  i
    margin-right: 5px
  .select, .btn
    width: 100%
  .inputSelect
    min-width: 60px
</style>
