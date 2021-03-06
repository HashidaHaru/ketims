<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="课题id">
          <el-input
            placeholder="搜索条件"
            v-model="searchInfo.ketiId"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
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
      <!-- <el-table-column type="selection" width="55"></el-table-column> -->
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{
          scope.row.CreatedAt | formatDate
        }}</template>
      </el-table-column>

      <el-table-column
        label="学生"
        prop="Student"
        width="120"
      ></el-table-column>

      <el-table-column label="课题" prop="Keti" width="120"></el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button
            class="table-button"
            @click="updateKetiApply(scope.row)"
            size="small"
            type="primary"
            icon="el-icon-edit"
            >详细信息</el-button
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
        <el-form-item label="学生"
          ><el-input
            v-model.number="formData.Student"
            clearable
            placeholder="请输入"
          ></el-input>
        </el-form-item>

        <el-form-item label="课题"
          ><el-input
            v-model.number="formData.Keti"
            clearable
            placeholder="请输入"
          ></el-input>
        </el-form-item>

        <el-form-item label="申请材料:">
          <el-upload
            :file-list="formData.pics"
            list-type="picture"
            :on-preview="handlePreview"
            action="http://101.132.104.14:8888/fileUploadAndDownload/upload"
          >
          </el-upload>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="denyApply" type="danger">拒 绝</el-button>
        <el-button @click="agreeApply" type="primary">同 意</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createKetiApply,
  deleteKetiApply,
  deleteKetiApplyByIds,
  updateKetiApply,
  findKetiApply,
  getKetiApplyList,
} from "@/api/ketiApply"; //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "KetiApply",
  mixins: [infoList],
  data() {
    return {
      listApi: getKetiApplyList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        studentId: 0,
        ketiId: 0,
        pics: "",
      },
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
  methods: {
    async denyApply() {
      let r = {};
      r.ID = this.formData.ID;
      r.status = 3;
      r.studentId = this.formData.studentId;
      r.ketiId = this.formData.ketiId;
      let res = await updateKetiApply(r);
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "操作成功",
        });
        this.closeDialog();
        this.getTableData();
      }
    },
    async agreeApply() {
      let r = {};
      r.ID = this.formData.ID;
      r.status = 2;
      r.studentId = this.formData.studentId;
      r.ketiId = this.formData.ketiId;
      let res = await updateKetiApply(r);
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "操作成功",
        });
        this.closeDialog();
        this.getTableData();
      }
    },
    handlePreview(file) {
      window.open(file.url);
    },
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
        this.deleteKetiApply(row);
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
      const res = await deleteKetiApplyByIds({ ids });
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
    async updateKetiApply(row) {
      const res = await findKetiApply({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reketiApply;
        this.dialogFormVisible = true;
        if (res.data.reketiApply.pics === "") {
          this.formData.pics = "[]";
        }
        if (res.data.reketiApply.pics !== "") {
          this.formData.pics = JSON.parse(res.data.reketiApply.pics);
        }
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        studentId: 0,
        ketiId: 0,
        pics: [],
      };
    },
    async deleteKetiApply(row) {
      const res = await deleteKetiApply({ ID: row.ID });
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
      let res;
      switch (this.type) {
        case "create":
          res = await createKetiApply(this.formData);
          break;
        case "update":
          res = await updateKetiApply(this.formData);
          break;
        default:
          res = await createKetiApply(this.formData);
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
  },
};
</script>

<style>
</style>
