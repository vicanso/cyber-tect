import { CONFIG_EDIT_MODE } from "../constants/common";
import { diff } from "../helpers/util";

export default {
  computed: {
    currentPage(): number {
      const { offset, limit } = this.query;
      return Math.floor(offset / limit) + 1;
    },
    editMode(): boolean {
      return this.$route.query.mode === CONFIG_EDIT_MODE;
    },
  },
  beforeMount(): void {
    this._currentRoute = this.$route.name;
    if (!this.editMode && !this.disableBeforeMountFetch) {
      this.fetch();
    }
  },
  methods: {
    handleCurrentChange(page: number): void {
      this.query.offset = (page - 1) * this.query.limit;
      this.fetch();
    },
    handleSizeChange(pageSize: number): void {
      this.query.limit = pageSize;
      this.query.offset = 0;
      this.fetch();
    },
    handleSortChange(params: { prop: string; order: string }): void {
      let key = params.prop.replace("Desc", "");
      if (params.order === "descending") {
        key = `-${key}`;
      }
      this.query.order = key;
      this.query.offset = 0;
      this.fetch();
    },
    add(): void {
      this.$router.push({
        query: {
          mode: CONFIG_EDIT_MODE,
        },
      });
    },
    // eslint-disable-next-line
    modify(item: any): void {
      this.$router.push({
        query: {
          mode: CONFIG_EDIT_MODE,
          id: item.id,
        },
      });
    },
    // eslint-disable-next-line
    filter(params: any): void {
      Object.assign(this.query, params);
      this.query.offset = 0;
      this.fetch();
    },
  },
  watch: {
    // eslint-disable-next-line
    "$route.query"(query: any, prevQuery: any): void {
      // 如果路由已更换，则直接返回
      if (this.$route.name !== this._currentRoute) {
        return;
      }
      if (!diff(query, prevQuery).modifiedCount) {
        return;
      }

      if (!this.editMode) {
        this.fetch();
      }
    },
  },
};
