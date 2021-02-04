<template lang="pug">
.dnsResults: el-card
  template(
    #header
  )
    span DNS检测结果列表
  
</template>
<script lang="ts">
import { defineComponent } from "vue";

import { useDetectorResultStore } from "../../store";
import BaseFilter from "../../components/base/Filter.vue";
import BaseTooltip from "../../components/Tooltip.vue";
import TimeFormater from "../../components/TimeFormater.vue";
import ResultStatus from "./Status.vue";
import { PAGE_SIZES } from "../../constants/common";
import FilterTable from "../../mixins/FilterTable";
</script>
