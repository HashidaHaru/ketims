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
  </div>
</template>

<script>
// import musicPlayer from "./component/musicPlayer";
// import TodoList from "./component/todoList";
import { mapGetters } from "vuex";
import { getLatestNotification } from "@/api/notification";
export default {
  name: "Dashboard",
  data() {
    return {
      notification: "",
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
          label: "通知管理",
          icon: "el-icon el-icon-cpu",
          name: "notification",
          color: "#ffd666",
        },
        {
          label: "课题管理",
          icon: "el-icon el-icon-document-checked",
          name: "shejiketi",
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
    // musicPlayer, //音乐播放器
    // TodoList, //TodoList
    // RaddarChart, //雷达图
    // stackMap, //堆叠图
    // Sunburst, //旭日图
  },
  methods: {
    toTarget(name) {
      this.$router.push({ name });
    },
  },
  async created() {
    let res = await getLatestNotification();
    if (res.code == 0) {
      let content = res.data.latest.content;
      this.notification = content;
    }
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
