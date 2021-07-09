<template>
  <el-container>

    <!--    <el-row>-->
    <!--      <el-carousel :interval="5000" arrow="always">-->
    <!--        <el-carousel-item v-for="item in 4" :key="item">-->
    <!--          <h3>{{ item }}</h3>-->
    <!--        </el-carousel-item>-->
    <!--      </el-carousel>-->
    <!--    </el-row>-->
    <el-row type="flex" justify="center">
      <el-col :sm="22" :md="20" :lg="18">
        <el-col :xs="24" :sm="16" :md="17">
          <el-card>
            <el-page-header @back="goBack" content="详情页面"/>
          </el-card>
          <div class="posts-list">
            <el-card
                v-for="comment in comments"
                :key="comment.content"
                :timestamp="comment.ptime"
                class="posts-list-card"
            >
              <!--              <el-dropdown class="management" v-if="right==1">-->
              <!--                <span class="el-dropdown-link">-->
              <!--                  管理<i class="el-icon-arrow-down el-icon&#45;&#45;right"></i>-->
              <!--                </span>-->
              <!--                <el-dropdown-menu slot="dropdown">-->
              <!--                  <el-dropdown-item><p @click="deletePost(blog.id)">删除</p></el-dropdown-item>-->
              <!--                </el-dropdown-menu>-->
              <!--              </el-dropdown>-->

              <!--              <div v-html="compiledMarkdown(blog.content)" class="post-content"/>-->
              <div v-html="comment.content" class="post-content" @click="goDetail(comment.pid)"/>
              <!--              <el-skeleton />-->
              <el-divider class="divider" content-position="right">{{comment.username}}</el-divider>
              <p class="post-time">{{comment.ptime }}</p>
            </el-card>
          </div>

          <el-pagination
              v-model="pageNum"

              layout="prev, pager, next"
              :page-count="totalPages"
              :hide-on-single-page="true"
              background
          />
        </el-col>
        <el-col :sm="8" :md="7" class="hidden-xs-only right-list">
          <el-card>
            <p>施工中，敬请期待</p>
          </el-card>
        </el-col>
      </el-col>
    </el-row>
  </el-container>
</template>

<script>
export default {
  name: "detail",
  data() {
    return {
      totalPages : 5,
      comments: []
    }
  },
  methods: {
    goBack() {
      window.history.length > 1 ? this.$router.go(-1) : this.$router.push('/')
    },
    getCommentList() {
      // this.$router.params.pid
      this.$ajax.get('/api/app/mock/data/179').then((response) => {
        this.comments = response.data.data.list;
      })
    },
    goDetail(pid) {
      this.$router.push({ name: 'detail', params: { pid: pid } })
    }
  },
  created() {
    this.getCommentList();
  }
}
</script>

<style scoped>
  .posts-list-card {
    margin-bottom: 1em;
  }
  .posts-list-card p {
    text-align: left;
    margin: 1em 0em;
  }
  .posts-list-card .divider {
    margin: 1em 0;
  }
  .posts-list-card .management {
    float: right;
  }
  .posts-list-card .post-time{
    text-align: right;
    color: #2c3e50;
    font-size: 50%;
    margin: 0;
  }
  .right-list {
    padding-left: 2em;
  }
</style>