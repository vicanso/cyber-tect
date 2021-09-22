import { NSelect } from "naive-ui";
import {
  Value,
  SelectMixedOption,
  Size,
} from "naive-ui/lib/select/src/interface";
import { defineComponent, PropType, ref } from "vue";

import { detectorListUser } from "../states/detector";
import { FormItem } from "./ExFormInterface";

export default defineComponent({
  name: "UserSelect",
  props: {
    formItem: {
      type: Object as PropType<FormItem>,
      required: true,
    },
    size: {
      type: String as PropType<Size>,
      required: true,
    },
    multiple: {
      type: Boolean,
      default: true,
    },
    onUpdateValue: {
      type: Function as PropType<(users: string[]) => void>,
      required: true,
    },
  },
  setup() {
    const options = ref([] as SelectMixedOption[]);
    const loading = ref(false);
    const onSearch = async (keyword: string) => {
      loading.value = true;
      try {
        const accounts = await detectorListUser(keyword);
        options.value = accounts.map((account) => {
          return {
            label: account,
            value: account,
          };
        });
      } finally {
        loading.value = false;
      }
    };
    return {
      options,
      loading,
      onSearch,
    };
  },
  render() {
    const { formItem, size, multiple } = this.$props;
    const { loading, options, onSearch, onUpdateValue } = this;
    return (
      <NSelect
        multiple={multiple}
        defaultValue={formItem.defaultValue as Value}
        size={size}
        filterable
        clearable
        options={options}
        placeholder={formItem.placeholder}
        loading={loading}
        onUpdateValue={onUpdateValue}
        remote
        onSearch={onSearch}
      />
    );
  },
});
