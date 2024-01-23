<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">
        新增
      </el-button>
    </div>
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

      <el-table-column min-width="300px" label="集群简称">
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
      <el-table-column width="150px" align="center" label="创建时间">
        <template slot-scope="scope">
          <span>{{ scope.row.ctime }}</span>
        </template>
      </el-table-column>
      <el-table-column width="150px" align="center" label="修改时间">
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

      <el-table-column class-name="status-col" label="操作" width="220px">
        <template slot-scope="{row}">
          <el-tag v-if="row.status === 1" @click="goPodList(row.id)">
            获取pod列表
          </el-tag>
          &nbsp;
          <el-tag v-if="row.status === 1" type="info" @click="statusChange(row)">
            禁用
          </el-tag>
          <el-tag v-else type="warning" @click="statusChange(row)">
            启用
          </el-tag>
          &nbsp;
          <el-tag type="success" @click="handleUpdate(row)">
            编辑
          </el-tag>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog title="新增/编辑" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :model="temp" label-position="left" label-width="70px" style="width: 400px; margin-left:50px;">
        <el-form-item label="集群简称" prop="name">
          <el-input v-model="temp.name" />
        </el-form-item>
        <el-form-item label="namespaces" prop="namespaces">
          <el-input v-model="temp.namespaces" />
        </el-form-item>
        <el-form-item label="服务地址" prop="server_api">
          <el-input v-model="temp.server_api" />
        </el-form-item>
        <el-form-item label="config" prop="config">
          <el-input v-model="temp.config" type="textarea" :rows="10" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">
          取消
        </el-button>
        <el-button type="primary" @click="upsert()">
          提交
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { clusterList, clusterStatusChange, clusterUpsert } from '@/api/api'

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
      dialogFormVisible: false,
      list: null,
      listQuery: {
        page: 1,
        limit: 5,
        type: this.type
      },
      temp: {
        id: 0,
        name: '',
        server_api: '',
        namespaces: '',
        config: ''
      },
      loading: false
    }
  },
  created() {
    this.getList()
    this.resetTemp()
  },
  methods: {
    resetTemp() {
      this.temp = {
        id: 0,
        name: '',
        server_api: '',
        namespaces: '',
        config: ''
      }
    },
    handleUpdate(row) {
      this.temp = Object.assign({}, row) // copy obj
      this.temp.timestamp = new Date(this.temp.timestamp)
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    handleCreate() {
      this.resetTemp()
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    getList() {
      this.loading = true
      clusterList({}).then(response => {
        this.list = response.data.list
        this.loading = false
      })
    },
    upsert() {
      clusterUpsert(this.temp).then(response => {
        this.getList()
      })
    },
    statusChange(row) {
      let status = row.status
      if (status === 1) {
        status = 2
      } else {
        status = 1
      }
      clusterStatusChange({ id: row.id, status: status }).then(response => {
        this.getList()
      })
    },
    goPodList(id) {
      this.$router.push(`/pods?id=` + id)
    }
  }
}
</script>

