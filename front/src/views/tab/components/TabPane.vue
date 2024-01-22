<template>
  <el-table :data="list" border fit highlight-current-row style="width: 100%">
    <el-table-column
      v-loading="loading"
      align="center"
      label="ID"
      width="65"
      element-loading-text="请给我点时间！"
    >
      <template slot-scope="scope">
        <span>{{ scope.row.id }}</span>
      </template>
    </el-table-column>

    <el-table-column min-width="300px" label="Title">
      <template slot-scope="{row}">
        <span>{{ row.name }}</span>
      </template>
    </el-table-column>

    <el-table-column width="110px" align="center" label="NameSpace">
      <template slot-scope="scope">
        <span>{{ scope.row.namespaces }}</span>
      </template>
    </el-table-column>

    <el-table-column width="210px" align="center" label="服务地址">
      <template slot-scope="scope">
        <span>{{ scope.row.server }}</span>
      </template>
    </el-table-column>
    <el-table-column width="180px" align="center" label="创建时间">
      <template slot-scope="scope">
        <span>{{ scope.row.ctime }}</span>
      </template>
    </el-table-column>
    <el-table-column width="180px" align="center" label="修改时间">
      <template slot-scope="scope">
        <span>{{ scope.row.mtime }}</span>
      </template>
    </el-table-column>

    <!--    <el-table-column width="120px" label="Importance">-->
    <!--      <template slot-scope="scope">-->
    <!--        <svg-icon v-for="n in +scope.row.importance" :key="n" icon-class="star" />-->
    <!--      </template>-->
    <!--    </el-table-column>-->

    <!--    <el-table-column align="center" label="Readings" width="95">-->
    <!--      <template slot-scope="scope">-->
    <!--        <span>{{ scope.row.pageviews }}</span>-->
    <!--      </template>-->
    <!--    </el-table-column>-->

    <el-table-column class-name="status-col" label="操作" width="110">
      <template slot-scope="{row}">
        <el-tag v-if="row.status === 1" @click="goPodList(row.id)">
          获取pod列表
        </el-tag>
        <span v-else>
          &nbsp;
        </span>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import { clusterList } from '@/api/api'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        1: '获取pod列表',
        2: 'danger'
      }
      return statusMap[status]
    }
  },
  props: {
    type: {
      type: String,
      default: 'CN'
    }
  },
  data() {
    return {
      list: null,
      listQuery: {
        page: 1,
        limit: 5,
        type: this.type
      },
      loading: false
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.loading = true
      clusterList({}).then(response => {
        this.list = response.data.list
        // console.log(this.list)
        this.loading = false
      })
      // this.list = [
      //   {
      //     id:1,
      //     timestamp: 1234567890,
      //     author: '@first',
      //     reviewer: '@first',
      //     title: 'xxxxxx',
      //     content_short: 'mock data',
      //     content: 'xxxxxx',
      //     forecast: 'xxxxxxxx',
      //     importance: '1',
      //     'type': 'CN',
      //     'status': 'published',
      //     display_time: '@datetime',
      //     comment_disabled: true,
      //     pageviews: '300',
      //     platforms: ['a-platform']
      //   },
      // ]
    },
    goPodList(id) {
      this.$router.push(`/pods?id=` + id)
    }
  }
}
</script>

