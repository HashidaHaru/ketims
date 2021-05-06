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

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateShejiKeti(scope.row)" type="primary">
            详细信息
          </el-button>
          <el-button @click="checkKetiApply(scope.row)" type="primary"
            >报名信息</el-button
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
      :before-close="closeBaomingDialog"
      :visible.sync="baomingFormVisible"
      title="报名信息"
    >
      <el-form :model="baoMingForm" label-position="right" label-width="80px">
        <el-form-item label="状态:" v-if="baoMingForm.status !== 0">
          <el-tag type="warning" v-if="baoMingForm.status === 1">审核中</el-tag>
          <el-tag type="success" v-if="baoMingForm.status === 2">通过</el-tag>
          <el-tag type="danger" v-if="baoMingForm.status === 3">拒绝</el-tag>
        </el-form-item>

        <el-form-item label="审核材料:">
          <el-upload
            class="upload-demo"
            action="http://101.132.104.14:8888/fileUploadAndDownload/upload"
            :file-list="baoMingForm.pics"
            :on-success="handleSuccess"
            :on-remove="handleRemove"
            :on-preview="handlePreview"
            list-type="picture"
          >
            <el-button
              size="small"
              type="primary"
              v-if="baoMingForm.status === 0"
              >点击上传</el-button
            >
            <div
              slot="tip"
              class="el-upload__tip"
              v-if="baoMingForm.status === 0"
            >
              只能上传jpg/png文件，且不超过500kb
            </div>
          </el-upload>
        </el-form-item>

        <template v-if="baoMingForm.status === 2">
          <el-form-item
            v-for="(r, index) in baoMingForm.shejiKeti.resources"
            :key="r.key"
            :label="'资源' + index"
          >
            <el-link :href="r.value" target="_blank">{{ r.value }}</el-link>
          </el-form-item>
          <el-form-item label="提交作品">
            <Upload @one-success="oneSuccess" />
          </el-form-item>
          <el-form-item label="我的作品">
            <el-link :href="baoMingForm.ketiGroup.project" target="_blank">
              {{ baoMingForm.ketiGroup.project }}
            </el-link>
          </el-form-item>
          <el-form-item label="我的成绩">
            <h3>
              {{ baoMingForm.ketiGroup.score }}
            </h3>
          </el-form-item>

          <el-form-item label="教学评价">
            <el-row :gutter="20">
              <el-col :span="8">
                <el-input-number
                  v-model="baoMingForm.ketiGroup.review"
                  :min="0"
                  :max="100"
                ></el-input-number>
              </el-col>
              <el-col :span="4">
                <el-button @click.prevent="updateReview">提交评价</el-button>
              </el-col>
            </el-row>
          </el-form-item>
        </template>
      </el-form>

      <div class="dialog-footer" slot="footer" v-if="baoMingForm.status === 0">
        <el-button @click="submitApply" type="primary">马上报名</el-button>
      </div>
    </el-dialog>

    <el-dialog
      :before-close="closeDialog"
      :visible.sync="dialogFormVisible"
      title="详细信息"
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
      </el-form>
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
import { createKetiApply, checkKetiApply } from "@/api/ketiApply";
import { checkKetiGroup, updateKetiGroup } from "@/api/ketiGroup";
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
      baomingFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      teacherArray: [],
      formData: {
        name: "",
        intro: "",
        teacherId: 0,
        teacher: "",
        teacherIntroduction: "",
      },
      baoMingForm: {
        pics: [],
        ketiId: 0,
        status: 0,
        shejiKeti: {},
        ketiGroup: {},
        questions: [],
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
  computed: {
    ...mapGetters("user", ["userInfo"]),
  },
  methods: {
    async updateReview() {
      let r = {
        ID: this.baoMingForm.ketiGroup.ID,
        review: this.baoMingForm.ketiGroup.review,
      };
      let res = await updateKetiGroup(r);
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "评价成功",
        });
      }
    },
    async oneSuccess(name) {
      const url = "http://101.132.104.14:8888/uploads/file/" + name;
      let r = {
        ID: this.baoMingForm.ketiGroup.ID,
        project: url,
      };
      let res = await updateKetiGroup(r);
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "上传成功",
        });
      }
    },
    async submitApply() {
      let r = {
        studentId: this.userInfo.ID,
        ketiId: this.baoMingForm.ketiId,
        status: 1,
      };
      r.pics = JSON.stringify(this.baoMingForm.pics);

      const res = await createKetiApply(r);
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "申请成功",
        });
        this.closeDialog();
      }
    },
    handlePreview(file) {
      window.open(file.url);
    },
    handleSuccess(response) {
      let { name, url } = response.data.file;
      url = "http://101.132.104.14:8888/" + url;
      console.log(name, url);
      console.log(this.baoMingForm.pics);
      this.baoMingForm.pics.push({ name, url });
    },
    handleRemove(file) {
      this.baoMingForm.pics = this.baoMingForm.pics.filter(
        (item) => item.name !== file.name
      );
      console.log(this.baoMingForm.pics);
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
    async checkKetiApply(row) {
      let res = await checkKetiApply({
        ketiId: row.ID,
        studentId: this.userInfo.ID,
      });

      this.baoMingForm.status = res.data.ketiApply.status;
      this.baoMingForm.ketiId = row.ID;
      if (res.data.ketiApply.pics !== "") {
        this.baoMingForm.pics = JSON.parse(res.data.ketiApply.pics);
      }

      res = await checkKetiGroup({
        ketiId: row.ID,
        studentId: this.userInfo.ID,
      });
      this.baoMingForm.ketiGroup = res.data.ketiGroup;

      res = await findShejiKeti({ ID: row.ID });
      this.baoMingForm.shejiKeti = res.data.reshejiketi;
      this.baoMingForm.shejiKeti.resources = JSON.parse(
        this.baoMingForm.shejiKeti.resources
      );
      this.baomingFormVisible = true;
    },
    async updateShejiKeti(row) {
      const res = await findShejiKeti({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reshejiketi;
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
    closeBaomingDialog() {
      this.baomingFormVisible = false;
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
      let res;
      switch (this.type) {
        case "create":
          res = await createShejiKeti(this.formData);
          break;
        case "update":
          res = await updateShejiKeti(this.formData);
          break;
        default:
          res = await createShejiKeti(this.formData);
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
