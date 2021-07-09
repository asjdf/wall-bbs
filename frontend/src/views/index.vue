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
          <div class="board-list">
            <el-row>
              <el-col
                  :xs="12" :sm="6" :md="6" :lg="4" :xl="1"
                  v-for="board in boards"
                  :key="board.bname"
              >
                <el-card class="board-list-card">
                  <el-avatar :size="70" :src="board.bpic">{{board.bname}}</el-avatar>
                  <el-divider class="divider">{{board.bname}}</el-divider>
                </el-card>
              </el-col>


            </el-row>
          </div>

          <div class="posts-list">
            <el-card
                v-for="blog in blogs"
                :key="blog.content"
                :timestamp="blog.ptime"
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
              <div v-html="blog.content" class="post-content"/>
<!--              <el-skeleton />-->
              <el-divider class="divider" content-position="right">{{blog.username}}</el-divider>
              <p class="post-time">{{ blog.ptime }}</p>
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
// import HelloWorld from "@/components/HelloWorld";
export default {
  name: "index",
  data() {
    return {
      totalPages : 5,
      blogs: [],
      boards: []
    }
  },
  components: {
    // HelloWorld
  },
  methods: {
    getPostList() {
      this.$ajax.get('/api/app/mock/data/179').then((response) => {
        this.blogs = response.data.data.list;
      })
    },
    getBoardList() {
      this.$ajax.get('/api/app/mock/data/180').then((response) => {
        this.boards = response.data.data.list;
      })
    }
  },
  created() {
    this.getPostList();
    this.getBoardList();
  }
}
</script>

<style scoped>
  .el-row {
    width: 100%;
  }
  .board-list-card {
    padding-bottom: 0px;
    margin: 0.5em;
  }
  .board-list-card .divider {
    margin: 1em 0 0 0;
  }
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