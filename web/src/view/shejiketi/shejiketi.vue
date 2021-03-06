<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="课题名称">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>
        <el-form-item label="教师">
          <el-select
            v-model="searchInfo.teacherId"
            placeholder="搜索条件"
            clearable
          >
            <el-option
              v-for="teacher in teacherArray"
              :key="teacher.ID"
              :label="teacher.nickName"
              :value="teacher.ID"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增设计课题</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button @click="deleteVisible = false" size="mini" type="text"
                >取消</el-button
              >
              <el-button @click="onDelete" size="mini" type="primary"
                >确定</el-button
              >
            </div>
            <el-button
              icon="el-icon-delete"
              size="mini"
              slot="reference"
              type="danger"
              >批量删除</el-button
            >
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
      <el-table-column type="selection" width="55"></el-table-column>
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{
          scope.row.CreatedAt | formatDate
        }}</template>
      </el-table-column>

      <el-table-column
        label="课题名称"
        prop="name"
        width="120"
      ></el-table-column>

      <el-table-column
        label="课题简介"
        prop="intro"
        width="120"
        show-overflow-tooltip
      ></el-table-column>

      <el-table-column
        label="教师"
        prop="teacher"
        width="120"
      ></el-table-column>

      <el-table-column label="按钮组" v-if="userInfo.authorityId === '1001'">
        <template slot-scope="scope">
          <el-button
            class="table-button"
            @click="updateShejiKeti(scope.row)"
            size="small"
            type="primary"
            icon="el-icon-edit"
          >
            详细信息
          </el-button>
        </template>
      </el-table-column>

      <el-table-column
        label="按钮组"
        v-if="userInfo.authorityId === '1000' || userInfo.authorityId === '888'"
      >
        <template slot-scope="scope">
          <el-button
            class="table-button"
            @click="updateShejiKeti(scope.row)"
            size="small"
            type="primary"
            icon="el-icon-edit"
            >变更</el-button
          >
          <el-button
            type="danger"
            icon="el-icon-delete"
            size="mini"
            @click="deleteRow(scope.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{ float: 'right', padding: '20px' }"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog
      :before-close="closeDialog"
      :visible.sync="dialogFormVisible"
      title="弹窗操作"
    >
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="课题名称:">
          <el-input
            v-model="formData.name"
            clearable
            placeholder="请输入"
          ></el-input>
        </el-form-item>

        <el-form-item label="课题简介:">
          <el-input
            v-model="formData.intro"
            clearable
            placeholder="请输入"
            type="textarea"
            autosize
          ></el-input>
        </el-form-item>

        <el-form-item label="教师"
          ><el-input v-model="formData.teacher" clearable></el-input>
        </el-form-item>

        <el-form-item label="教师信息"
          ><el-input
            v-model="formData.teacherIntroduction"
            clearable
            type="textarea"
            autosize
          ></el-input>
        </el-form-item>

        <el-form-item
          v-for="(r, index) in resources"
          :key="r.key"
          :label="'资源' + index"
        >
          <el-row :gutter="20">
            <el-col :span="16">
              <el-link :href="r.value" target="_blank">{{ r.value }}</el-link>
            </el-col>
            <el-col :span="4">
              <el-button @click.prevent="removeResource(r)">删除</el-button>
            </el-col>
          </el-row>
        </el-form-item>

        <el-form-item label="上传资源">
          <Upload @one-success="oneSuccess" />
        </el-form-item>
      </el-form>

      <div
        class="dialog-footer"
        slot="footer"
        v-if="userInfo.authorityId === '1001'"
      >
        <el-button @click="submitApply" type="primary">马上报名</el-button>
      </div>
      <div
        class="dialog-footer"
        slot="footer"
        v-if="userInfo.authorityId === '1000' || userInfo.authorityId === '888'"
      >
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createShejiKeti,
  deleteShejiKeti,
  deleteShejiKetiByIds,
  updateShejiKeti,
  findShejiKeti,
  getShejiKetiList,
} from "@/api/shejiketi"; //  此处请自行替换地址
import { getUsersByAuthorityId } from "@/api/user";
import { mapGetters } from "vuex";
import Upload from "@/view/example/simpleUploader/simpleUploader.vue";

import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "ShejiKeti",
  mixins: [infoList],
  components: {
    Upload,
  },
  data() {
    return {
      listApi: getShejiKetiList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      teacherArray: [],
      formData: {
        name: "",
        intro: "",
        teacherId: 0,
      },
      resources: [],
    };
  },
  filters: {
    formatDate: function (time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function (bool) {
      if (bool != null) {
        return bool ? "是" : "否";
      } else {
        return "";
      }
    },
  },
  computed: {
    ...mapGetters("user", ["userInfo"]),
  },
  methods: {
    removeResource(item) {
      var index = this.resources.indexOf(item);
      if (index !== -1) {
        this.resources.splice(index, 1);
      }
    },
    addResource(url) {
      this.resources.push({
        value: url,
        key: Date.now(),
      });
      console.log(this.resources);
    },
    oneSuccess(name) {
      const url = "http://101.132.104.14:8888/uploads/file/" + name;
      this.addResource(url);
    },
    submitApply() {},

    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    deleteRow(row) {
      this.$confirm("确定要删除吗?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(() => {
        this.deleteShejiKeti(row);
      });
    },
    async onDelete() {
      const ids = [];
      if (this.multipleSelection.length == 0) {
        this.$message({
          type: "warning",
          message: "请选择要删除的数据",
        });
        return;
      }
      this.multipleSelection &&
        this.multipleSelection.map((item) => {
          ids.push(item.ID);
        });
      const res = await deleteShejiKetiByIds({ ids });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功",
        });
        if (this.tableData.length == ids.length) {
          this.page--;
        }
        this.deleteVisible = false;
        this.getTableData();
      }
    },
    async updateShejiKeti(row) {
      const res = await findShejiKeti({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reshejiketi;
        if (res.data.reshejiketi.resources !== "") {
          this.resources = JSON.parse(res.data.reshejiketi.resources);
        }
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        name: "",
        intro: "",
        teacherId: 0,
      };
    },
    async deleteShejiKeti(row) {
      const res = await deleteShejiKeti({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功",
        });
        if (this.tableData.length == 1) {
          this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
      let r = this.formData;
      r.resources = JSON.stringify(this.resources);
      let res;
      switch (this.type) {
        case "create":
          res = await createShejiKeti(r);
          break;
        case "update":
          res = await updateShejiKeti(r);
          break;
        default:
          res = await createShejiKeti(r);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "创建/更改成功",
        });
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    },
  },
  async created() {
    await this.getTableData();
    let res = await getUsersByAuthorityId({
      authorityId: "1000",
    });
    if (res.code === 0) {
      this.teacherArray = res.data.list;
    }
  },
};
</script>

<style>
</style>
