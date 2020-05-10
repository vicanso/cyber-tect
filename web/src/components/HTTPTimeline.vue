<template lang="pug">
  .httpTimeline.clearfix
    el-tooltip(
      placement="right"
    )
      div(
        slot="content"
      )
        div(
          v-html="tips"
        )
      .timelineWrapper
        .timeline(
          v-for="duration in durations"
          :style="duration.style"
        )
    .durationDesc {{durationDesc}}
</template>

<script>
import {
  formatDuration
} from '@/helpers/util'

export default {
  name: 'HTTPTimeline',
  props: {
    dnsLookup: {
      type: Number,
      default: 0
    },
    tcpConnection: {
      type: Number,
      default: 0
    },
    tlsHandshake: {
      type: Number,
      default: 0
    },
    serverProcessing: {
      type: Number,
      default: 0
    },
    contentTransfer: {
      type: Number,
      default: 0
    },
    duration: {
      type: Number,
      default: 0
    }
  },
  data () {
    const {
      dnsLookup,
      tcpConnection,
      tlsHandshake,
      serverProcessing,
      contentTransfer,
      duration
    } = this.$props
    const durations = [
      {
        name: 'dns lookup',
        duration: dnsLookup,
        color: '#419488'
      },
      {
        name: 'tcp connection',
        duration: tcpConnection,
        color: '#f29c38'
      },
      {
        name: 'tls handshake',
        duration: tlsHandshake,
        color: '#9035aa'
      },
      {
        name: 'server processing',
        duration: serverProcessing,
        color: '#5ac462'
      },
      {
        name: 'content transfer',
        duration: contentTransfer,
        color: '#3d7cab'
      }
    ]
    const tips = []
    durations.forEach(item => {
      tips.push(`${item.name}：${item.duration} 毫秒`)
    })
    let percentCount = 0
    durations.forEach((item, index) => {
      let percent = Math.round(100 * item.duration / duration)
      if (index === durations.length - 1) {
        percent = 100 - percentCount
      }
      item.style = {
        width: `${percent}px`,
        backgroundColor: item.color
      }
      percentCount += percent
    })
    return {
      tips: tips.join('<br />'),
      durations,
      durationDesc: formatDuration(duration)
    }
  }
}
</script>
<style lang="sass" scoped>
@import "@/common.sass"
.timelineWrapper
  border: 1px solid $darkGray
  float: left
  margin-top: 6px
  margin-right: 3px
  min-width: 100px
  .timeline
    float: left
    height: 10px
.durationDesc
  float: left
</style>
