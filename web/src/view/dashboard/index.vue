<template>
  <div class="big">
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>最新公告</span>
      </div>
      <div>
        <p v-html="notification"></p>
      </div>
    </el-card>

    <!-- <el-row>
      <el-card shadow="hover">
        <h2>
          使用教学：<a
            style="color: #409eff"
            target="view_window"
            href="https://www.bilibili.com/video/BV1fV411y7dT/"
            >https://www.bilibili.com/video/BV1fV411y7dT/</a
          >
        </h2>
        <br />
        <h2>
          工作流教学：<a
            style="color: #409eff"
            target="view_window"
            href="https://www.bilibili.com/video/BV1Ka411F7Ji/"
            >https://www.bilibili.com/video/BV1Ka411F7Ji/</a
          >
        </h2>
        <div></div>
      </el-card>
    </el-row> -->
    <div class="shadow">
      <el-row :gutter="20">
        <el-col
          :span="4"
          v-for="(card, key) in toolCards"
          :key="key"
          @click.native="toTarget(card.name)"
          :xs="8"
        >
          <el-card shadow="hover" class="grid-content">
            <i :class="card.icon" :style="{ color: card.color }"></i>
            <p>{{ card.label }}</p>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <!-- <div class="bottom">
      <el-row :gutter="32">
        <el-col :xs="24" :sm="24" :lg="12">
          <div class="chart-player">
            <musicPlayer />
          </div>
        </el-col>
        <el-col :xs="24" :sm="24" :lg="12">
          <div class="chart-player">
            <todo-list />
          </div>
        </el-col>
      </el-row>
    </div> -->
  </div>
</template>

