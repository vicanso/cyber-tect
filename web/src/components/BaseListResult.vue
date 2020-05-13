<script>
import { mapActions } from 'vuex'

export default {
  name: 'BaseListResult',
  methods: {
    ...mapActions([
      'listDetectorResult'
    ]),
    handleSizeChange (pageSize) {
      this.query.limit = pageSize
      this.currentPage = 1
      this.fetch()
    },
    handleCurrentChange (page) {
      this.currentPage = page
      this.fetch()
    },
    filter (params) {
      Object.keys(params).forEach((key) => {
        this.query[key] = params[key]
      })
      this.currentPage = 1
      this.fetch()
    },
    async fetch () {
      const {
        query,
        category,
        currentPage
      } = this
      try {
        await this.listDetectorResult({
          category,
          params: Object.assign({
            offset: (currentPage - 1) * query.limit
          }, query)
        })
      } catch (err) {
        this.$message.error(err.message)
      }
    }
  }
}
</script>
