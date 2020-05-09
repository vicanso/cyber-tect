import request from '@/request'
import {
  DETECTORS,
  DETECTOR
} from '@/constants/url'
import {
  formatDate,
  delay
} from '@/helpers/util'

const mutationDetectorProcessing = 'detector.processing'
const mutationDetectorReset = 'detector.reset'
const mutationDetectorList = 'detector.list'
const mutationDetectorChangeCurrent = 'detector.changeCurrent'
const mutationDetectorUpdate = 'detector.update'

const statusDescList = ['未知', '启用', '禁用']

const state = {
  processing: false,
  currentCategory: '',
  currentDetectors: null,
  updateDetector: null,
  http: {
    count: -1,
    detectors: null
  },
  dns: {
    count: -1,
    detectors: null
  },
  ping: {
    count: -1,
    detectors: null
  },
  tcp: {
    count: -1,
    detectors: null
  }
}
export default {
  state,
  mutations: {
    [mutationDetectorProcessing] (state, value) {
      state.processing = value
    },
    [mutationDetectorReset] (state, category) {
      state[category].count = -1
      state[category].detectors = null
    },
    [mutationDetectorList] (state, { category, data }) {
      state.currentCategory = category
      const arr = (state[category].detectors || []).slice(0)
      if (data.count >= 0) {
        state[category].count = data.count
      }
      state[category].detectors = arr.concat(data.detectors || [])
    },
    [mutationDetectorChangeCurrent] (state, { category, limit, offset }) {
      state.currentCategory = category
      if (!limit) {
        state.currentDetectors = null
      } else {
        state.currentDetectors = state[category].detectors.slice(offset, offset + limit)
      }
    },
    [mutationDetectorUpdate] (state, data) {
      state.updateDetector = data
    }
  },
  actions: {
    async addDetector ({ commit }, { category, detector }) {
      commit(mutationDetectorProcessing, true)
      try {
        await request.post(DETECTORS.replace(':category', category), detector)
      } finally {
        commit(mutationDetectorProcessing, false)
      }
    },
    async listDetector ({ commit }, { category, params }) {
      commit(mutationDetectorProcessing, true)
      // 如果没有指定offset或者为0，则重置
      if (params && !params.offset) {
        commit(mutationDetectorReset, category)
        commit(mutationDetectorChangeCurrent, {
          category,
          limit: 0
        })
      }
      try {
        const {
          data
        } = await request.get(DETECTORS.replace(':category', category), {
          params
        })
        if (!data.count) {
          data.count = 0
        }
        if (data.detectors) {
          data.detectors.forEach((item) => {
            item.statusDesc = statusDescList[item.status]
            item.updatedAtDesc = formatDate(item.updatedAt)
          })
        }
        // 保存完整的列表，暂时未使用
        commit(mutationDetectorList, {
          category,
          data
        })
        commit(mutationDetectorChangeCurrent, {
          category,
          limit: params.limit,
          offset: params.offset
        })
      } finally {
        commit(mutationDetectorProcessing, false)
      }
    },
    async updateDetector ({ commit }, { id, category, data }) {
      const url = DETECTOR.replace(':category', category).replace(':id', id)
      commit(mutationDetectorProcessing, true)
      try {
        await request.patch(url, data)
      } finally {
        commit(mutationDetectorProcessing, false)
      }
    },
    async clearUpdateDetector ({ commit }) {
      commit(mutationDetectorUpdate, null)
    },
    async getUpdateDetector ({ commit }, { category, id }) {
      commit(mutationDetectorUpdate, null)
      commit(mutationDetectorProcessing, true)
      // 从缓存中获取
      if (state.currentCategory === category && state.currentDetectors) {
        let found = null
        state.currentDetectors.forEach((item) => {
          if (item.id === id) {
            found = item
          }
        })
        if (found) {
          await delay(30)
          commit(mutationDetectorUpdate, found)
          commit(mutationDetectorProcessing, false)
          return
        }
      }

      const url = DETECTOR.replace(':category', category).replace(':id', id)
      try {
        const {
          data
        } = await request.get(url)
        commit(mutationDetectorUpdate, data)
      } finally {
        commit(mutationDetectorProcessing, false)
      }
    }
  }
}