<script>
import musicPlayer from "./component/musicPlayer";
import TodoList from "./component/todoList";
import { mapGetters } from "vuex";
export default {
  name: "Dashboard",
  data() {
    return {
      notification: `<div id="id_4e34fbfb_0dc0_4798_87c9_312176aefde4"><p align="center"><strong><span style="FONT-SIZE: 16pt; FONT-FAMILY: 宋体;">关于课程设计的相关通知</span><span style="FONT-SIZE: 16pt"></span></strong></p>
<p><strong><span style="FONT-SIZE: 12pt; FONT-FAMILY: 宋体;">各系：</span><span style="FONT-SIZE: 12pt"></span></strong></p>
<p style="TEXT-INDENT: 30pt; LINE-HEIGHT: 200%;"><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%; FONT-FAMILY: 宋体;">课程设计是学生对主干专业课及所学知识的综合应用，是学院实现培养目标不可或缺的重要实践教学环节，更是本科教学水平评估的重要检查内容。根据学院工作的安排，现将相关工作通知如下：</span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%"></span></p>
<p style="LINE-HEIGHT: 200%"><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%; FONT-FAMILY: 宋体;">一、各系进一步清查所开课程设计的大纲和指导书，并完成补充和汇总工作。</span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%"></span></p>
<p style="LINE-HEIGHT: 200%"><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%; FONT-FAMILY: 宋体;">二、教务处根据教学任务中各系</span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%"><span style="font-family:Times New Roman;">06~07</span></span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%; FONT-FAMILY: 宋体;">学年上学期课程设计教学任务情况进行了统计，请各系根据培养方案和教学任务中课程开设情况认真核对和补全相关信息，如有统计遗漏的情况请按表格内容要求补全信息。并将纸质和电子文档于第</span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%"><span style="font-family:Times New Roman;">12</span></span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%; FONT-FAMILY: 宋体;">周星期五（</span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%"><span style="font-family:Times New Roman;">11</span></span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%; FONT-FAMILY: 宋体;">月</span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%"><span style="font-family:Times New Roman;">24</span></span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%; FONT-FAMILY: 宋体;">日</span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%; FONT-FAMILY: 宋体;">）之前返回实践教学管理科。（各系</span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%"><span style="font-family:Times New Roman;">06~07</span></span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%; FONT-FAMILY: 宋体;">年上学期课程设计计划一览表可于：</span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%"><span style="font-family:Times New Roman;">http://jwc.suse.edu.cn/download.asp?type=4</span></span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%; FONT-FAMILY: 宋体;">中下载）</span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%"></span></p>
<p style="LINE-HEIGHT: 200%"><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%; FONT-FAMILY: 宋体;">三、各系按《课程设计统计表模版》要求上报</span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%"><span style="font-family:Times New Roman;">05~06</span></span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%; FONT-FAMILY: 宋体;">学年所开课程设计一览表，并将纸质和电子文档于</span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%"><span style="font-family:Times New Roman;">13</span></span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%; FONT-FAMILY: 宋体;">周星期五（</span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%"><span style="font-family:Times New Roman;">12</span></span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%; FONT-FAMILY: 宋体;">月</span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%"><span style="font-family:Times New Roman;">1</span></span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%; FONT-FAMILY: 宋体;">日</span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%; FONT-FAMILY: 宋体;">）之前交实践教学管理科。</span><span style="FONT-SIZE: 12pt; LINE-HEIGHT: 200%"></span></p>
<p style="MARGIN: 0cm 10.5pt 0pt 0cm; TEXT-INDENT: 6pt;" align="right"><span style="FONT-SIZE: 12pt"><span style="font-family:Times New Roman;">&nbsp;</span></span></p>
<p style="MARGIN: 0cm 10.5pt 0pt 0cm; TEXT-INDENT: 6pt;" align="right"><span style="FONT-SIZE: 12pt; FONT-FAMILY: 宋体;">教</span><span style="FONT-SIZE: 12pt"><span><span style="font-family:Times New Roman;">&nbsp;&nbsp;&nbsp; </span></span></span><span style="FONT-SIZE: 12pt; FONT-FAMILY: 宋体;">务</span><span style="FONT-SIZE: 12pt"><span><span style="font-family:Times New Roman;">&nbsp;&nbsp;&nbsp; </span></span></span><span style="FONT-SIZE: 12pt; FONT-FAMILY: 宋体;">处</span><span style="FONT-SIZE: 12pt"></span></p>
</div>`,
      toolCards: [
        {
          label: "用户管理",
          icon: "el-icon el-icon-monitor",
          name: "user",
          color: "#ff9c6e",
        },
        {
          label: "角色管理",
          icon: "el-icon el-icon-setting",
          name: "authority",
          color: "#69c0ff",
        },
        {
          label: "菜单管理",
          icon: "el-icon el-icon-menu",
          name: "menu",
          color: "#b37feb",
        },
        {
          label: "代码生成器",
          icon: "el-icon el-icon-cpu",
          name: "autoCode",
          color: "#ffd666",
        },
        {
          label: "表单生成器",
          icon: "el-icon el-icon-document-checked",
          name: "formCreate",
          color: "#ff85c0",
        },
        {
          label: "关于我们",
          icon: "el-icon el-icon-user",
          name: "about",
          color: "#5cdbd3",
        },
      ],
    };
  },
  computed: {
    ...mapGetters("user", ["userInfo"]),
  },
  components: {
    musicPlayer, //音乐播放器
    TodoList, //TodoList
    // RaddarChart, //雷达图
    // stackMap, //堆叠图
    // Sunburst, //旭日图
  },
  methods: {
    toTarget(name) {
      this.$router.push({ name });
    },
  },
};
</script>

<style lang="scss" scoped>
.big {
  margin: 100px 0 0 0;
  padding-top: 0;
  background-color: rgb(243, 243, 243);
  padding-top: 15px;
  .top {
    width: 100%;
    height: 360px;
    margin-top: 20px;
    overflow: hidden;
    .chart-container {
      position: relative;
      width: 100%;
      height: 100%;
      padding: 20px;
      background-color: #fff;
    }
  }
  .mid {
    width: 100%;
    height: 380px;
    .chart-wrapper {
      height: 340px;
      background: #fff;
      padding: 16px 16px 0;
      margin-bottom: 32px;
    }
  }
  .bottom {
    width: 100%;
    height: 300px;
    // margin: 20px 0;
    .el-row {
      margin-right: 4px !important;
    }
    .chart-player {
      width: 100%;
      height: 270px;
      padding: 10px;
      background-color: #fff;
    }
  }
}
</style>
