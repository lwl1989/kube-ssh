<template>
  <div class="app-container">
    <el-table
      :key="tableKey"
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;"
      :span-method="reduceCell"
      :row-class-name="rowIndexStyle"
    >
      <el-table-column label="nodeName">
        <template slot-scope="{row}">
          <span>
            {{ row.nodeName }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label="podIp">
        <template slot-scope="{row}">
          <span>
            {{ row.podIp }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label="namespace">
        <template slot-scope="{row}">
          <span>
            {{ row.metadata_namespace }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label="容器信息">
        <el-table-column label="名称">
          <template slot-scope="{row}">
            <span style="font-weight: bold;">
              {{ row.name }}
            </span>
          </template>
        </el-table-column>
        <el-table-column label="镜像">
          <template slot-scope="{row}">
            <span>
              {{ row.container_image }}
            </span>
          </template>
        </el-table-column>
        <el-table-column label="重启次数">
          <template slot-scope="{row}">
            <span>
              {{ row.container_restartCount }}
            </span>
          </template>
        </el-table-column>
        <el-table-column label="状态" class-name="status-col">
          <template slot-scope="{row}">
            <span>
              {{ row | statusContainerStatusesFilter }}
            </span>
          </template>
        </el-table-column>
        <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
          <template slot-scope="{row}">
            <el-button v-if="row.container_ready" size="mini" type="success" @click="goTerminal(row)">
              进入命令行
            </el-button>
          </template>
        </el-table-column>
      </el-table-column>
      <el-table-column label="创建时间" align="center">
        <template slot-scope="{row}">
          <span>{{ row.startTime }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { podList, signCluster } from '@/api/api'
import waves from '@/directive/waves' // waves directive
import { parseTime } from '@/utils'

export default {
  name: 'PodListTable',
  directives: { waves },
  filters: {
    statusContainerStatusesFilter(st) {
      if (st.container_ready && st.container_started) {
        return 'running'
      }
      return 'other'
    }
  },
  data() {
    return {
      tableKey: 0,
      list: null,
      listLoading: true,
      listQuery: {
        id: this.$route.query.id
      }
    }
  },
  created() {
    if (this.$route.query.id === undefined || parseInt(this.$route.query.id) === 0) {
      this.$router.push(`/dashboard?tab=clusters`)
      return
    }
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      podList(this.listQuery).then(response => {
        this.list = this.reconstructionData(response.data.items)
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },
    goTerminal(row) {
      this.listLoading = true
      signCluster({
        id: parseInt(this.$route.query.id),
        pod: row.metadata_name,
        container: row.name,
        namespace: row.metadata_namespace
      }).then(response => {
        window.open('/terminal/?token=' + response.data.token)
        this.listLoading = false
      }).catch(() => {
        this.listLoading = false
      })
    },
    formatJson(filterVal) {
      return this.list.map(v => filterVal.map(j => {
        if (j === 'timestamp') {
          return parseTime(v[j])
        } else {
          return v[j]
        }
      }))
    },
    reconstructionData(data) {
      const tmpData = []
      data.forEach((item, i) => {
        let metadata = {}
        if (Object.prototype.hasOwnProperty.call(item, 'metadata')) {
          Object.keys(item.metadata).forEach((key) => {
            const tKey = 'metadata_' + key
            metadata[tKey] = item.metadata[key]
          })
        } else {
          metadata = {
            'metadata_name': '',
            'metadata_namespace': ''
          }
        }
        let base = { rowSpan: 1, rowIndex: 1, podIp: item.status.podIP, startTime: item.metadata.creationTimestamp, nodeName: item.spec.nodeName }
        base = Object.assign(base, metadata)
        if (Array.isArray(item.spec.containers) && item.spec.containers.length > 0) {
          item.spec.containers.forEach((sub, j) => {
            let subData = Object.assign({}, sub, item)
            const rowSpan = item.spec.containers.length
            if (j === 0) {
              base.rowSpan = rowSpan
              base.rowIndex = 0
              subData = Object.assign(subData, base)
            } else {
              base.rowSpan = rowSpan
              base.rowIndex = 1
              subData = Object.assign(subData, base)
            }
            if (Object.prototype.hasOwnProperty.call(item.status, 'containerStatuses') && Array.isArray(item.status.containerStatuses) && item.status.containerStatuses.length > 0) {
              item.status.containerStatuses.forEach((container, k) => {
                if (container.name === subData.name) {
                  const containerData = {}
                  Object.keys(container).forEach((key) => {
                    const tKey = 'container_' + key
                    containerData[tKey] = container[key]
                  })
                  subData = Object.assign(subData, containerData)
                }
              })
            }
            tmpData.push(subData)
          })
        }
        // else {
        //   tmpData.push(
        //     Object.assign(base, item)
        //   )
        // }
      })
      return tmpData
    },
    rowIndexStyle(row) {
      if (row.rowIndex === 0) {
        return ''
      }
      if (row.row.rowIndex === 0) {
        return 'row_index0_solid'
      }
    },
    reduceCell({ row, column, rowIndex, columnIndex }) {
      if (columnIndex < 4 || columnIndex === 9) {
        if (row.rowIndex === 1) {
          return {
            rowspan: 0,
            colspan: 0
          }
        }
        return {
          rowspan: row.rowSpan,
          colspan: 1
        }
      }
      return {
        rowspan: 1,
        colspan: 1
      }
    }
  }
}
</script>

<style>
.row_index0_solid td{
  border-top: 2px solid;
}
</style>
